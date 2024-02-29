package main

import (
	"taniyuu/csv-samples/modules"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	encoder := japanese.ShiftJIS.NewEncoder()
	data := [][]string{
		{"Name", "City", "Country", "年齢"},
		{"John", "Boston", "USA", "20"},
		{"太郎", "滋賀県\"大津市\"浜大津\n一丁目", "日本", ""},
	}
	if err := modules.EncodeStandard("out/standard.csv", encoder, data); err != nil {
		panic(err)
	}

	type Address struct {
		City    string
		Country string
	}

	type User struct {
		Name string
		Address
		Age int `csv:"年齢,omitempty"`
	}
	dataStruct := []*User{
		{"John", Address{"Boston", "USA"}, 20},
		{"太郎", Address{"滋賀県\"大津市\"浜大津\n一丁目", "日本"}, 0},
	}
	if err := modules.EncodeCSVUtil("out/csvutil.csv", encoder, dataStruct); err != nil {
		panic(err)
	}

}
