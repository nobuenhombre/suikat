package gt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPathGetDirectories(t *testing.T) {
	p := Path("test-data/html")
	list, err := p.GetSubDirectories()
	require.NoError(t, err)
	require.Len(t, list, 2)
}
