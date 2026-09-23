package main

import "fmt"

func register() {
	fmt.Println("注册")
}

func check() {
	fmt.Println("查找")
}

func usecenter() {
	fmt.Println("用户中心")
}

func main() {
	var index int
	fmt.Scan(&index)
	var usemap = map[int]func(){
		1: register,
		2: check,
		3: usecenter,
	}
	var value, ok = usemap[index]
	if ok {
		value()
	} else {
		fmt.Println("错误！")
	}
}

//可以试试改成循环执行的菜单
//闭包
