package main

import (
	"fmt"
	"log"
	"net/url"
)

var urls = []string{"google.com",
	"www.google.com",
	"http://google.com",
	"https://www.google.com",
	"file://www.google.com/a/b/c",
	"tcp://www.google.com/a/b?c=d&e=f",
	"google.com/a/b?c=d&e=f#info",
}

func main() {
	for _, u := range urls {
		uStruct, err := url.Parse(u)
		if err != nil {
			log.Fatalln(err)
		}
		UrlPirnt(u, uStruct)
	}
}

func UrlPirnt(raw string, u *url.URL) {
	fmt.Printf("Raw String:%s\n", raw)
	fmt.Printf("Scheme: %s\n", u.Scheme)
	fmt.Printf("Opaque: %s\n", u.Opaque)
	fmt.Printf("User: %v\n", u.User)
	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Path: %s\n", u.Path)
	fmt.Printf("Raw path: %s\n", u.RawPath)
	fmt.Printf("Force Query: %t\n", u.ForceQuery)
	fmt.Printf("Raw Query: %s\n", u.RawQuery)
	fmt.Printf("Fragment: %s\n", u.Fragment)

	fmt.Println()
	fmt.Println()
}
