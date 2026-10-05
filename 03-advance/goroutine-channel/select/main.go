package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "Hello" + fmt.Sprintf("%d", i)
	}

	//传统方法中，遍历管道时，不关闭管道会阻塞而导致deadlock
	for {
		select {
		//如果intChan一直不关闭，不会一直阻塞
		//会自动匹配下一个case匹配
		case v := <-intChan:
			fmt.Printf("从intChan读取数据：%d\n", v)
			time.Sleep(time.Millisecond * 200)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取数据：%s\n", v)
			time.Sleep(time.Millisecond * 200)
		default:
			fmt.Println("匹配不到case了")
			return
		}
	}
}
