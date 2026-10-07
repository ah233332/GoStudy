package main

import (
	"fmt"
	"reflect"
)

//通过反射修改，
//num int 的值
//student struct的值

func reflectNum(n interface{}) {
	rVal := reflect.ValueOf(n)
	fmt.Printf("rVal kind = %v\n", rVal.Kind())
	rVal.Elem().SetInt(3)

	/*
	   	rVal.Elem()
	   类似于：
	   		num := 9
	   		p *int = &num
	   		num1 := *p
	*/
}

func main() {
	var num int = 10
	reflectNum(&num)
	fmt.Println("num = ", num)

}
