package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel
// channel 关闭一个已经关闭的channel panic
// 从一个已经关闭的channel取值，会取到channel
// 从一个已经关闭并且没有值的channel取值，会取到对应类型的零值
// 往一个无缓冲的channel发送值会一直阻塞，直到有接收操作

func main() {
	ch := make(chan int, 1)
	ch <- 10
	close(ch)
	for x := range ch {
		fmt.Println(x)
	}
	a := GetNumber()
	sum := Sum(a)
	fmt.Println(sum)
}

// 线程池
// 用来控制并发数量
// 使用channel个goroutine实现一个计算int64随机数各个位数和的程序

func GetNumber() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Int63n(10000)
}

func Sum(n int64) (sum int64) {
	var a int64
	for n > 0 {
		a = n % int64(10)
		n = n / 10
		sum = sum + a
	}
	return sum
}
