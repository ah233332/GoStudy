package main

import (
	"fmt"
	"sync"
	"time"
)

// 利用structChan来完成需求，不依赖list[]int
type Res struct {
	index int //编号
	num   int //数据
}

var (
	numChan = make(chan int, 2000)
	resChan = make(chan Res, 2000)
)

// 计算1+...+n的值
func sumData(n int) int {
	ans := n * (n + 1) / 2
	return ans
}

// 从numChan取出数据，计算并存入ResChan
func MainChan(numChan chan int, resChan chan Res) {

	for v := range numChan {
		ans := sumData(v)
		var res = Res{
			index: v - 1,
			num:   ans,
		}
		resChan <- res
	}
}

// 将1-2000放入numChan
func AddData(numChan chan int) {
	defer close(numChan)
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
}

// 将ResChan数据打印出来
func PrintResChan(resChan chan Res) {
	var list = make([]int, 2000)
	for v := range resChan {
		list[v.index] = v.num
	}
	for index, n := range list {
		time.Sleep(time.Millisecond * 5)
		fmt.Printf("res[%d]=%d\n", index+1, n)
	}

}

func main() {
	go AddData(numChan)
	var ws sync.WaitGroup

	for i := 0; i < 8; i++ {
		ws.Add(1)
		go func() {
			MainChan(numChan, resChan)
			defer ws.Done()
		}()
	}

	ws.Wait()
	close(resChan)

	PrintResChan(resChan)

}
