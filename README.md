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
反射的实现原理:
[https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/]

1.4、go select 底层数据结构和一些特性
[https://blog.csdn.net/wys74230859/article/details/121879389]

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

如果删除的是值类型，如果int,float,bool,string以及数组和struct,map的内存不会自动释放
如果删除的元素是引用类型，如指针，slice,map,chan等，map的内存会自动释放，释放的内存是子元素引用类型的内存占用

将map置为nil后，内存会被回收

1.8、go 的 gc 回收针对堆还是针对栈？变量内存分配在堆还是在栈？
[https://blog.csdn.net/csdniter/article/details/103617531]

go的垃圾回收针对堆
引用类型的全局变量内存分配在堆上，值类型的全局变量内存分配在栈上
局部变量的内存分配可能在栈上也可能在堆上

实际上，go语言编译器会自动决定把一个变量放在栈上还是放在堆上，编译器会做逃逸分析，当发现变量的作用域没有抛出函数范围，就可以在栈上，反之则必须分配在堆上

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

2.1、go net/http

2.2、go template
[https://www.liwenzhou.com/posts/Go/template/]

2.3、gin 框架
[https://www.liwenzhou.com/posts/Go/gin/]

2.4、实现优雅关机和平滑重启
[https://www.liwenzhou.com/posts/Go/graceful-shutdown/]

2.5、参数校验库 validator
[https://darjun.github.io/2020/04/04/godailylib/validator/]

2.6、依赖注入库 wire
[https://darjun.github.io/2020/03/02/godailylib/wire/]

2.7、log 库 logrus
[https://darjun.github.io/2020/02/07/godailylib/logrus/]

2.8、邮件库
[https://darjun.github.io/2020/02/16/godailylib/email/]

2.9、日志库 zap
[https://darjun.github.io/2020/04/23/godailylib/zap/]

2.10、xorm
[https://darjun.github.io/2020/05/07/godailylib/xorm/]

2.11、gorm
[https://www.liwenzhou.com/posts/Go/gorm/]
[https://www.liwenzhou.com/posts/Go/gorm-crud/]

gorm 标准文档:
[https://gorm.io/zh_CN/docs/index.html]

2.12、热加载工具 air
[https://darjun.github.io/2020/09/27/godailylib/air/]

2.13、路由管理库 mux
[https://darjun.github.io/2021/07/19/godailylib/gorilla/mux/]
[https://darjun.github.io/2021/07/21/godailylib/gorilla/handlers/]
[https://darjun.github.io/2021/07/22/godailylib/gorilla/schema/]

2.14、net/http 标准库
[https://darjun.github.io/2021/07/13/in-post/godailylib/nethttp/]

2.15、swagger 接口注释文档
[https://blog.csdn.net/qq_57467091/article/details/123373790]

2.16、JWT 登录鉴权
[https://www.liwenzhou.com/posts/Go/json-web-token/]

2.17、gin 的框架源码解析
[https://www.liwenzhou.com/posts/Go/gin-sourcecode/]

2.18、http 常用的压力测试工具
[https://www.liwenzhou.com/posts/Go/benchmark-tools/]

2.19、viper 配置管理
[https://darjun.github.io/2020/01/18/godailylib/viper/]

2.20、ini 配置文件
[https://darjun.github.io/2020/01/15/godailylib/go-ini/]

2.21、http 响应状态码
[https://juejin.cn/post/6844904202863394830]

## 3、数据库复习

3.1、mysql

mysql 知识点:[https://github.com/HsiaoCz/go-program/blob/master/08day/mysql/MySQL.md]

mysql 面试题:[https://juejin.cn/post/6850037271233331208]

sqlx:[https://github.com/HsiaoCz/go-program/blob/master/08day/mysql/sqlx%E4%BD%BF%E7%94%A8.md]

3.2、redis

redis 文档:
[https://juejin.cn/post/6844903982066827277]

go-redis:
[https://juejin.cn/post/7027347979065360392#heading-43]

3.3、mongoDB

文档:
[https://juejin.cn/post/6844904150635921422]

go 操作 mongodb
[https://github.com/qiniu/qmgo/blob/master/README_ZH.md]

## 4、微服务

4.1、rpc/jsonrpc
[https://darjun.github.io/2020/05/08/godailylib/rpc/]
[https://darjun.github.io/2020/05/10/godailylib/jsonrpc/]

4.2、protobuf

protobuf v3 语法:
[https://www.liwenzhou.com/posts/Go/Protobuf3-language-guide-zh/]

go 使用 proto:
[https://www.liwenzhou.com/posts/Go/protobuf/]
[https://www.liwenzhou.com/posts/Go/oneof-wrappers-field_mask/]

4.3、grpc

grpc:
[https://www.liwenzhou.com/posts/Go/gRPC/]

grpc 名称解析和负载均衡:
[https://www.liwenzhou.com/posts/Go/name-resolving-and-load-balancing-in-grpc/]

grpc-gateway:
[https://www.liwenzhou.com/posts/Go/grpc-gateway/]

grpc transcoding:
[https://www.liwenzhou.com/posts/Go/grpc-transcoding/]

4.4、服务注册与发现

consul:
[https://www.liwenzhou.com/posts/Go/consul/]

etcd:
[https://juejin.cn/post/6844903984440803341]
[https://juejin.cn/post/6844904031186321416]

raft
[https://juejin.cn/post/6924468915724615688]

4.5、配置中心 apollo
[https://www.liwenzhou.com/posts/Go/apollo/]

4.6、负载均衡策略

[https://github.com/HsiaoCz/microservice/blob/master/grpc/grpc.md]

4.7、幂等性机制、分布式锁

[https://github.com/HsiaoCz/microservice/blob/master/miden/miden.md]

4.8、链路追踪 jaeger
[https://github.com/HsiaoCz/microservice/blob/master/jaeger/jaeger.md]

4.9、熔断、限流、降级 sentinel
[https://github.com/HsiaoCz/microservice/blob/master/sentinel/readmd.md]

4.10、api 网关 kong

这里需要学习

4.11、go-kit

[https://www.liwenzhou.com/posts/Go/go-kit-tutorial-01/]

4.12、go-kratos

[https://go-kratos.dev/docs/]

4.13、go-zero

[https://go-zero.dev/cn/]

## 5、其他

5.2、rabbitMQ

5.3、kafka

5.4、ElsticSearch

这个一般未必用得到，最后总复习的时候看一下

5.5、promithus

5.6、k8s

5.7、Jekins

5.8、nginx

5.9、docker

5.10、liunx 常用命令

5.11、git

5.12、markdown

5.13、vim

5.14、正则表达式

## 6、基础

6.1、数据结构和算法
[https://abelsu7.top/2019/03/24/go-algo-and-data-structure/]

6.2、操作系统

6.3、计算机网络

6.4、计算机组成原理

6.5、设计模式

## 7、前端

7.1、HTML

7.2、css

7.3、javascript

7.4、vue
