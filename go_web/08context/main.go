package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context
// context 上下文管理
// context 主要用来实现并发协调以及对groutine的生命周期控制
// context 的核心数据结构
//
//	type Context interface{
//	 Deadline()(time.Time,bool) # context在何时有一个生命的终点
//	 Done()<-chan struct{} # 只读channel 表示context结束
//	 Err()error  # 返回错误
//	 Value(key any)any # 返回context中对应的key值
//	}
//
// 官方实现的第一个context
// emptyCtx是一个空的context,本质上是一个整型
// context.Background()本质上也是指向这个context的指针
// 这个空context拥有四个方法
// Deadline()方法会返回一个公元元年时间以及false的flag，标识当前context
// 不存在过期时间
// Done方法返回一个nil值，用户无论往nil中写入或者读取，都会阻塞
// Err方法返回值永远为nil
// value方法的返回的value永远为nil

// context.Background() 和context.TODO()
// 这两个都是返回的emptyCtx的实例

var wg sync.WaitGroup

// 使用全局变量通知子goroutine退出
// 使用通道的方式通知退出
// 使用context进行控制退出
var notify bool

var exitchan chan bool

func main() {
	c, cancel := context.WithCancel(context.Background())
	wg.Add(3)
	go f()
	go f1()
	go f3(c)
	time.Sleep(time.Microsecond * 100)
	notify = true
	exitchan <- true
	cancel()
	wg.Wait()
}

// type cancelCtx struct{
// Context # cancelContext不能作为Cotext树的鼻祖，这是它的父context
// mu sync.Mutex # 互斥锁，作为并发锁
// done atomic.Value # 实际为空结构体的channel
// children map[canceler]struct{} # 子节点
// err error
// }
//
// context的使用
// 首先看一个例子，如何控制子goroutine退出

func f() {
	defer wg.Done()
	for {
		if notify {
			break
		}
		fmt.Println("Hello")
		time.Sleep(time.Second * 1)
	}
}

func f1() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("Hello")
		time.Sleep(time.Second * 1)
		select {
		case <-exitchan:
			break LOOP
		default:
		}
	}
}

func f3(c context.Context) {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 2)
		select {
		case <-c.Done():
			break LOOP
		default:
		}
	}
}
