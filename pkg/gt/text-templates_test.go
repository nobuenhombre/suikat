package gt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTextPathGetTemplate(t *testing.T) {
	p := TextPath("test-data/text/single-root")
	template, err := p.GetTemplate()
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

	p := TextPath("test-data/text/single-root")
	text, err := p.Text("index", textDocData, nil)
	require.NoError(t, err)
	require.NotEmpty(t, text)
	require.Equal(
		t,
		"\n"+
			"    \n"+
			"    Hello\n\n"+
			"    \n"+
			"    World\n\n",
		text,
	)
}
