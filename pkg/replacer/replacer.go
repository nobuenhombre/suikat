package replacer

import (
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

func ReplaceAll(data *Data, rules *Rules) (*Data, error) {
	for ruleIndex := range *rules {
		(*rules)[ruleIndex].Compile()
	}

	for dataIndex := range *data {
		after, err := ApplyRules((*data)[dataIndex].Before, rules)
		if err != nil {
			return nil, ge.Pin(err, ge.Params{"dataIndex": dataIndex})
		}

		(*data)[dataIndex].After = after
	}

	return data, nil
}

func ApplyRules(before string, rules *Rules) (after string, err error) {
	after = before
	for ruleIndex, rule := range *rules {
		after, err = ApplyRule(after, rule)
		if err != nil {
			return "", ge.Pin(err, ge.Params{"dataIndex": ruleIndex})
		}
	}

	return after, nil
}

func ApplyRule(before string, rule ReplaceRule) (after string, err error) {
	switch rule.SourceType {
	case SourceTypeRegexp:
		if rule.regExp == nil {
			return "", ge.Pin(&ge.RegExpIsNotCompiledError{})
		}

		after = before

		foundMatches := rule.regExp.FindAllString(before, -1)
		for _, match := range foundMatches {
			after = strings.ReplaceAll(after, match, rule.Replacement)
		}

		return after, nil

	case SourceTypeString:
		after = strings.ReplaceAll(before, rule.Source, rule.Replacement)

		return after, nil

	default:
		return "", ge.Pin(&ge.UndefinedSwitchCaseError{Var: rule.SourceType})
	}
}
