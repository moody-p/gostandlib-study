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
同步信号是程序运行期间发生错误触发的信号。
- SIGBUS 意味着指针所对应的地址是有效地址，但总线不能正常使用该指针。通常是未对齐的数据访问所致。
- SIGFPE 是当一个进程执行了一个错误的算术操作时发送给它的信号。
- SIGSEGV 味着指针所对应的地址是无效地址，没有物理内存对
   应该地址。

Note：go程序碰到此类信号，go程序会进入panic。
  
### 异步信号
kernel或者其他程序发送给此程序的信号是异步信号。
几个示例信号：
- SIGHUP 程序丢失终端
- SIGINT 在程序执行期间，交互终端按下CTRL+C
- SIGQUIT 在程序执行期间，交换终端键盘按下CTRL+/


## go语言对信号的处理
#### go 信号的默认行为
1. 同步信号会引发panic
2. SIGHUP,SIGINT,SIGTERM 程序退出
3. SIGQUIT,SIGILL,SIGTRAP,SIGABRT,SIGSTEFLT,SIGSYS 程序退出，并且dump出栈信息
4. SIGTSTP,SIGTTIN,SIGTTOU获取系统默认行为
5. SIGPROF 性能测试用

#### 改变信号的默认处理函数
**func Ignore(sig ...os.Signal)**
忽略某些信号，相当于C函数signal(sig, SIG_IGN)
**func Ignored(sig os.Signal) bool**
判断某个信号是否是被忽悠的
**func Notify(c chan<-, sig ...os.Signal)**
1. 收到的信号，发往channel c
2. c 必须有足够的缓存来储存信号，也就是buffer channel

    一个处理一个信号的c，可以设置buffer为1
3. 可以调用notify多次设置多个channel，接收同一信号。
4. 清空notify使用Stop

