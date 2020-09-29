package interceptor

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const MaxFormSizeMemory = 10 << 20

type UploadedFile struct {
	Name string
	Size int64
	Mime string
	Data []byte
}

func GetFile(key string, r *http.Request) (*UploadedFile, error) {
	err := r.ParseMultipartForm(MaxFormSizeMemory)
	if err != nil {
		return nil, err
	}

	file, handler, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &UploadedFile{
		Name: handler.Filename,
		Size: handler.Size,
		Mime: fmt.Sprintf("%+v", handler.Header),
		Data: fileData,
	}, nil
}
