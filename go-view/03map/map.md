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
    // B用来决定创建的桶的个数
	B         uint8
	// overflow 的 bucket 近似数
	noverflow uint16
	// 计算 key 的哈希的时候会传入哈希函数
    // 哈希因子，类似于md5加盐里面加的盐
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

go的map核心是由hhmap和bmap两个结构体来实现的

当我们创建一个map，会创建一个hmap，存储map的基础信息，具体的键值对存储在bmap里面
每个bmap里面可以存储8个键值对

```go
// bmap
type bmap struct{
    // 8个元素的数组，存储字典每个key的hash值得高八位
    tophash
    // 8个元素得数组，存储字典key
    keys
    // 8个元素的数组，存储字典value
    values
    // 当前桶存不下时创建的溢出桶
    overflow
}
```
bmap又叫桶，可以存储8个键值对

#### 4、初始化

```go
// 初始化一个可容纳10个元素的map
info=make(map[string]string,10)
```

1、第一步：创建一个hmap结构体对象
2、第二步：生成一个哈希因子hash0并赋值到hmap对象中，用于后续为key创建哈希值
3、第三步：根据hint=10，并根据算法规则来创建B，当前B应该为1
```bash
hint  B
0-8   0
9-13  1
14-26 2
...
```
4、第四步：根据B创建桶(bmap对象)并存放在buckets数组中，当前bmap的数量应为2
- 当B<4时，根据B创建的桶的个数规则为：2^B(标准桶)
- 当B>4时，根据B创建的桶的个数为：2^B+2^(B-4)

每个bmap可以存储8个键值对，当不够存储时需要使用溢出桶，并将当前bmap中的overflow字段指向溢出桶的位置

#### 5、写入数据

```go
info["name"]="张三"
```

在map中写入数据执行的流程：
1、结合哈希因子和键name生成哈希值011011110011000100101010
2、结合获取的哈希值的后8位，并根据后8位的值来决定将此键值对存放在哪个桶中

```bash
将hash值和桶掩码(B为1的二进制)进行&运算，最终得到哈希值的后8位的值，假设当B为1时，结果为0：
哈希值：0101010101010101111110101011
桶掩码: 0000000000000000000000000001

的到的结果为1

根据示例可以知道，找桶的原则是根据后B位的位运算来计算出的索引位置，然后再去buckets数组中根据索引找到目标桶
```

3、第三步：在上一步确定桶之后，接下来就在桶中写入数据

```bash
获取hash值的tophash也就是hash值最高的八位，将tophash和keys,values分别写入到桶的三个数组中

如果桶满了，就通过overflow找到溢出桶的，并在溢出桶中继续写入

注意：以后在桶中查找数据时，会基于tophash来找(tophash相同会去比较Key)
```

4、第四步：hmap的个数count++(map中的元素个数加1)

#### 6、读取数据

```go
value:=info["name"]
```
在map中读取数据时，内部执行的流程为：
- 第一步：结合哈希因子和键name生成hash值
- 第二步：获取哈希值的后8位，并根据后B位的值来决定将此键值对存放到哪个桶中(bmap)
- 第三步：确定桶之后，再根据key的哈希值计算出tophash(高8位)，根据tophash和key去桶中查找数据

> 如果桶中没找到，则根据overflow再去溢出桶中找，均未找到则表示key不存在

#### 7、map的扩容

在向map中添加数据时，当达到某个条件，则会引发字典扩容

扩容条件：
- map中数据总个数/桶个数>6.5，会引发翻倍扩容
- 使用了太多的溢出桶(溢出桶使用的太多会导致map处理速度降低)
 - B<=15，已使用的溢出桶个数>=2^B时，引发等量扩容（原来有多少桶，就新建多少桶，将溢出桶的数据整到新创建的桶中）
 - B>15，已使用的溢出桶个数>=2^15时，引发等量扩容

#### 8、map的数据迁移

1、翻倍扩容

如果时翻倍扩容，那么迁移规则就是将就桶中的数据分流到两个桶中（分流的比例时根据hash值动态的决定的），并且桶 编号的位置为：同编号位置和翻倍后对应的位置

那么问题来了，如何实现这种迁移呢？

首先，我们要知道如果翻倍扩容（数据总个数/桶个数>6.5），则新桶个数是旧桶的2倍，即map中的B的值要+1（因为桶的个数等于2的B次方，而翻倍扩容之后新桶的个数就是2^B*2，就是2^(B+1),所以新桶的B的值=原桶B+1）

迁移时会遍历某个旧桶中所有的key（包括溢出桶），并根据key重新生成hash值，根据哈希值的底B位来决定将此键值对分流到哪个新桶中、

扩容后，B的值在原来的基础上已加1，也就意味着通过多1位来计算此键值对要分流到新桶的位置
- 当新增的位的值为0，则数据迁移到于旧桶编号一致的位置
- 当新增的位的值为1，则数据迁移到翻倍后对应编号的位置

例如：
```go
旧桶个数为32个，翻倍后新桶的个数为64个
在重新计算旧桶中的所有key的hash值后，红色位只能是0或1，所以桶中的所有数据的后B位只能是以下两种情况：
- 000111[7]，意味着要迁移到旧桶编号一致的位置
- 100111[39],意味着迁移到翻倍后对应的编号位置

特别提醒：同一个桶中key的哈希值的低B位一定是相同的，不然不会放在同一个桶中，所以同一个桶中黄色标记的位都是相同的
```

等量扩容

如果是等量扩容(溢出桶)，那么数据迁移机制就会比较简单，就是将旧桶（含溢出桶）中的值迁移到新桶中。
这种扩容的意义在于：当溢出桶比较多而每个桶中的数据又不多时，可以通过等量扩容和迁移让数据更紧凑，从而减少溢出桶