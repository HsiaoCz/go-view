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

go是强类型的语言，编译时对每个变量的类型信息做强校验，所以每个类型的元信息需要使用一个结构体描述。
go的反射也是基于类型的元信息实现的，这里_type就是所有类型最原始的元信息

```go
type _type struct {
	size       uintptr // 类型占用内存大小
	ptrdata    uintptr // 包含所有指针的内存前缀大小
	hash       uint32  // 类型 hash
	tflag      tflag   // 标记位，主要用于反射
	align      uint8   // 对齐字节信息
	fieldAlign uint8   // 当前结构字段的对齐字节数
	kind       uint8   // 基础类型枚举值
	equal func(unsafe.Pointer, unsafe.Pointer) bool // 比较两个形参对应对象的类型是否相等
	gcdata    *byte    // GC 类型的数据
	str       nameOff  // 类型名称字符串在二进制文件段中的偏移量
	ptrToThis typeOff  // 类型元信息指针在二进制文件段中的偏移量
}
```
因为 Go 语言中函数方法是以包为单位隔离的。所以 interfacetype 除了保存 _type 还需要保存包路径等描述信息。mhdr 存的是各个 interface 函数方法在段内的偏移值 offset，知道偏移值以后才方便调用。

kind：描述如何解析基础类型
在go语言中，基础类型是一个枚举常量，有26个基础类型

```go
const (
	kindBool = 1 + iota
	kindInt
	kindInt8
	kindInt16
	kindInt32
	kindInt64
	kindUint
	kindUint8
	kindUint16
	kindUint32
	kindUint64
	kindUintptr
	kindFloat32
	kindFloat64
	kindComplex64
	kindComplex128
	kindArray
	kindChan
	kindFunc
	kindInterface
	kindMap
	kindPtr
	kindSlice
	kindString
	kindStruct
	kindUnsafePointer

	kindDirectIface = 1 << 5
	kindGCProg      = 1 << 6
	kindMask        = (1 << 5) - 1
)
```

### 2、空interface数据结构

空interface{}是没有方法集的接口，所有不需要itab数据结构。
它只需要存类型对应的值即可

```go
type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```

从这个数据结构可以看出，只有当 2 个字段都为 nil，空接口才为 nil。空接口的主要目的有 2 个，一是实现“泛型”，二是使用反射。