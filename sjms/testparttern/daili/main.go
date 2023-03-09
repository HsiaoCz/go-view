package main

import "fmt"

// 代理模式
// 比如说购物
type Goods struct {
	Kind string
	Fact bool
}

// 抽象的购物
type Shopping interface {
	Buy(good *Goods)
}

// 具体的购物
type KoreaShopping struct{}

func (k *KoreaShopping) Buy(goods *Goods) { fmt.Println("在韩国进行了购物") }

type AmericaShopping struct{}

func (a *AmericaShopping) Buy(goods *Goods) { fmt.Println("在美国进行了购物") }

type AfricaShopping struct{}

func (a *AfricaShopping) Buy(goods *Goods) { fmt.Println("在非洲进行了购物") }

// 代理
type ShoppingProxy struct {
	shopping Shopping
}

func NewProxy(shopping Shopping) Shopping {
	return &ShoppingProxy{
		shopping: shopping,
	}
}

func (op *ShoppingProxy) Buy(goods *Goods) {
	// 1、辨别真伪
	if op.distinguish(goods) {
		op.shopping.Buy(goods)
		op.check(goods)
	}
	// 2、调用具体要被代理的购物方式
	// 3、海关安检

}

func (op *ShoppingProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了检查")
	if !goods.Fact {
		fmt.Println("发现假货,", goods.Kind, "不应该购买")
	}
	return goods.Fact
}

func (op *ShoppingProxy) check(goods *Goods) {
	fmt.Println("对商品进行了安检...", goods.Kind)
}
func main() {

}
