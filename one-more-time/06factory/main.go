package main

import "fmt"

// 简单工厂方法模式完全不符合开闭
// 工厂方法模式在简单工厂方法模式的基础上添加了开闭原则

// 抽象层
type Fruit interface {
	Show()
}

// 抽象的工厂

type Factory interface {
	CreateFruit() Fruit
}

// 具体的类

type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple...") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana...") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is pear...") }

// 工厂类
type AppleFactory struct{}

func (a *AppleFactory) CreateFurit() Fruit {
	apple := new(Apple)
	return apple
}

type BananaFactory struct{}

func (b *BananaFactory) CreateFurit() Fruit {
	banana := new(Banana)
	return banana
}

type PearFactory struct{}

func (p *PearFactory) CreateFurit() Fruit {
	pear := new(Pear)
	return pear
}

// 业务逻辑层
func main() {
	appleFactory := new(AppleFactory)
	apple := appleFactory.CreateFurit()
	apple.Show()

	// banana
	bananaFactory := new(BananaFactory)
	banana := bananaFactory.CreateFurit()
	banana.Show()

	// pear
	PearFactory := new(PearFactory)
	pear := PearFactory.CreateFurit()
	pear.Show()
}
