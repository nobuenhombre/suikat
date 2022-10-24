package sfu

import (
	"errors"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/nobuenhombre/suikat/pkg/ge"
)

// Prepared Data
//===============================

// FormA
// -------------------------------
type FormA struct {
	ID   int64  `form:"id"`
	Name string `form:"name"`
}

func GetFormA() *url.Values {
	form := &url.Values{}
	form.Add("id", "12")
	form.Add("name", "Hello World 12")

	return form
}

// FormB
// -------------------------------
type Address struct {
	PostIndex   int64  `form:"postIndex"`
	Country     string `form:"country"`
	City        string `form:"city"`
	Street      string `form:"street"`
	HouseNumber string `form:"houseNumber"`
}

type FormB struct {
	ID      int64   `form:"id"`
	Name    string  `form:"name"`
	Address Address `form:"address"`
}

func GetFormB() *url.Values {
	form := &url.Values{}
	form.Add("id", "15")
	form.Add("name", "Riki tiki tavi")
	form.Add("address[postIndex]", "123654")
	form.Add("address[country]", "Russia")
	form.Add("address[city]", "Samara")
	form.Add("address[street]", "Leningradskaya str.")
	form.Add("address[houseNumber]", "321B")

	return form
}

// FormC
// -------------------------------
type Prices struct {
	Chicken  float64 `form:"chicken"`
	FishRice float64 `form:"fishRice"`
	Juice    float64 `form:"juice"`
}

type PriceList struct {
	MinPrices Prices `form:"min"`
	MaxPrices Prices `form:"max"`
}

type FormC struct {
	ID     int64     `form:"id"`
	Name   string    `form:"name"`
	Prices PriceList `form:"prices"`
}

func GetFormC() *url.Values {
	form := &url.Values{}
	form.Add("id", "27")
	form.Add("name", "AirFly")
	form.Add("prices[min][chicken]", "12.3")
	form.Add("prices[min][fishRice]", "23.4")
	form.Add("prices[min][juice]", "4.5")
	form.Add("prices[max][chicken]", "34.5")
	form.Add("prices[max][fishRice]", "45.6")
	form.Add("prices[max][juice]", "6.7")

	return form
}

// FormD
// -------------------------------
type Tunnel struct {
	IPFrom  string `form:"ipFrom"`
	IPTo    string `form:"ipTo"`
	Encoded bool   `form:"encoded"`
}

type FormD struct {
	ID       int64     `form:"id"`
	Name     string    `form:"name"`
	Robots   []string  `form:"robots"`
	Timeouts []int64   `form:"timeouts"`
	Sizes    []float64 `form:"sizes"`
	Allows   []bool    `form:"allows"`
	Tunnels  []Tunnel  `form:"tunnels"`
}

func GetFormD() *url.Values {
	form := &url.Values{}
	form.Add("id", "404")
	form.Add("name", "Robots")
	form.Add("robots[0]", "Mail.ru")
	form.Add("robots[1]", "Yandex-Bot")
	form.Add("robots[2]", "Google-Bot")
	form.Add("timeouts[0]", "234")
	form.Add("timeouts[1]", "567")
	form.Add("sizes[0]", "12.3")
	form.Add("sizes[1]", "45.6")
	form.Add("sizes[2]", "78.9")
	form.Add("allows[0]", "true")
	form.Add("allows[1]", "false")
	form.Add("allows[2]", "false")
	form.Add("allows[3]", "true")
	form.Add("tunnels[0][ipFrom]", "97.34.177.231")
	form.Add("tunnels[0][ipTo]", "97.34.199.231")
	form.Add("tunnels[0][encoded]", "true")
	form.Add("tunnels[1][ipFrom]", "227.34.177.231")
	form.Add("tunnels[1][ipTo]", "327.34.177.231")
	form.Add("tunnels[1][encoded]", "false")

	return form
}

type FormErrA struct {
	Unknown time.Time `form:"unknown"`
}

func GetFormErrA() *url.Values {
	return &url.Values{}
}

type FormErrB struct {
	Unknown byte `form:"unknown"`
}

func GetFormErrB() *url.Values {
	return &url.Values{}
}

