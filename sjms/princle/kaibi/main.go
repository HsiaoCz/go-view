package main

import "fmt"

// 开闭原则：对类的改动是通过添加代码来实现的，而不是修改代码
// 抽象层 假设有一个抽象的银行职员
// 他有一个DO方法
type Baker interface {
	Dobuss()
}

// 如果这个银行职员具有这些方法：存款，转账，股票
type SaveBanker struct{}

func (s *SaveBanker) Dobuss() { fmt.Println("进行了转账的业务") }

type CunBanker struct{}

func (c *CunBanker) Dobuss() { fmt.Println("进行了存款的业务") }

type StockBanker struct{}

func (s *StockBanker) Dobuss() { fmt.Println("进行了股票的业务") }

func main() {
	banker := new(SaveBanker)
	banker.Dobuss()
}
