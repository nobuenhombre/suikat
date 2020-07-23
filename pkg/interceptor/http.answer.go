package interceptor

import (
	"encoding/json"
	"github.com/nobuenhombre/suikat/pkg/mimes"
	"io"
	"net/http"
)

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
	case []byte:
		// Bytes - this is file
		// ContentType require
		outContent = string(v)
		outContentType = mimes.BinaryData
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
