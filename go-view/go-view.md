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

协程泄露事指协程创建之后没有得到释放，主要原因有：

1、缺少接收器，导致发送阻塞
2、缺少发送器，导致接收阻塞
3、死锁，多个线程由于竞争资源导致死锁
4、创建的协程没有回收

**14、可以限制操作系统的线程数量吗？常见的goroutine操作函数有哪些？**

使用runtime.GOMAXPROCS(num int)可以设置线程数目，该值默认为CPU的逻辑核心数，如果设的太大，会引起频繁的线程切换，降低性能


runtime.Gosched()，用于出让时间片，让出当前goroutine的执行权限，调度器安排其他的等待任务运行，并在下次的某个时候从该位置恢复执行。
runtime.Goexit()，调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。请注意千万别在主函数调用runtime.Goexit，因为会引发panic。

**15、如何控制协程的数目**

从官方文档的解释可以看到，GOMAXPROCS 限制的是同时执行用户态 Go 代码的操作系统线程的数量，但是对于被系统调用阻塞的线程数量是没有限制的。GOMAXPROCS 的默认值等于 CPU 的逻辑核数，同一时间，一个核只能绑定一个线程，然后运行被调度的协程。因此对于 CPU 密集型的任务，若该值过大，例如设置为 CPU 逻辑核数的 2 倍，会增加线程切换的开销，降低性能。对于 I/O 密集型应用，适当地调大该值，可以提高 I/O 吞吐率

对于协程，可以使用带缓冲的chennel来控制，下面的例子是协程数为1024的例子
```go
var wg sync.WaitGroup

ch :=make(chan struct{},1024)

for i:=o;i<20000;i++{
    wg.Add(1)
    ch<-struct{}{}
    go func(){
        defer wg.Done()
        <-ch
    }
}
wg.Wait()
```

此外，还可以使用线程池

**16、关于new 和 make 的区别?**

**17、uint型变量值分别为1,2,它们相减的结果为多少?**

结果会溢出，如果是32位系统，结果是2^32-1，如果是64位系统，结果2^64-1.

**18、go的内存管理**

**19、mutex有几种模式**

两种模式：普通模式和饥饿模式**

1、正常模式

所有goroutine按照FIFO的顺序进行锁获取，被唤醒的goroutine和新请求锁的goroutine同时进行锁获取，通常新请求锁的goroutine更容易获取锁(持续占有cpu)，被唤醒的goroutine则不容易获取到锁。公平性：否

2、饥饿模式

所有尝试获取锁的goroutine进行等待排队，新请求锁的goroutine不会进行锁获取(禁用自旋)，而是加入队列尾部等待获取锁。公平性：是

**20、切片扩容**

当使用append向slice追加元素，实际上是往底层数组添加元素。但是底层数组的长度是固定的，如果len-1指向的元素已经是最后一个元素，就会发生扩容，发生扩容的时候，会将slice迁移到新的内存位置，新底层数组的长度会增加，这样就可以放置新增的元素，同时为了应对未来可能再次发生的append操作，新的底层数组的长度，也就是新slice的容量预留了一部分的buffer

在go1.18之前，扩容策略：
> 当原 slice 容量小于 1024 的时候，新 slice 容量变成原来的 2 倍；原 slice 容量超过 1024，新 slice 容量变成原来的1.25倍。

在1.18之后，扩容策略:
> 当原slice容量(oldcap)小于256的时候，新slice(newcap)容量为原来的2倍；原slice容量超过256，新slice容量newcap = oldcap+(oldcap+3*256)/4

```go
package mian

func  main(){
    s:=make([]int,0)

    oldCap:=cap(s)

    for i:=0;i<2048;i++{
        s=append(s,i)
        newCap:=cap(s)
        if newCap!=oldCap{
            fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
            oldCap=newCap
        }
    }
}
```

不管是1.18之前还是1.18之后，在超过1024或者256，会对切片进行内存对齐，内存对齐之后新的切片容量要大于原来的切片，再根据这个新的切片的容量根据公式来算，会产生误差

**21、关于map详解**

这里先了解一下hash
所谓的hash算法：

>将任意长度的二进制值串映射为固定长度的二进制值串，这个映射的规则就是哈希算法，而通过原始数据映射之后得到的二进制值串就是哈希值。

常见的hash算法：MD5(MD5 Message-Digest Algorithm，MD5 消息摘要算法) SHA(Secure Hash Algorithm，安全散列算法)

hash算法有两个比较重要的特点：
1.很难根据hash值反向推导出原始数据
2.散列冲突的概率要很小

MD5本身其实是不可逆的，本质是因为MD5是一种散列函数，在计算过程中会丢失一部分信息
就好像马赛克

map 的设计也被称为 “The dictionary problem”，它的任务是设计一种数据结构用来维护一个集合的数据，并且可以同时对集合进行增删查改的操作。最主要的数据结构有两种：哈希查找表（Hash table）、搜索树（Search tree）。

哈希查找表一般会存在“碰撞”的问题，就是说不同的 key 被哈希到了同一个 bucket。一般有两种应对方法：链表法和开放地址法。链表法将一个 bucket 实现成一个链表，落在同一个 bucket 中的 key 都会插入这个链表。开放地址法则是碰撞发生后，通过一定的规律，在数组的后面挑选“空位”，用来放置新的 key。

搜索树法一般采用自平衡搜索树，包括：AVL 树，红黑树。面试时经常会被问到，甚至被要求手写红黑树代码，很多时候，面试官自己都写不上来，非常过分。

自平衡搜索树法的最差搜索效率是 O(logN)，而哈希查找表最差是 O(N)。当然，哈希查找表的平均查找效率是 O(1)，如果哈希函数设计的很好，最坏的情况基本不会出现。还有一点，遍历自平衡搜索树，返回的 key 序列，一般会按照从小到大的顺序；而哈希查找表则是乱序的。

