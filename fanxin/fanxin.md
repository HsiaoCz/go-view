## go泛型

在理解泛型之前，我们来看一个函数

```go
func add(a int,b int)int{return a+b}
```

这个函数可以计算两个int类型的值，
但这个函数无法计算两个float64类型的值
在这如果我们使用interface{}类型，那么我们就需要使用类型断言来获取到具体的值的类型
这样做代码是很脆弱的

一种方法是再添加一个计算float类型的函数

```go
func add(a float64,b float64){return a+b}
```

当我们需要计算多种类型的值的时候，我们需要写多个函数来实现计算

函数在定义时候参数叫作形参，函数调用时候传入的参数叫实参

如果我们把形参实参推广一下，来一套类型形参和类型实参

```go
func add(a T,b T)T{
    return a+b
}
```

这里T就是类型形参，当我们调用的时候传入具体的类型实参,类型形参，它不是具体的类型，在定义函数时类型并不确定。因为T类型并不确定

```go
add[T=int](100,200)

// 传入这种类型实参之后，可以将Add()函数看成这种
func Add(a int,b int)int{return a+b}

// 当我们需要string时
add[T=string]("hello","number")
```

泛型：如果经常要分别为不同的数据类型写相同的函数，那么使用泛型很好

### go的泛型

go的泛型里面有很多概念:

类型形参 (Type parameter)
类型实参(Type argument)
类型形参列表( Type parameter list)
类型约束(Type constraint)
实例化(Instantiations)
泛型类型(Generic type)
泛型接收器(Generic receiver)
泛型函数(Generic function)

**类型形参、类型实参、类型约束和泛型类型**

```go
type IntSlice []int

var a InSlice=[]int{1,2,3}  //这种是正确的，IntSlice的底层是int类型的

var b IntSlice=[]string{"1","2","3"}  // 这种就是不正确的，IntSlice底层是int类型的，想要这种的Slice就需要新定义一个slice

type StrSlice []string

```

我们使用泛型来实现上面的这两种slice

```go
type Slice[T int|float64|string] []T
```

在这里T就是**类型形参**，它类似一个占位符
`int|float64|string` 这一部分被称为**类型约束**
这里中括号里的`T int|float64|string` 一整套被称为**类型参数列表**
这里定义的新类型名称叫Slice[T]

我们称这种类型为**泛型**(泛型，类型定义的时候带类型形参的类型)

泛型必须传入类型实参才能使用，这种传入类型实参将其确定为具体的类型的操作叫作**实例化**

**对于map**

```go
type MyMap[KEY int|string,VALUE float32|float64] map[KEY]VALUE
```
KEY VALUE:类型形参
int|string flaot32|float64类型约束
中括号里的一整套，类型参数列表
MyMap[KEY VALUE] :泛型类型

```go
// 一个泛型类型的结构体。可用 int 或 sring 类型实例化
type MyStruct[T int | string] struct {  
    Name string
    Data T
}

// 一个泛型接口(关于泛型接口在后半部分会详细讲解）
type IPrintData[T int | float32 | string] interface {
    Print(data T)
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan[T int | string] chan T

```

一些错误的语法
```go
//✗ 错误。T *int会被编译器误认为是表达式 T乘以int，而不是int指针
type NewType[T *int] []T
// 上面代码再编译器眼中：它认为你要定义一个存放切片的数组，数组长度由 T 乘以 int 计算得到
type NewType [T * int][]T 

//✗ 错误。和上面一样，这里不光*被会认为是乘号，| 还会被认为是按位或操作
type NewType2[T *int|*float64] []T 

//✗ 错误
type NewType2 [T (int)] []T 

```

这种错误的避免方法
```go
type NewType[T interface{*int}] []T
type NewType2[T interface{*int|*float64}] []T 

// 如果类型约束中只有一个类型，可以添加个逗号消除歧义
type NewType3[T *int,] []T

//✗ 错误。如果类型约束不止一个类型，加逗号是不行的
type NewType4[T *int|*float32,] []T 

```

泛型类型的套娃

```go
// 先定义个泛型类型 Slice[T]
type Slice[T int|string|float32|float64] []T

// ✗ 错误。泛型类型Slice[T]的类型约束中不包含uint, uint8
type UintSlice[T uint|uint8] Slice[T]  

// ✓ 正确。基于泛型类型Slice[T]定义了新的泛型类型 FloatSlice[T] 。FloatSlice[T]只接受float32和float64两种类型
type FloatSlice[T float32|float64] Slice[T] 

// ✓ 正确。基于泛型类型Slice[T]定义的新泛型类型 IntAndStringSlice[T]
type IntAndStringSlice[T int|string] Slice[T]  
// ✓ 正确 基于IntAndStringSlice[T]套娃定义出的新泛型类型
type IntSlice[T int] IntAndStringSlice[T] 

// 在map中套一个泛型类型Slice[T]
type WowMap[T int|string] map[string]Slice[T]
// 在map中套Slice[T]的另一种写法
type WowMap2[T Slice[int] | Slice[string]] map[string]T

```

