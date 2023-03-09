package main

import (
	"fmt"
	"sync/atomic"
)

// 单例的懒汉式
// 之前的单例属于饿汗式，不管有没有使用，它本身就已经初始化了

// 单例：保证一个类永远只能有一个实例，这个对象还能被系统的其他模块使用

// 1、保证类是非公有化的，这样这个类就不能在其它包创建实例
type singleton struct{}

// 2、保证这个类只有一个实例，所以提供一个实例
// 由于要保证这个类只有这么一个实例，不能改变方向，所以这个实例应该是非公有的
var instance *singleton

// 只有在第一次调用的时候才初始化
// 我们可以使用原子操作来保证并发安全
var inintialized uint32

// 这个单例只能被读，不能被赋值
// 3、单例需要对外提供这个实例，所以需要一个函数，返回这个实例的指针
// 同时这个函数需要是公有的
func GetSingleton() *singleton {
	if atomic.LoadUint32(&inintialized) == 1 {
		return instance
	}
	if instance == nil {
		instance = new(singleton)
		atomic.AddUint32(&inintialized, 1)
	}
	return instance
}

func (s *singleton) Something() { fmt.Println("hello") }

func main() {
	ins := GetSingleton()
	ins.Something()
}
