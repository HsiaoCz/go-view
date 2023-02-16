package main

import (
	"fmt"
)

// defer 延迟调用函数
// 多个defer的执行顺序:后进先出
// defer执行顺序 先注册，后执行
// defer 即使函数执行出错，也会执行
// defer 在panic之后注册不行

// go 语言return之前会做一个操作，将返回值赋值，然后执行RET执行令
// defer 的执行时机在这之间

// 所以如果这时候修改返回值的值，如果是值传递，没用，但是如果是引用传递就可以修改
func HiDefer() {
	// defer recover
	defer fmt.Println("Hello")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("ccc")
}

func HelloDefer() {
	defer fmt.Println("这是defer----1")
	defer fmt.Println("这是defer----2")
	defer fmt.Println("这是defer----3")
}

func main() {
	HelloDefer()
	HiDefer()
	s := add()
	fmt.Println(s)
}

func add() (x int) {
	defer func(x int) {
		x++
		fmt.Println(x)
	}(x)
	return 5
}
