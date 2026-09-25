package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json:"PersonName"` //反射机制
	Age    int     //`json:"PersonAge"`
	Salary float64 //`json:"PersonSalary"`
}

func unStruct() {
	str := "{\"PersonName\":\"张三\",\"Age\":20,\"Salary\":10000}"

	var p Person

	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Println("反序列失败")
		return
	}
	fmt.Printf("struct反序列化后为:%v\n p.Name:%v\n", p, p.Name)
}

func unMap() {
	str := "{\"age\":18,\"hobby\":[\"看书\",\"跑步\"],\"name\":\"李超\"}"

	//反序列化时不需要make,make操作被封装到了unmarshal里面了
	var m map[string]interface{}

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("反序列失败")
		return
	}

	fmt.Printf("map反序列化后为:%v\n", m)
}

func unSlice() {
	str := "[{\"address\":[\"中国\",\"安徽省\",\"合肥\"],\"age\":18,\"name\":\"一一\"}," +
		"{\"address\":[\"北美\",\"美国\",\"纽约\"],\"age\":88,\"name\":\"二二\"}]"

	var s []map[string]interface{}

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("反序列失败")
		return
	}

	fmt.Printf("slice反序列化后为:%v\n", s)

}

func main() {
	unStruct()
	unMap()
	unSlice()
}
