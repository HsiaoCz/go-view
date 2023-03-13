package main

import "fmt"

// 依赖倒转原则
// 依赖于抽象而不是依赖于具体的类

// 来看一个例子
// 假如说有两个司机，张三李四
// 张三开奔驰，李四开宝马

// 我们使用平铺设计完成了这个需求
// 但是现在，张三不仅能开奔驰，还能开宝马
// 李四不仅能开宝马还能开奔驰
// 很容易想到，这种情况下，添加代码即可
// 但是显然添加代码，会造成很多的耦合
// 这时候就可以使用依赖倒转原则

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
	fmt.Println("李四 开宝马...")
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
