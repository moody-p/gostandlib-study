# golang(os/signal)阅读
## 信号和信号函数
在unix-like系统中，我们要经常用到信号和信号处理函数。我们在linux编程中，经常要处理程序接收到的函数。比如：
```
signal(SIGINT, SignalHandler)
```
上面的函数注册SIGINT的信号处理函数SignalHandler，当程序运行中，当收到SIGINT信号，程序相应的执行SignalHandler函数。

## 信号的分类
### 特殊信号
SIGKILL和SIGSTOP不能被程序修改。
- SIGKILL 杀死一个程序
- SIGSTOP 暂停一个程序
```
//运行一个程序
for {
		time.Sleep(1 * time.Second)
	}
//程序是个死循环
// 使用SIGINT 退出
^Csignal: interrupt
// 程序退出
// 忽略SIGINT，SIGSTOP和SIGKILL
signal.Ignore(syscall.SIGKILL, syscall.SIGINT, syscall.SIGSTOP)
//SIGINT ^C
//SIGSTOP
[3]  + 25295 suspended  go run non-caught-signal.go
//SIGKILL
mchen    25413 25295  0 06:04 pts/5    00:00:00 [non-caught-sign] <defunct>
```
### 同步信号
### 异步信号

## go语言对信号的处理
### go语言中信号的行为
## go信号的应用