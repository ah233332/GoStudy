package main

import (
	"testing"
)

//文件名要以_test.go结尾。
// 函数名要以Test开头，紧接着的一个字母要大写(Add)中的A.

//go test -v cal_test.go cal.go 测试单个文件
//go test -v -run TestSub 测试单个文件里面的某个函数

// 编写测试用例，用来测试add是否正确
func TestAdd(t *testing.T) {
	res := add(10)
	if res != 55 {
		//fmt.Printf("Add(10) 执行错误,期望值=%v 实际值=%v\n",55,res)
		t.Fatalf("Add(10) 执行错误,期望值=%v 实际值=%v\n", 55, res)
	}

	t.Logf("执行正确")

}

func TestSub(t *testing.T) {
	ans := sub(7, 3)
	if ans != 4 {
		t.Fatalf("Sub(7,3) 执行错误,期望值为=%v 实际值=%v\n", 4, ans)
	}
	t.Logf("执行正确")
}
