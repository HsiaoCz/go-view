package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享内存
// 多线程共享内存来进行通信
// 通过加锁来访问共享数据，如数组、map或结构体
// go语言也实现了这种并发模型
// 共享内存，实际就是全局变量
var mp sync.Map

func rwGlobalMemory() {
	if value, exists := mp.Load("mykey"); exists {
		fmt.Println(value)
	} else {
		mp.Store("mykey", "myvalue")
	}
}

func main() {
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()
	go rwGlobalMemory()

	time.Sleep(time.Second)
}
