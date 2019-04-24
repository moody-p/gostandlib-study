package main

import (
	"encoding/json"
	"fmt"

	jdata "./jdata"
)

func main() {
	v := jdata.TagExample{"test", 10, 1, 0}
	data, _ := json.Marshal(v)
	fmt.Printf("%s\n", data)
}
