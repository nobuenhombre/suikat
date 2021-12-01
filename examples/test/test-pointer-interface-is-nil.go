package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/nobuenhombre/suikat/pkg/adapt"
)

func getReaderA() io.Reader {
	return nil
}

func getReaderB() io.Reader {
	var data *bytes.Buffer
	data = nil

	return data
}

func TestPointerInterfaceIsNil() {
	fmt.Printf("nil = %#v\n", nil)
	fmt.Printf("readerA = %#v\n", getReaderA())
	fmt.Printf("readerB = %#v\n", getReaderB())
	fmt.Printf("--------------------------------\n\n")

	fmt.Printf("isNil nil = %#v\n", adapt.IsNil(nil))
	fmt.Printf("isNil readerA = %#v\n", adapt.IsNil(getReaderA()))
	fmt.Printf("isNil readerB = %#v\n", adapt.IsNil(getReaderB()))
	fmt.Printf("--------------------------------\n\n")

	fmt.Printf("nil = nil ? %#v\n", true)
	fmt.Printf("readerA = nil ? %#v\n", getReaderA() == nil)
	fmt.Printf("readerB = nil ? %#v\n", getReaderB() == nil)
}
