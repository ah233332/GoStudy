package main

import (
	"fmt"
	"reflect"
)

//使用反射来遍历结构体的字段，调用结构体的方法，获取结构体标签的值，并修改结构体的值

// 定义Monster
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

// 给Monster绑定方法,打印内容
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Print("---=end----\n\n")
}

// 给Monster绑定方法，返回两数相加的结果
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 给Monster绑定方法，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// 用反射对struct字段进行处理
func TestStructField(a interface{}) {
	//获取reflect.Type,Value类型
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	//获取a对应的kind
	kd := val.Kind()
	//如果不是结构体就退出不继续了
	if kd != reflect.Struct {
		fmt.Println("not struct")
		return
	}

	//获取该结构体的字段数量
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("val.field %d : 值=%v\n", i, val.Field(i))
		//val.field()返回的是值
	}

	fmt.Println()

	for i := 0; i < num; i++ {
		tagVal := typ.Field(i).Tag.Get("json")
		//typ.field()返回的是一个结构体方法
		if tagVal != "" {
			fmt.Printf("typ.field %d : tag=%v\n", i, tagVal)
		} else {
			fmt.Printf("typ.field %d : tag=\n", i)
		}
	}

}

// 用反射对struct方法进行处理
func TestStructMethod(b interface{}) {
	//获取reflect.Type,Value类型
	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)

	//获取到结构体里面的方法数量
	num := val.NumMethod()
	fmt.Printf("struct has %d Methods\n", num)

	//获取结构体中方法的排列顺序
	for i := 0; i < num; i++ {
		fmt.Printf("序列%d MethodName=%v\n", i, typ.Method(i).Name)
	}

	fmt.Println()

	//结构体的方法排序，按照函数名的ASCII码排序的
	val.Method(1).Call(nil) //选取第二个方法，然后调用

	//调用第一个方法，来两数之和
	var parse []reflect.Value
	parse = append(parse, reflect.ValueOf(10))
	parse = append(parse, reflect.ValueOf(20))
	res := val.Method(0).Call(parse)
	fmt.Printf("%v + %v = %v\n", parse[0].Int(), parse[1].Int(), res[0].Int())

}

// 用反射对struct的值进行修改
func StructUpdata(c interface{}) {
	val := reflect.ValueOf(c)
	val.Elem().FieldByName("Name").SetString("奖杯比")
	val.Elem().FieldByName("Score").SetFloat(69.78)

}

func main() {
	var v Monster = Monster{
		Name:  "奖杯",
		Age:   78,
		Score: 91,
		Sex:   "牛比",
	}

	TestStructField(v)
	fmt.Println("------------------------------------------")
	TestStructMethod(v)
	fmt.Println("------------------------------------------")
	StructUpdata(&v)
	fmt.Println(v)
}
