package sfu

import (
	"net/url"
	"reflect"
	"testing"
)

// Prepared Data
//===============================

// FormA
//-------------------------------
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
//-------------------------------
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
//-------------------------------
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
			ID:   27,
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
}

func TestConvert(t *testing.T) {
	for i := 0; i < len(convertTests); i++ {
		test := &convertTests[i]

		form := &url.Values{}

		err := Convert(test.structData, test.parent, form)

		if !(reflect.DeepEqual(form, test.form) && reflect.DeepEqual(err, test.err)) {
			t.Errorf(
				"Convert(\n\t%#v,\n\t%#v),\n Expected (\n\t%#v,\n\t%#v),\n Actual (\n\t%#v,\n\t%#v).\n",
				test.structData, test.parent, test.form, test.err, form, err,
			)
		}
	}
}
