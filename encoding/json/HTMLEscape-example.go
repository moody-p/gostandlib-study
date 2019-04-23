package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	jdata "./jdata"
)

func main() {
	var b bytes.Buffer
	myTeams := []jdata.Person{
		{"Joe", 18, []string{"<basketball>", "<football>", "<reading>"}},
		{"Jason", 20, []string{"pingpang&basketball", "Swim"}},
		{"Smith", 19, []string{"game&game"}},
	}
	_, err := b.Write([]byte(`<html><h1>Test<\h1><body>Run HTML Escape Example<script>`))
	if err != nil {
		log.Fatal(err)
	}
	jsonData, err := json.Marshal(myTeams)
	if err != nil {
		log.Fatal(err)
	}
	json.HTMLEscape(&b, jsonData)
	_, err = b.Write([]byte(`</script></body></html>`))
	b.WriteTo(os.Stdout)
}
