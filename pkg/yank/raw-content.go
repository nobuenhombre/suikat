package yank

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nobuenhombre/suikat/pkg/adapt"
)

type RawContent struct {
	Body          string
	MimeType      string
	ContentType   string
	ContentLength int
	buffer        *bytes.Buffer
}

func NewRawContent(buffer *bytes.Buffer, mimeType string, contentType string) *RawContent {
	bb := &BodyBuffer{Buffer: buffer}

	return &RawContent{
		Body:          bb.String(),
		MimeType:      mimeType,
		ContentType:   contentType,
		ContentLength: bb.Len(),
		buffer:        buffer,
	}
}

func (rc *RawContent) GetBuffer() io.Reader {
	if adapt.IsNil(rc.buffer) {
		return nil
	}

	return rc.buffer
}

func (rc *RawContent) AddHeaders(r *HTTPRequest) {
	if rc.buffer != nil {
		// Set the Boundary in the Content-Type
		r.Header.Set("Content-Type", rc.ContentType)

		// Set Content-Length
		r.Header.Set("Content-Length", fmt.Sprintf("%d", rc.ContentLength))
	}
}
