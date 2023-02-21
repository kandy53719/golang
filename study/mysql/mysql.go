// go get -u github.com/go-sql-driver/mysql
package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 初始化 全局变量
var mysql *sql.DB

func init() {
	if db, err := sql.Open("mysql", "root:123456@/go_db"); err != nil {
		fmt.Println("mysql连接失败：", err)
	} else {
		//最大连接时长
		db.SetConnMaxLifetime(time.Minute * 3)
		//最大连接数
		db.SetMaxOpenConns(10)
		//最大空闲数
		db.SetMaxIdleConns(10)
		//fmt.Println(db.Ping())
		mysql = db
	}
	//fmt.Println("初始化！")
}

func main() {
	//CreateTable()
	//ModifyTable()
	//DropTable()
	//TableInsert()
	//TableDelete()
	//TableUpdate()
	TableSelect()
}

// 数据库
func DbOperate() {
	var sql string = ""
	//创建数据库
	sql = "create database if not exists go_db"
	//使用数据库
	sql = "use go_db"
	//删除数据库
	sql = "drop database go_db"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println("数据库错误：", err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

func TableOperate() {
	var sql = ""
	//创建表
	sql = `create table if not exists user(
		id int not null auto_increment,
		name varchar(20) not null,
		password varchar(20) not null,
		primary_key(id)
		)engine=go_db default charset=utf8`
	//修改表名
	sql = "alter table user rename to user1"
	//修改表字段和类型
	sql = "alter table user add nickname varchar(20) not null"     //默认加最后
	sql = "alter table user add nickname varchar(20) first"        //加在最前
	sql = "alter table user add nickname varchar(20) after id"     //加在指定位置
	sql = "alter table user change nickname nickname2 varchar(30)" //修改字段类型及名称
	sql = "alter table user drop nickname2"                        //删除字段
	//索引
	sql = "alter table user add index id_index(id)"
	sql = "alter table user drop index id_index"
	//唯一索引
	sql = "alter table user add unique unique_id(id)"
	sql = "alter table user drop unique unique_id"
	//主键
	sql = "alter table user add primary key(id)"
	sql = "alter table user drop primary key"

	//删除表
	sql = "drop table user"

	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println("数据库错误：", err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 表数据
func TableDataOperate() {
	var sql = ""
	//插入
	sql = "insert into user(name, password) values('张三','12345')"
	//修改
	sql = "update user set password = '133' where name = '张三'"
	//删除
	sql = "delete from user where name = '张三'"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println("数据库错误：", err)
	} else {
		fmt.Println(r.LastInsertId())
		fmt.Println(r.RowsAffected())
	}
}

// 查询
func Search() {
	//var sql = ""

}

// 建表
func CreateTable() {
	var sql = `create table user(
				id int primary key auto_increment,
				name varchar(20) not null,
				password varchar(20) not null
			)`
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 改表
func ModifyTable() {
	//添加字段
	//var sql = "alter table user add nickname varchar(20)"                // 字段nickname在最后
	//var sql = "alter table user add nickname varchar(20) first"          // 字段nickname在最前
	var sql = "alter table user add nickname varchar(20) after password" // 字段nickname在id后面

	//修改字段
	sql = "alter table user change nickname nickname2 varchar(30) not null" // 修改字段名和字段类型

	//删除字段
	sql = "alter table user drop nickname2"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 删表
func DropTable() {
	var sql = "drop table user"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 增
func TableInsert() {
	var sql = "insert into user(name, password) values(?,?)"
	if r, err := mysql.Exec(sql, "王五", "1234567890"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 删
func TableDelete() {
	var sql = "delete from user where id = 3"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 改
func TableUpdate() {
	var sql = "update user set password = 123333 where id = 2"
	if r, err := mysql.Exec(sql); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

// 查
func TableSelect() {
	var sql = "select * from user where name is not null"
	if rows, err := mysql.Query(sql); err != nil {
		fmt.Println(err)
	} else {
		var user User
		for rows.Next() {
			rows.Scan(&user.id, &user.name, &user.password)
			fmt.Println(user)
		}
	}
}

type User struct {
	id       int
	name     string
	password string
}
