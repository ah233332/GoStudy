package main

import (
	"fmt"
	"io"
	"os"
)

//可以拷贝一切二进制文件

func CopyFile(dstName string, srcName string) (written int64, err error) {

	srcFile, err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open srcFile error: %v\n", err)
		return
	}
	defer srcFile.Close()

	//reader := bufio.NewReader(srcFile) // 问题4：多余，io.Copy 自带 32KB 缓冲。

	dstFile, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open dstFile error: %v\n", err)
		return
	}
	defer dstFile.Close()

	//writer := bufio.NewWriter(dstFile)
	//defer writer.Flush()

	return io.Copy(dstFile, srcFile)

}

func main() {
	dstFileName := "02-stdlib/file/copy-dir/srcdata/iii.txt" //目标文件
	srcFileName := "02-stdlib/file/copy-dir/srcdata/11.txt"  //源文件
	_, err := CopyFile(dstFileName, srcFileName)
	if err != nil {
		return // 问题11：静默 return、退出码 0，应打到 stderr 并 os.Exit(1)。
	}

}
