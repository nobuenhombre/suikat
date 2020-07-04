package interceptor

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	ContentTypeJSON     = "application/json"
	ContentTypeHTML     = "text/html"
	ContentTypeImagePNG = "image/png"
	ContentTypeCSS      = "text/css"
	ContentTypeJS       = "text/javascript"
)

type HTTPAnswer struct {
	ResponseCode int
	Content      interface{}
	ContentType  string
}

func (answer *HTTPAnswer) Send(w http.ResponseWriter) {
	var outContent, outContentType string

	w.WriteHeader(answer.ResponseCode)

	switch v := answer.Content.(type) {
	case nil:
		// Empty content
		outContent = ""
	case string:
		// Just String
		outContent = v
		outContentType = ContentTypeHTML
	case []byte:
		// Bytes - this is file
		// ContentType require
		outContent = string(v)
	default:
		// Struct or map
		outBytes, outError := json.Marshal(answer.Content)
		if outError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		outContent = string(outBytes)
		outContentType = ContentTypeJSON
	}

	if len(outContent) > 0 {
		if len(answer.ContentType) > 0 {
			outContentType = answer.ContentType
		}

		w.Header().Add("content-type", outContentType)

		_, err := io.WriteString(w, outContent)
		if err != nil {
			return
		}
	}
}
