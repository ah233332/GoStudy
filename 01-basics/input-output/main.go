package main

import "fmt"

func main() {
	var name string
	fmt.Print(`输入
你的
名
字: `)
	fmt.Scan(&name)
	fmt.Println(name) //自带一个换行printline

	//fmt.Printf("%d\n", 3)
	//fmt.Printf("%.2f\n", 3.1415)
	//fmt.Printf("%T\n", 3.33)
	//fmt.Printf("%v\n", "任意")
	// %t布尔型 %v默认型 %#v值加类型（调试用的）

	fmt.Print("\"输入\"你\\的年\\龄: ")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)

	var f = fmt.Sprintf("%.2f\n", 3.1415)
	fmt.Println(f)

}
