package main

import (
	"fmt"
	"reflect"
)

/*给一个变量var v float64 = 1.2,
使用反射来得到它的reflect.Value,
然后获取对应的Type,Kind和值，
并将reflect.Value转成interface{},
再将interface{}转换成float64 */

//练手小case

func reflectFloat(b interface{}) {
	//1.获取到reflect.Value
	rVal := reflect.ValueOf(b)

	//2.获取Type,Kind,值
	fmt.Println(rVal.Type())
	fmt.Println(rVal.Kind())
	fmt.Println(rVal.Float())

	//3.reflect.value --> interface{}
	Iv := rVal.Interface()

	//4.interface{} --> float64
	num := Iv.(float64)
	fmt.Println(num)
}

func main() {
	var v float64 = 1.2
	reflectFloat(v)
	fmt.Println("-------------------")
	//用reflect修改string类型变量的值
	var name string = "tom"
	rVal := reflect.ValueOf(&name)
	rVal.Elem().SetString("alan")
	fmt.Println(name)
}