type FormErrC struct {
	Unknown []byte `form:"unknown"`
}

func GetFormErrC() *url.Values {
	return &url.Values{}
}

type FormErrD struct {
	Unknown []FormErrA `form:"unknown"`
}

func GetFormErrD() *url.Values {
	return &url.Values{}
}

func GetFormErrF() *url.Values {
	return &url.Values{}
}

// Tests
//-------------------------------

type convertTest struct {
	structData interface{}
	parent     string
	form       *url.Values
	err        error
}

var convertTests = []convertTest{
	{
		structData: &FormA{
			ID:   12,
			Name: "Hello World 12",
		},
		parent: "",
		form:   GetFormA(),
		err:    nil,
	},
	{
		structData: &FormB{
			ID:   15,
			Name: "Riki tiki tavi",
			Address: Address{
				PostIndex:   123654,
				Country:     "Russia",
				City:        "Samara",
				Street:      "Leningradskaya str.",
				HouseNumber: "321B",
			},
		},
		parent: "",
		form:   GetFormB(),
		err:    nil,
	},
	{
		structData: &FormC{
			ID:   27,
			Name: "AirFly",
			Prices: PriceList{
				MinPrices: Prices{
					Chicken:  12.3,
					FishRice: 23.4,
					Juice:    4.5,
				},
				MaxPrices: Prices{
					Chicken:  34.5,
					FishRice: 45.6,
					Juice:    6.7,
				},
			},
		},
		parent: "",
		form:   GetFormC(),
		err:    nil,
	},
	{
		structData: &FormD{
			ID:   404,
			Name: "Robots",
			Robots: []string{
				"Mail.ru",
				"Yandex-Bot",
				"Google-Bot",
			},
			Timeouts: []int64{
				234,
				567,
			},
			Sizes: []float64{
				12.3,
				45.6,
				78.9,
			},
			Allows: []bool{
				true,
				false,
				false,
				true,
			},
			Tunnels: []Tunnel{
				{
					IPFrom:  "97.34.177.231",
					IPTo:    "97.34.199.231",
					Encoded: true,
				},
				{
					IPFrom:  "227.34.177.231",
					IPTo:    "327.34.177.231",
					Encoded: false,
				},
			},
		},
		parent: "",
		form:   GetFormD(),
		err:    nil,
	},
	{
		structData: &FormErrA{
			Unknown: time.Now(),
		},
		parent: "",
		form:   GetFormErrA(),
		err: &ge.PrivateStructFieldError{
			Name: "time.Time",
		},
	},
	{
		structData: &FormErrB{
			Unknown: 15,
		},
		parent: "",
		form:   GetFormErrB(),
		err:    &ge.UnknownTypeError{Type: "uint8"},
	},
	{
		structData: &FormErrC{
			Unknown: []byte{1, 2, 3},
		},
		parent: "",
		form:   GetFormErrC(),
		err:    &ge.UnknownTypeError{Type: "uint8"},
	},
	{
		structData: &FormErrD{
			Unknown: []FormErrA{
				{
					Unknown: time.Now(),
				},
			},
		},
		parent: "",
		form:   GetFormErrD(),
		err: &ge.PrivateStructFieldError{
			Name: "time.Time",
		},
	},
	{
		structData: 123,
		parent:     "",
		form:       GetFormErrF(),
		err:        &ge.MismatchError{Expected: "struct", Actual: "int"},
	},
	{
		structData: FormA{
			ID:   12,
			Name: "Hello World 12",
		},
		parent: "",
		form:   GetFormErrF(),
		err:    &ge.CantBeSetError{},
	},
}

func TestConvert(t *testing.T) {
	for i := 0; i < len(convertTests); i++ {
		test := &convertTests[i]

		form := &url.Values{}

		err := Convert(test.structData, test.parent, form)

		if !(reflect.DeepEqual(form, test.form) && errors.Is(err, test.err)) {
			t.Errorf(
				"Convert(\n\t%#v,\n\t%#v),\n Expected (\n\t%#v,\n\t%#v),\n Actual (\n\t%#v,\n\t%#v).\n",
				test.structData, test.parent, test.form, test.err, form, err,
			)
		}
	}
}
