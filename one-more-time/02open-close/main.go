package main

import "fmt"

// 开闭原则
// 所谓的开闭原则
// 对类的修改是通过增加代码来实现的，而不是通过修改代码

// 抽象的banker

type Banker interface {
	DoBus()
}

// 具体的banker
// 比如这有一个存款的业务
// 那我们就创建一个具体的存款的银行职员
type SaveBanker struct{}

func (s *SaveBanker) DoBus() { fmt.Println("进行了存款业务...") }

// 那么现在有一个转账的业务
type TransBanker struct{}

func (t *TransBanker) DoBus() { fmt.Println("进行了转账的业务...") }

// 现在银行职员还有一个股票的业务
type StackBanker struct{}

func (s *StackBanker) DoBus() { fmt.Println("进行了股票的业务") }

// 中间层
type MiddleMan struct {
	banker Banker
}

func NewMiddleMan(banker Banker) *MiddleMan {
	return &MiddleMan{
		banker: banker,
	}
}

func (m *MiddleMan) DoBus() {
	m.banker.DoBus()
}

// 实现层
func main() {
	saveBanker := &SaveBanker{}
	saveBanker.DoBus()

	middleMan := NewMiddleMan(saveBanker)
	middleMan.DoBus()
}
