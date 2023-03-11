## go channel

### 1、什么是csp

所谓的csp就是通过通信来实现内存共享

### 2、channel的底层数据结构

```go
type hchan struct {
	// chan 里元素数量
	qcount   uint
	// chan 底层循环数组的长度
	dataqsiz uint
	// 指向底层循环数组的指针
	// 只针对有缓冲的 channel
	buf      unsafe.Pointer
	// chan 中元素大小
	elemsize uint16
	// chan 是否被关闭的标志
	closed   uint32
	// chan 中元素类型
	elemtype *_type // element type
	// 已发送元素在循环数组中的索引
	sendx    uint   // send index
	// 已接收元素在循环数组中的索引
	recvx    uint   // receive index
	// 等待接收的 goroutine 队列
	recvq    waitq  // list of recv waiters
	// 等待发送的 goroutine 队列
	sendq    waitq  // list of send waiters

	// 保护 hchan 中所有字段
	lock mutex
}
```

buf指向底层循环数组，只有缓冲型的channel才有
sendx，recvx均指向底层循环数组，表示当前可以发送和接受的元素位置索引值
sendq，和recvq分别表示被阻塞的groutine，这些groutine由于尝试读取channel或向channel发送数据而被阻塞

waitq是sudog的一个双向链表，而sudog实际是对groutine的一个封装

```go
type waitq struct {
	first *sudog
	last  *sudog
}
```
lock用来保证每一个读或写channel操作都是原子的

**创建channel**

```go
// 无缓冲的通道
ch1:=make(chan int)
// 有缓冲的通道
ch2:=make(chan int,10) //这里的10就是dataqsize的数值
```
在底层创建channel的函数是
```go
func makechan(t *chantype,size int64)*hchan
```

### 3、向channel发送数据的过程是什么样的？

