package main

import "fmt"

var ch = make(chan bool, 1)

// 使用channel 来控制并发执行的顺序

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello my man")
		<-ch
	}
}

func baz() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello my woman")
	}
}

func main() {
	select {
	case ch <- true:
		foo()
	case <-ch:
		baz()
	}
}
