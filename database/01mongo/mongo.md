## mongo

mongo 的安装；

```bash
# 拉取镜像
docker pull mongo:4.2.5

# 启动
docker run -p 27017:27017 --name mongo \
-v /mydata/mongo/db:/data/db \
-d mongo:4.2.5

# 设置账号 (推荐使用这种方式启动)
docker run -p 27017:27017 --name mongo \
-v /mydata/mongo/db:/data/db \
-d mongo:4.2.5 --auth

# 进入容器
docker exec -it mongo mongo

# 之后在admin集合中创建一个账号用于连接，这里创建的是基于root角色的超级管理员账号
use admin
db.createUser({
    user: 'mongoadmin',
    pwd: 'secret',
    roles: [ { role: "root", db: "admin" } ] });

# 创建完成之后验证是否可以登录
db.auth("mongoadmin","secret")
```

mongo 的一些概念:
|SQL 概念|Mongo 概念|解释/说明|
|:----:|:-------:|:-------:|
|database|database|数据库|
|table|collection|数据表/集合|
|row|document|数据记录/文档|
|column|field|数据字段/域|
|index|index|索引|
|primary key|primary key|主键/mongo 会将\_id 自动设置为主键|

### 操作数据库

1、创建数据库

mongo 创建数据库使用的是 use 命令，当插入第一条数据时创建数据库，例如创建一个 test 数据库

```mongo
> use test
switched to db test
> db.article.insert({name:"MongoDB 教程"})
WriteResult({ "nInserted" : 1 })
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
test    0.000GB
```

2、删除数据库

mongo 删除数据库使用 db.dropDatabase()来删除

```mongo
> db.dropDatabase()
{ "dropped" : "test", "ok" : 1 }
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
```

### 集合操作

1、创建集合

使用 db.createCollection()创建集合

```mongo
> use test
switched to db test
> db.createCollection("article")
{ "ok" : 1 }
> show collections
article
```

2、删除集合

使用 db.collection.drop()删除集合

```mongo
> db.article.drop()
true
> show collections
```

### 文档操作

1、插入文档

使用 collection 对象的 insert 方法插入文档

```mongo
db.article.insert({title: 'MongoDB 教程',
    description: 'MongoDB 是一个 Nosql 数据库',
    by: 'Andy',
    url: 'https://www.mongodb.com/',
    tags: ['mongodb', 'database', 'NoSQL'],
    likes: 100
})
```

2、获取文档

使用 collection 对象的 find 方法获取对象

```mongo
db.article.find({})
```

更新文档

MongoDB 通过 collection 对象的 update()来更新文档

```mongo
db.collection.update(
   <query>,
   <update>,
   {
     multi: <boolean>
   }
)
# query：修改的查询条件，类似于SQL中的WHERE部分
# update：更新属性的操作符，类似与SQL中的SET部分
# multi：设置为true时会更新所有符合条件的文档，默认为false只更新找到的第一条
```

例如将 title 为 Mongo 教程的所有文档的 title 换成 MongoDB

```mongo
db.article.update({'title':'mongo'},{$set:{'title':"MongoDB"}},{multi:true})
```

collection.save()这个方法也可以用来更新文档

```mongo
db.article.save({
    "_id" : ObjectId("5e9943661379a112845e4056"),
    "title" : "MongoDB 教程",
    "description" : "MongoDB 是一个 Nosql 数据库",
    "by" : "Andy",
    "url" : "https://www.mongodb.com/",
    "tags" : [
        "mongodb",
        "database",
        "NoSQL"
    ],
    "likes" : 100.0
})
```

3、删除文档

collection.remove()方法可以删除文档

```mongo
db.collection.remove(
   <query>,
   {
     justOne: <boolean>
   }
)
# query：删除的查询条件，类似于SQL中的WHERE部分
# justOne：设置为true只删除一条记录，默认为false删除所有记录

```

删除 title 为 MongoDB 的文档

`db.article.remove({'title':'MongoDB'})`

4、查询文档

```mongo
<!-- 使用collection对象的find()方法来查询文档 -->
db.collection.find(query,projection)
<!-- query:查询的条件，类似于SQL中的where部分 -->
<!-- projection:可选，使用投影操作符指定返回的键 -->
```

例如查询 article 集合的所有文档:

```mongo
db.article.find()
```

mongoBD 中的条件操作符

