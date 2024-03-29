package main

import "fmt"

// 模板方法模式
// 比如制作一杯咖啡，需要煮水，冲咖啡，倒入杯中，加料
// 用在流程相同，但是具体的内容不同的不一样的事情上

// 抽象类，制作饮料
type Beverage interface {
	// 煮开水
	BoilWater()

	// 冲泡
	Brew()

	// 倒入杯中
	PourInCup()

	//添加佐料
	Addthing()
}

// 封装一套流程
type template struct {
	b Beverage
}

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	// 固定的流程
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	t.b.Addthing()
}

// 具体的制作饮料的流程
type MakeConffee struct {
	template //继承模板
}

func NewMakeCoffee() *MakeConffee {
	makeCoffee := new(MakeConffee)
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeConffee) BoilWater() { fmt.Println("煮开水....") }
func (mc *MakeConffee) Brew()      { fmt.Println("冲泡咖啡..") }
func (mc *MakeConffee) PourInCup() { fmt.Println("导入杯中..") }
func (mc *MakeConffee) Addthing()  { fmt.Println("加点糖..") }

func main() {
	// 制作咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.template.MakeBeverage()
}
