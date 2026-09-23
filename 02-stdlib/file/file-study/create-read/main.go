package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//创建与写入
	filepath := "02-stdlib/file/file-study/test.txt"
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	//perm(权限控制)在window不起作用，用在Linux
	if err != nil {
		fmt.Println(err)
		return
	}

	str1 := "hello world\r\n"
	//带缓存的写入方式
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str1)
	}
	//内容先写进缓存空间，再利用flush函数来写入磁盘
	writer.Flush()
	file.Close()

	fmt.Println("-----------------------------------------------------------")

	//清空并写入
	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer = bufio.NewWriter(file)
	str2 := "你好你好\r\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str2)
	}
	writer.Flush()
	file.Close()

	fmt.Println("-----------------------------------------------------------")

	//继续在下面写入
	file, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer = bufio.NewWriter(file)
	str3 := "这是写入在下方的新内容\r\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str3)
	}
	writer.Flush()
	file.Close()

	fmt.Println("-----------------------------------------------------------")

	//读出一个文件，并追加内容
	file, err = os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	str4 := "1111111111追加内容\r\n"
	writer = bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str4)
	}
	writer.Flush()
	file.Close()

}
