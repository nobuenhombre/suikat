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

type reverseTest struct {
	in  []byte
	out []byte
}

var reverseTests = []reverseTest{
	{
		in:  []byte{1, 5, 4, 99, 11, 2, 7, 85, 33, 0},
		out: []byte{0, 33, 85, 7, 2, 11, 99, 4, 5, 1},
	},
}

func TestReverseFullBytes(t *testing.T) {
	for i := 0; i < len(reverseTests); i++ {
		test := &reverseTests[i]
		out := ReverseFullBytes(test.in)

		if !reflect.DeepEqual(out, test.out) {
			t.Errorf(
				"ReverseFullBytes(%v), Expected %v, Actual %v",
				test.in, test.out, out,
			)
		}
	}
}

type swapBytesTest struct {
	in  []byte
	out []byte
	err error
}

var swapBytesTests = []swapBytesTest{
	{
		in:  []byte{1, 5, 4, 99, 11, 2, 7, 85, 33, 0},
		out: []byte{2, 7, 85, 33, 0, 1, 5, 4, 99, 11},
		err: nil,
	},
}

func TestSwapHalfBytes(t *testing.T) {
	for i := 0; i < len(swapBytesTests); i++ {
		test := &swapBytesTests[i]
		out, err := SwapHalfBytes(test.in)

		if !(reflect.DeepEqual(out, test.out) && reflect.DeepEqual(err, test.err)) {
			t.Errorf(
				"SwapHalfBytes(%v), Expected (%v, %v) Actual (%v, %v)",
				test.in, test.out, test.err, out, err,
			)
		}
	}
}
