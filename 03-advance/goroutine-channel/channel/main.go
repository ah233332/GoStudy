package main

import (
	"fmt"
	"sync" //synchornized同步
)

//计算1-200各个数的阶乘，并把各个数的阶乘放入map中

var (
	myMap = make(map[int]int, 10)
	lock  sync.Mutex //全局的互斥锁

)

// 计算n!,将结果放入myMap内
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
	//concurrent map writes?
	//多进程写入矛盾错误
	//竞态
	//资源竞争问题
}

func main() {

	//开了200个协程来做事情
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	// time.Sleep(time.Second * 10)

	lock.Lock() //这里为什么要加锁？？ 新手用全局互斥锁，高手用channel
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
