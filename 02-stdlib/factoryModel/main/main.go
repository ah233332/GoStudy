package main

import (
	"fmt"
	"study/02-stdlib/factoryModel/model"
)

func main() {
	var stu = model.NewStudent("李超", 69, 78.91)
	fmt.Println(*stu)
	fmt.Printf("成绩为：%v", stu.GetScore())
}
