package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}
	json, err := json.MarshalIndent(data, ">", "\t\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", json)
}
