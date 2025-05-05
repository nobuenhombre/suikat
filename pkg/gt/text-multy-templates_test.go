package gt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTextPathsGetTemplate(t *testing.T) {
	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/index")
	p.AddPath("test-data/text/multy-roots/components")

	template, err := p.GetTemplate()
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

	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/index")
	p.AddPath("test-data/text/multy-roots/components")

	html, err := p.HTML("page", htmlPageData)
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
			"    P Hello World\n\n\n",
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

	p := NewTextPaths()
	p.AddPath("test-data/text/multy-roots/docs/contact")
	p.AddPath("test-data/text/multy-roots/components")

	html, err := p.HTML("page", htmlPageData)
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
			"    Card\n"+
			"    \n"+
			"    [SUBMIT]\n\n"+
			"    \n"+
			"    [RESET]\n\n\n\n",
		html,
	)
}
