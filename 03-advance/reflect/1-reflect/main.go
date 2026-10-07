package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest0(b interface{}) {
	//通过反射获取变量的type(类型),kind(类别),值
	//kind类别 比 type类型 范围要大
	//kind抽象描述，type是具体描述
	//例: kind:struct  type:pkg.Student

	//1.获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	fmt.Printf("rVal类型=%T\n", rVal)
	//rVal显示是int，但本质是reflect.value类型

	fmt.Println(2 + rVal.Int())
	//.int()要匹配类型，如果用.float()会报错

	//3.获取变量对应的kind
	kind1 := rTyp.Kind()
	kind2 := rVal.Kind()
	fmt.Printf("kind = %v , kind = %v\n", kind1, kind2)
	fmt.Printf("kind type = %T,%T\n", kind1, kind2)

	//4.将rVal 转成 interface{}
	iv := rVal.Interface()

	//将 iv 类型断言 转换成需要的类型
	num1 := iv.(int)
	fmt.Println(num1)

}

// 对结构体的反射演示
func reflectTest1(b interface{}) {
	//1.获取reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=", rTyp)

	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	fmt.Printf("rVal类型=%T\n", rVal)

	//3.将rVal 转成 interface{}
	iv := rVal.Interface()
	fmt.Printf("iv=%v , iv type=%T\n", iv, iv)

	student, ok := iv.(Stu)
	if ok {
		fmt.Printf("stuName= %v", student.Name)
	}
}

type Stu struct {
	Name string
	Age  int
}

func main() {

	// 对(基本数据类型、interface{}、reflect.Value)进行反射基本操作
	var num int = 100

	reflectTest0(num)

	fmt.Println("-------------------------")

	//对结构体进行反射基本操作
	var student = Stu{
		Name: "tom",
		Age:  20,
	}
	reflectTest1(student)

}
