### go内存分配的原理

内存空间包含两个重要的区域：栈区和堆区

函数调用的参数、返回值以及局部变量大都会被分配到栈上

这部分内存会由编译器进行管理

堆中的对象由内存分配器分配并由垃圾收集器回收

go的内存分配借鉴了tcmalloc的思想，尽量减少在多线程模型下，锁的竞争开销
来提高内存分配的效率

TcMAlloc

再说吧这