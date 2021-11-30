package yank

type BaseContent struct {
	MimeType string
	Data     interface{}
}

type Content interface {
	GetRawContent() (raw *RawContent, err error)
}
