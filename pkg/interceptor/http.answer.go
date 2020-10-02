package interceptor

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/nobuenhombre/suikat/pkg/mimes"
)

type FileData struct {
	Name string
	Size int64
	Data []byte
}

const (
	BrowserCacheLifeTimeWeek = 604800 // week 7 days
)

type HTTPAnswer struct {
	// Configure
	//--------------------------
	GZipped              bool
	BrowserCached        bool
	ETagUsed             bool
	GZipLevel            int // gzip.BestCompression
	BrowserCacheLifeTime int

	// Data
	//--------------------------
	ResponseCode int
	Content      interface{}
	ContentType  string
	ETag         string
}

func (answer *HTTPAnswer) gzipData(data *[]byte) (gzData []byte, err error) {
	var (
		tmp      bytes.Buffer
		gzWriter *gzip.Writer
	)

	gzWriter, err = gzip.NewWriterLevel(&tmp, answer.GZipLevel)
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

func (answer *HTTPAnswer) enableBrowserCacheHeaders(w http.ResponseWriter) {
	if answer.ETagUsed {
		w.Header().Add(
			"Cache-Control",
			fmt.Sprintf("private, max-age=%v, must-revalidate", answer.BrowserCacheLifeTime),
		)
	} else {
		w.Header().Add("Pragma", "public")
		w.Header().Add(
			"Cache-Control",
			fmt.Sprintf("private, max-age=%v", answer.BrowserCacheLifeTime),
		)
	}

	lastModifiedTime := time.Now()
	expiredTime := lastModifiedTime.Add(time.Second * time.Duration(answer.BrowserCacheLifeTime))

	w.Header().Add("Last-Modified", lastModifiedTime.Format(time.RFC1123))
	w.Header().Add("Expires", expiredTime.Format(time.RFC1123))
}

func (answer *HTTPAnswer) disableBrowserCacheHeaders(w http.ResponseWriter) {
	if answer.ETagUsed {
		w.Header().Add("Cache-Control", "no-cache, must-revalidate")
	} else {
		w.Header().Add("Cache-Control", "no-cache")
	}

	lastModifiedTime := time.Date(1997, 7, 26, 5, 0, 0, 0, time.UTC)
	w.Header().Add("Expires", lastModifiedTime.Format(time.RFC1123))
}

func (answer *HTTPAnswer) setBrowserCacheHeaders(w http.ResponseWriter) {
	if answer.BrowserCached {
		answer.enableBrowserCacheHeaders(w)

		return
	}

	answer.disableBrowserCacheHeaders(w)
}

func (answer *HTTPAnswer) SetETag(w http.ResponseWriter) {
	w.Header().Add("Etag", answer.ETag)
}

func (answer *HTTPAnswer) CheckETag(r *http.Request, w http.ResponseWriter) {
	match := r.Header.Get("If-None-Match")
	if match == answer.ETag {
		w.WriteHeader(http.StatusNotModified)

		return
	}
}

func (answer *HTTPAnswer) sendData(data *[]byte, w http.ResponseWriter) (err error) {
	if len(*data) == 0 {
		return
	}

	if answer.GZipped {
		w.Header().Set("Content-Encoding", "gzip")

		*data, err = answer.gzipData(data)
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
	answer.setBrowserCacheHeaders(w)
	w.WriteHeader(answer.ResponseCode)

	err = answer.sendData(&data, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
