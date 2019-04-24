package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Animal struct {
	Name  string `json:"name"`
	Order string `json:"order"`
}

func main() {
	var jsonBlob = []byte(`[
		{"name": "Platypus", "order": "Monotremata"},
		{"name": "Auoll", "order": "Dasyuromorphia"}
		]`)
	var animals = []Animal{}
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", animals)
}
