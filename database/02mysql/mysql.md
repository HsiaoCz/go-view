## mysql的面试题

>  1、MySQL的执行流程？或者一条语句在Mysql是如何执行的？

客户端请求 ---> 连接器（验证用户身份，给予权限）  ---> 查询缓存（存在缓存则直接返回，不存在则执行后续操作） ---> 分析器（对SQL进行词法分析和语法分析操作）  ---> 优化器（主要对执行的sql优化选择最优的执行方案方法）  ---> 执行器（执行时会先看用户是否有执行权限，有才去使用这个引擎提供的接口） ---> 去引擎层获取数据返回（如果开启查询缓存则会缓存查询结果）

> 2、mysql有哪些存储引擎？都有哪些区别？

**查看存储引擎**

```mysql
-- 查看支持的存储引擎
SHOW ENGINES

-- 查看默认存储引擎
SHOW VARIABLES LIKE 'storage_engine'

--查看具体某一个表所使用的存储引擎，这个默认存储引擎被修改了！
show create table tablename

--准确查看某个数据库中的某一表所使用的存储引擎
show table status like 'tablename'
show table status from database where name="tablename"
```

**设置存储引擎**

```mysql
-- 建表时指定存储引擎。默认的就是INNODB，不需要设置
CREATE TABLE t1 (i INT) ENGINE = INNODB;
CREATE TABLE t2 (i INT) ENGINE = CSV;
CREATE TABLE t3 (i INT) ENGINE = MEMORY;

-- 修改存储引擎
ALTER TABLE t ENGINE = InnoDB;

-- 修改默认存储引擎，也可以在配置文件my.cnf中修改默认引擎
SET default_storage_engine=NDBCLUSTER;
```

常见的存储引擎有 InnoDB、MyISAM、Memory、NDB。

InnoDB 现在是 MySQL 默认的存储引擎，支持事务、行级锁定和外键

在 MySQL中建立任何一张数据表，在其数据目录对应的数据库目录下都有对应表的 .frm 文件，.frm 文件是用来保存每个数据表的元数据(meta)信息，包括表结构的定义等，与数据库存储引擎无关，也就是任何存储引擎的数据表都必须有.frm文件，命名方式为 数据表名.frm，如user.frm。

查看MySQL 数据保存在哪里：show variables like 'data%'

**MyISAM 物理文件结构为：**

.frm文件：与表相关的元数据信息都存放在frm文件，包括表结构的定义信息等
.MYD (MYData) 文件：MyISAM 存储引擎专用，用于存储MyISAM 表的数据
.MYI (MYIndex)文件：MyISAM 存储引擎专用，用于存储MyISAM 表的索引相关信息

**InnoDB 物理文件结构为：**

.frm 文件：与表相关的元数据信息都存放在frm文件，包括表结构的定义信息等

.ibd 文件或 .ibdata 文件： 这两种文件都是存放 InnoDB 数据的文件，之所以有两种文件形式存放 InnoDB 的数据，是因为 InnoDB 的数据存储方式能够通过配置来决定是使用共享表空间存放存储数据，还是用独享表空间存放存储数据。
独享表空间存储方式使用.ibd文件，并且每个表一个.ibd文件
共享表空间存储方式使用.ibdata文件，所有表共同使用一个.ibdata文件（或多个，可自己配置）

关于这个问题的面试回答：

InnoDB 支持事务，MyISAM 不支持事务。这是 MySQL 将默认存储引擎从 MyISAM 变成 InnoDB 的重要原因之一；
InnoDB 支持外键，而 MyISAM 不支持。对一个包含外键的 InnoDB 表转为 MYISAM 会失败；
InnoDB 是聚簇索引，MyISAM 是非聚簇索引。聚簇索引的文件存放在主键索引的叶子节点上，因此 InnoDB 必须要有主键，通过主键索引效率很高。但是辅助索引需要两次查询，先查询到主键，然后再通过主键查询到数据。因此，主键不应该过大，因为主键太大，其他索引也都会很大。而 MyISAM 是非聚集索引，数据文件是分离的，索引保存的是数据文件的指针。主键索引和辅助索引是独立的。
InnoDB 不保存表的具体行数，执行 select count(*) from table 时需要全表扫描。而 MyISAM 用一个变量保存了整个表的行数，执行上述语句时只需要读出该变量即可，速度很快；
InnoDB 最小的锁粒度是行锁，MyISAM 最小的锁粒度是表锁。一个更新语句会锁住整张表，导致其他查询和更新都会被阻塞，因此并发访问受限。这也是 MySQL 将默认存储引擎从 MyISAM 变成 InnoDB 的重要原因之一；

> 一张表，里面有ID自增主键，当insert了17条记录之后，删除了第15,16,17条记录，再把Mysql重启，再insert一条记录，这条记录的ID是18还是15 ？

如果表的类型是MyISAM，那么是18。因为MyISAM表会把自增主键的最大ID 记录到数据文件中，重启MySQL自增主键的最大ID也不会丢失；
如果表的类型是InnoDB，那么是15。因为InnoDB 表只是把自增主键的最大ID记录到内存中，所以重启数据库或对表进行OPTION操作，都会导致最大ID丢失。

> 哪个存储引擎执行 select count(*) 更快，为什么?

