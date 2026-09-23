package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Count struct {
	english int
	number  int
	space   int
	other   int
	chinese int
}

func main() {
	//统计文件内数字、英文、空格、字符数量
	var count Count
	filename := "03-practice/char-count/作用文件.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range line {

			switch {
			case v >= 'a' && v <= 'z', v >= 'A' && v <= 'Z':
				count.english++
			case v >= '\u4e00' && v <= '\u9fff':
				count.chinese++
			case v == ' ' || v == '\t':
				count.space++
			case v >= '0' && v <= '9':
				count.number++
			default:
				count.other++
			}
		}
	}
	fmt.Println("统计结果如下：")
	fmt.Println()
	fmt.Printf("中文：%v,英文：%v,数字：%v,空格：%v,其他：%v", count.chinese, count.english, count.number, count.space, count.other)
}
