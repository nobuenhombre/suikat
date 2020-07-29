# CSVARSER

компонент предназначен для парсинга строк CSV файлов в структуру Go.

## Пример использования

```go
package importCSV

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	...
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/nobuenhombre/suikat/pkg/csvarser"
	"github.com/nobuenhombre/suikat/pkg/fico"
)

type SomeCSV struct {
	A pgtype.Int4    `order:"0"`
	B pgtype.Int8    `order:"1"`
	C pgtype.Varchar `order:"2"`
}

type service struct {
	FileName string
	Parser   csvarser.CsvParser
}

type ImportCSVService interface {
	MakeImport() (int64, error)
}

func NewService(fileName string) ImportCSVService {
	parser := csvarser.CsvParser{}
	return &service{
		FileName: fileName,
		Parser:   parser,
	}
}

type CsvFileReadError struct {
	FileName string
	Parent   error
}

func (e *CsvFileReadError) Error() string {
	return fmt.Sprintf("csv file [%v] read error [%v]", e.FileName, e.Parent)
}

func (srv *service) MakeImport() (int64, error) {
	// Open File
	csvFile := fico.TxtFile(srv.FileName)
	csvFileContent, err := csvFile.Read()
	if err != nil {
		return 0, &CsvFileReadError{
			FileName: srv.FileName,
			Parent:   err,
		}
	}

	srv.Parser.Init()
	srv.Parser.AddTypeParser("pgtype.Int4", srv.ParseInt4)
	srv.Parser.AddTypeParser("pgtype.Int8", srv.ParseInt8)
	srv.Parser.AddTypeParser("pgtype.Varchar", srv.ParseVarchar)

	r := csv.NewReader(bytes.NewReader([]byte(csvFileContent)))
	r.Comma = ';'
	r.Comment = '#'
	r.LazyQuotes = true
	lineNumber := int64(0)

	for {
		record, readErr := r.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return 0, &ge.IdentityPlaceError{
				Place:  "r.Read",
				Parent: readErr,
			}
		}

		log.Printf("LINE [%#v] \n", lineNumber)
		log.Printf("RECORD [%#v] \n", record)

		// Пропускаем заголовок
		if lineNumber == 0 {
			lineNumber++
			continue
		}

		csvRecord, csvRecordErr := srv.GetCSVRecord(record)
		if csvRecordErr != nil {
			return 0, &ge.IdentityPlaceError{
				Place:  "srv.GetCSVRecord",
				Parent: csvRecordErr,
			}
		}

        // Here some DB operations

		log.Printf("Document: [%#v] \n", csvRecord)
		lineNumber++
	}

	return lineNumber, nil
}

// Парсим Числовые строки = на выходе [pgtype.Int4]
func (srv *service) ParseInt4(s string) (interface{}, error) {
	var result pgtype.Int4

	normal := s
	normal = strings.Trim(normal, " ")
	if normal == "" {
		result = pgtype.Int4{
			Status: pgtype.Null,
		}
	} else {
		int4, parseErr := strconv.Atoi(normal)
		if parseErr != nil {
			return nil, parseErr
		}
		result = pgtype.Int4{
			Int:    int32(int4),
			Status: pgtype.Present,
		}
	}

	return result, nil
}

// Парсим Числовые строки = на выходе [pgtype.Int8]
func (srv *service) ParseInt8(s string) (interface{}, error) {
	var result pgtype.Int8

	normal := s
	normal = strings.Trim(normal, " ")
	if normal == "" {
		result = pgtype.Int8{
			Status: pgtype.Null,
		}
	} else {
		int8, parseErr := strconv.Atoi(normal)
		if parseErr != nil {
			return nil, parseErr
		}
		result = pgtype.Int8{
			Int:    int64(int8),
			Status: pgtype.Present,
		}
	}

	return result, nil
}

// Парсим строки Varchar
func (srv *service) ParseVarchar(s string) (interface{}, error) {
	var result pgtype.Varchar

	normal := s
	normal = strings.Trim(normal, " ")
	if normal == "" {
		result = pgtype.Varchar{
			Status: pgtype.Null,
		}
	} else {
		result = pgtype.Varchar{
			String: normal,
			Status: pgtype.Present,
		}
	}

	return result, nil
}

func (srv *service) GetCSVRecord(record []string) (*PhonesCSV, error) {
	doc := &SomeCSV{}

	err := srv.Parser.FillStructFromSlice(doc, record)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

``` 