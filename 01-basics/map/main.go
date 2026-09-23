package main

import "fmt"

func main() {
	var usermap = map[int]string{
		1: "一一",
		2: "李二",
		3: "张三",
		4: "",
	}
	fmt.Println(usermap)
	value, ok := usermap[4]
	fmt.Println(value, ok)

	value, ok = usermap[5]
	fmt.Println(value, ok)

	usermap[1] = "李超"
	delete(usermap, 4)
	fmt.Println(usermap)

	//空map和未初始化map是两个东西
	//var amap = map[string]string{}
	//var amap = make(map[string]string)
	//两种初始化写法
}
