// Package futi provides util functions to work with files
package futi

import (
	"fmt"
	"io"
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

func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
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

type IsNotRegularFileError struct {
	File string
}

func (e *IsNotRegularFileError) Error() string {
	return fmt.Sprintf("%s is not a regular file", e.File)
}

func Copy(inFile, outFile string) error {
	if inFile == outFile {
		return nil
	}

	sourceFileStat, err := os.Stat(inFile)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return &IsNotRegularFileError{
			File: inFile,
		}
	}

	source, err := os.Open(inFile)
	if err != nil {
		return err
	}

	defer source.Close()

	destination, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	return nil
}

func Move(inFile, outFile string) error {
	if inFile == outFile {
		return nil
	}

	err := Copy(inFile, outFile)
	if err != nil {
		return err
	}

	err = os.Remove(inFile)
	if err != nil {
		return err
	}

	return nil
}
