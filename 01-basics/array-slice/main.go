package main

import (
	"fmt"
	"sort"
)

func main() {
	var namelist = [3]string{"一一", "李二", "张三"} //事先定义好长度的是数组
	fmt.Println(namelist[len(namelist)-1])

	var name1list []string //不固定长度的是切片（数组plus）
	//name1list := make([]string, 0)
	//make([]type,长度,容量)

	name1list = append(name1list, "一个一个")
	name1list = append(name1list, "大傻逼")
	fmt.Println(name1list[0])
	fmt.Println(name1list[1])

	//nil是空的意思，类似与null

	array := [3]int{1, 2, 3}
	fmt.Println(array[:])
	fmt.Println(array[1:3]) //左闭右开区间

	var it = []int{4, 1, 7, 6, 3}
	//sort.Ints(i)  //从小到大排序
	//sort.Sort(sort.Reverse(sort.IntSlice(it)))  //从大到小排序
	//先将it包装成intslice类型，然后给reverse处理翻转，最后再sort排序
	//不通用，建议用slice函数

	sort.Slice(it, func(i, j int) bool {
		return it[i] > it[j]
	})
	fmt.Println(it)
	sort.Slice(it, func(i, j int) bool { return it[i] < it[j] })
	//sort.slice是通用排序工具，return it[i] > it[j] 降序，return it[i] < it[j] 升序
	fmt.Println(it)

}
