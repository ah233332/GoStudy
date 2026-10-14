package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Monster struct {
	name  string
	age   int
	skill string
}

func main() {
	//链接redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial err=", err)
		return
	}

	//写入数量
	fmt.Print("你要存放几个monster:")
	var num int
	fmt.Scan(&num)

	//先存入list内
	var list []Monster
	for i := 0; i < num; i++ {
		var (
			name  string
			age   int
			skill string
		)
		fmt.Print("name:")
		fmt.Scan(&name)
		fmt.Print("age:")
		fmt.Scan(&age)
		fmt.Print("skill:")
		fmt.Scan(&skill)
		var mon = Monster{
			name:  name,
			age:   age,
			skill: skill,
		}
		list = append(list, mon)
		fmt.Println()
	}

	//遍历取出list数据，存入redis
	for i := 0; i < num; i++ {
		hname := fmt.Sprintf("mon%d", i)
		monname := list[i].name
		monage := list[i].age
		monskill := list[i].skill
		_, err = conn.Do("hmset", hname, "name", monname, "age", monage, "skill", monskill)
		if err != nil {
			fmt.Println("hmset err=", err)
			return

		}
	}

	//遍历取出redis数据并输出
	for i := 0; i < num; i++ {
		hname := fmt.Sprintf("mon%d", i)
		map1, err := redis.StringMap(conn.Do("hgetall", hname))
		if err != nil {
			fmt.Println("hgetall err=", err)
			return
		}
		fmt.Println(map1)
	}

	//清空redis库
	conn.Do("flushall")

}
