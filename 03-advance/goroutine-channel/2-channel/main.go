package main

import "fmt"

//接着学习channel关闭，遍历

func main() {
	var intChan chan int
	intChan = make(chan int, 101)
	for i := 1; i <= 100; i++ {
		intChan <- i
	}

	/*
		for i := 1; i <= len(intChan); i++ {
			fmt.Println(<-intChan)
		}
		//不能这么遍历管道，每遍历一次，长度减一
		//最终只能输出50个数据
		//len(intChan)也不建议换成管道容量，管道有可能未满
		//输出空管道会报错哦
	*/

	//用for range 来遍历管道
	//遍历管道前要close(channel)
	close(intChan) //关闭管道，不可写入，可读取
	for v := range intChan {
		fmt.Println(v)
	}
}
