package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func testStruct() {
	p := Person{
		Name:   "张三",
		Age:    20,
		Salary: 10000.0,
	}

	// 将结构体序列化为 JSON
	//这里的&p是将结构体p的地址传递给json.Marshal函数，表示要序列化的是结构体p本身，
	// 而不是它的副本。这样可以避免在序列化过程中对结构体进行拷贝，提高性能。
	jsondata, err := json.Marshal(&p) //返回[]byte类型的json数据
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}
	fmt.Println("struct序列化后的 JSON 数据:", string(jsondata))

}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	//make函数。map类型不make直接写会报错，读不会报错返回0值
	//看到 slice / map / chan 就用 make；
	// 要指针就用 new（结构体更常见 &T{}）；其余情况两个都用不上
	a["name"] = "李超"
	a["age"] = 18
	a["hobby"] = []string{"看书", "跑步"}
	mapdata, err := json.Marshal(a)
	if err != nil {
		return
	}
	fmt.Println("map序列化后的 JSON 数据:", string(mapdata))
	//map是无序的，序列化后无法保证顺序
}

func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "一一"
	m1["age"] = 18
	m1["address"] = []string{"中国", "安徽省", "合肥"}

	m2 := make(map[string]interface{})
	m2["name"] = "二二"
	m2["age"] = 88
	m2["address"] = []string{"北美", "美国", "纽约"}

	slice = append(slice, m1)
	slice = append(slice, m2)

	sliceData, err := json.Marshal(slice)
	if err != nil {
		return
	}
	fmt.Println("slice序列化后的 JSON 数据:", string(sliceData))
}

func main() {
	//将结构体，map，切片进行序列化
	testStruct()
	fmt.Println()
	testMap()
	fmt.Println()
	testSlice()
}
