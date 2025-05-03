package gotemplates

import (
	"bytes"
	htmlTemplate "html/template"
	"os"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

const MaskGoHTML = "*.gohtml"

type HTMLPath Path

func (hp HTMLPath) GetMask() string {
	return MaskGoHTML
}

func (hp HTMLPath) GetTemplate() (*htmlTemplate.Template, error) {
	// Create a template for parsing all directories
	fullPath := string(hp) + string(os.PathSeparator) + hp.GetMask()
	t, err := htmlTemplate.ParseGlob(fullPath)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"path": fullPath})
	}

	subDirectories, err := Path(hp).GetSubDirectories()
	if err != nil {
		return nil, ge.Pin(err)
	}

	// Parse all directories (without the root directory)
	for _, path := range subDirectories {
		fullPath = path + string(os.PathSeparator) + hp.GetMask()
		t, err = t.ParseGlob(fullPath)
		if err != nil {
			return nil, ge.Pin(err, ge.Params{"path": fullPath})
		}
	}

	return t, nil
}

func (hp HTMLPath) HTML(data interface{}) (string, error) {
	t, err := hp.GetTemplate()
	if err != nil {
		return "", ge.Pin(err)
	}

	buf := new(bytes.Buffer)

	err = t.ExecuteTemplate(buf, "index", data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
