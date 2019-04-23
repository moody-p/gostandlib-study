package main

import "fmt"

type Person struct {
	Name string
	age  int
}

var SimpleData = map[string]interface{}{
	"int":     42,
	"string":  "hello",
	"float":   123456789.123456789,
	"complex": complex(1, 2),
	"bool":    0 == 0,
}
var ComplexData = map[string]interface{}{
	"slice": []int{1, 2, 3},
	"map": map[string]int{
		"a": 1,
		"b": 2,
	},
	"struct": Person{"Moon", 33},
}
var IntegerFormat = map[string]string{
	"base10":   "%d",
	"base8":    "%o",
	"base16":   "%x",
	"unicode":  "%U",
	"unicodec": "%c",
}

func main() {
	//%v 打印默认格式
	/*
		for k, v := range SimpleData {
			fmt.Printf("%v\t\t%v\n", k, v)
			fmt.Printf("%+v\t\t%+v\n", k, v)
			fmt.Printf("%#v\t\t%#v\n", k, v)
			fmt.Printf("%v\t\t%T\n", k, v)
		}

		for k, v := range ComplexData {
			fmt.Printf("%v\t\t%v\n", k, v)
			fmt.Printf("%+v\t\t%+v\n", k, v)
			fmt.Printf("%#v\t\t%#v\n", k, v)
			fmt.Printf("%v\t\t%T\n", k, v)
		}
	*/
	/*
		fmt.Printf("Bool print:%t\n", SimpleData["bool"])
		for k, v := range IntegerFormat {
			line := k + ":" + v + "\n"
			fmt.Printf(line, SimpleData["int"])
		}
	*/
	/*
		fmt.Printf("%b\n", SimpleData["float"])
		fmt.Printf("%f\n", SimpleData["float"])
		fmt.Printf("%e\n", SimpleData["float"])
	*/

	s := "hello, world"
	b := []byte("good morning")

	fmt.Printf("string normal: %s\n", s)
	fmt.Printf("string go syntax:%q\n", s)

	fmt.Printf("string normal: %s\n", b)
	fmt.Printf("string go syntax:%q\n", b)
	fmt.Printf("string base 16:%x\n", b)
	fmt.Printf("%.4s %.5s\n", s, b)

	/*
		fmt.Printf("slice %p\n", ComplexData["slice"])
		a := 5
		p := &a
		fmt.Printf("Point %p\n", p)
	*/
	/*
		fmt.Printf("%f\n", SimpleData["float"])
		fmt.Printf("%6.2f\n", SimpleData["float"])
		fmt.Printf("%  10f\n", SimpleData["float"])
		fmt.Printf("%.5f\n", SimpleData["float"])
	*/
	/*
		fmt.Printf("%5.2f", 1.2345)
		fmt.Printf("%*.*f", 5, 2, 1.2345)
		fmt.Printf("%[2]*.[3]*[1]f", 1.2345, 5, 2)
	*/

}