MyISAM更快，因为MyISAM内部维护了一个计数器，可以直接调取。
在 MyISAM 存储引擎中，把表的总行数存储在磁盘上，当执行 select count(*) from t 时，直接返回总数据。


在 InnoDB 存储引擎中，跟 MyISAM 不一样，没有将总行数存储在磁盘上，当执行 select count(*) from t 时，会先把数据读出来，一行一行的累加，最后返回总数量。
InnoDB 中 count(*) 语句是在执行的时候，全表扫描统计总数量，所以当数据越来越大时，语句就越来越耗时了，为什么 InnoDB 引擎不像 MyISAM 引擎一样，将总行数存储到磁盘上？这跟 InnoDB 的事务特性有关，由于多版本并发控制（MVCC）的原因，InnoDB 表“应该返回多少行”也是不确定的。


### mysql的数据类型

> CHAR 和 VARCHAR 的区别？

char是固定长度，varchar长度可变：
char(n) 和 varchar(n) 中括号中 n 代表字符的个数，并不代表字节个数，比如 CHAR(30) 就可以存储 30 个字符。
存储时，前者不管实际存储数据的长度，直接按 char 规定的长度分配存储空间；而后者会根据实际存储的数据分配最终的存储空间
相同点：

char(n)，varchar(n)中的n都代表字符的个数
超过char，varchar最大长度n的限制后，字符串会被截断。

不同点：

char不论实际存储的字符数都会占用n个字符的空间，而varchar只会占用实际字符应该占用的字节空间加1（实际长度length，0<=length<255）或加2（length>255）。因为varchar保存数据时除了要保存字符串之外还会加一个字节来记录长度（如果列声明长度大于255则使用两个字节来保存长度）。
能存储的最大空间限制不一样：char的存储上限为255字节。
char在存储时会截断尾部的空格，而varchar不会。

char是适合存储很短的、一般固定长度的字符串。例如，char非常适合存储密码的MD5值，因为这是一个定长的值。对于非常短的列，char比varchar在存储空间上也更有效率。

> 列的字符串类型可以是什么？

字符串类型是：SET、BLOB、ENUM、CHAR、TEXT、VARCHAR

> blob 和text有什么区别

BLOB是一个二进制对象，可以容纳可变数量的数据。有四种类型的BLOB：TINYBLOB、BLOB、MEDIUMBLO和 LONGBLOB
TEXT是一个不区分大小写的BLOB。四种TEXT类型：TINYTEXT、TEXT、MEDIUMTEXT 和 LONGTEXT。
BLOB 保存二进制数据，TEXT 保存字符数据。

###、索引

```
说说你对 MySQL 索引的理解？
数据库索引的原理，为什么要用 B+树，为什么不用二叉树？
聚集索引与非聚集索引的区别？
InnoDB引擎中的索引策略，了解过吗？
创建索引的方式有哪些？
聚簇索引/非聚簇索引，mysql索引底层实现，为什么不用B-tree，为什么不用hash，叶子结点存放的是数据还是指向数据的内存地址，使用索引需要注意的几个地方？

为什么MySQL 索引中用B+tree，不用B-tree 或者其他树，为什么不用 Hash 索引
聚簇索引/非聚簇索引，MySQL 索引底层实现，叶子结点存放的是数据还是指向数据的内存地址，使用索引需要注意的几个地方？
使用索引查询一定能提高查询的性能吗？为什么?

那为什么推荐使用整型自增主键而不是选择UUID？

为什么非主键索引结构叶子节点存储的是主键值？
为什么Mysql索引要用B+树不是B树？

面试官：为何不采用Hash方式？
```

### mysql查询

```mysql
count(*) 和 count(1)和count(列名)区别 ps：这道题说法有点多
MySQL中 in和 exists 的区别？

UNION和UNION ALL的区别?

mysql 的内连接、左连接、右连接有什么区别？

什么是内连接、外连接、交叉连接、笛卡尔积呢？

mysql的执行顺序?
```

### mysql事物

```
事务的隔离级别有哪些？MySQL的默认隔离级别是什么？

什么是幻读，脏读，不可重复读呢？

MySQL事务的四大特性以及实现原理

MVCC熟悉吗，它的底层原理？

事务是如何通过日志来实现的，说得越深入越好。
```

### mysql的锁机制

```
数据库的乐观锁和悲观锁？

MySQL 中有哪几种锁，列举一下？

MySQL中InnoDB引擎的行锁是怎么实现的？

MySQL 间隙锁有没有了解，死锁有没有了解，写一段会造成死锁的 sql 语句，死锁发生了如何解决，MySQL 有没有提供什么机制去解决死锁

select for update有什么含义，会锁表还是锁行还是其他
```

### mysql调优

```
日常工作中你是怎么优化SQL的？
SQL优化的一般步骤是什么，怎么看执行计划（explain），如何理解其中各个字段的含义？
如何写sql能够有效的使用到复合索引？
一条sql执行过长的时间，你如何优化，从哪些方面入手？
什么是最左前缀原则？什么是最左匹配原则？

```

### 分区分表分库

```
随着业务的发展，业务越来越复杂，应用的模块越来越多，总的数据量很大，高并发读写操作均超过单个数据库服务器的处理能力怎么办？

说说分库与分表的设计
为什么要分库?
分库是什么？
```

### 三范式，百万级别的表如何删除

