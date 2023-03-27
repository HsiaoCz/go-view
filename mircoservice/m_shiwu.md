### 1、本地事务

本地事务，是指传统的单机数据库事务，必须具备ACID原则

原子性(A):所谓的原子性，指的是在整个事务中的所有操作，要么全部完成，要么全部不完成，没有中间状态，对于事务在执行中发生错误，所有的操作都会被回滚，整个事务就像从来没有执行过一样

一致性(C):事务的执行必须保证系统的一致性，在事务开始之前和事务结束以后，数据库的完整性没有被破坏，就拿转账为例，A有500,B有500，如果在一个事务里A成功转给B50元，那么不管发生什么，那么最后A账户里和b账户里的数据之和必须是1000元

隔离性(I):所谓的隔离性就是说，事务与事务之间不会互相影响，一个事务的中间状态不会被其他事务感知，数据库保证隔离性包括四种不同的隔离级别

- 读未提交(脏读)
- 读已提交
- 可重复读
- 串行化

持久性(D):事务只要提交了，事务对数据库的变更就完全保存在了数据库

传统的额项目中，项目部署基本是单节点的，即单个服务器和单个数据库，这种情况下，数据库本身的事务机制就能保证ACID的原则，这样的事务就是本地事务

单个服务与单个数据库的架构中，产生的事务就是本地事务

其中原子性和持久性就是就要靠undo和redo日志来实现

#### 1.1、undo和redo

在数据库系统中，既有存放数据的文件，也有存放日志的文件，日志在内存中也有缓存log buffer，也有磁盘文件log file

Mysql中的日志文件，有这么两种与事务有关，undo日志与redo日志

**undo日志**

原子性利用undo日志来实现

undo log的原理

为了满足事务的原子性，在操作任何数据库之前，首先将数据备份到ubdo log，然后进行数据的修改，如果出现了错误或者用户执行了rollback语句，系统可以利用undo log中的备份将数据恢复到事务开始之前的版本

数据库写入到磁盘之前，会把数据先缓存到内存中，事务提交时才会写入磁盘中，用undo log实现原子性和持久性的简化过程:

假设有A、B两个数据，值分别为1,2:
A.事务开始
B.记录A=1到Undo log
c.修改A=3
D.记录B=2到undo log
E.修改b=4
F.将undo log写到磁盘
G.将数据写到磁盘
H.事务提交

- 如何保证持久性？
  事务提交前，会将数据写到磁盘，然后等待事务提交，也就是先持久再提交事务，所以事务只要提交了，肯定持久了

- 如何保证原子性
  - 每次对数据库修改，都会把数据记录到undo log，那么需要回滚时，可以读取undo log，恢复数据

  - 若系统在G和H之间崩溃
  此时事务并未提交，需要回滚，而undo log已经被持久化，可以根据undo log来恢复数据

  - 若系统在G之前崩溃
  此时数据并未持久化到磁盘，依然保持事务之前的状态

缺陷：每个事务提交前将数据和Undo log 写入磁盘，这样导致大量的磁盘IO，因此性能很低

如果能够将数据缓存一段时间，就能减少ID提高性能，但是这样会丧失事务的持久性，因此引入了另外一种机制实现持久化，即Redo log

**redo log**

和undo log 相反，redo log 记录的时新数据的备份，在事务提交前，只要将redo log持久化即可，不需要将数据持久化，减少了IO的次数

先看一下基本原理:

假如有A,B两个数据，值分别为1,2
A.事务开始
B.记录A=1到undo log buffer
C.修改A=3
D.记录A=3到redo log buffer
E.记录B=2到undo log buffer
F.修改B=4
G.记录B=4到redo log buffer
H.将undo log写入到redo log
I.将redo log 写入到磁盘
J.事务提交

数据写入到磁盘两种方式:

随机写:先寻址再写

顺序写:在连续空间写入，减少了寻址的时间，效率更高

