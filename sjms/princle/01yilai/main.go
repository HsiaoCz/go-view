package main

import "fmt"

// 依赖倒转原则
// 依赖于抽象，而不是依赖于具体的类

// 假如现在有司机类，汽车类
// 有司机张三，李四
// 有汽车宝马，奔驰

// 抽象层
type Car interface {
	Run()
}
type Driver interface {
	Drive(Car)
}

// 业务层
// car
type Benz struct{}

func (b *Benz) Run() {
	fmt.Println("benz is running")
}

type BMW struct{}

func (b *BMW) Run() {
	fmt.Println("bwm is running")
}

type Fent struct{}

func (f *Fent) Run() {
	fmt.Println("fent is running")
}

// driver

type Zhangs struct{}

func (z *Zhangs) Drive(car Car) {
	fmt.Println("张三在开车")
	car.Run()
}

type Lisi struct{}

func (z *Lisi) Drive(car Car) {
	fmt.Println("李四在开车")
	car.Run()
}

type Wangwu struct{}

func (w *Wangwu) Drive(car Car) {
	fmt.Println("王五在开车")
	car.Run()
}

// 实现层
func main() {
	benz := new(Benz)
	zs := new(Zhangs)
	zs.Drive(benz)
}
