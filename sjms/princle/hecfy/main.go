package main

import "fmt"

// 合成复用原则
// 优先使用组合

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

// 给小猫来个睡觉的能力
// 使用继承来实现
type CatB struct {
	Cat
}

func (c *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 使用组合来实现
type CatC struct {
	cat Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}

// 只对接口依赖实现
// 这依赖很显然耦合度更低
type CatM struct{}

func (cm *CatM) Eat(c *Cat) {
	c.Eat()
}

func (cm *CatM) Sleep() {
	fmt.Println("小猫睡觉")
}
func main() {
	catb := new(CatB)
	catb.Eat()
	catb.Sleep()

	cc := new(CatC)
	cc.cat.Eat()
	cc.Sleep()

	cm := new(CatM)
	cm.Eat(new(Cat))
	cm.Sleep()
}
