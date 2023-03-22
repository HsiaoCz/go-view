package main

import "fmt"

// 依赖倒转原则，依赖于抽象而不是依赖于具体的类
type Driver interface {
	Drive(Car)
}

type Car interface {
	Run()
}

// 实现类
type BMW struct{}

func (b *BMW) Run() { fmt.Println("宝马 is running...") }

type Benz struct{}

func (b *Benz) Run() { fmt.Println("Benz is running...") }

type Toyoto struct{}

func (t *Toyoto) Run() { fmt.Println("丰田 is running...") }

// 司机类
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
	benz := new(Benz)
	bmw := new(BMW)
	toyoto := new(Toyoto)

	zhangsan := new(Zhangsan)
	lisi := new(Lisi)

	zhangsan.Drive(benz)
	zhangsan.Drive(bmw)
	zhangsan.Drive(toyoto)

	lisi.Drive(benz)
	lisi.Drive(bmw)
	lisi.Drive(toyoto)
}
