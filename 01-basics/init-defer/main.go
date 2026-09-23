package main

import "fmt"

// init函数一般用于初始化，引入数据库等。
func init() {
	fmt.Println("上面的先执行")
}

func init() {
	fmt.Println("比main函数提前执行的函数")
}

func main() {
	defer fmt.Println("defer2")
	defer fmt.Println("defer1")
	return
	//defer函数，离return近的先执行,用作资源清理。
}
