package gt

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"testing"
	"text/template"
)

func TestTextPathsGetTemplate(t *testing.T) {
	getFuncMap := func() template.FuncMap {
		return template.FuncMap{
			"formatPercent": func(value float64) string {
				p := message.NewPrinter(language.Russian)
				return p.Sprintf("%d%%", int64(math.Round(value*100)))
			},
		}
	}

	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/index")
	p.AddPath("test-data/text/multy-roots/components")

	template, err := p.GetTemplate(getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, template)
}

func TestTextPathsHTMLIndex(t *testing.T) {

	type HtmlHeader struct {
		H1 string
	}

	type HtmlParagraph struct {
		P string
	}

	type HtmlPageBody struct {
		Header    HtmlHeader
		Paragraph HtmlParagraph
	}

	type HtmlPageHead struct {
		Title string
	}

	type HtmlPage struct {
		Head HtmlPageHead
		Body HtmlPageBody
	}

	htmlPageData := HtmlPage{
		Head: HtmlPageHead{
			Title: "Hello",
		},
		Body: HtmlPageBody{
			Header: HtmlHeader{
				H1: "H1 Hello World",
			},
			Paragraph: HtmlParagraph{
				P: "P Hello World",
			},
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

	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/index")
	p.AddPath("test-data/text/multy-roots/components")

	html, err := p.Text("page", htmlPageData, getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, html)
	require.Equal(
		t,
		"\n"+
			"    \n"+
			"    Hello\n\n"+
			"    \n"+
			"    \n"+
			"    H1 Hello World\n\n"+
			"    \n"+
			"    P Hello World 25%\n\n\n",
		html,
	)
}

func TestTextPathsHTMLContact(t *testing.T) {

	type HtmlCard struct {
	}

	type HtmlPageBody struct {
		Card HtmlCard
	}

	type HtmlPageHead struct {
		Title string
	}

	type HtmlPage struct {
		Head HtmlPageHead
		Body HtmlPageBody
	}

	htmlPageData := HtmlPage{
		Head: HtmlPageHead{
			Title: "Hello",
		},
		Body: HtmlPageBody{
			Card: HtmlCard{},
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

	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/contact")
	p.AddPath("test-data/text/multy-roots/components")

	html, err := p.Text("page", htmlPageData, getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, html)
	require.Equal(
		t,
		"\n"+
			"    \n"+
			"    Hello\n\n"+
			"    \n"+
			"    \n"+
			"    ## Hello\n"+
			"    Card 25%\n"+
			"    \n"+
			"    [SUBMIT]\n\n"+
			"    \n"+
			"    [RESET]\n\n\n\n",
		html,
	)
}
