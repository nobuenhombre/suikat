package gt

import (
	"bytes"
	"errors"
	"github.com/nobuenhombre/suikat/pkg/ge"
	htmlTemplate "html/template"
	"os"
	"strings"
	"text/template"
)

type HTMLPaths struct {
	List []Path
}

var NoPathsDefinedError = errors.New("no paths defined")

func NewHTMLPaths() *HTMLPaths {
	return &HTMLPaths{
		List: []Path{},
	}
}

func (htmlPaths *HTMLPaths) AddPath(path Path) {
	htmlPaths.List = append(htmlPaths.List, path)
}

func (htmlPaths *HTMLPaths) GetMask() string {
	return MaskGoHTML
}

func (htmlPaths *HTMLPaths) GetTemplate(funcMap template.FuncMap) (*htmlTemplate.Template, error) {
	if len(htmlPaths.List) == 0 {
		return nil, ge.Pin(NoPathsDefinedError)
	}

	t := htmlTemplate.New("")

	if funcMap != nil {
		t = t.Funcs(funcMap)
	}

	// Create a template for parsing all directories
	fullPath := string(htmlPaths.List[0]) + string(os.PathSeparator) + htmlPaths.GetMask()
	t, err := t.ParseGlob(fullPath)
	if err != nil {
		return nil, ge.Pin(err, ge.Params{"path": fullPath})
	}

	subDirectories := make([]string, 0)

	for _, dir := range htmlPaths.List {
		dirSubDirectories, err := dir.GetSubDirectories()
		if err != nil {
			return nil, ge.Pin(err)
		}

		subDirectories = append(subDirectories, dirSubDirectories...)
	}

	// Parse all directories (without the root directory)
	for _, path := range subDirectories {
		fullPath = path + string(os.PathSeparator) + htmlPaths.GetMask()
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

func (htmlPaths *HTMLPaths) HTML(name string, data interface{}, funcMap template.FuncMap) (string, error) {
	t, err := htmlPaths.GetTemplate(funcMap)
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
