package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Message struct {
	Name, Text string
}

func main() {
	jsonStream := `[
	{"Name": "Ed", "Text": "Knock knock"},
	{"Name": "Sam", "Text": "Who's there?"},
	{"Name": "Ed", "Text": "Go fmt."},
	{"Name": "Sam", "Text": "Go fmt who?"},
	{"Name": "Ed", "Text": "Go fmt yourself!"},
	]`
	dec := json.NewDecoder(strings.NewReader(jsonStream))

	t, err := dec.Token()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	for dec.More() {
		var m Message
		err := dec.Decode(&m)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%v: %v\n", m.Name, m.Text)
	}

	t, err = dec.Token()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}
