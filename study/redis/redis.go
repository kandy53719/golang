//go get -u github.com/go-redis/redis
//go get github.com/gomodule/redigo/redis

package main

import (
	"fmt"

	//goredis "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
)

var pool redis.Pool

func init() {
	// 基础连接
	if conn, err := redis.Dial("tcp", "127.0.0.1:6379"); err != nil {
		fmt.Println("redigo连接失败")
	} else {
		defer conn.Close()
		conn.Do("", "")
	}

	//连接池
	pool = redis.Pool{
		MaxIdle:     10,   //最初的连接数量
		MaxActive:   1000, //连接池最大连接数量,（0表示自动定义），按需分配
		IdleTimeout: 300,  //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	if conn := pool.Get(); conn != nil {
		fmt.Println("数据库连接失败")
	} else {
		defer conn.Close()
	}
}

func main() {
	string()
	list()
	hash()
	set()
	zset()
}

// string
func string() {
	var sql = ""
	sql = "set name1 zhangsan"                          //添加 更新
	sql = "get name1"                                   //获取
	sql = "mset name1 zhangsan name2 lisi name3 wangwu" //同是添加多个
	sql = "mget name1 name2 name3"                      //多值获取
	sql = "append name1 ~"                              //值追加

	var conn = pool.Get()
	if s, err := redis.String(conn.Do(sql)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

// list
func list() {
	var sql = ""
	sql = "lpush namelist zhangsan lisi"  //从左边插入张三
	sql = "rpush namelist wangwu zhaoliu" //从右边插入
	sql = "lrange namelist 0 1"           //从左边获取从0到1个元素
	sql = "lpop namelist"                 //从左边取出第一个元素
	sql = "lrem namelist 1 zhangsan"      //从左边删除1个张三

	var conn = pool.Get()
	if s, err := redis.Strings(conn.Do(sql)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

// hash
func hash() {
	var sql = ""
	sql = "hset userhash name zhangsan"         //单个赋值
	sql = "hmset userhash name zhangsan age 18" //多个赋值
	sql = "hget userhash name"                  //获取单个值
	sql = "hget userhash name age"              //获取多个值
	sql = "hdel userhash name age"              //删除指定字段
	sql = "hkeys userhash"                      //获取userhash里面所有字段
	sql = "hvals userhash"                      //获取userhash所有值
	sql = "hexists userhash name"               //查询userhash中是否包含name字段

	var conn = pool.Get()
	if s, err := redis.Strings(conn.Do(sql)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

// set
func set() {
	var sql = ""
	sql = "sadd nameset zhangsan lisi wangwu" //集合添加成员
	sql = "ismember nameset zhangsan"         //zhangsan是否为nameset的成员
	sql = "smembers nameset"                  //返回nameset里的所有成员
	sql = "srem nameset lisi zhangsan"        //移除nameset里面的张三和李四

	var conn = pool.Get()
	if s, err := redis.Strings(conn.Do(sql)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}

// zset
func zset() {
	var sql = ""
	sql = "sadd userzset 1 zhangsan 2 lisi 3 wangwu" //添加分数及成员
	sql = "zrange userzset 0 -1"                     //遍历有序集合所有成员
	sql = "zscore userzset zhangsan"                 //返回张三的分数
	sql = "srem userzset lisi zhangsan"              //移除nameset里面的张三和李四

	var conn = pool.Get()
	if s, err := redis.Strings(conn.Do(sql)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
