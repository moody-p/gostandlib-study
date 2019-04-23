# golang 标准库（encoding/json）学习
## 什么是json
json的全称是 JavaScript Object Notation，是一种轻量级的数据交换格式。易于人阅读和编写，同时易于机器解析和生成。
json有两种类型的数据，用go语言来描述这两种类型：
- map “名称/值”对，用{}标识
- slice 也可以说是数组，用[]标识
  ```
  {
      "Name": "Joe", 
      "Age": 18,
  }
  ["apple", "banana", "pear"]
  ```
## golang如何处理json
golang标准库encoding/json是用来处理json数据的，它有两组函数或者方法来处理json数据。
| 函数名    | 输入         | 输出        | 描述                        |
|:---------|:-------------|:------------|:--------------------------:|
|Marshal   |v interface{} |[]byte, error|编码v成json数据     To Json  |
|Unmarshal |[]byte, v(指针)| error      |解码json，储存在v中 From Json|
|Encode    |v interface{}  | error      | 编码v成json数据 输出到 io.Writer|
|Decode    |v (指针)       | error       | 解码json， 输入来自io.Reader |
### 编码
### 解码
### 辅助函数
**function Compact(* bytes.Buffer, []byte) err**

compact的意思是把...紧压在一起， 我们可以这样理解，buffer之中已经有一些数据，我们希望把编码后的json数据，加入到buffer中去。
代码示例：https://github.com/MoonNan/gostandlib-study/blob/master/encoding/json/compact-example.go
```
var b bytes.Buffer \\ 声明变量
n, err := b.Write([]byte("Employee info: \n")) \\buffer中填充一些数据
err = json.Compact(&b, jData) //把json数据添加到buffer中去
b.WriteTo(os.Stdout) //输出到标准输出
/*
jdata的内容 {"name":"joe","age":20,"hobby":["Basketball","Football","Reading"]}
标准输出的内容
Employee info: 
{"name":"joe","age":20,"hobby":["Basketball","Football","Reading"]}
*/
```

**func HTMLEscape(dst *bytes.Buffer, src []byte)**

我们知道<,>,&在html的语法中有自己的意义，如果我们编码的json中有这些字符，html的<script></script>使用json数据，这容易引起页面的问题。这个函数就是把这几个字符转化成unicode字符。
|Source|Destination|
|------|:---------:|
|<|\u003c|
|>|\u003e|
|&|\u0026|
|\u+2028|\u2028|
|\u+2029|\u2029|
```
实例代码：https://github.com/MoonNan/gostandlib-study/blob/master/encoding/json/HTMLEscape-example.go
// 编码前的数据：
myTeams := []jdata.Person{
		{"Joe", 18, []string{"<basketball>", "<football>", "<reading>"}},
		{"Jason", 20, []string{"pingpang&basketball", "Swim"}},
		{"Smith", 19, []string{"game&game"}},
	}
[{"name":"Joe","age":18,"hobby":["\u003cbasketball\u003e","\u003cfootball\u003e","\u003creading\u003e"]},{"name":"Jason","age":20,"hobby":["pingpang\u0026basketball","Swim"]},{"name":"Smith","age":19,"hobby":["game\u0026game"]}]
```