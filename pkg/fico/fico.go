package fico

import (
	"bufio"
	"compress/gzip"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mholt/archiver/v3"
)

const EmptyString = ""
const ExtGZ = ".gz"

type TxtFile string

func (f *TxtFile) Read() (string, error) {
	file, openErr := os.Open(string(*f))
	if openErr != nil {
		return EmptyString, openErr
	}
	defer file.Close()

	b, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return EmptyString, readErr
	}

	return string(b), nil
}

func (f *TxtFile) Write(content string) error {
	file, createErr := os.Create(string(*f))
	if createErr != nil {
		return createErr
	}

	defer file.Close()

	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		return writeErr
	}

	return nil
}

func (f *TxtFile) WriteGZ(content string) error {
	file, createErr := os.Create(string(*f) + ExtGZ)
	if createErr != nil {
		return createErr
	}

	defer file.Close()

	gz := archiver.Gz{
		CompressionLevel: gzip.BestCompression,
		SingleThreaded:   false,
	}

	contentReader := strings.NewReader(content)
	gzWriter := bufio.NewWriter(file)

	compressErr := gz.Compress(contentReader, gzWriter)
	if compressErr != nil {
		return compressErr
	}

	gzWriter.Flush()

	return nil
}
