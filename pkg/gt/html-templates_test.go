package gt

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTMLPathGetTemplate(t *testing.T) {
	getFuncMap := func() template.FuncMap {
		return template.FuncMap{
			"htmlSafe": func(html string) template.HTML {
				return template.HTML(html)
			},
		}
	}

	p := HTMLPath("test-data/html/single-root")
	template, err := p.GetTemplate(getFuncMap())
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

	getFuncMap := func() template.FuncMap {
		return template.FuncMap{
			"htmlSafe": func(html string) template.HTML {
				return template.HTML(html)
			},
		}
	}

	p := HTMLPath("test-data/html/single-root")
	html, err := p.HTML("index", htmlPageData, getFuncMap())
	require.NoError(t, err)
	require.NotEmpty(t, html)
	require.Equal(
		t,
		"\n"+
			"    <!DOCTYPE html>\n"+
			"    <html>\n"+
			"        <!-- test funcMap -->\n"+
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
