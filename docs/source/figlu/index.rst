FIGLU - FIles GLUe
==================

функции для склеивания содержимого нескольких файлов в один

func (path *Path) GlueContent(onlyExt string, ignoreScanErr bool) (string, error)

сканирует путь, получает список файлов
если указано расширение, т.е. len(onlyExt) > 0 то читаются только файлы с указанным расширением
их контент в виде строк склеивается в одну строку которую и возвращает данная функция.
для одного пути с одним корнем.


func (path *Path) Glue(outFile fico.TxtFile, onlyExt string, ignoreScanErr bool) error

сканирует путь, получает список файлов
если указано расширение, т.е. len(onlyExt) > 0 то читаются только файлы с указанным расширением
их контент в виде строк склеивается в одну строку которая записывается в файл outFile.


func (pathList *PathList) GlueContent(onlyExt string, ignoreScanErr bool) (string, error)

обходит список путей pathList в цикле и делает для каждого пути GlueContent.
а потом все результаты склеивает и возвращает строкой. Для путей с разными корнями.



func (pathList *PathList) Glue(outFile fico.TxtFile, onlyExt string, ignoreScanErr bool) error

обходит список путей pathList в цикле и делает для каждого пути GlueContent.
а потом все результаты склеивает и возвращает строку которая записывается в файл outFile.
