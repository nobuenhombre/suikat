// Package chunks provides
// functions for dividing a one-dimensional slice into several pieces
package chunks

// Split dividing a one-dimensional slice []int64 into several pieces
func Split(in []int64, limit int) [][]int64 {
	var chunk []int64

	chunks := make([][]int64, 0, len(in)/limit+1)

	for len(in) >= limit {
		chunk, in = in[:limit], in[limit:]
		chunks = append(chunks, chunk)
	}

	if len(in) > 0 {
		chunks = append(chunks, in)
	}

	return chunks
}
