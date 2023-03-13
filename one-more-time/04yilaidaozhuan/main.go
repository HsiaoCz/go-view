package main

import "fmt"

// 使用依赖倒转原则实现司机开车

// 抽象的车类
type Car interface {
	Run()
}

// 抽象的司机

type Driver interface {
	Drive(Car)
}

// 具体的车类
type BMW struct{}

func (b *BMW) Run() { fmt.Println("BMW is running...") }

type Bnez struct{}

func (b *Bnez) Run() { fmt.Println("Benz is running...") }

// 具体的司机
type Zhangsan struct{}

func (z *Zhangsan) Drive(car Car) {
	fmt.Println("zhangsan is driving")
	car.Run()
}

type Lisi struct{}

func (l *Lisi) Drive(car Car) {
	fmt.Println("Lisi is driving")
	car.Run()
}

func main() {
	// car
	benz := new(Bnez)
	bmw := new(BMW)

	// 驾驶员
	zhangsan := new(Zhangsan)
	lisi := new(Lisi)

	// 具体的业务
	// 张三开奔驰 李四开宝马
	zhangsan.Drive(benz)
	lisi.Drive(bmw)

	// 现在张三想开宝马，李四想开奔驰
	zhangsan.Drive(bmw)
	lisi.Drive(benz)

}
