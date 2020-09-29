package fina

import (
	"os"
	"path/filepath"
	"strings"
)

type FilePartsInfo struct {
	Path           string
	Name           string
	Ext            string
	NameWithoutExt string
}

// Return new filename with Prefix
// was:  /some/path/file.ext
// will: /some/path/<prefix>file.ext
// ---------------------------------
// prefix: "demo-"
// will: /some/path/demo-file.ext
func (fpi *FilePartsInfo) Prefix(prefix string) string {
	return fpi.Path + prefix + fpi.Name
}

// Return new filename with new extension
// was:  /some/path/file.ext
// will: /some/path/file.<newext>
func (fpi *FilePartsInfo) NewExt(ext string) string {
	return fpi.Path + fpi.NameWithoutExt + ext
}

// Return new filename with new extension and prefix
// was:  /some/path/file.ext
// will: /some/path/<prefix>file.<newext>
func (fpi *FilePartsInfo) PrefixWithNewExt(prefix, ext string) string {
	return fpi.Path + prefix + fpi.NameWithoutExt + ext
}

func GetFilePartsInfo(file string) *FilePartsInfo {
	return &FilePartsInfo{
		Path:           filepath.Dir(file) + string(os.PathSeparator),
		Name:           filepath.Base(file),
		Ext:            filepath.Ext(file),
		NameWithoutExt: strings.TrimSuffix(filepath.Base(file), filepath.Ext(filepath.Base(file))),
	}
}