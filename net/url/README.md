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
