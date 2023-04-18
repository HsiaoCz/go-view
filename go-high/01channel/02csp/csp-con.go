package main

import "fmt"

// csp
type Task struct {
	A int
	B int
}

var ch = make(chan Task, 100)

func Producer() {
	for i := 0; i < 10; i++ {
		ch <- Task{A: i + 3, B: i - 8}
	}
}

func Consumer() {
	for i := 0; i < 10; i++ {
		task := <-ch
		sum := task.A + task.B
		fmt.Println(sum)
	}
}

func main() {
	go Producer() // 生产者协程
	go Consumer() // 消费者协程
}
