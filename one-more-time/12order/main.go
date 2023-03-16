package main

import (
	"fmt"
)

// order 命令模式
// 有命令的收集者，命令的执行者
// 命令的收集者收集命令，命令的执行者执行命令

type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("烤串师傅烤了羊肉串...") }
func (c *Cooker) MakeJiChi()   { fmt.Println("烤串师傅烤了鸡翅...") }

// 抽象的命令
type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeJiChi()
}

type CommandCookYang struct {
	cooker *Cooker
}

func (cmd *CommandCookYang) Make() {
	cmd.cooker.MakeChuaner()
}

type WaiterMM struct {
	CmdList []Command
}

func (w WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}
func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker: cooker}
	cmdYang := CommandCookYang{cooker: cooker}

	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &cmdChicken)
	mm.CmdList = append(mm.CmdList, &cmdYang)

	mm.Notify()

}
