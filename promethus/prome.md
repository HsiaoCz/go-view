## 监控告警

**prometheus**

Prometheus 是一款基于时序数据库的开源监控告警系统
Prometheus的基本原理是通过HTTP协议周期性抓取被监控组件的状态，任意组件只要提供对应的HTTP接口就可以接入监控

为什么需要监控?


zabbix和promethus

zabix：
1. 图形页面友好
2. 成熟，资料也很多
3. 告警，分析，完善
4. 架构比较成熟

prometheus:
1. 不是很友好，各种配置都需要手写
2. 对docker，k8s监控有成熟的解决方案

主要使用场景区别是，Zabbix适合用于虚拟机、物理机的监控，因为每个监控指标是以 IP 地址作为标识进行区分的。而Prometheus的监控指标是由多个 label 组成，IP地址并不是唯一的区分指标，Prometheus 强大在可以支持自动发现规则，因此适合于容器环境。
从自定义监控项角度而言，Prometheus 开发难度较大，zabbix配合shell脚本更加方便。Prometheus在监控虚拟机上业务时，可能需要安装多个 exporter，而zabbix只需要安装一个 Agent。
Prometheus 采用拉数据方式，即使采用的是push-gateway，prometheus也是从push-gateway拉取数据。而Zabbix可以推可以拉。

目前的运维：容器运维

Prometheus特点：
- 多维数据模型，由度量名称和键值对标识的时间序列函数
- PromQL:一种灵活的查询语言，可以利用多维数据完成复杂的查询
- 不依赖分布式存储，单个服务器节点可以直接工作
- 基于HTTP的pull方式采集时间序列数据
- 推送时间序列数据通过PushGateway组件支持
- 通过服务发现或静态配置发现目标
- 多种图形模式以及仪表盘支持(grafana)


### prometheus的架构

![prometheus架构](https://pic1.zhimg.com/80/v2-6c625d85a52daed32891c13bc4b00710_720w.webp)

#### 1、prometheus Server

主要负责数据采集和存储，提供PromQL查询语言的支持
包含了三个组件:

- Retrieval：获取监控数据
- TSDB：时间序列数据库(Time Series Database)，我们可以简单的理解为一个优化后用来处理时间序列数据的软件，并且数据中的数组是由时间进行索引的，具备如下特点：
- 大部分时间都是顺序写入操作，很少涉及修改数据
- 删除操作都是删除一段时间的数据，而不涉及到删除无规律的数据
- 读操作一般都是升序或者降序

