package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//channel放满了不能继续放了，取空了也不能接着取了
	var intChan chan int
	intChan = make(chan int, 3)
	num1 := 2
	intChan <- 1
	intChan <- num1
	intChan <- 5
	fmt.Println(intChan, &intChan)

	num := <-intChan
	<-intChan
	//相当于扔了一个数据
	fmt.Println(num, <-intChan, cap(intChan), len(intChan)) //cap是容量

	cat := Cat{
		Name: "奖杯猫",
		Age:  3,
	}
	//存放任意类型的管道
	allChan := make(chan interface{}, 3)
	allChan <- "hello"
	allChan <- 34
	allChan <- cat

	<-allChan
	<-allChan

	newcat := <-allChan
	fmt.Printf("类型=%T  数据=%v\n", newcat, newcat)
	cat111 := newcat.(Cat) //类型断言
	// newcat刚拿出来是interface{}空接口，可以是任何类型，需要.(type)判断一下
	fmt.Printf("cat.name=%v\n", cat111.Name)

}
