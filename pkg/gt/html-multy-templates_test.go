package gt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHTMLPathsGetTemplate(t *testing.T) {
	p := NewHTMLPaths()
	p.AddPath("test-data/html/multy-roots/pages/index")
	p.AddPath("test-data/html/multy-roots/components")

	template, err := p.GetTemplate()
	require.NoError(t, err)
	require.NotEmpty(t, template)
}

func TestHTMLPathsHTMLIndex(t *testing.T) {

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

	p := NewHTMLPaths()
	p.AddPath("test-data/html/multy-roots/pages/index")
	p.AddPath("test-data/html/multy-roots/components")

	html, err := p.HTML("page", htmlPageData, nil)
	require.NoError(t, err)
	require.NotEmpty(t, html)
	require.Equal(
		t,
		"\n"+
			"    <!DOCTYPE html>\n"+
			"    <html>\n"+
			"        \n"+
			"    <head>\n"+
			"        <meta charset=\"UTF-8\">\n"+
			"        <title>Hello</title>\n"+
			"    </head>\n\n"+
			"        \n"+
			"    <body>\n"+
			"        \n"+
			"    <h1>\n"+
			"        H1 Hello World\n"+
			"    </h1>\n\n"+
			"        \n"+
			"    <p>\n"+
			"        P Hello World\n"+
			"    </p>\n\n"+
			"    </body>\n\n"+
			"    </html>\n",
		html,
	)
}

func TestHTMLPathsHTMLContact(t *testing.T) {

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

	p := NewHTMLPaths()
	p.AddPath("test-data/html/multy-roots/pages/contact")
	p.AddPath("test-data/html/multy-roots/components")

	html, err := p.HTML("page", htmlPageData, nil)
	require.NoError(t, err)
	require.NotEmpty(t, html)
	require.Equal(
		t,
		"\n"+
			"    <!DOCTYPE html>\n"+
			"    <html>\n"+
			"        \n"+
			"    <head>\n"+
			"        <meta charset=\"UTF-8\">\n"+
			"        <title>Hello</title>\n"+
			"    </head>\n\n"+
			"        \n"+
			"    <body>\n"+
			"        \n"+
			"    <div class=\"card\">\n"+
			"        <h2>Hello</h2>\n"+
			"        <p>Card</p>\n"+
			"        \n"+
			"    <button class=\"blue\" type=\"submit\">\n"+
			"        SUBMIT\n"+
			"    </button>\n\n"+
			"        \n"+
			"    <button class=\"red\" type=\"reset\">\n"+
			"        RESET\n"+
			"    </button>\n\n"+
			"    </div>\n\n"+
			"    </body>\n\n"+
			"    </html>\n",
		html,
	)
}
