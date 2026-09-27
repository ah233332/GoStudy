package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU() //获取本机CPU核心数量
	runtime.GOMAXPROCS(1)   //设置进程使用最大CPU核心数量
	fmt.Println(num)
}
