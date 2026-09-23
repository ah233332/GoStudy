package main

import "fmt"

func main() {
	/*
		var sum int = 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	/*
		for {
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	*/

	var a = []string{"李一", "王二", "张三", "李四"}
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	fmt.Print("\n")

	for index, item := range a {
		fmt.Println(index, item)
	}
	// for 下标，值 := range 切片{ }

}
