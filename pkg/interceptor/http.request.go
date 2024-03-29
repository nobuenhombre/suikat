package interceptor

import (
	"fmt"
	"io"
	"net/http"

	"github.com/nobuenhombre/suikat/pkg/ge"
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
		return nil, ge.Pin(err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = ge.Pin(closeErr, ge.Params{ge.BaseError: err})
		}
	}()

	fileData, err := io.ReadAll(file)
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
