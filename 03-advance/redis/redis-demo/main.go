package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 通过go链接redis，写入数据与读取数据
func main() {
	//连接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	} else {
		fmt.Println("成功连接redis")
	}
	defer conn.Close()

	//每次运行前清空redis库
	conn.Do("flushall")

	//向redis写入hash数据
	_, err = conn.Do("set", "name", "tome")
	if err != nil {
		fmt.Println("set err=", err)
	}

	//从redis获取hash数据
	data, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
	}
	fmt.Println(data)

	conn.Do("hmset", "user1", "name", "tim", "age", "11", "address", "beijing")
	//一大串内容需要用map来接住
	data1, _ := redis.StringMap(conn.Do("hgetall", "user1"))
	fmt.Println(data1)
	//一大串内容用strings接住，再遍历输出
	data2, _ := redis.Strings(conn.Do("hmget", "user1", "name", "age"))
	for _, v := range data2 {
		fmt.Println(v)
	}

	//给数据设置有效时间 10s
	conn.Do("expire", "name", "10")

	//设置list数据
	//每次重复前先删除list，防止数据堆叠
	conn.Do("del", "list1")
	conn.Do("lpush", "list1", "aaa", "bbb", "ccc", "ddd", "eee", "fff")
	//这边pop指令,取走了
	datal1, _ := redis.String(conn.Do("rpop", "list1"))
	fmt.Println(datal1)
	datal2, _ := redis.Strings(conn.Do("lrange", "list1", "0", "-1"))
	fmt.Println(datal2)

	//stringmap适合键值对，strings适合列表
}
