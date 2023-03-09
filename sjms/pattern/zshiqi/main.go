package main

import "fmt"

// 抽象层
type Phone interface {
	Show()
}

// 抽象的装饰器类，装饰器的基础类
type Decorator struct {
	phone Phone //组合进来的Phone
}

func (d *Decorator) Show() {
}

// 实现类

type HuaWei struct{}

func (h *HuaWei) Show() { fmt.Println("秀出了华为手机") }

type XiaoMi struct{}

func (x *XiaoMi) Show() { fmt.Println("秀出了小米的手机") }

// 具体的装饰器

type MoDectory struct {
	Decorator // 继承基础的装饰器类
}

func (md *MoDectory) Show() {
	md.phone.Show() // 调用被装饰的构建
	fmt.Println("贴膜的手机")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDectory{Decorator{phone: phone}}
}

type KeDectory struct {
	Decorator // 继承自基础的类
}

func (kd *KeDectory) Show() {
	kd.phone.Show() // 调用被装饰的构建
	fmt.Println("带壳的手机")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDectory{Decorator{phone: phone}}
}

// 业务逻辑层
func main() {
	huawei := new(HuaWei)
	huawei.Show()
	// 装饰器，编程贴膜的华为手机
	modectory := NewMoDecorator(huawei)
	modectory.Show()
}
