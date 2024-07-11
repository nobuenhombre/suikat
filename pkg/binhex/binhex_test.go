package binhex

import (
	"encoding/hex"
	"errors"
	"testing"
)

type bin2hexTest struct {
	input    BinString
	expected HexString
}

func TestBin2Hex(t *testing.T) {
	tests := map[string]bin2hexTest{
		"Hello World Test": {
			input:    "Hello World",
			expected: "48656c6c6f20576f726c64",
		},
	}

	for index, test := range tests {
		t.Run(index, func(t *testing.T) {
			result := test.input.ToHex()

			if result != test.expected {
				t.Errorf("exp %v but got %v", test.expected, result)
			}
		})
	}
}

type hex2binTest struct {
	input    HexString
	expected BinString
	err      error
}

func TestHex2Bin(t *testing.T) {
	tests := map[string]hex2binTest{
		"Hello World Test": {
			input:    "48656c6c6f20576f726c64",
			expected: "Hello World",
			err:      nil,
		},
		"Hello World Test with Error": {
			input:    "48656c6c6f20576f726c6",
			expected: "",
			err:      hex.ErrLength,
		},
	}

	for index, test := range tests {
		t.Run(index, func(t *testing.T) {
			result, err := test.input.ToBin()

			if result != test.expected || !errors.Is(err, test.err) {
				t.Errorf("exp %v but got %v, error exp %v but got %v", test.expected, result, test.err, err)
			}
		})
	}
}
