package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s完成了减法运算，%v-%v=%v", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func NewCal(num1, num2 int) Cal {
	var cal = Cal{
		Num1: num1,
		Num2: num2,
	}
	return cal
}

func reflectCal(b interface{}) {
	val := reflect.ValueOf(b)
	num := val.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("第%d个字段为：%v\n", i+1, val.Field(i))
	}

	var parse []reflect.Value
	var name string
	fmt.Print("请输入你的名字：")
	fmt.Scan(&name)
	parse = append(parse, reflect.ValueOf(name))
	val.Method(0).Call(parse)
}

func main() {
	var num1, num2 int
	fmt.Print("请输入两个字段：")
	fmt.Scan(&num1, &num2)
	cal := NewCal(num1, num2)
	reflectCal(cal)
}
