package replacer

import "regexp"

const (
	SourceTypeRegexp = "regexp"
	SourceTypeString = "string"
)

type ReplaceRule struct {
	// nolint: tagliatelle
	SourceType string `json:"SourceType"`
	// nolint: tagliatelle
	Source string `json:"Source"`
	// nolint: tagliatelle
	Replacement string `json:"Replacement"`

	regExp *regexp.Regexp
}

func (rr *ReplaceRule) Compile() {
	if rr.SourceType == SourceTypeRegexp {
		rr.regExp = regexp.MustCompile(rr.Source)
	}
}

type ReplaceData struct {
	ID     int64
	Before string
	After  string
}

type Rules []ReplaceRule

type Data []ReplaceData
