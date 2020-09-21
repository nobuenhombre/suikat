package figlu

import (
	"github.com/nobuenhombre/suikat/pkg/fico"
	"github.com/nobuenhombre/suikat/pkg/fina"
	"github.com/nobuenhombre/suikat/pkg/fitree"
)

type Path string

// Возвращает только склеенный контент
// onlyExt - если не пустая строка - тогда фильтр по расширению, например ".css"
func (path *Path) GlueContent(onlyExt string, ignoreScanErr bool) (string, error) {
	// Сканируем Дерево каталогов
	list := fitree.TreeNodeListStruct{}

	scanErr := list.Scan(string(*path), fitree.StartDepth, ignoreScanErr)
	if scanErr != nil && !ignoreScanErr {
		return fico.EmptyString, scanErr
	}

	s := fico.EmptyString

	for _, dir := range list.List {
		for _, file := range dir.Files {
			allowGlue := true
			fileName := dir.Path + file.Name()

			if len(onlyExt) > 0 {
				fpi := fina.GetFilePartsInfo(fileName)
				if fpi.Ext != onlyExt {
					allowGlue = false
				}
			}

			if allowGlue {
				txtFile := fico.TxtFile(fileName)

				fileContent, err := txtFile.Read()
				if err != nil {
					return fico.EmptyString, err
				}

				s += fileContent
			}
		}
	}

	return s, nil
}

// Создает файл out-file.ext с содержимым - склеенный контент
// onlyExt - если не пустая строка - тогда фильтр по расширению, например ".css"
func (path *Path) Glue(outFile fico.TxtFile, onlyExt string, ignoreScanErr bool) error {
	content, err := path.GlueContent(onlyExt, ignoreScanErr)
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
// onlyExt - если не пустая строка - тогда фильтр по расширению, например ".css"
func (pathList *PathList) GlueContent(onlyExt string, ignoreScanErr bool) (string, error) {
	s := fico.EmptyString

	for _, path := range *pathList {
		glueContent, err := path.GlueContent(onlyExt, ignoreScanErr)
		if err != nil {
			return s, err
		}

		s += glueContent
	}

	return s, nil
}

// Создает файл out-file.ext с содержимым - склеенный контент
// onlyExt - если не пустая строка - тогда фильтр по расширению, например ".css"
func (pathList *PathList) Glue(outFile fico.TxtFile, onlyExt string, ignoreScanErr bool) error {
	content, err := pathList.GlueContent(onlyExt, ignoreScanErr)
	if err != nil {
		return err
	}

	writeErr := outFile.Write(content)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
