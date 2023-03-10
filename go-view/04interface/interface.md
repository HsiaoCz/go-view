### go interface 详解

接口是高级语言中的一个规约，是一组方法签名的集合。GO的interface是非侵入式的

所谓的侵入式和非侵入式指的是：不需要显式的声明实现了哪个接口
只需要实现接口中的所有方法，就认为是实现了该接口

#### 1、接口的数据结构

1、非空interface数据结构

非空interface初始化的底层数据结构是iface

```go
type iface  struct{
    tab *itab
    data unsafe.Pointer
}
```
其中tab存放的是类型、方法等信息。
data指针指向的iface绑定对象的原始数据的副本。
tab是itab类型的指针

```go
type itab struct {
    // interface自己的静态类型
	inter *interfacetype
    // interface对应具体对象的类型
	_type *_type

	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

itab 中的 _type 和 iface 中的 data 能简要描述一个变量。
_type 是这个变量对应的类型，data 是这个变量的值。
这里的 hash 字段和 _type 中存的 hash 字段是完全一致的，这么做的目的是为了类型断言(下文会提到)。
fun 是一个函数指针，它指向的是具体类型的函数方法。
虽然这里只有一个函数指针，但是它可以调用很多方法。
在这个指针对应内存地址的后面依次存储了多个方法，利用指针偏移便可以找到它们