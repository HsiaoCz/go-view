## kafka


### 消息队列的通信模式

**点对点模式**

消息生产者生产消息发送到queue中，然后消费者从queue中取出并且消费消息，一条消息被消费以后，queue中就没有了，不存在重复消费

**发布订阅模式**

消息的发布者将消息发布到topic中，同时多个消息订阅者订阅该消息，和点对点实现不同
，发布到topic的消息会被所有订阅者消费(类似于关注了微信公众号的人都能收到推送的文章)

发布订阅模式下，当发布者消息量很大时，显然单个订阅者的处理能力是不足的。实际上现实场景中，多个订阅者节点组成一个订阅组负载
均衡消费topic消息，即分组订阅，这样订阅者很容易实现消费

kafka是领英公司开发

kafka具有高吞吐，低延迟，高容错等特点

### kafka的架构

左边为生产者，中间为集群，右边为消费者

一个kafka集群由多个borker组成，每个borker由多个分区Partition组成，每个分区有各自的主题，多个分区也可以有相同的主题
所谓的主题就是消息的title，举个例子，比如nginx的日志可以写一个nginx_log，这就是一个主题
分区由leader和follower

broker:是指部署了kafka实例的服务器节点。每个服务器上有一个或多个kafka的实例，我们姑且认为每个broker对应一台服务器。每个kafka集群内的broker都有一个不重复的编号，如broker-1

Topic：消息的主题，可以理解为消息的分类，kafka的数据就保存在topic中。每个broker上都可以创建多个topic。实际应用中，通常是一个业务线建一个topic

partition：topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞吐量。同一个topic在不同的分区的数据是不重复的，partition的表现形式就是一个一个文件夹

Repartition：每一个分区都有多个副本，副本的作用就相当于是备胎。当主分区故障的时候会选择一个备胎上位，成为Leader。在kafka中默认的副本最大数量是10个，且副本的数量不能大于Broker的数量，follower和leader绝对在不同的机器，同一机器对同一个分区也只能存放一个副本

producer就是生产者，是数据的入口。producer在写入数据的时候会把数据写入到leader中，不会直接将数据写入follower，那么leader怎么找呢？写入的流程是什么样的呢？

**producer要往集群写入数据：**

1、消息的生产者从Kafka集群获取分区的leader信息
2、生产者将消息发送到leader
3、leader将消息写入本地磁盘
4、follower从leader拉取消息数据
5、follower将消息写入本地磁盘后向leader发送ACK
6、leader收到所有的follower的ACK之后向生产者发送ACK

**选择partition的原则**
在kafka中，如果某个topic有多个partition，producer又怎么知道该将数据发往那个partition呢？
1、producer在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition
2、如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition
3、如果既没有指定partition，又没有设置key，则会采用轮询的方式，即每次取一小段时间的数据写入某个partition，下一小段实现写入下一个partition

**ACK应答机制**

producer在向 kafka写入消息的时候，可以设置参数来确定是否确认收到，这个参数可以设置为0,1,all

- 0代表往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效率最高
- 1代表producer往集群发送数据只需要leader应答就可以发送下一条，只确保Leader发送成功
- all代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保leader发送成功和所有的副本都完成备份。安全性最高，但是效率最低

最后需要注意的是，如果往不存在的topic写数据，kafka会自动创建topic，partition和replication的数量默认配置都是

**Topic和数据日志**

topic是同一类别的消息的记录的集合。在kafka中，一个主题通常有多个订阅者，对于每个主题，kafka维护了一个分区数据日志文件
每个partition都是一个有序并且不可变的消息记录集合。当新的数据写入时，就被追加到partition的末尾。在每个partition中，
每条消息都会被分配一个顺序的唯一标识，这个标识被称为offset，即偏移量

partition在服务器上的表现形式就是一个一个文件夹，每个partition的文件夹下面会有多组seament文件，每组segement文件包含.index文件,.log文件.timeindex文件，其中.log文件就是实际存储message的地方，而.index和.timeindex文件为索引文件，用于检索信息


**kafka的面试题**

1、kafka的架构

2、生产者往kafka发送数据的流程

3、kafka选择分区的模式

4、生产者往kafka发送数据的模式

5、分区存储文件的原理

6、消费者消费数据的原理

7、为什么kafka快

8、kafka的使用场景
