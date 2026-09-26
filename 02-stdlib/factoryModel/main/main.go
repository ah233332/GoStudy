package main

import (
	"factoryModel/model"
	"fmt"
)

func main() {
	var stu = model.NewStudent("李超", 69, 78.91)
	fmt.Println(*stu)
}
