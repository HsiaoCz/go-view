## go的面试题

这里记录一些不太知道的

**1.如何高效的拼接字符串**

1.使用`+`号进行拼接

使用+操作符进行拼接时，会对字符串进行遍历，计算并开辟一个新的空间来存储原来的两个字符串。


2.使用fmt.Sprintf

由于采用了接口参数，必须要用反射获取值，因此有性能损耗

3.strings.Builder
用WriteString()进行拼接，内部实现是指针+切片，同时String()返回拼接后的字符串，它是直接把[]byte转换为string，从而避免变量拷贝。

4.bytes.Buffer
bytes.Buffer是一个缓冲byte类型的缓冲器，这个缓冲器里存放着都是byte，

bytes.buffer底层也是一个[]byte切片

5.strings.join
strings.join也是基于strings.builder来实现的,并且可以自定义分隔符，在join方法内调用了b.Grow(n)方法，这个是进行初步的容量分配，而前面计算的n的长度就是我们要拼接的slice的长度，因为我们传入切片长度固定，所以提前进行容量分配可以减少内存分配，很高效。

性能比较：

strings.Join ≈ strings.Builder > bytes.Buffer > "+" > fmt.Sprintf

拼接字符串的代码：
```go
a:=[]string{"a","b","c"}

//方式1
ret:=a[0]+a[1]+a[2]

// 方式2
ret:=fmt.Sprintf("%s%s%s",a[0],a[1],a[2])

// 方式3
var sb strings.Builder
sb.WriteString(a[0])
sb.WriteString(a[1])
sb.WriteString(a[2])

ret:=sb.String()

//方式4
buf :new(bytes.Buffer)
buf.Write(a[0])
buf.Write(a[1])
buf.Write(a[2])
ret:=buf.String()

// 方式5
ret:=strings.join(a,"")
```

**2.关于defer**
[https://www.cnblogs.com/traditional/p/11440728.html]

**3.如何获取一个tag**

```go
import reflect

type Author struct{
    Name int  `json:"name"`
    Publication []string `json:"publication,omitempty"`
}

func main(){
    t:=reflect.TypeOf(Author{})
    for i:=0;i<t.NumField();i++{
        name:=t.Field(i).Name
        s,_:=t.FieldByName(name)
        fmt.Println(name,s.Tag)
    }
}
```

reflect.TypeOf方法获取对象的类型，之后NumField()获取结构体成员的数量。 通过Field(i)获取第i个成员的名字。 再通过其Tag 方法获得标签。

**4.空struct{}的用途**

1.使用map来模拟set，把map的值置为struct{},可以减少内存分配的

```go
type Set map[string]struct{}

func main() {
	set := make(Set)

	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}
	fmt.Println(len(set)) // 3
	if _, ok := set["A"]; ok {
		fmt.Println("A exists") // A exists
	}
}
```

2.给通道发送一个空的结构体
3.仅有方法的结构体

**5.如何知道一个对象是分配在栈上还是堆上**

go的局部变量会进行逃逸分析，如果变量在离开作用域后没有被引用，那么会被分配到栈上，否则会被分配到堆上

```go
go build -gcflags "-m -m -l" xxx.go
```

**6.2个interface可以比较吗?**

interface内部维护了两个字段,类型T和值V,interface可以进行比较，这两种情况下interface会相等

1.两个interface均等于nil(此时V和T都处于unset状态)
2.类型T相同，且对应的值V相等

**7.两个nil可能不相等吗？**

可能不相等，interface在运行时绑定值，只有值为nil的接口才为nil,但与指针的nil并不相等，两个nil只有在类型相同时才相等

**8.go的GC工作原理**

**9.函数的局部变量的指针是否安全**

go的局部变量的指针是安全的。因为go会进行逃逸分析，如果发现局部变量的作用域超过该函数则会把指针分配到堆区，避免内存泄漏

**10.非接口的任意类型T()可以调用*T的方法吗？反过来呢？**

一个T类型的值可以调用*T类型声明的方法，当且仅当T是可寻址的。

反之：*T 可以调用T()的方法，因为指针可以解引用。

**11.go切片的扩容问题**

go在1.18之前，当切片的容量小于1024时，先判断所需容量是否是大于容量的二倍，是则使用当前容量加上所需容量，否则2倍扩容
如果超过1024,每次按1.25倍扩容

但在1.18之后，这种扩容策略发生了改变

**12.无缓冲的channel和有缓冲的channel的区别**

对于无缓冲的channel：
如果发送的数据没有被接收方接受，那么发送方会阻塞，如果接受方一直接收不到发送方的数据，那么接收方会阻塞

有缓冲的channel：
发送方在缓冲区满的时候阻塞，接受方不阻塞，接受方在缓冲区为空的时候阻塞，发送方不阻塞

**13.为什么有协程泄露**