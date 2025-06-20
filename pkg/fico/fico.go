// Package fico provides access to files content
package fico

import (
	"bufio"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"

	"github.com/andybalholm/brotli"

	"github.com/nobuenhombre/suikat/pkg/fina"
	"github.com/nobuenhombre/suikat/pkg/mimes"

	"github.com/mholt/archiver/v3"
)

const EmptyString = ""
const ExtGZ = ".gz"
const ExtBR = ".br"
const ExtB64 = ".b64"
const ExtHEX = ".hex"

type TxtFile string

func (f *TxtFile) ReadBytes() ([]byte, error) {
	file, openErr := os.Open(string(*f))
	if openErr != nil {
		return []byte{}, openErr
	}
	defer file.Close()

	b, readErr := io.ReadAll(file)
	if readErr != nil {
		return []byte{}, readErr
	}

	return b, nil
}

func (f *TxtFile) Read() (string, error) {
	b, err := f.ReadBytes()
	if err != nil {
		return EmptyString, err
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

func (f *TxtFile) WriteJSON(content interface{}) error {
	outBytes, err := json.Marshal(content)
	if err != nil {
		return err
	}

	fileContent := string(outBytes)

	return f.Write(fileContent)
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

func (f *TxtFile) GZ() error {
	data, err := f.Read()
	if err != nil {
		return err
	}

	err = f.WriteGZ(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *TxtFile) WriteBR(content string) error {
	file, createErr := os.Create(string(*f) + ExtBR)
	if createErr != nil {
		return createErr
	}

	defer file.Close()

	gz := archiver.Brotli{
		Quality: brotli.BestCompression,
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

func (f *TxtFile) BR() error {
	data, err := f.Read()
	if err != nil {
		return err
	}

	err = f.WriteBR(data)
	if err != nil {
		return err
	}

	return nil
}

func (f *TxtFile) WriteAndCompress(content string) error {
	err := f.Write(content)
	if err != nil {
		return ge.Pin(err)
	}

	err = f.WriteGZ(content)
	if err != nil {
		return ge.Pin(err)
	}

	err = f.WriteBR(content)
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}

func (f *TxtFile) ReadAsB64String() (string, error) {
	data, err := f.Read()
	if err != nil {
		return "", err
	}

	strB64 := base64.StdEncoding.EncodeToString([]byte(data))

	return strB64, nil
}

func (f *TxtFile) B64() error {
	strB64, err := f.ReadAsB64String()
	if err != nil {
		return err
	}

	txtOutFile := TxtFile(string(*f) + ExtB64)

	err = txtOutFile.Write(strB64)
	if err != nil {
		return err
	}

	return nil
}

// glue = " "
// isUpper = true
func StrBytes(in, glue string, isUpper bool) string {
	out := ""
	bytes := []byte(in)

	for _, b := range bytes {
		out += hex.EncodeToString([]byte{b}) + glue
	}

	out = strings.TrimRight(out, glue)

	if isUpper {
		out = strings.ToUpper(out)
	}

	return out
}

func (f *TxtFile) ReadAsHexString() (string, error) {
	data, err := f.Read()
	if err != nil {
		return "", err
	}

	strHex := StrBytes(data, " ", true)

	return strHex, nil
}

func (f *TxtFile) Hex() error {
	strHex, err := f.ReadAsHexString()
	if err != nil {
		return err
	}

	txtOutFile := TxtFile(string(*f) + ExtHEX)

	err = txtOutFile.Write(strHex)
	if err != nil {
		return err
	}

	return nil
}

func (f *TxtFile) DataURI() (string, error) {
	fpi := fina.GetFilePartsInfo(string(*f))

	mime := mimes.GetByExt(fpi.Ext)

	b64data, err := f.ReadAsB64String()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("data:%v;base64,%v", mime, b64data), nil
}

// Пачка файлов [имя файла]контент
type TxtFilesPack map[string]string

func (p *TxtFilesPack) Read() error {
	for fileName := range *p {
		f := TxtFile(fileName)

		c, err := f.Read()
		if err != nil {
			return err
		}

		(*p)[fileName] = c
	}

	return nil
}

func (p *TxtFilesPack) Write() error {
	for fileName, content := range *p {
		f := TxtFile(fileName)

		err := f.Write(content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *TxtFilesPack) WriteGZ() error {
	for fileName, content := range *p {
		f := TxtFile(fileName)

		err := f.WriteGZ(content)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *TxtFilesPack) Remove() error {
	for fileName := range *p {
		err := os.Remove(fileName)
		if err != nil {
			return err
		}
	}

	return nil
}
