package main

import (
	"fmt"
	"net/url"
	"reflect"
)

var escapeMap = map[string]interface{}{
	"path":  url.PathEscape,
	"query": url.QueryEscape,
}

var unescapeMap = map[string]interface{}{
	"path":  url.PathUnescape,
	"query": url.QueryUnescape,
}

func Call(m map[string]interface{}, name string, param ...interface{}) (result []reflect.Value) {
	f := reflect.ValueOf(m[name])
	in := make([]reflect.Value, len(param))
	for i, v := range param {
		in[i] = reflect.ValueOf(v)
	}
	result = f.Call(in)
	return
}

func main() {
	a := "hello, 世界" //contain non-ascii code
	var urlValue map[string]string = make(map[string]string)
	fmt.Printf("Raw value is %s\n", a)
	for k, v := range escapeMap {
		t := v.(func(string) string) // convert interface to type
		temp := t(a)
		if _, ok := urlValue[k]; !ok {
			urlValue[k] = temp
		}

		fmt.Printf("%s escape: %v\n", k, temp)
	}
	fmt.Println("urlValue:%v\n", urlValue)

	for k, v := range urlValue {
		if _, ok := unescapeMap[k]; ok {
			t := unescapeMap[k].(func(string) (string, error))
			val, _ := t(v)
			fmt.Printf("%s unescape to raw %s.\n", k, val)
		}
	}

}
