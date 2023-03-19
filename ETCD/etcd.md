## etcd 

### 1、什么是ETCD

etcd 是一款分布式存储中间件，是用go语言编写的并通过raft一致性算法处理和确保分布式一致性，解决了分布式数据中数据一致性的问题

etcd常用于分布式中服务注册于发现

各个微服务之间通过etcd实现通信

**etcd到底是一款什么组件？**

etcd是一个用于配置共享和服务发现的键值存储的组件

etcd的核心架构
              gRPC Server
              ^        ^
mvcc  <------>etcd sercer <---------> snapshot / WAL
              |        |
             \|/      \|/
                  raft

etcd server:对外接受和处理客户端的请求
grpc server:etcd和其他etcd节点之间的通信和信息同步
MVCC :多版本控制，etcd的存储模块，键值对的每一次操作行为都会被记录存储，这些数据底层存储在BoltDB数据库中
WAL：预写式日志，etcd中的数据提交前都会记录到日志
Snapshot:快照，以防WAL日志过多，用于存储某一时刻etcd的所有数据
通过Snapshot和WAL结合，etcd可以有效地进行数据存储和节点故障恢复等操作

通过etcdctl客户端命令行和访问etcd中的数据
通过HTTP API接口直接访问etcd
etcd中的数据结构很简单，它的数据存储其实就是键值对的有序映射

### 2、ETCD的安装部署

etcd集群的三种启动方式:
1、静态启动，2、etcd动态发现，3、DNS发现

1、静态方式启动etcd集群，使用goreman,是一个go语言编写的多进程管理工具
`go get github.com/mattn/goreman`
