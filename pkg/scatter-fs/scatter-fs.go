package scatterfs

import (
	"errors"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

var (
	NoDirectoriesWithFreeSpaceAvailableError = errors.New("no directories with free space available")
	InvalidPathError                         = errors.New("invalid path")
	CantRemoveRootError                      = errors.New("cannot remove root")
	DirectoryNotFoundInAnyBaseDirectoryError = errors.New("directory not found in any base directory")
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
	MoveDir(oldDir, newDir string) error
}

// FileSystem реализация виртуальной файловой системы
type FileSystem struct {
	dirs        []string
	freePercent float64
	rand        *rand.Rand
}

// New конструктор
func New(dirs []string, freePercent float64) IFileSystem {
	return &FileSystem{
		dirs:        dirs,
		freePercent: freePercent,
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
	for i, dir := range sfs.dirs {
		realPath := filepath.Join(dir, path)
		if _, err := os.Stat(realPath); err == nil {
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

func (sfs *FileSystem) MoveDir(oldDir, newDir string) error {
	oldDir = filepath.Clean(oldDir)
	newDir = filepath.Clean(newDir)

	// 1. Собираем все реальные файлы из всех базовых каталогов
	type fileInfo struct {
		realPath string
		baseIdx  int
		relPath  string // Относительный путь от oldDir
	}

	var files []fileInfo

	// Сканируем каждый базовый каталог
	for i, baseDir := range sfs.dirs {
		realOldDir := filepath.Join(baseDir, oldDir)

		_ = filepath.Walk(realOldDir, func(realPath string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil // Пропускаем ошибки и директории
			}

			// Вычисляем относительный путь внутри виртуальной директории
			relPath, err := filepath.Rel(realOldDir, realPath)
			if err != nil {
				return nil
			}

			files = append(files, fileInfo{
				realPath: realPath,
				baseIdx:  i,
				relPath:  relPath,
			})

			return nil
		})
	}

	if len(files) == 0 {
		return ge.Pin(DirectoryNotFoundInAnyBaseDirectoryError)
	}

	// 2. Перемещаем файлы
	for _, file := range files {
		newVirtualPath := filepath.Join(newDir, file.relPath)

		// Выбираем новый базовый каталог
		available, err := sfs.getAvailableDirs()
		if err != nil {
			return ge.Pin(err)
		}
		newBaseIdx := available[sfs.rand.Intn(len(available))]
		newRealPath := filepath.Join(sfs.dirs[newBaseIdx], newVirtualPath)

		// Создаём целевую директорию
		if err := os.MkdirAll(filepath.Dir(newRealPath), 0755); err != nil {
			return ge.Pin(err)
		}

		// Перемещаем файл
		if err := os.Rename(file.realPath, newRealPath); err != nil {
			return ge.Pin(err)
		}
	}

	// 3. Удаляем пустые директории в исходном расположении
	sfs.cleanupEmptyDirs(oldDir)

	return nil
}

func (sfs *FileSystem) cleanupEmptyDirs(virtualDir string) {
	for _, baseDir := range sfs.dirs {
		realDir := filepath.Join(baseDir, virtualDir)
		_ = filepath.Walk(realDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() {
				return nil
			}

			// Пытаемся удалить, если директория пуста
			_ = os.Remove(path)

			return nil
		})
	}
}
