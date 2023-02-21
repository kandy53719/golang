package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Client

func init() {
	if cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost")); err != nil {
		fmt.Println(err)
	} else if cli.Ping(context.TODO(), nil) != nil {
		fmt.Println("mongodb连接失败")
	} else {
		mongodb = cli
		fmt.Println("mongodb初始化成功")
	}
}

func main() {
	//AddDoc()
	//FindDoc()
	//UpdateDoc()
	DeleteDoc()
}

type Student struct {
	Name string
	Age  int
}

// // 创建数据库
// func dbOperate() {
// 	//创建数据库

// 	//删除数据库

// }

// 添加文档
func AddDoc() {
	fmt.Println("测试")
	c := mongodb.Database("go_db").Collection("student")

	//单条插入
	S1 := Student{
		Name: "张三",
		Age:  18,
	}
	if ior, err := c.InsertOne(context.TODO(), S1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("单条插入：", ior.InsertedID)
	}
	//多条插入
	S2 := Student{
		Name: "李四",
		Age:  19,
	}
	S3 := Student{
		Name: "王五",
		Age:  20,
	}
	if ior, err := c.InsertMany(context.TODO(), []interface{}{S2, S3}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ior.InsertedIDs...)
	}
}

// 查找文档
func FindDoc() {
	col := mongodb.Database("go_db").Collection("student")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{{"name", "张三"}}
	if cur, err := col.Find(ctx, filter); err != nil {
		fmt.Println(err)
	} else {
		defer cur.Close(ctx)
		var result bson.D
		for cur.Next(ctx) {
			if err2 := cur.Decode(&result); err2 != nil {
				fmt.Println(err2)
			} else {
				fmt.Println(result, result.Map()["name"])
			}
		}
	}
}

// 更新文档
func UpdateDoc() {
	col := mongodb.Database("go_db").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if ur, err := col.UpdateMany(ctx, bson.D{{"name", "张三"}}, bson.D{{"$set", bson.D{{"name", "ceshi"}, {"age", "10"}}}}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ur.ModifiedCount)
	}
}

// 删除文档
func DeleteDoc() {
	var col = mongodb.Database("go_db").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if re, err := col.DeleteMany(ctx, bson.D{{"name", "ceshi"}}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(re.DeletedCount)
	}
}
