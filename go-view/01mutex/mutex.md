## mutex 详解

### 1、mutex的数据结构

```go
type Mutex struct{
    state int32 // 表示互斥锁的状态，比如是否被锁定等
    sema  uint32 // 表示信号量，协程阻塞等待该信号量，解锁的协程释放信号量从而唤醒等待信号量的协程
}
```
