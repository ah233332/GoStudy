package main

import (
	"testing"
)

// 编写测试用例，用来测试add是否正确
func TestAdd(t *testing.T) {
	res := add(10)
	if res != 55 {
		//fmt.Printf("Add(10) 执行错误,期望值=%v 实际值=%v\n",55,res)
		t.Fatalf("Add(10) 执行错误,期望值=%v 实际值=%v\n", 55, res)
	}

	t.Logf("执行正确")

}
