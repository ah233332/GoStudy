package main

import "fmt"

func main() {
	var age int
	fmt.Print("请输入你的年龄:")

	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("已成年")
	}

	if age < 18 {
		fmt.Println("萝莉")
		return
	}
	if age < 30 {
		fmt.Println("青年")
		return
	}
	if age < 50 {
		fmt.Println("中年")
		return
	}
	fmt.Println("老年")

	// && || !
	// &&中，如果第一个条件是false的话，不走后面的条件语句
	// ||中，如果第一个条件是true的话，不走后面的条件语句

}
