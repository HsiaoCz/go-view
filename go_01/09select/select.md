## select 的实现原理

select 是 go 语言层面提供的多路 IO 复用机制
可以检测多个 channel 是否 ready(即是否可读或可写)

实现原理:
golang 实现 select 时，定义了一个数据结构表示每个 case 语句(含 default,default 实际上是一种特殊的 case)，这个 scase 在 runtime 包里

select 执行过程可以类比成一个函数

- 输入 case 数组
- 输出选中的 case
- 然后程序流程转到选中的 case 块

case 数据结构

```go
type scase struct{
    c *hchan
    kind uint16
    elem unsafe.Pointer
}
```

这里的 scase.c 为当前 case 语句所操作的 channel 指针，这也说明了一个 case 语句只能操作一个 channel

scase.kind 表示 case 的类型，分别表示读 channel，写 channel 和 default
三种类型分别由常量定义:

- caseRecv:case 语句中尝试读取 scase.c 中的数据
- caseSend:case 语句中尝试向 scase.c 中写入数据
- caseDefault:default 语句

scase.elem 表示缓冲区地址
根据 scase.kind 不同，有不同的用途

- scase.kind==caseRecv:scase.elem 表示读出的 channel 的数据存放地址
- scase.kind==caseSend:scase.elem 表示要写入的 channel 的数据存放地址

select 的实现逻辑

```go
func selectgo(cas0 *scase,order0 *uint16,ncases int)(int,bool)
```

函数的参数:

case0 为 case 数组的首地址
- selectgo 从这些 scase 中找出一个并执行

order0 为一个两倍 cas0 数组长度的 buffer
- 报错 scase 随机序列 pollorder
- 和 scase 中 channel 地址序列 lockorder
   - pollorder:每次 selectgo 执行都会把 scase 序列打乱，以达到随机检测 case 的目的
   - lockorder:所有 case 语句中 channel 序列，以达到去重防止对 channel 加锁时重复加锁的目的
ncases 表示 scase 数组的长度

函数返回值:
int:选中 case 的编号
这个 case 编号跟代码一致
bool：是否成功从 channel 中读取数据
如果选中的 case 是从 channel 中读取数据
则该返回值表示是否读取成功

selectgo 的实现机制

[https://blog.csdn.net/wys74230859/article/details/121879389]
