package yank

import (
	"bytes"
	"net/url"

	"github.com/gorilla/schema"
	"github.com/nobuenhombre/suikat/pkg/ge"
	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type FormURLEncodedData struct {
	BaseContent
}

func NewFormURLEncoded(data interface{}) Content {
	return &FormURLEncodedData{
		BaseContent: BaseContent{
			MimeType: mimes.FormUrlencoded,
			Data:     data,
		},
	}
}

func (d *FormURLEncodedData) GetRawContent() (raw *RawContent, err error) {
	var (
		q   url.Values
		ok  bool
		tmp []byte
	)

	q, ok = d.Data.(url.Values)
	if !ok {
		var encoder = schema.NewEncoder()

		q = url.Values{}

		err = encoder.Encode(d.Data, q)
		if err != nil {
			err = ge.Pin(err, ge.Params{"d.Data": d.Data, "q": q})

			return
		}
	}

	tmp = []byte(q.Encode())

	buffer := bytes.NewBuffer(tmp)

	return NewRawContent(buffer, d.MimeType, d.MimeType), nil
}
