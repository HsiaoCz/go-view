package main

import "fmt"

// 工厂方法模式

// 抽象层
// 水果接口
type Fruit interface {
	Show()
}

// 工厂接口
type Factory interface {
	CretaeFruit() Fruit
}

// 实现层
type Apple struct{}

func (a *Apple) Show() { fmt.Println("this is apple") }

type Banana struct{}

func (b *Banana) Show() { fmt.Println("this is banana") }

type Pear struct{}

func (p *Pear) Show() { fmt.Println("this is pear") }

// 工厂实现层
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
	// apple
	appleFactory := new(AppleFactory)
	apple := appleFactory.CreateFurit()
	apple.Show()

	// banana
	bananaFactory := new(BananaFactory)
	banana := bananaFactory.CreateFurit()
	banana.Show()

	//pear
	pearFactory := new(PearFactory)
	pear := pearFactory.CreateFurit()
	pear.Show()
}
