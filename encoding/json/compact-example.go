package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	jdata "./jdata"
)

func main() {
	joe := jdata.Person{"joe", 20, []string{"Basketball", "Football", "Reading"}}
	var b bytes.Buffer
	n, err := b.Write([]byte("Employee info: \n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write %d bytes to buffer.\n", n)
	jData, err := json.Marshal(joe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", jData)
	err = json.Compact(&b, jData)
	if err != nil {
		log.Fatal(err)
	}
	b.WriteTo(os.Stdout)
}
