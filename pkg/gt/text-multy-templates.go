package gt

import (
	"bytes"
	"html/template"
	"os"
	"strings"
	textTemplate "text/template"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type TextPaths struct {
	List []Path
}

func NewTextPaths() *TextPaths {
	return &TextPaths{
		List: []Path{},
	}
}

func (textPaths *TextPaths) AddPath(path Path) {
	textPaths.List = append(textPaths.List, path)
}

func (textPaths *TextPaths) GetMask() string {
	return MaskGoTPL
}

func (textPaths *TextPaths) GetTemplate(funcMap template.FuncMap) (*textTemplate.Template, error) {
	if len(textPaths.List) == 0 {
		return nil, ge.Pin(NoPathsDefinedError)
	}

	t := textTemplate.New("")

	if funcMap != nil {
		t = t.Funcs(funcMap)
	}

	// Create a template for parsing all directories
	fullPath := string(textPaths.List[0]) + string(os.PathSeparator) + textPaths.GetMask()
	t, err := t.ParseGlob(fullPath)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"path": fullPath})
	}

	subDirectories := make([]string, 0)

	for _, dir := range textPaths.List {
		dirSubDirectories, err := dir.GetSubDirectories()
		if err != nil {
			return nil, ge.Pin(err)
		}

		subDirectories = append(subDirectories, dirSubDirectories...)
	}

	// Parse all directories (without the root directory)
	for _, path := range subDirectories {
		fullPath = path + string(os.PathSeparator) + textPaths.GetMask()
		tPath, err := t.ParseGlob(fullPath)
		if err != nil {
			if strings.Contains(err.Error(), "pattern matches no files:") {
				continue
			}

			return nil, ge.Pin(err, ge.Params{"path": fullPath})
		}

		t = tPath
	}

	return t, nil
}

func (textPaths *TextPaths) Text(name string, data interface{}, funcMap template.FuncMap) (string, error) {
	t, err := textPaths.GetTemplate(funcMap)
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
