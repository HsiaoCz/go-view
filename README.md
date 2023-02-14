## go 后端开发工程师面试复习

## 1.基础复习

1.go 基础
基础知识点:
[https://www.liwenzhou.com/posts/Go/golang-menu/]

1.1、关于 defer 的解析:
[https://www.cnblogs.com/traditional/p/11440728.html]

1.2、uint 的类型溢出问题，像 uint8 最大 255,超过就会发生问题，具体是怎样的问题呢?

1.3、go 解析 tag 是怎么实现的?反射的原理
[https://www.liwenzhou.com/posts/Go/reflect/]

1.4、go select 底层数据结构和一些特性[https://blog.csdn.net/wys74230859/article/details/121879389]

1.5、go 的 defer 的底层数据结构和一些特性

> defer 底层其实是一个单链表，每个 defer 是\_defer 实例，放在链表的头部，每个 defer 通过指针串联起来

1.6、单引号，双引号，反引号的区别？

Go 语言的字符串类型 string 在本质上就与其他语言的字符串类型不同：

- Java 的 String、C++的 std::string 以及 Python3 的 str 类型都只是定宽字符序列
- Go 语言的字符串是一个用 UTF-8 编码的变宽字符序列，它的每一个字符都用一个或多个字节表示
  即：一个 Go 语言字符串是一个任意字节的常量序列。

Golang 的双引号和反引号都可用于表示一个常量字符串，不同在于：

- 双引号用来创建可解析的字符串字面量(支持转义，但不能用来引用多行)
- 反引号用来创建原生的字符串字面量，这些字符串可能由多行组成(不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML 以及正则表达式

而单引号则用于表示 Golang 的一个特殊类型：rune，类似其他语言的 byte 但又不完全一样，是指：码点字面量（Unicode code point），不做任何转义的原始内容。

1.7、map 中删除一个 key，它的内存会释放吗？
[https://blog.csdn.net/csdniter/article/details/103611783]

1.8、go 的 gc 回收针对堆还是针对栈？变量内存分配在堆还是在栈？
[https://blog.csdn.net/csdniter/article/details/103617531]

1.9、map 的数据结构是什么?怎么实现扩容?
[https://blog.csdn.net/weixin_45743893/article/details/122927041]

1.11、关于 context 很重要
[https://www.cnblogs.com/juanmaofeifei/p/14439957.html]

1.12、go channel 的底层数据结构
[https://juejin.cn/post/7037656471210819614]

1.13、关于 GMP，抢占式调度
[https://www.bilibili.com/video/BV19r4y1w7Nx/?spm_id_from=333.337.search-card.all.click&vd_source=3a35126ffec4d04a8cf2c2532b09d9b5]

1.14、go 中 GC 回收机制三色标记与混合读写屏障
[https://www.bilibili.com/video/BV1wz4y1y7Kd/?spm_id_from=333.337.search-card.all.click&vd_source=3a35126ffec4d04a8cf2c2532b09d9b5]

1.15、Mutex 是悲观锁还是乐观锁？
[https://kingjcy.github.io/post/golang/go-mutex/]

1.15、mutex 有几种模式?
[https://www.cnblogs.com/tsxylhs/p/15042871.html]

1.16、goroutine 的自旋占用资源如何解决
自旋锁是指当一个线程在获取锁的时候，如果锁已经被其他线程获取，那么该线程将循环等待，然后不断地判断是否能够被成功获取，直到获取到锁才会退出循环。

自旋的条件如下：

- 还没自旋超过 4 次,
- 多核处理器，
- GOMAXPROCS > 1，
- p 上本地 goroutine 队列为空。

mutex 会让当前的 goroutine 去空转 CPU，在空转完后再次调用 CAS 方法去尝试性的占有锁资源，直到不满足自旋条件，则最终会加入到等待队列里。

1.17、怎么控制并发数？
[https://juejin.cn/post/6845166890571005959]
[https://zhuanlan.zhihu.com/p/471490292]
1.18、如何优雅的实现协程池
[https://juejin.cn/post/7086443265309818894]

1.19、go 的内存泄露问题
[https://blog.csdn.net/m0_37290103/article/details/116493163]

1.20、go 的其他问题
[https://blog.csdn.net/m0_37290103/article/details/116493163]

2.python 基础

## 2.web 复习

go 框架 goFrame
[https://goframe.org/pages/viewpage.action?pageId=1114399]
