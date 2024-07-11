package binhex

import (
	"encoding/hex"
	"fmt"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

type BinString string
type HexString string

func (binStr *BinString) ToHex() HexString {
	return HexString(hex.EncodeToString([]byte(*binStr)))
}

func (binStr *BinString) ToString() string {
	return string(*binStr)
}

func (hexStr *HexString) ToBin() (BinString, error) {
	decoded, err := hex.DecodeString(string(*hexStr))
	if err != nil {
		return BinString(""), ge.Pin(err)
	}

	return BinString(fmt.Sprintf("%s", decoded)), nil
}

func (hexStr *HexString) ToString() string {
	return string(*hexStr)
}
