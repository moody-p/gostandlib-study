***go标准库（fmt）学习***

    每种编程语言都有自己的格式化输入和输出。c语言是通过标准输入输出库（stdio），python语言是语言的一部分（print）。go语言是通过库（fmt）来实现格式化输入输出的功能。
**1 打印**

*1.1 打印格式*

| 格式符号  |描述                      |
|:---------|:------------------------:|
|%v        |打印默认格式               |
|%+v       |当打印结构体时，会添加字段名 |
|%#v       |go语法展示数据             |
|%T        |go语法展示数据的类型        |
|%%        |打印%                       |
测试数据：

```
type Person struct {
	Name string
	age  int
}

var SimpleData = map[string]interface{}{
	"int":     1,
	"string":  "hello",
	"float":   1.23,
	"complex": complex(1, 2),
    "bool": 0 == 0,
}
var ComplexData = map[string]interface{}{
	"slice": []int{1, 2, 3},
	"map": map[string]int{
		"a": 1,
		"b": 2,
	},
	"struct": Person{"Moon", 33},
}
```
%v和%+v打印都是一样的，只有struct打印是不同的。
```
struct		{Moon 33} //%v 打印结构体
struct		{Name:Moon age:33} // %+v打印结构体
```
%#v的打印
```
"slice"		[]int{1, 2, 3} //go语法的字符串会添加双引号， slice是go语法定义的模式
"map"		map[string]int{"a":1, "b":2} //map的打印
```
%T打印类型
```
map		map[string]int //map的类型
slice		[]int//slice的类型
```
上面是一般模式的打印，下面我们考虑精确到具体的类型的打印。
- 如何打印boolean类型
  - %t可以打印boolean值，值是true或者false
- 如何打印interger
  - 二进制%b
  - 十进制%d
  - 八进制%o
  - 十六进制%x或者%X
  - unicode 格式%U
  - %c unicode code point
  ```
    var IntegerFormat = map[string]string{
        "base10":   "%d",
        "base8":    "%o",
        "base16":   "%x",
        "unicode":  "%U",
        "unicodec": "%c",
    }

    base10:42
    base8:52
    base16:2a
    unicode:U+002A
    unicodec:*
  ```
- 如何打印float
  - 指数形式 2为底的指数形式%b，e为底的指数形式%e， %E
  - 非指数形式， %f 或者%F
  - 通用形式 %g，或者%G 根据实际情况使用%e或者%f
    ```
    fmt.Printf("%b\n", SimpleData["float"])
	fmt.Printf("%f\n", SimpleData["float"])
	fmt.Printf("%e\n", SimpleData["float"])
    //Output
    //5539427541665710p-52
    //1.230000
    //1.230000e+00
    ```
- 如何打印string
  - go语法形式 %q
  - 普通形式 %s
  - 16进制模式 %x或者%X
  ```
  /*
    string normal: hello, world
    string go syntax:"hello, world"
    string normal: good morning
    string go syntax:"good morning"
    string base 16:676f6f64206d6f726e696e67
  */
  ```
- 如何打印slice
  - 打印首元素的地址 %p
- 如何打印point
  - 打印指针地址
```
fmt.Printf("slice %p\n", ComplexData["slice"])
a := 5
p := &a
fmt.Printf("Point %p\n", p)
//slice 0xc000064140
//Point 0xc000066088
```
**1.2 width precision**
duration .表示duration
实例最好说明这两个东西， 比如"%9.2f",宽度是9，精度是2，我们用实例来展示这两个东西。
```
fmt.Printf("%f\n", SimpleData["float"])
fmt.Printf("%6.2f\n", SimpleData["float"])
fmt.Printf("%  10f\n", SimpleData["float"])
fmt.Printf("%.5f\n", SimpleData["float"])
/*
Output:
123456789.123457 //默认的精度是6
123456789.12
 123456789.123457 //10的宽度，实际有11，舍弃一个空隔
123456789.12346 
*/
```
    宽度和精度都是通过unicode码点的单位数量来度量的，也就是rune个数。其中"%5.2f"中的5和2这种整数类型可以用*代替，通过后面的参数传递。
    ```
    fmt.Printf("%5.2f", 1.2345)
	fmt.Printf("%*.*f", 5, 2, 1.2345)
    //两个的输出都是1.23
    ```
- 大部分的数值的宽度，是能够展示这个数据的最小runes加上格式化指定的空格数。
- 字符串和byte切片，精度控制的输入的长度，如果是%x使用字节数来度量
  ```
  s := "hello, world"
  b := []byte("good morning")
  fmt.Printf("%.4s %.5s\n", s, b)
  //Output: 
  //hell good 
  ```
- 浮点数，width代表的是整数部分，precision代表的小数点后的位数
- 复数， width和precision对实部和虚部都有效

其他标记：
- + 打印符号
- - 右边填充空格
- ‘#’ 八进制填充O， 16进制填充0x ...
- ‘ ’填充空格
- 0 填充0

Print默认使用%v， Pirntln自动填充空格并且添加换行符。
**1.3 接口**
传入的操作数是接口，打印的是运行时该变量的值而不是接口本身。
- 如果操作数是一个reflect.value，打印将使用它运行时的值
- 如果传入接口的类型实现了Formater，它将会被调用。
- 如果使用%#v,类型实现了GoString方法，这个将会被调用
- 如果实现了error的interface，Error方法将会被调用（需要string的verb）
- 如果实现了String该方法，它将会被调用（需要string的verb）

**1.4显式的位置参数**
%[index]verb 通过后面的传入操作数的位置来匹配verb
```
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
// 16 先传给第一个verb 17 传第二个verb， 后面的%#[1]x又把16传过来，后面紧跟
```
**1.5错误类型**
- 错误类型或者未知verb
- 太多参数
- 太少参数
- 非整形的width和precision
- index的错误使用

**2 scan**
scan就是从格式话的文本中获取值放入变量中去。
Scan Scanf Scanln 是从标准输入读取
Fscan Fscanf Fscanln 是从io.reader中读取
Sscan Sscanf Sscanln 是从字符串中读取

函数对"\n"是怎么处理的？
- Scan Fscan Sscan 把"\n"当作空格处理
- Scanln Fscanln Sscanln 中"\n"是它的结束标志，Fscanln同样把EOF作为结束标志
- Scanf Fscanf Sscanf取决于格式化字符串要求

Scan系列函数接受指向特定类型的指针或者实现Scan方法的类型。
