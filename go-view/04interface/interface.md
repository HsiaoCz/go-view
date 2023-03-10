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


#### 3、接口的动态类型和动态值

iface包含两个字段:tab是接口表指针，指向类型信息；data是数据指针，指向具体的数据
它们分别被称为动态类型和动态值。而接口值包括动态类型和动态值

接口值的零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为 nil 的情况下，这个接口值就才会被认为 接口值 == nil。

例题：

```go
package main

import "fmt"

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

func main() {
	var c Coder
	fmt.Println(c == nil)  //这里会输出true 这里接口的动态值和动态类型都为nil
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	fmt.Println(g == nil) //这里也会输出nil  因为结构体没有初始化，为空

	c = g
	fmt.Println(c == nil) // 这里会输出false，因为当把g赋值给c后，c的动态类型就变成了*main.Gopher，即便c的动态值仍然为nil，它也不是nil了
	fmt.Printf("c: %T, %v\n", c, c)
}
```

再看一个例子：

```go
package main

import "fmt"

type MyError struct {}

func (i MyError) Error() string {
	return "MyError"
}

func main() {
	err := Process()
	fmt.Println(err)  //nil

	fmt.Println(err == nil) //false
}

func Process() error {
	var err *MyError = nil
	return err
}
```

如何打印出接口的动态类型和动态值？

```go
package main

import (
	"unsafe"
	"fmt"
)

type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)
	
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}
```

#### 4、编译器自动检测类型是否实现接口

```go
var _ io.Writer=(*myWriter)(nil)
```

这么写的目的是为了让编译器自动检查*myWriter类型是否实现了io.Writer接口

```go
package main

import "io"

type myWriter struct {

}

/*func (w myWriter) Write(p []byte) (n int, err error) {
	return
}*/

func main() {
    // 检查 *myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = (*myWriter)(nil)

    // 检查 myWriter 类型是否实现了 io.Writer 接口
    var _ io.Writer = myWriter{}
}
```

解除注释后，运行程序不报错。

实际上，上述赋值语句会发生隐式地类型转换，在转换的过程中，编译器会检测等号右边的类型是否实现了等号左边接口所规定的函数。

#### 5、接口的构造过程

```go
package main

import "fmt"

type Person interface {
	growUp()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

func main() {
	var qcrao = Person(Student{age: 18})

	fmt.Println(qcrao)
}
```

接口的构造过程：


#### 6、类型断言和类型转换

fmt.Println 函数的参数是 interface。对于内置类型，函数内部会用穷举法，得出它的真实类型，然后转换为字符串打印。而对于自定义类型，首先确定该类型是否实现了 String() 方法，如果实现了，则直接打印输出 String() 方法的结果；否则，会通过反射来遍历对象的成员进行打印。

```go
package main

import "fmt"

type Student struct {
	Name string
	Age int
}

func main() {
	var s = Student{
		Name: "qcrao",
		Age: 18,
	}

	fmt.Println(s)
}
```

因为 Student 结构体没有实现 String() 方法，所以 fmt.Println 会利用反射挨个打印成员变量：{qcrao 18}

增加一个String()方法的实现:

```go
func (s Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}
```

输出:[Name: qcrao], [Age: 18]

这里，如果将这个方法改为

```go
func (s *Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}
```

打印的结果会是:{qcrao 18}

这里之所以会出现这样的结果，
因为对于任意类型T只有接受者是T的方法，而类型*T拥有接受者是T和*T的方法。
语法上T能直接调*T的方法仅仅是GO的语法糖

所以， Student 结构体定义了接受者类型是值类型的 String() 方法时，通过

```go
fmt.Println(s)
fmt.Println(&s)
```

均可以按照自定义的格式来打印。

如果 Student 结构体定义了接受者类型是指针类型的 String() 方法时，只有通过

```go
fmt.Println(&s)
```
才能按照自定义的格式打印。

#### 7、接口转换的原理

通过前面提到的 iface 的源码可以看到，实际上它包含接口的类型 interfacetype 和 实体类型的类型 _type，这两者都是 iface 的字段 itab 的成员。也就是说生成一个 itab 同时需要接口的类型和实体的类型。

当判定一种类型是否满足某个接口时，Go 使用类型的方法集和接口所需要的方法集进行匹配，如果类型的方法集完全包含接口的方法集，则可认为该类型实现了该接口。

例如某类型有 m 个方法，某接口有 n 个方法，则很容易知道这种判定的时间复杂度为 O(mn)，Go 会对方法集的函数按照函数名的字典序进行排序，所以实际的时间复杂度为 O(m+n)。

当把实体类型赋值给接口的时候，会调用 conv 系列函数，例如空接口调用 convT2E 系列、非空接口调用 convT2I 系列。这些函数比较相似：

- 具体类型转空接口时，_type 字段直接复制源类型的 _type；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。

- 具体类型转非空接口时，入参 tab 是编译器在编译阶段预先生成好的，新接口 tab 字段直接指向入参 tab 指向的 itab；调用 mallocgc 获得一块新内存，把值复制进去，data 再指向这块新内存。

- 而对于接口转接口，itab 调用 getitab 函数获取。只用生成一次，之后直接从 hash 表中获取。