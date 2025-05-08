package gt

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

func TestTextPathGetTemplate(t *testing.T) {
	getFuncMap := func() template.FuncMap {
		return template.FuncMap{
			"formatPercent": func(value float64) string {
				p := message.NewPrinter(language.Russian)
				return p.Sprintf("%d%%", int64(math.Round(value*100)))
			},
		}
	}

	p := TextPath("test-data/text/single-root")
	template, err := p.GetTemplate(getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, template)
}

func TestTextPathText(t *testing.T) {
	type TextHead struct {
		Title string
	}

	type TextBody struct {
		Content string
	}

	type TextDoc struct {
		Head TextHead
		Body TextBody
	}

	textDocData := TextDoc{
		Head: TextHead{
			Title: "Hello",
		},
		Body: TextBody{
			Content: "World",
		},
	}

	getFuncMap := func() template.FuncMap {
		return template.FuncMap{
			"formatPercent": func(value float64) string {
				p := message.NewPrinter(language.Russian)
				return p.Sprintf("%d%%", int64(math.Round(value*100)))
			},
		}
	}

	p := TextPath("test-data/text/single-root")
	text, err := p.Text("index", textDocData, getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, text)
	require.Equal(
		t,
		"\n"+
			"    \n"+
			"    Hello\n\n"+
			"    \n"+
			"    World\n"+
			"    25%\n\n",
		text,
	)
}
