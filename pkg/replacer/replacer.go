package replacer

import (
	"strings"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

const ErrorsIdent = "SUIKAT.Replacer"

func ReplaceAll(data *Data, rules *Rules) (*Data, error) {
	for ruleIndex := range *rules {
		(*rules)[ruleIndex].Compile()
	}

	for dataIndex := range *data {
		after, err := ApplyRules((*data)[dataIndex].Before, rules)
		if err != nil {
			return nil, &ge.IdentityError{
				Package: ErrorsIdent,
				Caller:  "ReplaceAll()",
				Place:   "c.ApplyRules()",
				Params: map[string]interface{}{
					"dataIndex": dataIndex,
				},
				Parent: err,
			}
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
			return "", &ge.IdentityError{
				Package: ErrorsIdent,
				Caller:  "ApplyRules()",
				Place:   "c.ApplyRule()",
				Params: map[string]interface{}{
					"dataIndex": ruleIndex,
				},
				Parent: err,
			}
		}
	}

	return after, nil
}

func ApplyRule(before string, rule ReplaceRule) (after string, err error) {
	switch rule.SourceType {
	case SourceTypeRegexp:
		if rule.regExp == nil {
			return "", &ge.IdentityError{
				Package: ErrorsIdent,
				Caller:  "ApplyRule()",
				Place:   "rule.regExp == nil",
			}
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
		return "", &ge.IdentityError{
			Package: ErrorsIdent,
			Caller:  "ApplyRule()",
			Place:   "Undefined rule.SourceType",
			Params: map[string]interface{}{
				"rule.SourceType": rule.SourceType,
			},
		}
	}
}
