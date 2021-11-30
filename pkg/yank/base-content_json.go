package yank

import (
	"bytes"
	"encoding/json"

	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type JSONData struct {
	BaseContent
}

func NewJSON(data interface{}) Content {
	return &JSONData{
		BaseContent: BaseContent{
			MimeType: mimes.JSON,
			Data:     data,
		},
	}
}

func (d *JSONData) GetRawContent() (raw *RawContent, err error) {
	var tmp []byte

	tmp, err = json.Marshal(d.Data)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"d.Data": d.Data})
	}

	buffer := bytes.NewBuffer(tmp)

	return NewRawContent(buffer, d.MimeType, d.MimeType), nil
}
