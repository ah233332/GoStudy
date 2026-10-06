package main

import "fmt"

const (
	a = iota
	b
	c = 100
	d
	e    = iota
	f, g = iota, iota
	h    = iota
	i
)

func main() {
	//const常量,不能修改，必须初始化赋值
	//const name [type] = value

	const pi float64 = 3.1415926
	const name = "tom"
	//const只能用在bool,int,float,string
	fmt.Println(pi, name)

	//编译前完成的
	//例如 var num = 9
	// const a = num/3 这么写是错误的

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
