package main

import (
	"fmt"
	"sync"
	"time"
)

type res struct {
	num int //编号
	ans int //结果
}

var numChan = make(chan int, 2000)
var resChan = make(chan res, 2000)

// 向numChan写入2000个数据
func writeNumChan() {
	for i := 1; i <= 2000; i++ {
		numChan <- i
	}
	close(numChan)
}

// 计算1~n之和,返回ans
func calcaulate(n int) int {
	ans := (n * (n + 1)) / 2
	return ans
}

// 创建结构体，并将结构体写入writerResChan
func writerResChan() {
	for v := range numChan {
		Res := res{
			num: v,
			ans: calcaulate(v),
		}
		resChan <- Res
	}
	//close(resChan) 后面n个协程启动会重复关闭，会报错
}

// 读取readerResChan,并输出
func readerResChan() {
	/*	//range resChan 是“取走并处理”，不是“查看但保留”。
		// 代码遇到不匹配的结果时没有保存，数据被丢掉了
		//而且再次使用range时不是重头开始遍历的，而是接着上次继续遍历的

		for i := 1; i <= 2000; i++ {
			for v := range resChan {
				if v.num == i {
					fmt.Printf("res[%v]=%v", v.num, v.ans)
					break
				}
			}
		}	*/
	answers := make([]int, 2001)
	for v := range resChan {
		answers[v.num] = v.ans
	}

	for i := 1; i <= 2000; i++ {
		fmt.Printf("res[%v]=%v\n", i, answers[i])
		time.Sleep(time.Millisecond * 5)
	}

}

func main() {
	go writeNumChan()

	var wg sync.WaitGroup //创建一个并发任务计数器

	//任务量太小了，8个协程反而不是最优解
	//调度协程，分配任务反而更加浪费时间
	//ai测试，2~4个协程是该任务最优方案
	for i := 1; i <= 8; i++ {
		wg.Add(1) //增加一个待完成任务

		//go func 相当于一个匿名函数
		go func() {
			defer wg.Done() //任务完成，计数器减一
			writerResChan()
		}()
	}

	go func() {
		wg.Wait() //一直等待，直到计数器为0
		close(resChan)
	}()

	readerResChan()
	//这里没用go readerResChan()
	//因为这样协程刚开始，main就结束了，根本来不及打印

}
