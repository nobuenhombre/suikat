package gt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTMLPathGetTemplate(t *testing.T) {
	p := HTMLPath("test-data/html/single-root")
	template, err := p.GetTemplate()
	require.NoError(t, err)
	require.NotEmpty(t, template)
}

func TestHTMLPathHTML(t *testing.T) {
	type HtmlPageHead struct {
		Title string
	}

	type HtmlPageBody struct {
		Content string
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
			Content: "World",
		},
	}

	p := HTMLPath("test-data/html/single-root")
	html, err := p.HTML("index", htmlPageData, nil)
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
			"        <main>\n"+
			"            World\n"+
			"        </main>\n"+
			"    </body>\n\n"+
			"    </html>\n",
		html,
	)
}
