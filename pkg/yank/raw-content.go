package yank

import (
	"bytes"
	"fmt"
	"io"
)

type RawContent struct {
	Body          string
	MimeType      string
	ContentType   string
	ContentLength int
	reader        io.Reader
}

func NewRawContent(buffer *bytes.Buffer, mimeType string, contentType string) *RawContent {
	bb := &BodyBuffer{Buffer: buffer}

	return &RawContent{
		Body:          bb.String(),
		MimeType:      mimeType,
		ContentType:   contentType,
		ContentLength: bb.Len(),
		reader:        buffer,
	}
}

func (rc *RawContent) GetReader() io.Reader {
	return rc.reader
}

func (rc *RawContent) AddHeaders(r *HTTPRequest) {
	if rc.reader != nil {
		// Set the Boundary in the Content-Type
		r.Header.Set("Content-Type", rc.ContentType)

		// Set Content-Length
		r.Header.Set("Content-Length", fmt.Sprintf("%d", rc.ContentLength))
	}
}
