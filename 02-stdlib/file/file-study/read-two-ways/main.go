package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("02-stdlib/file/file-study/111.txt")
	if err != nil {
		fmt.Println(err)
	}

	//文件file是指针类型
	fmt.Printf("file=%v\n", file)

	defer file.Close()

	//默认缓冲区为4096字节  带缓冲区的读取方法，适合大型文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件末尾
			break
		}
		fmt.Print(str)
		//文件的末尾没有换行符的话，最后一行读取不到
	}

	//一次性读取文件,不需要手动打开与关闭,适合小文件
	//content以[]byte类型输出
	content, err := os.ReadFile("02-stdlib/file/file-study/111.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("content=%v\n", string(content))
}
