package futi

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

type DeleteFileError struct {
	FileName string
	Parent   error
}

func (e *DeleteFileError) Error() string {
	return fmt.Sprintf("delete file [%v] error [%v]", e.FileName, e.Parent)
}

func Delete(fileName string) error {
	_, err := os.Stat(fileName)
	if err != nil {
		return &DeleteFileError{
			FileName: fileName,
			Parent:   err,
		}
	}

	err = os.Remove(fileName)
	if err != nil {
		return &DeleteFileError{
			FileName: fileName,
			Parent:   err,
		}
	}

	return nil
}

type CreateTempFileError struct {
	Dir     string
	Pattern string
	Data    *[]byte
	Parent  error
}

func (e *CreateTempFileError) Error() string {
	return fmt.Sprintf("create temp file error %#v", e)
}

func CreateTempFile(dir, pattern string, data *[]byte) (string, error) {
	tempFile, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		return "", &CreateTempFileError{
			Dir:     dir,
			Pattern: pattern,
			Data:    data,
			Parent:  err,
		}
	}

	_, err = tempFile.Write(*data)
	if err != nil {
		return "", &CreateTempFileError{
			Dir:     dir,
			Pattern: pattern,
			Data:    data,
			Parent:  err,
		}
	}

	err = tempFile.Close()
	if err != nil {
		return "", &CreateTempFileError{
			Dir:     dir,
			Pattern: pattern,
			Data:    data,
			Parent:  err,
		}
	}

	return tempFile.Name(), nil
}
