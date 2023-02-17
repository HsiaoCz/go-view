## map 的底层实现

hmap 哈希表

hmap 是 Go map 的底层实现，每个 hmap 内部都含有多个 bmap(buckets 桶，oldbuckets 旧桶、overflow 溢出桶)，也就是说每个哈希表由多个桶组成

```go
type hmap struct {
    count int  //元素的个数
    flags  uint8  //状态标志
    B  uint8  // 可以容纳最多6.5*2^B个元素，6.5为装载因子
    noverflow  uint16  //溢出的个数
    hash0 uint32  //哈希种子

    buckets  unsafe.Pointer  //指向一个桶数组
    oldbuckets unsafe.Pointer  //指向一个旧桶数组，用于扩容
    nevacuate uinptr   //搬迁进度，小于noverflow的已经搬迁
    overflow *[2]*[]*bmap   //指向溢出桶的指针
}
```

buckets
buckets 是一个指针，指向一个 bmap 数组，存储多个桶

oldbuckets
oldbuckets 是一个指针，指向一个 bmap 数组，存储多个旧桶，用于扩容

overflow
overflow 是一个指针，指向一个元素个数为 2 的数组，数组的类型是一个指针，指向一个 slice,slice 的元素是桶(bmap)的地址，这些桶都是溢出桶，为什么有两个？因为 go 在哈希冲突过多时，会发生扩容操作。[0]表示当前使用的溢出桶集合，[1]是在发生扩容时，保存了旧的溢出桶集合。
overflow 存在的意义在于防止溢出桶被 gc

bmap 哈希桶

bmap 是一个隶属于 hmap 的结构体，一个桶可以存储 8 个键值对。如果有第 9 个键值对被分配到该桶，那就需要再创建一个桶，通过 overflow 指针将两个桶连接起来。在 hmap 中，多个 bmap 桶通过 overflow 指针项链，组成一个链表

```go
type bmap struct {
    // 元素hash值得高8位代表它在桶中得位置，如果tophash[0]<minaTopHash，表示这个桶的搬迁状态
    tophash [bucketCnt]uint8
    // 接下来8个key 8个value，但是我们不能直接看到，为了优化对其，go采用了key放在一起
    keys  [8]keytype  //key单独存储
    values [8]valuetype  // value单独存储
    pad  uintptr
    overflow  uintptr  //指向溢出桶的指针
}
```

map的扩容机制

增量扩容

go采用的是增量扩容方案，当map开始扩容后，每一次map操作都会触发一部分扩容搬迁工作(每进行一次赋值，会做至少一次搬迁工作)。由hmap中的nevacute成员记录当前搬迁的进度

注:在map进行扩容迁移的期间，不会触发第二次扩容。只有在前一个扩容迁移工作完成后，map才能进行下一次扩容

触发扩容的条件：
(1)存储的键值对数量过多(负载因子已经达到当前界限)
(2)由overflow指针所连接的溢出桶数量过多

go的负载因子界限：6.5
负载因子=哈希表中元素数量/桶的数量

扩容情况1：存储的键值对数量过多
这种情况下map会进行翻倍扩容

Go创建一个新的buckets数组，这个buckets数组的容量是旧buckets数组的两倍，并将旧数组的数据迁移至新数组

旧的buckets数组不会被直接删除，而是会把原来对旧数组的引用去掉，让GC来清除内存

扩容情况2：
溢出桶数量过多
如果出现这种情况，可能因为hash表里有过多的空键值对，很多桶的内部出现了空洞(装不满)。
这个时候就需要通过map扩容做内存整理，目的是为了清除bmap桶中空闲的键值对

这种情况下map扩容步骤与情况基本相同，只不过扩容后map容量还是原来的大小，go会创建一个与原来buckets数组容量相同的buckets数组，并将旧数组的数据逐步迁移至新数组。再去除旧数组的引用，让GC来清除内存