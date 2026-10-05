package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
		time.Sleep(time.Millisecond * 200)
	}

}

func test() {
	//可以使用defer + recover
	defer func() {
		//使用recover捕获panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "Goland"
	//没make,会报错
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("mainGo:", i)
		time.Sleep(time.Millisecond * 200)
	}
}
