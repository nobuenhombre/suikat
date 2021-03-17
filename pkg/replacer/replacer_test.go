package replacer

import (
	"reflect"
	"testing"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

type inputApplyRuleParams struct {
	before string
	rule   ReplaceRule
}

type outputApplyRuleParams struct {
	after string
	err   error
}

type applyRuleTest struct {
	compileRule bool
	in          inputApplyRuleParams
	out         outputApplyRuleParams
}

var applyRuleTests = []applyRuleTest{
	{
		compileRule: true,
		in: inputApplyRuleParams{
			before: "Hello %%name%% World",
			rule: ReplaceRule{
				SourceType:  SourceTypeString,
				Source:      "%%name%%",
				Replacement: "Mr.",
				regExp:      nil,
			},
		},
		out: outputApplyRuleParams{
			after: "Hello Mr. World",
			err:   nil,
		},
	},
	{
		compileRule: false,
		in: inputApplyRuleParams{
			before: "Hello data-uri(/img/welcome.jpg) World",
			rule: ReplaceRule{
				SourceType:  SourceTypeRegexp,
				Source:      "data-uri\\(.*?\\)",
				Replacement: "Mr.",
				regExp:      nil,
			},
		},
		out: outputApplyRuleParams{
			after: "",
			err: &ge.IdentityError{
				Package: ErrorsIdent,
				Caller:  "ApplyRule()",
				Place:   "rule.regExp == nil",
			},
		},
	},
	{
		compileRule: true,
		in: inputApplyRuleParams{
			before: "Hello data-uri(/img/welcome.jpg) World",
			rule: ReplaceRule{
				SourceType:  SourceTypeRegexp,
				Source:      "data-uri\\(.*?\\)",
				Replacement: "Mr.",
				regExp:      nil,
			},
		},
		out: outputApplyRuleParams{
			after: "Hello Mr. World",
			err:   nil,
		},
	},
	{
		compileRule: true,
		in: inputApplyRuleParams{
			before: "RegExr was created by gskinner.com, and is proudly hosted by Media Temple.",
			rule: ReplaceRule{
				SourceType:  SourceTypeRegexp,
				Source:      "([A-Z])\\w+",
				Replacement: "BZZ",
				regExp:      nil,
			},
		},
		out: outputApplyRuleParams{
			after: "BZZ was created by gskinner.com, and is proudly hosted by BZZ BZZ.",
			err:   nil,
		},
	},
}

func TestApplyRule(t *testing.T) {
	for i := 0; i < len(applyRuleTests); i++ {
		test := &applyRuleTests[i]
		if test.compileRule {
			test.in.rule.Compile()
		}

		after, err := ApplyRule(test.in.before, test.in.rule)

		if !(reflect.DeepEqual(after, test.out.after) && reflect.DeepEqual(err, test.out.err)) {
			t.Errorf(
				"ApplyRule(%v, %v),\n"+
					"Expected (%v, %v),\n"+
					"Actual (%v, %v)\n",
				test.in.before, test.in.rule,
				test.out.after, test.out.err,
				after, err,
			)
		}
	}
}

//===============================================================

type inputApplyRulesParams struct {
	before string
	rules  *Rules
}

type outputApplyRulesParams struct {
	after string
	err   error
}

type applyRulesTest struct {
	compileRules bool
	in           inputApplyRulesParams
	out          outputApplyRulesParams
}

var applyRulesTests = []applyRulesTest{
	{
		compileRules: true,
		in: inputApplyRulesParams{
			before: "Hello %%name%% World, data-uri(/img/welcome.jpg), hoho ZuZu",
			rules: &Rules{
				{
					SourceType:  SourceTypeString,
					Source:      "%%name%%",
					Replacement: "Mr.",
					regExp:      nil,
				},
				{
					SourceType:  SourceTypeRegexp,
					Source:      "data-uri\\(.*?\\)",
					Replacement: "uri",
					regExp:      nil,
				},
				{
					SourceType:  SourceTypeRegexp,
					Source:      "([A-Z])\\w+",
					Replacement: "BZZ",
					regExp:      nil,
				},
			},
		},
		out: outputApplyRulesParams{
			after: "BZZ BZZ. BZZ, uri, hoho BZZ",
			err:   nil,
		},
	},
}

func TestApplyRules(t *testing.T) {
	for i := 0; i < len(applyRulesTests); i++ {
		test := &applyRulesTests[i]
		if test.compileRules {
			for j := range *test.in.rules {
				(*test.in.rules)[j].Compile()
			}
		}

		after, err := ApplyRules(test.in.before, test.in.rules)

		if !(reflect.DeepEqual(after, test.out.after) && reflect.DeepEqual(err, test.out.err)) {
			t.Errorf(
				"ApplyRules(%v, %v),\n"+
					"Expected (%v, %v),\n"+
					"Actual (%v, %v)\n",
				test.in.before, test.in.rules,
				test.out.after, test.out.err,
				after, err,
			)
		}
	}
}

//===============================================================

type inputReplaceAllParams struct {
	data  *Data
	rules *Rules
}

type outputReplaceAllParams struct {
	data *Data
	err  error
}

type replaceAllTest struct {
	in  inputReplaceAllParams
	out outputReplaceAllParams
}

var replaceAllTests = []replaceAllTest{
	{
		in: inputReplaceAllParams{
			data: &Data{
				{
					ID:     1,
					Before: "hello %%name%% world",
					After:  "",
				},
				{
					ID:     2,
					Before: "Call data-uri(/img/welcome.jpg) This Site!",
					After:  "",
				},
				{
					ID:     3,
					Before: "Board hmm Band %%name%%",
					After:  "",
				},
			},
			rules: &Rules{
				{
					SourceType:  SourceTypeString,
					Source:      "%%name%%",
					Replacement: "Mr.",
					regExp:      nil,
				},
				{
					SourceType:  SourceTypeRegexp,
					Source:      "data-uri\\(.*?\\)",
					Replacement: "uri",
					regExp:      nil,
				},
				{
					SourceType:  SourceTypeRegexp,
					Source:      "([A-Z])\\w+",
					Replacement: "BZZ",
					regExp:      nil,
				},
			},
		},
		out: outputReplaceAllParams{
			data: &Data{
				{
					ID:     1,
					Before: "hello %%name%% world",
					After:  "hello BZZ. world",
				},
				{
					ID:     2,
					Before: "Call data-uri(/img/welcome.jpg) This Site!",
					After:  "BZZ uri BZZ BZZ!",
				},
				{
					ID:     3,
					Before: "Board hmm Band %%name%%",
					After:  "BZZ hmm BZZ BZZ.",
				},
			},
			err: nil,
		},
	},
}

func TestReplaceAll(t *testing.T) {
	for i := 0; i < len(replaceAllTests); i++ {
		test := &replaceAllTests[i]

		outData, err := ReplaceAll(test.in.data, test.in.rules)

		if !(reflect.DeepEqual(outData, test.out.data) && reflect.DeepEqual(err, test.out.err)) {
			t.Errorf(
				"ApplyRules(%v, %v),\n"+
					"Expected (%v, %v),\n"+
					"Actual (%v, %v)\n",
				test.in.data, test.in.rules,
				test.out.data, test.out.err,
				outData, err,
			)
		}
	}
}
