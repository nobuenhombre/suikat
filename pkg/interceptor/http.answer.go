package interceptor

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type FileData struct {
	Name string
	Size int64
	Data []byte
}

type HTTPAnswer struct {
	ResponseCode int
	Content      interface{}
	ContentType  string
}

func (answer *HTTPAnswer) Send(w http.ResponseWriter) {
	var outContent, outContentType string

	switch v := answer.Content.(type) {
	case nil:
		// Empty content
		outContent = ""
		outContentType = mimes.Text
	case string:
		// Just String
		outContent = v
		outContentType = mimes.HyperTextMarkupLanguage
	case FileData:
		// Bytes - this is file
		// ContentType require
		outContent = string(v.Data)
		outContentType = mimes.BinaryData

		//Send the headers
		w.Header().Set("Content-Disposition", "attachment; filename="+v.Name)
		w.Header().Set("Content-Length", strconv.FormatInt(v.Size, 10))
	default:
		// Struct or map
		outBytes, outError := json.Marshal(answer.Content)
		if outError != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		outContent = string(outBytes)
		outContentType = mimes.JSON
	}

	if len(answer.ContentType) > 0 {
		outContentType = answer.ContentType
	}

	w.Header().Add("content-type", outContentType)
	w.WriteHeader(answer.ResponseCode)

	if len(outContent) > 0 {
		_, err := io.WriteString(w, outContent)
		if err != nil {
			return
		}
	}
}
