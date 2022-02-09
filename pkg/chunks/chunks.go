// Package chunks provides
// functions for dividing a one-dimensional slice into several pieces
package chunks

// SplitInt64 dividing a one-dimensional slice []int64 into several pieces
func SplitInt64(in []int64, limit int) [][]int64 {
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

// SplitStr dividing a string into several pieces
func SplitStr(s string, limit int) []string {
	if len(s) == 0 {
		return nil
	}

	if limit >= len(s) {
		return []string{s}
	}

	chunks := make([]string, 0, (len(s)-1)/limit+1)
	currentLen := 0
	currentStart := 0

	for i := range s {
		if currentLen == limit {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}

		currentLen++
	}

	chunks = append(chunks, s[currentStart:])

	return chunks
}
