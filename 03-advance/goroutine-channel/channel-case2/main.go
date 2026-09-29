package main

import (
	"fmt"
	"time"
)

// 单个管道的写入与读取
// 两个协程，一个写入数据，一个读取数据

//利用一个chan bool，来解决主线程快速结束从而影响其他线程的问题

var intChan = make(chan int, 50)
var exitChan = make(chan bool, 1)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writrData ", i)
		time.Sleep(time.Millisecond * 200)
		// 1000毫秒 = 1秒
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for v := range intChan {
		fmt.Println("readData ", v)
		time.Sleep(time.Millisecond * 200)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		if <-exitChan {
			break
		}
	}
}
