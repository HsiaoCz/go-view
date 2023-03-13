package main

import "fmt"

// 依赖倒转原则
// 依赖于抽象而不是依赖于具体的类

// 来看一个例子
// 假如说有两个司机，张三李四
// 张三开奔驰，李四开宝马

// 汽车
type Bnez struct{}

func (b *Bnez) Run() { fmt.Println("Benz is running...") }

type BMW struct{}

func (b *BMW) Run() { fmt.Println("BMW is running...") }

// 驾驶员
type Zhangsan struct{}

func (z *Zhangsan) Drive(benz *Bnez) {
	fmt.Println("张三 开奔驰...")
	benz.Run()
}

type Lisi struct{}

func (l *Lisi) Drive(bmw *BMW) {
	fmt.Println("李四 开奔驰...")
	bmw.Run()
}

// 实现层

func main() {
	// zhangsan 开奔驰
	benz := &Bnez{}

	zhangsan := &Zhangsan{}
	zhangsan.Drive(benz)

	// lisi 开奔驰
	bmw := &BMW{}

	lisi := &Lisi{}
	lisi.Drive(bmw)
}
