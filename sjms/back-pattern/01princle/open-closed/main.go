package main

import "fmt"

// 开闭原则
// 对类的改动通过添加代码而不是修改代码来实现
// 比如有一个银行职员类
// 他有存款，转账，股票等方法

type Banker interface{
	Dobuss()
}

// 针对具体的业务创造具体的类
type SaveBanker struct{}

func (s *SaveBanker)Dobuss(){fmt.Println("进行了转账业务...")}
type CunBanker struct{}

func (c *CunBanker)Dobuss(){fmt.Println("进行了存款业务...")}

type StockBanker struct{}

func (s *StockBanker)Dobuss(){fmt.Println("进行了股票业务...")}

func main(){
	banker:=new(SaveBanker)
	banker.Dobuss()
}