| 操作       | 格式                   | sql 中类似的语句             |
| ---------- | ---------------------- | ---------------------------- |
| 等于       | {<key>:<value>}        | where title = 'MongoDB 教程' |
| 小于       | {<key>:{$lt:<value>}}  | where likes < 50             |
| 小于或等于 | {<key>:{$lte:<value>}} | where likes <= 50            |
| 大于       | {<key>:{$gt:<value>}}  | where likes > 50             |
| 大于或等于 | {<key>:{$gte:<value>}} | where likes >= 50            |
| 不等于     | {<key>:{$ne:<value>}}  | where likes != 50            |

条件查询，查询 title 为 mongo 教程的所有文档

```mongo
<!-- 条件查询，查询title为mongo教程的所有文档 -->
db.article.find({'title':'MongoDB'})

<!-- 条件查询，查询likes大于50的所有文档 -->
db.article.find({'title':{$gt:50}})

<!-- AND条件可以通过在find()方法传入多个键，以逗号隔开来实现，例如查询title为MongoDB 教程并且by为Andy的所有文档； -->

db.article.find({'title':'mongoDB 教程','by':'Andy'})

<!-- OR条件可以通过使用$or操作符实现，例如查询title为Redis 教程或MongoDB 教程的所有文档； -->

db.article.find({$or:[{"title":"Redis 教程"},{"title": "MongoDB 教程"}]})

<!-- AND 和 OR条件的联合使用，例如查询likes大于50，并且title为Redis 教程或者"MongoDB 教程的所有文档。 -->

db.article.find({"likes": {$gt:50}, $or: [{"title": "Redis 教程"},{"title": "MongoDB 教程"}]})
```

### 其他操作

1、limit 与 Skip 操作

读取指定数量的文档，可以使用 limit()方法，语法如下:

```mongo
db.collection.find().limit(NUMBER)
```

只查询 article 集合中的 2 条数据；

```mongo
db.article.find().limit(2)
```

跳过指定数量的文档来读取，可以使用 skip()方法，语法如下；

```mongo
db.collection.find().limit(NUMBER).skip(NUMBER)
```

从第二条开始，查询 article 集合中的 2 条数据；

```mongo
db.article.find().limit(2).skip(1)
```

2、排序

在 MongoDB 中使用 sort()方法对数据进行排序，sort()方法通过参数来指定排序的字段，并使用 1 和-1 来指定排序方式，1 为升序，-1 为降序；

```mongo
db.collection.find().sort({KEY:1})

<!-- 按article集合中文档的likes字段降序排列； -->
db.article.find().sort({'likes':-1})
```

3、索引

索引通常能够极大的提高查询的效率，如果没有索引，MongoDB 在读取数据时必须扫描集合中的每个文件并选取那些符合查询条件的记录。

MongoDB 使用 createIndex()方法来创建索引，语法如下；

```mongo
db.collection.createIndex(keys, options)
# background：建索引过程会阻塞其它数据库操作，设置为true表示后台创建，默认为false
# unique：设置为true表示创建唯一索引
# name：指定索引名称，如果没有指定会自动生成
```

给 title 和 description 字段创建索引，1 表示升序索引，-1 表示降序索引，指定以后台方式创建；

```mongo
db.article.createIndex({"title":1,"description":-1}, {background: true})
```

查看已经创建的索引:

```mongo
db.article.getIndexes()
```

4、聚合

MongoDB 中的聚合使用 aggregate()方法，类似于 SQL 中的 group by 语句，语法如下；

```mongo
db.collection.aggregate(AGGREGATE_OPERATION)
```

聚合操作$sum:计算总和
$avg：计算平均值，$min计算最小值，$max 计算最大值

根据 by 字段聚合文档并计算文档数量，类似与 SQL 中的 count()函数；

```mongo
db.article.aggregate([{$group : {_id : "$by", sum_count : {$sum : 1}}}])
```

根据 by 字段聚合文档并计算 likes 字段的平局值，类似与 SQL 中的 avg()语句；

```mongo
db.article.aggregate([{$group : {_id : "$by", avg_likes : {$avg : "$likes"}}}])
```

5、正则表达式

MongoDB 使用$regex 操作符来设置匹配字符串的正则表达式，可以用来模糊查询，类似于 SQL 中的 like 操作；

例如查询 title 中包含教程的文档；

```mongo
db.article.find({title:{$regex:"教程"}})
```

不区分大小写的模糊查询，使用$options 操作符；

```mongo
不区分大小写的模糊查询，使用$options操作符；
```
