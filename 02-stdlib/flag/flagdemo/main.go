package main

import (
	"flag"
	"fmt"
)

// flag包用于解析命令行参数
func main() {
	var user string
	var pwd string
	var host string
	var port int

	// 指定参数  默认值  备注说明
	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "localhost", "主机地址")
	flag.IntVar(&port, "port", 3306, "端口号")
	flag.Parse() // 解析命令行参数
	fmt.Println("user:", user)
	fmt.Println("pwd:", pwd)
	fmt.Println("host:", host)
	fmt.Println("port:", port)
}
