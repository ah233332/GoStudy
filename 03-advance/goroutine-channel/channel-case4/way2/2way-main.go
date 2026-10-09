package main

import "fmt"

func putNum(intChan chan int) {
	for i := 0; i <= 8000; i++ {
		intChan <- i
	}

	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true //假设是素数
		//判断是否为素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//说明不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据,退出")
	//不能关闭primeChan
	//向exitChan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000) //放入结果
	exitChan := make(chan bool, 4)    //标识退出类型

	//向intChan放入数据
	go putNum(intChan)

	//多个协程，从intChan中取出数据，并判断是否为素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	//遍历primeChan,把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d\n", res)
	}

}
