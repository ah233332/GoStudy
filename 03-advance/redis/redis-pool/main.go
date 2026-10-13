package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:   8, //最大空闲链接数
		MaxActive: 0, //表示和数据库的最大链接数，0为不限制
		//操作系统依旧会限制，但代码不限制
		IdleTimeout: 100, //最大空闲时间

		//初始化链接的代码
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//从链接池取一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "猫猫")
	if err != nil {
		fmt.Println("conn do err=", err)
		return
	}

	data, err := redis.String(conn.Do("get", "name"))
	fmt.Println(data)

	//pool.Close()
	//从pool中取出链接，要保证链接池没有关闭
}
