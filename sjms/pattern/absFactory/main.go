package main

import "fmt"

// 抽象工厂方法模式
// 产品等级结构和产品族

// 抽象层
type AppleFruit interface {
	ShowApple()
}
type BananaFruit interface {
	ShowBanana()
}
type PearFruit interface {
	ShowPear()
}

// 抽象的工厂

type AbstractFactory interface {
	CreateApple() AppleFruit
	CreateBanana() BananaFruit
	CreatePear() PearFruit
}

// 实现层

// 中国的产品族

type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() { fmt.Println("china apple") }

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() { fmt.Println("china banana") }

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() { fmt.Println("china pear") }

// 中国工厂

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AppleFruit {
	apple := new(ChinaApple)
	return apple
}
func (cf *ChinaFactory) CreateBanana() BananaFruit {
	banana := new(ChinaBanana)
	return banana
}
func (cf *ChinaFactory) CreatePear() PearFruit {
	pear := new(ChinaPear)
	return pear
}

// 日本产品族，日本工厂和中国产品族和中国工厂代码一样
// 逻辑层

func main() {
	chinaFactory := new(ChinaFactory)
	apple := chinaFactory.CreateApple()
	apple.ShowApple()

	banana := chinaFactory.CreateBanana()
	banana.ShowBanana()

	pear := chinaFactory.CreatePear()
	pear.ShowPear()
}
