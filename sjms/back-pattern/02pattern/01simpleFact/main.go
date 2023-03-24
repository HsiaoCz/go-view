package main

import "fmt"

// 简单工厂方法模式

// 抽象层

type Fruit interface {
	Show()
}

// 实现层

type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is a Apple") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is a Banana") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is a Pear") }

// 工厂类
type Factory struct{}

func (f *Factory) CreateFurit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple) // 满足多态 父类指针指向子类对象
	}
	if kind == "banana" {
		fruit = new(Banana)
	}
	if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit

}
func main() {
	factory := new(Factory)

	apple := factory.CreateFurit("apple")
	apple.Show()
	banana := factory.CreateFurit("banana")
	banana.Show()
	pear := factory.CreateFurit("pear")
	pear.Show()
}
