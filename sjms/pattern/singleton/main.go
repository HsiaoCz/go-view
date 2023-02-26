package main

import "fmt"

// 单例模式，保证一个类只有一个实例
// 它必须能够自行创建这个实例
// 它必须自行向整个系统提供这个实例

// 单例：保证一个类永远只能有一个实例，这个对象还能被系统的其他模块使用

// 1、保证类是非公有化的，这样这个类就不能在其它包创建实例
type singleton struct{}

// 2、保证这个类只有一个实例，所以提供一个实例
// 由于要保证这个类只有这么一个实例，不能改变方向，所以这个实例应该是非公有的
var instance = new(singleton)

// 这个单例只能被读，不能被赋值
// 3、单例需要对外提供这个实例，所以需要一个函数，返回这个实例的指针
// 同时这个函数需要是公有的
func GetSingleton() *singleton {
	return instance
}

func (s *singleton) Something() { fmt.Println("hello") }

func main() {
	ins := GetSingleton()
	ins.Something()
}

// 这种单例 叫作饿汉式的单例
// 无论执不执行，这个单例都分配了内存