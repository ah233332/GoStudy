package main

import (
	"fmt"
	"math/rand"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

// 简单的结构体管道写入与读取
//很基础的写入与读取，没什么难度

func randName() string {
	names := []string{"一一", "二二", "三三", "四四", "五五", "六六", "七七",
		"八八", "九九", "事事", "elel", "twtw", "thirthir", "ff", "fwfw", "qiqi"}
	return names[rand.Intn(len(names))]
}

func randAddress() string {
	addresses := []string{
		"北京", "上海", "广州", "深圳", "杭州", "成都",
	}
	return addresses[rand.Intn(len(addresses))]
}

func main() {
	PerChan := make(chan Person, 10)
	for i := 1; i <= 10; i++ {
		p := Person{
			Name:    randName(),
			Age:     rand.Intn(100),
			Address: randAddress(),
		}
		PerChan <- p
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(<-PerChan)
	}

}
