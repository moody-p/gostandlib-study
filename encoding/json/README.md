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
| :-------- | :------------- | :------------ | :------------------------: |
| Marshal   | v interface{} | []byte, error | 编码v成json数据     To Json  |
| Unmarshal | []byte, v(指针) | error      | 解码json，储存在v中 From Json |
| Encode    | v interface{}  | error     | 编码v成json数据 输出到 io.Writer |
| Decode    | v (指针)       | error       | 解码json， 输入来自io.Reader |

### 编码函数

编码函数： **func Marshal(v interface{}) ([]byte, error)**
编码规则： 递归遍历v，v中的一个值实现了**Marshaler interface**，调用**MarshalJSON method**去编码，如果没有这种方法，但是实现了**encoding.TextMarshl**那么使用**MarshalText**去编码。

编码对应表：

go value | JSON value
-------- | ----------
布尔型(true false)  | 布尔型(true false)
浮点数  | 数字型
整数    | 数字型
字符型  | 字符型
数组 切片 | 数组
byte 数组 | base-64 encoded 字符串
nil 数组  | null
结构体     | JSON 对象
map        | JSON 对象

JSON 结构体中TAG
使用json key来标识这个结构体中的field，对应编码后json中key值。
omitempty 如果值是false， 0， 空指针，空接口，空数组，空切片，空映射，空字符串，将被忽略。
“-” 这个field将会被忽略。
示例代码：https://github.com/MoonNan/gostandlib-study/blob/master/encoding/json/tag-exmple.go

其他编码函数： **func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)**
格式化编码输出
示例代码：https://github.com/MoonNan/gostandlib-study/blob/master/encoding/json/marshalindent-example.go

### 解码函数
解码函数： **func Unmarshal(data []byte, v interface{}) error**
- data 是一个JSON数据
- v 是存储解码数据的指针
- error 返回的错误类型

各类型解码
- 指针
  - null解码成空指针
  - 指针非空，解码成指针指向的数据
  - 指针为空，解码函数为他创建新指针，并且为其填充数据。
- 接口
  - UnmarshalJSON method
  - UnmarshlText method
- 结构体
  - field（tags）相对应解码
  - 碰到不匹配field，忽略掉（默认行为）
- 切片
  - reset slice length to 0， 添加元素到slice
  - 空的JSON 数组，重新生成一个空slice
- 数组
  - go数组小于JSON数组，多余的将忽略
  - go数组大于JSON数组，多余填充0值
- 映射
  - 映射为nil， 重新生成map，填充数据返回
  - 映射非空，直接填充数据
  
### 编码器
流程：
1. 调用**func NewEncoder(w io.Writer) *Encoder**生成一个编码器
   这里input是一个io.Writer， 返回的是一个Encoder的指针
2. 调用Encode方法，把编码数据发送到io.Writer, 后面会跟一个换行符

辅助方法：
SetEscapeHTML 和 函数HTMLEscape 同样的效果
SetIndent 和 MarshalIndent 同样的效果
### 解码器
流程：
1. 调用**func NewDecoder(r io.Reader) *Decoder**生成一个解码器
   输入是一个io.Reader，返回一个解码器
2. 调用Decode方法解码

辅助方法：
Buffered 数据暂存在解码器的buffer中，等下一次调用Decode再重新解码到io.Reader
DisallowUnknownField 如果解码的field和接收的field不匹配将会报错
More 是否有新的数组或者对象可以解析
Token 下一个JSON的token
UseNumber 使用interface{}处理数字型而不是float
### RawMessage
RawMessage 是已经预先编码好的JSON数据，他可以实现延迟解码，或者联合其他数据重新编码。
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

Source | Destination
------ | ---------
< | \u003c
| > | \u003e
& | \u0026
\u+2028 | \u2028
\u+2029 | \u2029

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

**func Valid(data []byte) bool**
data是不是一个可用的JSON编码


