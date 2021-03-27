package main

import (
	"fmt"
	"net/http"
	"reflect"
)

func Hello(outData interface{}) error {
	respHeader := http.Header{}
	respHeader.Add("Hello", "MyDarling")

	rh := reflect.ValueOf(respHeader)
	rv := reflect.ValueOf(outData)
	if rv.Kind() == reflect.Map {
		for _, k := range rh.MapKeys() {
			rv.SetMapIndex(k, rh.MapIndex(k))
		}
	}

	return nil
}

func main() {
	gettedHeader := http.Header{}

	err := Hello(gettedHeader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(gettedHeader)
}
