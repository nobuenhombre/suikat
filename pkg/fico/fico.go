package fico

import (
	"bufio"
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/fina"
	"github.com/nobuenhombre/suikat/pkg/mimes"

	"github.com/mholt/archiver/v3"
)

const EmptyString = ""
const ExtGZ = ".gz"

type TxtFile string

func (f *TxtFile) ReadBytes() ([]byte, error) {
	file, openErr := os.Open(string(*f))
	if openErr != nil {
		return []byte{}, openErr
	}
	defer file.Close()

	b, readErr := ioutil.ReadAll(file)
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

	txtOutFile := TxtFile(string(*f) + ".b64")

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

	txtOutFile := TxtFile(string(*f) + ".hex")

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
