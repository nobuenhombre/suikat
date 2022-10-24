package yank

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type FormMultipartData struct {
	BaseContent
}

func NewFormMultipart(data interface{}) Content {
	return &FormMultipartData{
		BaseContent: BaseContent{
			MimeType: mimes.JSON,
			Data:     data,
		},
	}
}

func (d *FormMultipartData) GetRawContent() (raw *RawContent, err error) {
	body := &bytes.Buffer{}

	w := multipart.NewWriter(body)

	for k, v := range d.Data.(url.Values) {
		for _, iv := range v {
			if strings.HasPrefix(k, "@") {
				// file
				//-----
				err = addFile(w, k[1:], iv)
				if err != nil {
					err = ge.Pin(err, ge.Params{"d.Data": d.Data})

					return
				}
			} else {
				// form value
				//-----------
				err = w.WriteField(k, iv)
				if err != nil {
					err = ge.Pin(err, ge.Params{"d.Data": d.Data, "k": k, "iv": iv})

					return
				}
			}
		}
	}

	err = w.Close()
	if err != nil {
		err = ge.Pin(err, ge.Params{"d.Data": d.Data})

		return
	}

	return NewRawContent(body, d.MimeType, w.FormDataContentType()), nil
}

// addFile - Add a file to the multipart request
// ----------------------------------------------
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return ge.Pin(err, ge.Params{"path": path})
	}

	defer file.Close()

	fileBasePath := filepath.Base(path)

	part, err := w.CreateFormFile(fieldName, fileBasePath)
	if err != nil {
		return ge.Pin(err, ge.Params{"fieldName": fieldName, "fileBasePath": fileBasePath})
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return ge.Pin(err, ge.Params{"part": part, "file": file})
	}

	return err
}
