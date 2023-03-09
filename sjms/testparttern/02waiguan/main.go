package main

import "fmt"

// waiguan parttern

type SubsystemA struct{}

func (s *SubsystemA) MethodA() { fmt.Println("系统A的方法A") }

type SubsystemB struct{}

func (s *SubsystemB) MethodB() { fmt.Println("系统B的方法B") }

type SubsystemC struct{}

func (s *SubsystemC) MethodC() { fmt.Println("系统C的方法C") }

// 提供一个外观模式
// 它拥有这些类
type WaiGuan struct {
	systemA *SubsystemA
	systemB *SubsystemB
	systemC *SubsystemC
}

func (w *WaiGuan) MethodOne() {
	w.systemA.MethodA()
	w.systemB.MethodB()
}

func (w *WaiGuan) MethodTwo() {
	w.systemA.MethodA()
	w.systemC.MethodC()
}

func main() {
	wg := new(WaiGuan)
	wg.MethodOne()
	wg.MethodTwo()
}

// 所谓的外观模式，比如说有电视机的遥控器，空调遥控器，冰箱遥控器等等
// 提供一个外观类，拥有这些类的，可以调用这些类的方法，当有一种模式需要调用几种