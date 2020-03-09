package webula

import (
	"reflect"
	"testing"
)

// StrLen Test
//===================================================
type webulaStrLenTest struct {
	in  string
	out int
}

var strLenTests = []webulaStrLenTest{
	{
		in:  "Hello",
		out: 5,
	},
	{
		in:  "Привет",
		out: 6,
	},
	{
		in:  "Hello World",
		out: 11,
	},
	{
		in:  "Привет Мир",
		out: 10,
	},
	{
		in:  "Hello Мир",
		out: 9,
	},
}

func TestStrLen(t *testing.T) {
	for i := 0; i < len(strLenTests); i++ {
		test := &strLenTests[i]
		out := StrLen(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"StrLen(%v), Expected %v, Actual %v",
				test.in, test.out, out,
			)
		}
	}
}

// StripHTML Test
//===================================================
type webulaStripHTMLTest struct {
	in  string
	out string
}

var stripHTMLTests = []webulaStripHTMLTest{
	{
		in:  "<h1>Заголовок</h1><p>Параграф <i>курсив</i> текст</p>",
		out: "ЗаголовокПараграф курсив текст",
	},
	{
		in:  "<h1>Заголовок</h1><p>Параграф <i>курсив" + HTMLSpace + "hello</i> текст</p>",
		out: "ЗаголовокПараграф курсив" + HTMLSpaceInUtf8 + "hello текст",
	},
}

func TestStripHTML(t *testing.T) {
	for i := 0; i < len(stripHTMLTests); i++ {
		test := &stripHTMLTests[i]
		out := StripHTML(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"StripHTML(%v), Expected %v, Actual %v",
				test.in, test.out, out,
			)
		}
	}
}

// Trim Test
//===================================================
type inputTrimParams struct {
	s        []string
	trimmers []string
}

type webulaTrimTest struct {
	in  inputTrimParams
	out []string
}

var trimTests = []webulaTrimTest{
	{
		in: inputTrimParams{
			s:        []string{"  Арбуз", "Дыня   ", HTMLSpaceInUtf8 + " Персик " + HTMLSpaceInUtf8},
			trimmers: []string{" ", HTMLSpaceInUtf8},
		},
		out: []string{"Арбуз", "Дыня", "Персик"},
	},
	{
		in: inputTrimParams{
			s:        []string{"  _Арбуз=", "==Дыня__   ", HTMLSpaceInUtf8 + " - Персик - " + HTMLSpaceInUtf8},
			trimmers: []string{" ", HTMLSpaceInUtf8, "_", "=", "-"},
		},
		out: []string{"Арбуз", "Дыня", "Персик"},
	},
}

func TestTrim(t *testing.T) {
	for i := 0; i < len(trimTests); i++ {
		test := &trimTests[i]
		out := Trim(test.in.s, test.in.trimmers)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Trim(%v, %v), Expected %v, Actual %v",
				test.in.s, test.in.trimmers, test.out, out,
			)
		}
	}
}

// Words Test
//===================================================
type webulaWordsTest struct {
	in  string
	out []string
}

var wordsTests = []webulaWordsTest{
	{
		in: "  Арбуз Дыня  " +
			CarriageReturn + " " + HTMLSpaceInUtf8 + " Персик" + NewLine + "Яблоко " + HTMLSpaceInUtf8 + " " +
			Tab + " Вишня ",
		out: []string{"Арбуз", "Дыня", "Персик", "Яблоко", "Вишня"},
	},
}

func TestWords(t *testing.T) {
	for i := 0; i < len(wordsTests); i++ {
		test := &wordsTests[i]
		out := Words(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"Words(%v), Expected %v, Actual %v",
				test.in, test.out, out,
			)
		}
	}
}

// NormalizeText Test
//===================================================
type inputNormalizeTextParams struct {
	text string
	glue string
}

type webulaNormalizeTextTest struct {
	in  inputNormalizeTextParams
	out string
}

var normalizeTextTests = []webulaNormalizeTextTest{
	{
		in: inputNormalizeTextParams{
			text: "  Арбуз Дыня  " +
				CarriageReturn + " " + HTMLSpaceInUtf8 + " Персик" + NewLine + "Яблоко " + HTMLSpaceInUtf8 + " " +
				Tab + " Вишня ",
			glue: Space,
		},
		out: "Арбуз Дыня Персик Яблоко Вишня",
	},
	{
		in: inputNormalizeTextParams{
			text: "  Арбуз Дыня  " +
				CarriageReturn + " " + HTMLSpaceInUtf8 + " Персик" + NewLine + "Яблоко " + HTMLSpaceInUtf8 + " " +
				Tab + " Вишня ",
			glue: Comma + Space,
		},
		out: "Арбуз, Дыня, Персик, Яблоко, Вишня",
	},
}

func TestNormalizeText(t *testing.T) {
	for i := 0; i < len(normalizeTextTests); i++ {
		test := &normalizeTextTests[i]
		out := NormalizeText(test.in.text, test.in.glue)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"NormalizeText(%v, %v), Expected %v, Actual %v",
				test.in.text, test.in.glue, test.out, out,
			)
		}
	}
}

// IsHTML Test
//===================================================
type webulaIsHTMLTest struct {
	in  string
	out bool
}

var isHTMLTests = []webulaIsHTMLTest{
	{
		in: "  Арбуз Дыня  " +
			CarriageReturn + " " + HTMLSpaceInUtf8 + " Персик" + NewLine + "Яблоко " + HTMLSpaceInUtf8 + " " +
			Tab + " Вишня ",
		out: false,
	},
	{
		in: "  Арбуз Дыня  " +
			CarriageReturn + " " + HTMLSpace + " Персик" + NewLine + "Яблоко " + HTMLSpace + " " +
			Tab + " Вишня ",
		out: true,
	},
	{
		in:  "<h1>Заголовок</h1><p>Параграф <i>курсив</i> текст</p>",
		out: true,
	},
}

func TestIsHTML(t *testing.T) {
	for i := 0; i < len(isHTMLTests); i++ {
		test := &isHTMLTests[i]
		out := IsHTML(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"IsHTM(%v), Expected %v, Actual %v",
				test.in, test.out, out,
			)
		}
	}
}
