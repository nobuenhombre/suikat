package fitree

import (
	"io/ioutil"
	"os"
)

// Один Узел Дерева
type TreeNodeStruct struct {
	Path         string
	Name         string
	Depth        int
	Files        []os.FileInfo
	FilesCount   int
	SubDirs      []os.FileInfo
	SubDirsCount int
}

// Заполняем структуру при сканировании дерева
func (node *TreeNodeStruct) Fill(path string, depth int) error {
	DirInfo, DirInfoErr := os.Lstat(path)
	if DirInfoErr != nil {
		return DirInfoErr
	}

	files, readDirErr := ioutil.ReadDir(path)
	if readDirErr != nil {
		return readDirErr
	}

	node.Path = path
	node.Name = DirInfo.Name()
	node.Depth = depth
	for _, f := range files {
		if f.IsDir() {
			node.SubDirs = append(node.SubDirs, f)
		} else {
			node.Files = append(node.Files, f)
		}
	}
	node.FilesCount = len(node.Files)
	node.SubDirsCount = len(node.SubDirs)

	return nil
}
