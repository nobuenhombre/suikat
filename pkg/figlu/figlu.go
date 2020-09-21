package figlu

import (
	"github.com/nobuenhombre/suikat/pkg/fico"
	"github.com/nobuenhombre/suikat/pkg/fitree"
)

type Path string

// Возвращает только склеенный контент
func (path *Path) GlueContent(ignoreScanErr bool) (string, error) {
	// Сканируем Дерево каталогов
	list := fitree.TreeNodeListStruct{}

	scanErr := list.Scan(string(*path), fitree.StartDepth, ignoreScanErr)
	if scanErr != nil && !ignoreScanErr {
		return fico.EmptyString, scanErr
	}

	s := fico.EmptyString

	for _, dir := range list.List {
		for _, file := range dir.Files {
			txtFile := fico.TxtFile(dir.Path + file.Name())

			fileContent, err := txtFile.Read()
			if err != nil {
				return fico.EmptyString, err
			}

			s += fileContent
		}
	}

	return s, nil
}

// Создает файл out-file.ext с содержимым - склеенный контент
func (path *Path) Glue(outFile fico.TxtFile, ignoreScanErr bool) error {
	content, err := path.GlueContent(ignoreScanErr)
	if err != nil {
		return err
	}

	writeErr := outFile.Write(content)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

type PathList []Path

// Возвращает только склеенный контент
func (pathList *PathList) GlueContent(ignoreScanErr bool) (string, error) {
	s := fico.EmptyString
	for _, path := range *pathList {
		s, err := path.GlueContent(ignoreScanErr)
		if err != nil {
			return s, err
		}
	}

	return s, nil
}

// Создает файл out-file.ext с содержимым - склеенный контент
func (pathList *PathList) Glue(outFile fico.TxtFile, ignoreScanErr bool) error {
	content, err := pathList.GlueContent(ignoreScanErr)
	if err != nil {
		return err
	}

	writeErr := outFile.Write(content)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
