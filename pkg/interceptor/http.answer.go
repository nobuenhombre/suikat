package interceptor

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
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
	GZipped      bool
	GZipLevel    int // gzip.BestCompression
}

func (answer *HTTPAnswer) gzipData(data *[]byte, w http.ResponseWriter) (gzData []byte, err error) {
	w.Header().Set("Content-Encoding", "gzip")

	var tmp bytes.Buffer

	gzWriter, err := gzip.NewWriterLevel(&tmp, answer.GZipLevel)
	if err != nil {
		return
	}

	defer func() {
		err = gzWriter.Close()
	}()

	_, err = gzWriter.Write(*data)
	if err != nil {
		return
	}

	gzData = tmp.Bytes()

	return
}

func (answer *HTTPAnswer) getData() (outBytes []byte, err error) {
	switch v := answer.Content.(type) {
	case nil:
		// Empty content
		outBytes = []byte("")
	case string:
		// Just String
		outBytes = []byte(v)
	case FileData:
		// Bytes - this is file
		// ContentType require
		outBytes = v.Data
	default:
		// Struct or map
		outBytes, err = json.Marshal(answer.Content)
		if err != nil {
			return
		}
	}

	return
}

func (answer *HTTPAnswer) setContentTypeHeaders(w http.ResponseWriter) {
	var outContentType string

	switch v := answer.Content.(type) {
	case nil:
		// Empty content
		outContentType = mimes.Text
	case string:
		// Just String
		outContentType = mimes.HyperTextMarkupLanguage
	case FileData:
		// Bytes - this is file
		// ContentType require
		outContentType = mimes.BinaryData

		//Send the headers
		w.Header().Set("Content-Disposition", "attachment; filename="+v.Name)
		w.Header().Set("Content-Length", strconv.FormatInt(v.Size, 10))
	default:
		// Struct or map
		outContentType = mimes.JSON
	}

	if len(answer.ContentType) > 0 {
		outContentType = answer.ContentType
	}

	w.Header().Add("Content-Type", outContentType)
}

func (answer *HTTPAnswer) sendData(data *[]byte, w http.ResponseWriter) (err error) {
	if len(*data) == 0 {
		return
	}

	if answer.GZipped {
		*data, err = answer.gzipData(data, w)
		if err != nil {
			return
		}
	}

	_, err = w.Write(*data)
	if err != nil {
		return
	}

	return
}

func (answer *HTTPAnswer) Send(w http.ResponseWriter) {
	data, err := answer.getData()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	answer.setContentTypeHeaders(w)

	w.WriteHeader(answer.ResponseCode)

	err = answer.sendData(&data, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
