// Package chunks provides
// functions for dividing a one-dimensional slice into several pieces
package chunks

import (
	"fmt"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

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

// SplitBytes dividing a one-dimensional slice []byte into several pieces
func SplitBytes(in []byte, limit int) [][]byte {
	var chunk []byte

	chunks := make([][]byte, 0, len(in)/limit+1)

	for len(in) >= limit {
		chunk, in = in[:limit], in[limit:]
		chunks = append(chunks, chunk)
	}

	if len(in) > 0 {
		chunks = append(chunks, in)
	}

	return chunks
}

// ReverseFullBytes reverse full slice []byte{1,2,3,4,5,6} -> []byte{6,5,4,3,2,1}
func ReverseFullBytes(in []byte) []byte {
	if len(in) == 0 {
		return in
	}

	return append(ReverseFullBytes(in[1:]), in[0])
}

const TwoParts = 2

// SwapHalfBytes reverse two halfs slice []byte{1,2,3,4,5,6} -> []byte{4,5,6,1,2,3}
func SwapHalfBytes(in []byte) ([]byte, error) {
	l := len(in)

	if (l % TwoParts) > 0 {
		return nil, ge.New(fmt.Sprintf("cant dive slice by two parts len(in)=%v", l))
	}

	halfs := SplitBytes(in, l/TwoParts)

	out := make([]byte, 0, l)
	out = append(out, halfs[1]...)
	out = append(out, halfs[0]...)

	return out, nil
}
