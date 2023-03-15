package main

import "fmt"

// 适配器模式

// 适配的目标
type V5 interface {
	Use5V()
}

// 被适配的角色，适配者

type V220 struct{}

func (v *V220) Use220() { fmt.Println("使用220v进行充电...") }

// 适配器的类
type Adapter struct {
	v220 *V220
}

func NewAdopter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	a.v220.Use220()
}

// 具体的业务类
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v: v}
}

func (p *Phone) Charge() {
	fmt.Println("对phone 进行充电...")
	p.v.Use5V()
}
