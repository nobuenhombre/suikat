package chunks

import (
	"reflect"
	"testing"
)

type splitTest struct {
	in    []int64
	limit int
	out   [][]int64
}

var splitTests = []splitTest{
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 1,
		out: [][]int64{
			{1},
			{3},
			{5},
			{7},
			{9},
			{0},
			{2},
			{4},
			{6},
			{8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 2,
		out: [][]int64{
			{1, 3},
			{5, 7},
			{9, 0},
			{2, 4},
			{6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 3,
		out: [][]int64{
			{1, 3, 5},
			{7, 9, 0},
			{2, 4, 6},
			{8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 4,
		out: [][]int64{
			{1, 3, 5, 7},
			{9, 0, 2, 4},
			{6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 5,
		out: [][]int64{
			{1, 3, 5, 7, 9},
			{0, 2, 4, 6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 6,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0},
			{2, 4, 6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 7,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0, 2},
			{4, 6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 8,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0, 2, 4},
			{6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 9,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0, 2, 4, 6},
			{8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 10,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		},
	},
	{
		in:    []int64{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		limit: 11,
		out: [][]int64{
			{1, 3, 5, 7, 9, 0, 2, 4, 6, 8},
		},
	},
}

func TestSplit(t *testing.T) {
	for i := 0; i < len(splitTests); i++ {
		test := &splitTests[i]
		out := SplitInt64(test.in, test.limit)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"SplitInt64(%v, %v), Expected %v, Actual %v",
				test.in, test.limit, test.out, out,
			)
		}
	}
}
