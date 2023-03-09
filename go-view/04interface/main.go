package main

import "fmt"

// 这里关于值接受者和指针接受者的问题

type Hello interface {
	SayHello()
}

type HelloMan struct{}

func (s *HelloMan) SayHello() {
	fmt.Println("hello my man")
}

func SayHelloByHello(h Hello) {
	h.SayHello()
}

type HelloWoman struct{}

func (h HelloWoman) SayHello() { fmt.Println("hello my woman") }

func main() {
	// 使用指针接受者声明的方法
	// 他的值对象没有实现接口

	// 但是使用值接受者声明的方法，他的指针对象实现了接口

	hman := HelloMan{}
	SayHelloByHello(&hman)

	hwoman := &HelloWoman{}
	SayHelloByHello(hwoman)
}
