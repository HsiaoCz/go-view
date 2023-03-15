package main

import "fmt"

// 外观模式

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
	systenB *SubsystemB
	systemC *SubsystemC
}

func (w *WaiGuan) MethodOne() {
	w.systemA.MethodA()
	w.systenB.MethodB()
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
