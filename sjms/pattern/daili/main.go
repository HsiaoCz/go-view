package main

import (
	"fmt"
)

// 代理模式

type Goods struct {
	Kind string //商品的种类
	Fact bool   // 商品的真伪
}

// 抽象层
type Shopping interface {
	Buy(goods *Goods) // 买东西
}

// 实现层
type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) { fmt.Println("去韩国进行了购物") }

type AmericaShopping struct{}

func (a *AmericaShopping) Buy(goods *Goods) { fmt.Println("去美国进行了购物") }

type AffracShopping struct{}

func (a *AffracShopping) Buy(goods *Goods) { fmt.Println("去非洲进行了购物") }

// 代理
type OverSeasProxy struct {
	shopping Shopping // 代理某个主题
}

func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{
		shopping: shopping,
	}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	// 1、辨别真伪
	if op.distinguish(goods) {
		op.shopping.Buy(goods)
		op.check(goods)
	}
	// 2、调用具体要被代理的购物方式
	// 3、海关安检

}

func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了检查")
	if !goods.Fact {
		fmt.Println("发现假货,", goods.Kind, "不应该购买")
	}
	return goods.Fact
}

func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对商品进行了安检...", goods.Kind)
}

func main() {
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: false,
	}

	g2 := Goods{
		Kind: "CET4",
		Fact: true,
	}
	shopping := new(KoreaShopping)

	proxy := NewProxy(shopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)

}
