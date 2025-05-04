package gt

import (
	"bytes"
	"os"
	textTemplate "text/template"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

const MaskGoTPL = "*.go.tpl"

type TextPath Path

func (tp TextPath) GetMask() string {
	return MaskGoTPL
}

func (tp TextPath) GetTemplate() (*textTemplate.Template, error) {
	// Create a template for parsing all directories
	fullPath := string(tp) + string(os.PathSeparator) + tp.GetMask()
	t, err := textTemplate.ParseGlob(fullPath)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"path": fullPath})
	}

	subDirectories, err := Path(tp).GetSubDirectories()
	if err != nil {
		return nil, ge.Pin(err)
	}

	// Parse all directories (without the root directory)
	for _, path := range subDirectories {
		fullPath = path + string(os.PathSeparator) + tp.GetMask()
		t, err = t.ParseGlob(fullPath)
		if err != nil {
			return nil, ge.Pin(err, ge.Params{"path": fullPath})
		}
	}

	return t, nil
}

func (tp TextPath) Text(name string, data interface{}) (string, error) {
	t, err := tp.GetTemplate()
	if err != nil {
		return "", ge.Pin(err)
	}

	buf := new(bytes.Buffer)

	err = t.ExecuteTemplate(buf, name, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
