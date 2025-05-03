package gotemplates

import (
	"os"
	"path/filepath"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type Path string

func (p Path) GetSubDirectories() ([]string, error) {
	var directories []string

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			directories = append(directories, path)
		}

		return nil
	}

	err := filepath.Walk(string(p), walkFunc)
	if err != nil {
		return nil, ge.Pin(err)
	}

	return directories[1:], nil
}