匿名结构体不支持泛型

### 泛型的接收者

定义一个泛型并提供一个方法

```go
type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
    var sum T
    for _, value := range s {
        sum += value
    }
    return sum
}

```

这里接收者是泛型名称，方法的返回值使用类型形参，方法的形参也可以使用类型形参

泛型的调用需要先实例化类型形参
```go
var s MySlice[int]=[]int{1,2,3,4}
sum:=s.Sum()
println(sum)
```

### 基于泛型的队列

队列是一种先入先出的数据结构，它和现实中的排队一样，数据只能从队尾放入，从队首取出，先放入的数据会被优先取出

泛型不支持类型断言 switch value.(type)

### 泛型函数

```go
func add[T int|float64|float32](a T,b T)T{
    return a+b
}

// 泛型函数必须先实例化再调用
add[int](1,2)

// 不过go语言也支持对类型形参进行类型推导
add(1,2)
```

匿名函数不止持泛型
方法不支持泛型

### 变得复杂的接口

```go
type IntUintFloat interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Slice[T IntUintFloat] []T

```

可以使用接口来定义类型约束
接口和接口，接口和普通类型之间也可以通过|进行组合

```go
type Int interface {
    int | int8 | int16 | int32 | int64
}

type Uint interface {
    uint | uint8 | uint16 | uint32
}

type Float interface {
    float32 | float64
}

type Slice[T Int | Uint | Float] []T  // 使用 '|' 将多个接口类型组合

```

接口里面也可以嵌套接口

```go
type SliceElement interface {
    Int | Uint | Float | string // 组合了三个接口类型并额外增加了一个 string 类型
}

type Slice[T SliceElement] []T 
```

还可以通过~指定底层类型
```go
type Int interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
    ~uint | ~uint8 | ~uint16 | ~uint32
}
type Float interface {
    ~float32 | ~float64
}

type Slice[T Int | Uint | Float] []T 

var s Slice[int] // 正确

type MyInt int
var s2 Slice[MyInt]  // MyInt底层类型是int，所以可以用于实例化

type MyMyInt MyInt
var s3 Slice[MyMyInt]  // 正确。MyMyInt 虽然基于 MyInt ，但底层类型也是int，所以也能用于实例化

type MyFloat32 float32  // 正确
var s4 Slice[MyFloat32]

```

`~`时有一定的限制
~后面的类型不能为接口
~后面的类型必须为基本类型

```go
type MyInt int

type _ interface {
    ~[]byte  // 正确
    ~MyInt   // 错误，~后的类型必须为基本类型
    ~error   // 错误，~后的类型不能为接口
}
```

在go1.18之前，接口定义了一套方法集，所有实现了方法集里所有的方法的类型就可以视为该接口

关于这句话也可以这样理解：
接口看成一个类型集合，所有实现了该类型集合中的方法的类型都在接口代表的类型集合当中

在go1.18之后，接口的定义发生了改变

一个接口类型定义了一个类型集

接口定义发生了改变，接口实现也发生了改变
当满足以下两个条件时，类型T实现了接口:
- T 不是接口时：类型 T 是接口 I 代表的类型集中的一个成员 (T is an element of the type set of I)
- T 是接口时： T 接口代表的类型集是 I 代表的类型集的子集(Type set of T is a subset of the type set of I)

### 类型的交并集

```go
type AllInt interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint32
}

type Uint interface {
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type A interface { // 接口A代表的类型集是 AllInt 和 Uint 的交集
    AllInt
    Uint
}

type B interface { // 接口B代表的类型集是 AllInt 和 ~int 的交集
    AllInt
    ~int
}

```

下面这种也代表交集

```go
type C interface {
    ~int
    int
}
```

**空集**

当多个类型的交集为空的时候，这个接口代表的类型集为一个空集
```go
type Bad interface{
    int
    flaot32
}
```

没有任何一种类型属于空集，空集没有任何意义

**空接口**

空接口代表了所有类型的集合

虽然空接口没有写入任何类型，但是它代表的是所有类型的集合，而非一个空集
类型约束中指定空接口的意思是指定一个包含所有类型的类型集，并不是类型约束限定了只能使用空接口来做类型形参

**可比较的类型和可排序的类型**

