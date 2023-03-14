package main

import "fmt"

// 简单工厂方法模式

// 抽象层
// 水果接口
type Fruit interface {
	Show()
}

// 实现层

type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple...") }

type Pear struct{}

func (a *Pear) Show() { fmt.Println("this is pear...") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana...") }

// 工厂类
type Factory struct{}

func (f *Factory) CreateFurit(kind string) Fruit {
	var friut Fruit
	if kind == "apple" {
		friut = new(Apple)
	}
	if kind == "banana" {
		friut = new(Banana)
	}
	if kind=="pear"{
		friut=new(Pear)
	}
	return friut
}

func main(){
	factory:=new(Factory)

	apple:=factory.CreateFurit("apple")
	apple.Show()
}