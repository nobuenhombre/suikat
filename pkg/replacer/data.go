package replacer

import "regexp"

const (
	SourceTypeRegexp = "regexp"
	SourceTypeString = "string"
)

type ReplaceRule struct {
	SourceType  string `json:"SourceType"`
	Source      string `json:"Source"`
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
