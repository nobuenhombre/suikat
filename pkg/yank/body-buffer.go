package yank

import "bytes"

type BodyBuffer struct {
	*bytes.Buffer
}

func (buffer *BodyBuffer) String() string {
	if buffer.Buffer == nil {
		return "nil"
	}

	return buffer.Buffer.String()
}

func (buffer *BodyBuffer) Len() int {
	if buffer.Buffer == nil {
		return 0
	}

	return buffer.Buffer.Len()
}
