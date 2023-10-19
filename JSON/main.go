package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string
	Price int
	Tags  string
}

func main() {
	// Create JSON
	Encodejson()
	Decodejson()

}

func Encodejson() {
	data := []course{
		{"Mello", 10, "tag"},
		{"Hello", 20, "tag"},
		{"hello", 30, "tag"},
		{"mello", 40, "tag"},
	}

	Finaljson, _ := json.MarshalIndent(data, "", "/t")
	fmt.Println(string(Finaljson))
}

func Decodejson() {
	//Fake json
	jsonfromweb := []byte(`
	{
		"name" : "MELLO",
		"price" : 100,
		"tags" : "tag"
	}
	`)

	var data course
	json.Unmarshal(jsonfromweb, &data)
	fmt.Printf("%#v", data)
}