类似于map它的key必须要是可比较的，也就是可以`!=`和`==`

go内置了一个接口`comparable`接口，代表所有课比较的类型(!=和==)

```go
type Mymap[KEY comparable,VALUE any] map[KEY]VALUE
```

可以比较的类型不代表可以使用`<、>、>=、<=`比较，可以这样比较的类型叫作可排序的类型(ordered)


**接口的两种类型**

假如有如下一个接口
```go
type ReadWriter interface {
    ~string | ~[]rune

    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}

```

这个接口的意思是:
接口类型 ReadWriter 代表了一个类型集合，所有以 string 或 []rune 为底层类型，并且实现了 Read() Write() 这两个方法的类型都在 ReadWriter 代表的类型集当中

```go
// 类型 StringReadWriter 实现了接口 Readwriter
type StringReadWriter string 

func (s StringReadWriter) Read(p []byte) (n int, err error) {
    // ...
}

func (s StringReadWriter) Write(p []byte) (n int, err error) {
 // ...
}

//  类型BytesReadWriter 没有实现接口 Readwriter
type BytesReadWriter []byte 

func (s BytesReadWriter) Read(p []byte) (n int, err error) {
 ...
}

func (s BytesReadWriter) Write(p []byte) (n int, err error) {
 ...
}

```

go1.18之后，接口被分为了两种类型

- 基本接口
- 一般接口

基本接口：
接口中只有方法，那么这种接口称为基本接口
这就是我们平常使用的最多的那种接口

一般接口：
接口中不仅有方法，还有类型的接口，称为一般接口
或者接口里面只有类型的接口被称为一般接口

```go
type Uint interface { // 接口 Uint 中有类型，所以是一般接口
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ReadWriter interface {  // ReadWriter 接口既有方法也有类型，所以是一般接口
    ~string | ~[]rune

    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
}
```

一般接口类型不能用来定义变量，只能用于泛型的类型约束

**泛型接口**

接口也可以泛型化

```go
type DataProcessor[T any] interface {
    Process(oriData T) (newData T)
    Save(data T) error
}

type DataProcessor2[T any] interface {
    int | ~struct{ Data interface{} }

    Process(data T) (newData T)
    Save(data T) error
}

```

接口引入了类型形参，这两个接口时泛型接口
泛型接口需要实例化之后才能使用

### 接口的种种限制

1、用|连接多个类型的时候，类型之间不能有相交的部分(即必须是不交集)

```go
type MyInt int

// 错误，MyInt的底层类型是int,和 ~int 有相交的部分
type _ interface {
    ~int | MyInt
}

// 相交地类型是接口地化，不受这一限制
type MyInt int

type _ interface {
    ~int | interface{ MyInt }  // 正确
}

type _ interface {
    interface{ ~int } | MyInt // 也正确
}

type _ interface {
    interface{ ~int } | interface{ MyInt }  // 也正确
}

```

2、类型的并集不能有类型形参

```go
type MyInf[T ~int | ~string] interface {
    ~float32 | T  // 错误。T是类型形参
}

type MyInf2[T ~int | ~string] interface {
    T  // 错误
}
```
3、接口不能直接或间接地并入自己

```go
type Bad interface {
    Bad // 错误，接口不能直接并入自己
}

type Bad2 interface {
    Bad1
}
type Bad1 interface {
    Bad2 // 错误，接口Bad1通过Bad2间接并入了自己
}

type Bad3 interface {
    ~int | ~string | Bad3 // 错误，通过类型的并集并入了自己
}

```
4、接口的并集成员个数大于一的时候，不能直接或间接地并入comparable接口

```go
type OK interface {
    comparable // 正确。只有一个类型的时候可以使用 comparable
}

type Bad1 interface {
    []int | comparable // 错误，类型并集不能直接并入 comparable 接口
}

type CmpInf interface {
    comparable
}
type Bad2 interface {
    chan int | CmpInf  // 错误，类型并集通过 CmpInf 间接并入了comparable
}
type Bad3 interface {
    chan int | interface{comparable}  // 理所当然，这样也是不行的
}

```
5、带方法地接口（无论是基本接口还是一般接口），都不能写入接口地并集中

```go
type _ interface {
    ~int | ~string | error // 错误，error是带方法的接口(一般接口) 不能写入并集中
}

type DataProcessor[T any] interface {
    ~string | ~[]byte

    Process(data T) (newData T)
    Save(data T) error
}

// 错误，实例化之后的 DataProcessor[string] 是带方法的一般接口，不能写入类型并集
type _ interface {
    ~int | ~string | DataProcessor[string] 
}

type Bad[T any] interface {
    ~int | ~string | DataProcessor[T]  // 也不行
}
```