package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	qStrings := "a=b&c=e"
	val, err := url.ParseQuery(qStrings)
	if err != nil {
		log.Fatalln(err)
	}
	// Print Values of Query
	QueryPrint(val)
	// add
	val.Add("d", "df")
	QueryPrint(val)
	// Get
	s := val.Get("c")
	fmt.Print(s)
	//Set
	val.Set("c", "good")
	QueryPrint(val)
	//delete
	val.Del("a")
	QueryPrint(val)

	//encode
	s = val.Encode()
	fmt.Println(s)
}

func QueryPrint(val url.Values) {
	for k, v := range val {
		fmt.Println(k, ":", v)
	}
	fmt.Println()
	fmt.Println()
}
