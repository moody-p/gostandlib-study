# go标准库（net/url）阅读

## url 介绍

- 问题1： 什么是url？
  
  当我们使用搜索引擎搜索东西的，我们第一步先在浏览器的地址栏里输入 google.com， www.google.com， http://www.google.com。我们在地址栏里输入的东西就是URL。
  URL 又叫统一资源定位符， 用于定位我们要访问的文档或者其他资源。

- 问题2： URL有什么样的格式？
  
  scheme://[userinfo]@[host]:[port]/path?key1=value1&key2=value2#fragment
    协议 （http, https, file, ftp)

    用户信息， 是可选的

    主机名字或者ip地址，定位网络位置

    port 服务端口， 一般端口表示提供了某项服务

    path 主机上的目录

    ？后的问query信息， key1=value1是表示key值和value值， &是连接符

    ‘#’ 后面的是fragment信息

- 问题3： URL如何处理非ascii编码？
  
  非ascii编码，使用%后跟两位16进制数表示，如%AB
  
  URL中不能有空格， 空格用“+”表示。

## url 库

我们知道url中不能有空格和非ascii字符。当我们的url path字段中出现这样的字符，我们该如何处理呢。

- url path字段中有空格和非ascii字符
  
  **func PathEscape(s string) string**
  返回的string将是url可以使用的%后跟两位16进制数的形式

- 如何把url中的path字段还原成原始模式

  **func PathUnescape(s string) (string, err)**

```
    a := "hello, 世界" //contain non-ascii code
	b := url.PathEscape(a)
	fmt.Printf("%v\n", b)
	// Output: hello%2C%20%E4%B8%96%E7%95%8C
	c, _ := url.PathUnescape(b)
	fmt.Printf("%v\n", c)
	// Output: hello, 世界
```
Note: path中的空格和非ascii字符使用同样的方式处理。

- query字段中出现非ascii字符和空格如何处理
  
  **func QueryEscape(s string) string**
  **func QueryEscape(s string) (string, error)**
  Note：空格的处理和path不太一样，' '将会编程'+'
  示例: https://github.com/MoonNan/gostandlib-study/blob/master/net/url/unicode-example.go

- 如何解析URL

主要通过Parse函数来解析URL地址。
示例代码：https://github.com/MoonNan/gostandlib-study/blob/master/net/url/parse-example.go

- 如何处理Query 数据

Query字段可以通过ParseQuery函数来处理。ParseQuery根据传入的字符串，生成一个Values字典。
```
type Values map[string][]string
```
方法：
Encode 把Values生成字符串
Get Set Del Add
示例代码：https://github.com/MoonNan/gostandlib-study/blob/master/net/url/query-example.go

- 如何处理userinfo

type Userinfo用来处理用户数据
User和UserPassword函数生成Userinfo 结构体
方法：Userinfo Password Username