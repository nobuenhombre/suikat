CSVARSER - CSV Parser
=====================

Моя реализация CSV парсера. Сделано на рефлексии.

Есть вот какой то CSV файл, с колонками.
И есть какая то SQL Табличка с ДРУГИМИ колонками.

Т.е. например в csv в колонке только два значения - 'фрукт', 'сиська'
А в БД мне надо записать false, true

Или вообще в колонке csv строки такого вида 'Цена 1 234.56 (EUR)'
А в БД надо сохранить только цифру 1234.56

И хочется сделать это написав не очень много строк,
но компонент можно сказать в зачаточном состоянии - он работает, но писать относительно много, но меньше чем без него.

в принципе есть несколько идей для сахарка в коде.

А пока - что же нам потребуется?

Определение CSV документа
-------------------------

Документ определяется структурой с тегами `order:"0"`, `order:"1"`, `order:"2"`...
Цифра это номер колонки в CSV.
Нумерация начинается с 0.

.. code-block:: go
  :linenos:

    import "github.com/jackc/pgtype"

    ...

    type DocumentNumber pgtype.Varchar

    // Документ - одна строка из CSV
    type DocumentCSV struct {
      Inn               pgtype.Int8      `order:"0"`
      Kpp               pgtype.Int8      `order:"1"`
      OrganizationName  pgtype.Varchar   `order:"2"`
      DocumentName      pgtype.Varchar   `order:"3"`
      DocumentNumber    DocumentNumber   `order:"4"`
      DocumentDate      pgtype.Date      `order:"5"`
      Totals            pgtype.Float8    `order:"6"`
      Vat               pgtype.Float8    `order:"7"`
    }

Определение парсеров колонок
----------------------------

В самом модуле функция парсера колонки определена следующим образом

.. code-block:: go
  :linenos:

    type ParserFunc func(s string) (interface{}, error)

Т.е. чтобы распарсить нужную колонку для нее нужно написать функцию парсер

Например

.. code-block:: go
  :linenos:

    import "github.com/jackc/pgtype"

    ...

    // Парсим Числовые строки = на выходе [pgtype.Int8]
    func ParseInt8(s string) (interface{}, error) {
      normal := strings.ReplaceAll(s, "=", "")
      normal = strings.ReplaceAll(normal, "\"", "")
      normal = strings.Trim(normal, " ")

      if normal == "" {
        return pgtype.Int8{
          Status: pgtype.Null,
        }, nil
      }

      int8, parseErr := strconv.Atoi(normal)
      if parseErr != nil {
        return nil, parseErr
      }

      return pgtype.Int8{
        Int:    int64(int8),
        Status: pgtype.Present,
      }
    }

    // Парсим Дату
    func ParseDate(s string) (interface{}, error) {
      var result pgtype.Date

    ...

Разбор файла CSV
----------------

пример

.. code-block:: go
  :linenos:

    import (
	  "encoding/csv"
	  "fmt"
	  "github.com/jackc/pgtype"
	  "github.com/jackc/pgx/v4"
	  "github.com/nobuenhombre/suikat/pkg/csvarser"
	  "io"
    )

    // Объявляем Документ - одна строка из CSV
    type DocumentCSV struct {
      Inn               pgtype.Int8      `order:"0"`
      Kpp               pgtype.Int8      `order:"1"`
      ...
    }

    // Пишем парсеры колонок
    func ParseInt8(s string) (interface{}, error) {...}
    func Parse...

    ...

    type CSVFile struct {
      FileName    string
      Parser      csvarser.CsvParser
    }

    // Добавляем функции парсеры по типу колонок
    func (f *CSVFile) InitParsers() {
	  f.Parser.Init()
	  f.Parser.AddTypeParser("pgtype.Int8", ParseInt8)
	  f.Parser.AddTypeParser("pgtype.Varchar", ParseVarchar)
	  f.Parser.AddTypeParser("pgtype.Date", ParseDate)
	  f.Parser.AddTypeParser("pgtype.Timestamp", ParseTimestamp)
	  f.Parser.AddTypeParser("pgtype.Float8", ParseFloat8)
	  f.Parser.AddTypeParser("pgtype.Bool", ParseBool)
	  f.Parser.AddTypeParser("diadoc.DocumentNumber", ParseDocumentNumber)
    }

    func (f *CSVFile) Read() error {
      // читаем содержимое файла
      csvFileContent := ...

      f.InitParsers()

      r := csv.NewReader(bytes.NewReader(csvFileContent))
      r.Comma = ';'
      r.Comment = '#'
      lineNumber := int64(0)

      for {
        record, readErr := r.Read()
        if readErr == io.EOF {
          break
        }

        if readErr != nil {
          return readErr
        }

        fmt.Printf("LINE [%#v] \n", lineNumber)
        fmt.Printf("RECORD [%#v] \n", record)

        // Пропускаем заголовок
        if lineNumber == 0 {
          lineNumber++

          continue
        }

        doc, docErr := f.GetDocument(record)
        if docErr != nil {
          return docErr
        }

        fmt.Printf("Document: [%#v] \n", doc)

        ...
        // Тут какая то обработка полученной структуры doc
        ...

        lineNumber++
      }

    }

    // Превращаем одну строку CSV в структуру DocumentCSV
    func (f *CSVFile) GetDocument(record []string) (*DocumentCSV, error) {
      doc := &DocumentCSV{}

      err := f.Parser.FillStructFromSlice(doc, record)
      if err != nil {
        return nil, err
      }

      return doc, nil
    }
