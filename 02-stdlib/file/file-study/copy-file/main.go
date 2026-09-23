package main

import (
	"fmt"
	"os"
)

func main() {
	//读取到原文件数据
	filepath := "02-stdlib/file/file-study/test.txt"
	newfilepath := "02-stdlib/file/file-study/newtest.txt"
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
	}

	//创建与写入内容
	//newfilepath := "02-stdlib/file/file-study/newtest.txt"
	//newfile, err := os.OpenFile(newfilepath, os.O_RDWR|os.O_CREATE, 0666)
	//writer := bufio.NewWriter(newfile)
	//writer.WriteString(string(content))
	//writer.Flush()
	//defer newfile.Close()

	err1 := os.WriteFile(newfilepath, content, 0777)
	if err1 != nil {
		fmt.Println(err1)
	}

	info, err := os.Stat("02-stdlib/file/file-study")
	if err != nil {
		fmt.Println("获取文件信息错误", err)
		return
	}
	fmt.Println(info.Size())
	fmt.Println(info.ModTime())
	fmt.Println(info.IsDir())
	fmt.Println(info.Name())
	fmt.Print(info.Mode())
}
