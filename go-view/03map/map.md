### 关于map

#### 1.hash表的存储原理

以取模+拉链法来快速了解下哈希表的存储原理

比如说有一个键值对：{"name":"zhangsan","age":18}，想让这个键值对存储到hash表中

第一步：先计算键和值的hash值
第二步：根据得到的hash值进行取模，比如说hash表里面有四个位置
这里就由得到的hash 值对4进行取模，根据取的模得到数据在hash表的存储位置

比如说"name"的hash值对4取模，得到的值是1，那么就将值存储到hash表1的位置
但是这里会存在一个问题，就是假如有一个key它的hash值的取模也是1那怎么办？

解决办法是也把它存储到这个位置，只不过使用链表的方式链出一个位置

这是hash表的基础存储原理

map的特点：键不能重复，键必须是可hash的，无序

键必须是可hash的，go语言中int/bool/string/array
引用类型不能hash

第三步：根据这个位置到hash表中存储数据

#### 2、map的增删改查

```go
// 删除键值对
data:=map[string]string{"n1":"zhangsan"}
delete(data,"nl")
```

#### 3、map的底层数据结构

之前的map的几种实现方式中,go语言使用的hash查找表的解决办法，并且使用的是链表法来解决
hash冲突

```go
// A header for a Go map.
type hmap struct {
    // 元素个数，调用 len(map) 时，直接返回此值
	count     int
	flags     uint8
	// buckets 的对数 log_2
	B         uint8
	// overflow 的 bucket 近似数
	noverflow uint16
	// 计算 key 的哈希的时候会传入哈希函数
	hash0     uint32
    // 指向 buckets 数组，大小为 2^B
    // 如果元素个数为0，就为 nil
	buckets    unsafe.Pointer
	// 等量扩容的时候，buckets 长度和 oldbuckets 相等
	// 双倍扩容的时候，buckets 长度会是 oldbuckets 的两倍
	oldbuckets unsafe.Pointer
	// 指示扩容进度，小于此地址的 buckets 迁移完成
	nevacuate  uintptr
	extra *mapextra // optional fields
}
```