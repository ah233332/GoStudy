package main

import (
	"fmt"
	"sync" //synchornized同步
	"time"
)

//计算1-20各个数的阶乘，并把各个数的阶乘放入map中

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
	//多goroutine写入矛盾错误
	//竞态
	//资源竞争问题
}

func main() {

	//开了20个协程来做事情
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5) //这个属于猜时间,后续学waitgroup

	//lock.Lock() 这地方加锁是防止myMap[n] = res与range矛盾，
	// 都访问myMap全局变量
	//新手用全局互斥锁，高手用channel
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	//lock.Unlock()

}
