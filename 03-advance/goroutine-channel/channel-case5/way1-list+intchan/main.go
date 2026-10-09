package main

//利用[]int + Chan int 来完成需求
import (
	"fmt"
	"sync"
	"time"
)

var numChan = make(chan int, 2000)
var resChan = make(chan int, 2000)

// 将1-2000放入numChan
func AddData(numChan chan int) {
	defer close(numChan)
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
}

// 打印出resChan结果
func PrintResChan() {
	index := 0
	for v := range resChan {
		index++
		time.Sleep(time.Millisecond * 2)
		fmt.Printf("res[%d]=%d\n", index, v)
	}
}

// 计算1+...+n的值
func sumData(n int) int {
	ans := n * (n + 1) / 2
	return ans
}

// 从numChan取数据，运算，存入list
func Calculate(list *[]int) {
	for v := range numChan {
		ans := sumData(v)
		(*list)[v] = ans
	}

}

// 将list数据存入resChan
func AddResChan(list []int) {
	defer close(resChan)
	for index, v := range list {
		if index == 0 {
			continue
		}
		resChan <- v
	}
}

func main() {
	var list = make([]int, 2001)
	var sw sync.WaitGroup
	go AddData(numChan)

	for i := 0; i < 8; i++ {
		sw.Add(1)
		go func() {
			defer sw.Done()
			Calculate(&list)
		}()
	}

	sw.Wait()

	go AddResChan(list)
	PrintResChan()
}
