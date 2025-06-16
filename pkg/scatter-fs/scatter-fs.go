package scatterfs

import (
	"errors"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

var (
	NoDirectoriesWithFreeSpaceAvailableError = errors.New("no directories with free space available")
	InvalidPathError                         = errors.New("invalid path")
	CantRemoveRootError                      = errors.New("cannot remove root")
)

// IFileSystem интерфейс виртуальной файловой системы
type IFileSystem interface {
	Create(path string) (*os.File, error)
	Open(path string) (*os.File, error)
	MkdirAll(path string, perm fs.FileMode) error
	Remove(path string) error
	RemoveAll(path string) error
	Stat(path string) (fs.FileInfo, error)
	ReadDir(path string) ([]fs.DirEntry, error)
}

// FileSystem реализация виртуальной файловой системы
type FileSystem struct {
	dirs        []string
	freePercent float64
	fileMap     map[string]int
	fileMapLock sync.RWMutex
	rand        *rand.Rand
}

// New конструктор
func New(dirs []string, freePercent float64) IFileSystem {
	return &FileSystem{
		dirs:        dirs,
		freePercent: freePercent,
		fileMap:     make(map[string]int),
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sfs *FileSystem) getAvailableDirs() ([]int, error) {
	var available []int

	for i, dir := range sfs.dirs {
		stat, err := getDiskUsage(dir)
		if err != nil {
			return nil, ge.Pin(err)
		}

		freePercent := float64(stat.Bavail) / float64(stat.Blocks) * 100
		if freePercent >= sfs.freePercent {
			available = append(available, i)
		}
	}

	if len(available) == 0 {
		return nil, ge.Pin(NoDirectoriesWithFreeSpaceAvailableError)
	}

	return available, nil
}

func getDiskUsage(path string) (*syscall.Statfs_t, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return nil, ge.Pin(err)
	}
	return &stat, nil
}

func (sfs *FileSystem) resolveRealPath(path string) (string, int, error) {
	sfs.fileMapLock.RLock()
	if idx, ok := sfs.fileMap[path]; ok {
		sfs.fileMapLock.RUnlock()
		return filepath.Join(sfs.dirs[idx], path), idx, nil
	}
	sfs.fileMapLock.RUnlock()

	for i, dir := range sfs.dirs {
		realPath := filepath.Join(dir, path)
		if _, err := os.Stat(realPath); err == nil {
			sfs.fileMapLock.Lock()
			sfs.fileMap[path] = i
			sfs.fileMapLock.Unlock()
			return realPath, i, nil
		}
	}

	return "", -1, ge.Pin(fs.ErrNotExist)
}

func (sfs *FileSystem) Create(path string) (*os.File, error) {
	if path == "" {
		return nil, ge.Pin(InvalidPathError)
	}

	available, err := sfs.getAvailableDirs()
	if err != nil {
		return nil, ge.Pin(err)
	}

	selectedIdx := available[sfs.rand.Intn(len(available))]
	selectedDir := sfs.dirs[selectedIdx]
	realPath := filepath.Join(selectedDir, path)

	if err := os.MkdirAll(filepath.Dir(realPath), 0755); err != nil {
		return nil, ge.Pin(err)
	}

	file, err := os.Create(realPath)
	if err != nil {
		return nil, ge.Pin(err)
	}

	sfs.fileMapLock.Lock()
	sfs.fileMap[path] = selectedIdx
	sfs.fileMapLock.Unlock()

	return file, nil
}

func (sfs *FileSystem) Open(path string) (*os.File, error) {
	realPath, _, err := sfs.resolveRealPath(path)
	if err != nil {
		return nil, ge.Pin(err)
	}

	file, err := os.Open(realPath)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return file, nil
}

func (sfs *FileSystem) MkdirAll(path string, perm fs.FileMode) error {
	if path == "" {
		return nil
	}

	available, err := sfs.getAvailableDirs()
	if err != nil {
		return ge.Pin(err)
	}

	selectedIdx := available[sfs.rand.Intn(len(available))]
	selectedDir := sfs.dirs[selectedIdx]
	realPath := filepath.Join(selectedDir, path)

	if err := os.MkdirAll(realPath, perm); err != nil {
		return ge.Pin(err)
	}

	sfs.fileMapLock.Lock()
	sfs.fileMap[path] = selectedIdx
	sfs.fileMapLock.Unlock()

	return nil
}

func (sfs *FileSystem) Remove(path string) error {
	if path == "" {
		return ge.Pin(CantRemoveRootError)
	}

	realPath, _, err := sfs.resolveRealPath(path)
	if err != nil {
		return ge.Pin(err)
	}

	if err := os.Remove(realPath); err != nil {
		return ge.Pin(err)
	}

	sfs.fileMapLock.Lock()
	delete(sfs.fileMap, path)
	sfs.fileMapLock.Unlock()

	return nil
}

func (sfs *FileSystem) RemoveAll(path string) error {
	if path == "" {
		return ge.Pin(CantRemoveRootError)
	}

	realPath, _, err := sfs.resolveRealPath(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return ge.Pin(err)
	}

	if err := os.RemoveAll(realPath); err != nil {
		return ge.Pin(err)
	}

	sfs.fileMapLock.Lock()
	delete(sfs.fileMap, path)
	sfs.fileMapLock.Unlock()

	return nil
}

func (sfs *FileSystem) Stat(path string) (fs.FileInfo, error) {
	realPath, _, err := sfs.resolveRealPath(path)
	if err != nil {
		return nil, ge.Pin(err)
	}

	info, err := os.Stat(realPath)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return info, nil
}

func (sfs *FileSystem) ReadDir(path string) ([]fs.DirEntry, error) {
	entriesMap := make(map[string]fs.DirEntry)

	for _, dir := range sfs.dirs {
		realPath := filepath.Join(dir, path)
		entries, err := os.ReadDir(realPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, ge.Pin(err)
		}

		for _, entry := range entries {
			if _, exists := entriesMap[entry.Name()]; !exists {
				entriesMap[entry.Name()] = entry
			}
		}
	}

	result := make([]fs.DirEntry, 0, len(entriesMap))
	for _, entry := range entriesMap {
		result = append(result, entry)
	}

	return result, nil
}
