package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const basePath = "D:/code/go/study/03-advance/goroutine-channel/channel-case4/"

// 打开文件，并写入随机数据
func writeDataToFile(path string) {
	//打开文件path
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("creatsrcfile err:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 1000; i++ {
		//生成随机数据
		v := rand.IntN(100)
		//数据写入文件
		_, err := file.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			fmt.Println("writefile err:", err)
		}
	}
}

// 取出文件数据，进行排序整理，再重新写入文件
func SortData(srcPath, dstPath string) {
	srcFileData, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println("readfile err:", err)
		return
	}

	//将sreFileData数据转成String类型
	//然后以"\n"来分割,输出[]string
	lines := strings.Split(string(srcFileData), "\n")

	var numlist []int

	//遍历[]string切片
	for _, v := range lines {
		//排除末尾的空值,使其不影响Atoi
		if v == "" {
			continue
		}

		//转成Int类型，并存入numlist
		intLine, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("atoi err:", err, "原始内容:", strconv.Quote(v))
			continue
		}
		numlist = append(numlist, intLine)
	}

	//排序
	sort.Ints(numlist)

	//创建文件dstPAth
	file, err := os.Create(dstPath)
	if err != nil {
		fmt.Println("creatdstfile err:", err)
		return
	}
	defer file.Close()

	//写入新数据
	for _, v := range numlist {
		_, err := file.WriteString(strconv.Itoa(v) + "\n")
		if err != nil {
			fmt.Println("writefile err:", err)
		}
	}

}

func main() {
	//确保环境没问题，根据路径建造文件夹的
	os.MkdirAll(basePath+"srcdata", 0755)
	os.MkdirAll(basePath+"dstdata", 0755)

	var writeWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		srcPath1 := fmt.Sprintf("%ssrcdata/srcdata-%d.txt", basePath, i)
		writeWg.Add(1)
		go func() {
			defer writeWg.Done()
			writeDataToFile(srcPath1)
		}()
	}

	//利用一个structChan来控制协程
	//空的struct是0字节,相当于一个信号线
	ready := make(chan struct{})

	//不起协程的话会因为要等WriteWg结束，从来影响main继续执行
	//起一个协程监工writeWg,这样起协程不影响main继续往下执行
	go func() {
		//一旦识别到writeWg结束，就关闭ready
		writeWg.Wait()
		close(ready)
	}()

	var sortWg sync.WaitGroup
	for j := 0; j < 10; j++ {
		srcPath2 := fmt.Sprintf("%ssrcdata/srcdata-%d.txt", basePath, j)
		dstPath := fmt.Sprintf("%sdstdata/dstdata-%d.txt", basePath, j)
		sortWg.Add(1)
		go func() {
			defer sortWg.Done()
			//从ready取数据,来卡住协程继续执行
			//管道没东西,就一直在等,直到管道关闭
			//也可以给管道发一个值,ready<-struct{}{},但是只能醒来一个协程
			//close管道可以一键唤醒所以协程

			//此时十个协程全都准备好了,就等close(ready),一键启动了
			<-ready
			SortData(srcPath2, dstPath)
		}()

	}

	sortWg.Wait()

}
