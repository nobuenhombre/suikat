package yank

import "bytes"

type BodyBuffer struct {
	*bytes.Buffer
}

func (buffer *BodyBuffer) String() string {
	if buffer == nil {
		return "nil"
	}

	return buffer.String()
}

func (buffer *BodyBuffer) Len() int {
	if buffer == nil {
		return 0
	}

	return buffer.Len()
}
