package yank

import (
	"bytes"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type BytesData struct {
	BaseContent
}

func NewBytes(data interface{}, mimeType string) Content {
	return &BytesData{
		BaseContent: BaseContent{
			MimeType: mimeType,
			Data:     data,
		},
	}
}

func (d *BytesData) GetRawContent() (raw *RawContent, err error) {
	tmp, ok := d.Data.([]byte)
	if !ok {
		return nil, ge.Pin(&ge.UndefinedTypeError{}, ge.Params{"d.Data": d.Data})
	}

	buffer := bytes.NewBuffer(tmp)

	return NewRawContent(buffer, d.MimeType, d.MimeType), nil
}
