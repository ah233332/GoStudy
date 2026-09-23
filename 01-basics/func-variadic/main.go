package main

import "fmt"

func jiujiu() {
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Print("\n")
	}
}

func add(numberlist ...int) {
	var sum, sum1 int

	for i := 0; i < len(numberlist); i++ {
		sum += numberlist[i]
	}

	// for 下标，值 := range 切片{ }
	for _, item := range numberlist {
		sum1 += item
	}
	fmt.Println(sum, sum1)
}

//试试闭包形式，把相加改成间隔两秒后相加

func main() {
	jiujiu()
	fmt.Println()
	add(1, 2, 3, 4, 5)
}
