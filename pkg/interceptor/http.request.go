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

func GetFile(key string, r *http.Request) (upFile *UploadedFile, err error) {
	err = r.ParseMultipartForm(MaxFormSizeMemory)
	if err != nil {
		return nil, err
	}

	file, fileHeader, err := r.FormFile(key)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = file.Close()
	}()

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &UploadedFile{
		Name: fileHeader.Filename,
		Size: fileHeader.Size,
		Mime: fmt.Sprintf("%+v", fileHeader.Header),
		Data: fileData,
	}, nil
}