- 如何保证原子性?
  如果事务提交前故障，通过undo log日志恢复数据,如果undo log都还没有写入，那么数据尚未持久化，无需回滚

- 如何保证持久化
  大家会发现，这里并没有出现数据的持久化，因为数据以及写入redo log持久化到了硬盘，因此只要到了步骤i之后，事务时可以提交的

- 内存中的数据何时持久化的到硬盘
  因为redo log已经持久化，因此数据库数据写入磁盘与否影响不大，不过为了避免出现脏数据(内存中与磁盘不一致)，事务提交后也会将内存数据刷入磁盘(也可以按照固定的频率写入到磁盘)

- redo log何时写入磁盘
  redo log会在事务提交走之前，或者redo log buffer满了的时候写入到磁盘
  
### 1.3、分布式事务

分布式事务，就是指不是在单个服务或单个数据库架构下，产生的事务：

- 跨数据源的分布式事务
- 跨服务的分布式事务
- 综合情况

(1) 跨数据源的分布式事务

主要是数据库的分库分表产生的

(2) 跨服务的分布式事务

服务拆分产生的分布式事务

(3) 分布式事务的数据一致性问题

在分布式环境下，肯定会出现部分操作成功，部分操作失败的问题，比如，订单生成了，库存也扣减了，但是用户账户的余额不足，这就造成数据不一致

订单的创建，库存的扣减，账户扣款在每一个服务和数据库内是一个本地事务，可以保证ACID原则

但是当我们把三个事情看作是一个事情的时候，要满足保证"业务"的原子性，要么所有操作全部成功，要么全部失败，不允许出现部分成功部分失败的现象，这就是分布式系统下的事务

此时的ACID难以满足，这就是分布式事务要解决的问题


### 分布式事务的产生

核心问题(1):在调用过程中，支付成功了，但是后续库存服务异常，就会造成用户付款了，但是库存没扣，会造成超卖

核心问题(2):在调用过程中，支付成功了，但是订单服务异常，就会造成用户付款了，订单状态未更新，造成用户投诉

#### 2、分阶段提交

##### 2.1、DTP和XA

分布式事务的解决手段之一，就是两阶段提交协议(2PC)

什么是两阶段提交?

1994年,X/open组织，定义了分布式事务的DTP模型，该模型包括这样几个角色

- 应用程序(AP)：我们的微服务
- 事务管理器(TM)：全局事务管理者
- 资源管理器(RM)：一般是数据库
- 通信资源管理器(CRM)：是TM和RM间的通信中间件

在该模型中，一个分布式事务可以被拆分成多个本地事务，运行在不同的AR和RM上，每一个本地事务的ACID很好实现，但是全局事务必须保证其中包含的每一个本地事务都能同时成功，若一个本地事务失败，则所有其他事务都必须回滚，但问题是，本地事务处理过程中，并不知道其他事务的运行状态，因此，就需要通过CRM来通知各个本地事务，同步事务执行的状态

因此，各个本地事务的通信必须有统一的标准，否则不同数据库间就无法通信，XA就是X/Open DTP中通信中间件与TM间联系的接口规范，定义了用于通知事件开始、提交、终止、回滚等接口，各个数据库厂商都必须实现这些接口


##### 2.2、二阶段提交

二阶段提交将全局事务拆分成为两个阶段来执行:

- 阶段一：准备阶段：各个本地事务完成本地事务的准备工作
- 阶段二：执行阶段，各个本地事务根据上一阶段的执行结果，进行提交或回滚

### 2、三阶段提交

### 3、TCC

它本质是一种补偿的思路，事务运行过程中包括三个方法
- try :资源的检测和预留
- Confirm:执行的业务操作提交;要求try成功confirm一定成功
- cancel：预留资源释放

执行分为两个阶段：
- 准备阶段(try)：资源的检测和预留
- 执行阶段(confirm/cancel)；根据上一步结果，判断下面的执行方法，如果上一步中所有事务参与者都成功，则这里执行confirm，反之，执行cancel