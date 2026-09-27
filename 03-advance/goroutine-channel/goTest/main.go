package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test hello world", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main hello goland", i)
		time.Sleep(time.Second)
	}
}
