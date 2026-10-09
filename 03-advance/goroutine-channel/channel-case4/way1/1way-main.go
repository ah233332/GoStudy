package main

import (
	"fmt"
	"sort"
	"sync"
)

var (
	numChan   = make(chan int, 1000)  //写入1-200000个数字
	primeChan = make(chan int, 10000) //存入素数
)

// 向numChan写入数据
func writeChan() {
	defer close(numChan) //发送方负责关闭管道
	for i := 1; i <= 200000; i++ {
		numChan <- i
	}
}

// 判断素数函数
func isprime(n int) bool {
	if n < 2 {
		return false

	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 结果写入primeChan
func calcaulate() {

	for v := range numChan {
		if isprime(v) {
			primeChan <- v
		}
	}
}

func readChan() {
	var s = make([]int, 0)
	for v := range primeChan {
		s = append(s, v)
		//不清楚具体数据长度，优先使用append
	}

	sort.Ints(s)

	for i, n := range s {
		fmt.Printf("第%d个素数是:%d\n", i+1, n)
	}

}

// 统计1-200000的数字中，哪些是素数
func main() {
	go writeChan()
	var wg sync.WaitGroup
	for i := 1; i <= 12; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			calcaulate()
		}()

	}

	go func() {
		wg.Wait()
		close(primeChan)
	}()

	readChan()
}
