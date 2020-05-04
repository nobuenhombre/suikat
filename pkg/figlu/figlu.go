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

// Создает два файла out-file.ext и out-file.ext.gz
// с содержимым - склеенный контент
func (path *Path) Glue(outFile string, ignoreScanErr bool) error {
	content, err := path.GlueContent(ignoreScanErr)
	if err != nil {
		return err
	}

	out := fico.TxtFile(outFile)

	writeErr := out.Write(content)
	if writeErr != nil {
		return writeErr
	}

	writeGZErr := out.WriteGZ(content)
	if writeGZErr != nil {
		return writeGZErr
	}

	return nil
}
