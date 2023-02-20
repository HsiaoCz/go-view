package main

import (
	"context"
	"fmt"
	"log"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// qmgo
// 七牛云开源的mongo操作mongo的库

// 安装
// go get -u github.com/qiniu/qmgo

func ConnMongo() (*qmgo.QmgoClient, error) {
	// 连接mongo
	// ctx := context.Background()
	// credential := qmgo.Credential{
	// 	Username: "hsiaocz",
	// 	Password: "shaw123",
	// }
	// client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://127.0.0.1:27017", Auth: &credential})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db := client.Database("test")
	// coll := db.Collection("article")

	// 如果连接指向固定的collection 可以使用这种方式
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://127.0.0.1:27017", Database: "test", Coll: "article", Auth: &qmgo.Credential{Username: "hsiaocz", Password: "shaw123"}})
	return cli, err
}

func main() {
	cli, err := ConnMongo()
	if err != nil {
		log.Fatalf("连接失败:%v\n", err)
	}
	fmt.Println("connect mongodb is successed!")
	type UserInfo struct {
		Name   string `bson:"name"`
		Age    uint16 `bson:"age"`
		Weight uint32 `bson:"weight"`
	}
	userInfo := UserInfo{
		Name:   "xm",
		Age:    7,
		Weight: 40,
	}
	// 创建单条索引
	cli.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"name"}})
	// 创建多条索引
	// cli.CreateIndexes(context.Background(),[]options.IndexModel{{Key: []string{"id2","id3"}}})

	// 插入文档
	// result, err := cli.InsertOne(context.Background(), userInfo)
	_, err = cli.InsertOne(context.Background(), userInfo)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(result.InsertedID)

	// 插入多条数据

	var userInfos = []UserInfo{
		{Name: "a1", Age: 6, Weight: 20},
		{Name: "b2", Age: 6, Weight: 25},
		{Name: "c3", Age: 6, Weight: 30},
		{Name: "d4", Age: 6, Weight: 35},
		{Name: "a1", Age: 7, Weight: 40},
		{Name: "a1", Age: 8, Weight: 45},
	}
	// results, err := cli.Collection.InsertMany(context.Background(), userInfos)
	_, err = cli.Collection.InsertMany(context.Background(), userInfos)
	if err != nil {
		log.Fatal(err)
	}
	// 查找一个文档
	one := UserInfo{}
	err = cli.Find(context.Background(), bson.M{"name": userInfo.Name}).One(&one)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(one)

	//批量查找
	batch := []UserInfo{}
	cli.Find(context.Background(), bson.M{"age": 6}).Sort("weight").Limit(7).All(&batch)
	fmt.Println(batch)

	// 计数
	// count, err := cli.Find(ctx, bson.M{"age": 6}).Count()

	// 更新
	// UpdateOne one
	// err := cli.UpdateOne(ctx, bson.M{"name": "d4"}, bson.M{"$set": bson.M{"age": 7}})

	// UpdateAll
	// result, err := cli.UpdateAll(ctx, bson.M{"age": 6}, bson.M{"$set": bson.M{"age": 10}})

	// 条件查询，查询单条
	// err := cli.Find(ctx, bson.M{"age": 10}).Select(bson.M{"age": 1}).One(&one)

	// 分组查询
	// matchStage := bson.D{{"$match", []bson.E{{"weight", bson.D{{"$gt", 30}}}}}}
	// groupStage := bson.D{{"$group", bson.D{{"_id", "$name"}, {"total", bson.D{{"$sum", "$age"}}}}}}
	// var showsWithInfo []bson.M
	// err = cli.Aggregate(context.Background(), Pipeline{matchStage, groupStage}).All(&showsWithInfo)

	
}
