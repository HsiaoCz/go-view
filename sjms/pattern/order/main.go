package main

import "fmt"

// 命令模式：假如在路边烤串，有烤羊肉，烤鸡翅，有烤串师傅，和服务员
// 还有服务员

type Cooker struct{}

func (c *Cooker) MakeChuaner() { fmt.Println("烤串师傅烤了羊肉串") }
func (c *Cooker) MakeJichi()   { fmt.Println("烤串师傅烤了鸡翅") }

// 抽象的命令
type Command interface {
	Make()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeJichi()
}

type CommandCookYang struct {
	cooker *Cooker
}

func (cmd *CommandCookYang) Make() {
	cmd.cooker.MakeChuaner()
}

// 服务员，命令收集者
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
