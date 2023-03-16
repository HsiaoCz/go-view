package main

import "fmt"

// 策略模式
// 武器策略（抽象的策略)

type WeaponStartegy interface {
	UseWeapon()
}

// 具体的策略
type Ak47 struct{}

func (a *Ak47) UseWeapon() { fmt.Println("使用Ak47去战斗...") }

type Knife struct{}

func (kn *Knife) UseWeapon() { fmt.Println("使用小刀去战斗") }

// 环境类
type Hero struct {
	strategy WeaponStartegy // 拥有一个抽象的策略
}

// 设置一个策略的方法
func (h *Hero) SetWeaponStrategy(s WeaponStartegy) {
	h.strategy = s
}
// 战斗方法
func (h *Hero) Fight() {
	h.strategy.UseWeapon() // 调用具体的策略
}
func main() {
	hero := Hero{}

	// 更换策略
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	// 更换策略2
	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}