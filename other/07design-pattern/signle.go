package main

import "fmt"

// 单例，一个类只能有一个实例
// 它必须自行创建这个实例
// 它必须自行向外界提供这个实例

type mysignle struct{}

var since = new(mysignle)

func NewSignle() *mysignle {
	return since
}

func (m *mysignle) show() {
	fmt.Println("hello")
}

func main() {
	since := NewSignle()
	since.show()
}

// 这种单例在程序还没执行之前就分配内存了
// 这叫饿汉式单例