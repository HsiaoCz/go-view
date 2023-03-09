package main

import "fmt"

// 装饰器模式
// 抽象层

type Phone interface {
	Show()
}

// 抽象的装饰器
type ShiPei struct {
	phone Phone
}

func (s *ShiPei) Show() {}

// 实现类，需要装饰的类
type XiaoMi struct{}

func (x *XiaoMi) Show() { fmt.Println("秀出了小米手机") }

type HuaWei struct{}

func (h *HuaWei) Show() { fmt.Println("秀出了华为手机") }

// 装饰器
type JuTi struct {
	shipei ShiPei
}

func (j *JuTi) Show() {
	j.shipei.Show() // 被装饰的构建
	fmt.Println("这是贴膜的手机")
}

func NewJuTi(phone Phone) *JuTi {
	return &JuTi{shipei: ShiPei{phone: phone}}
}

type KeDectory struct {
	shipei ShiPei
}

func (k *KeDectory) Show() {
	k.shipei.Show() // 被装饰的构建
	fmt.Println("这是戴壳的手机")
}

func NewKeDecorator(phone Phone) *KeDectory {
	return &KeDectory{shipei: ShiPei{phone: phone}}
}

func main() {
	huawei := new(HuaWei)
	huawei.Show()

	ju := NewJuTi(huawei)
	ju.Show()
}
