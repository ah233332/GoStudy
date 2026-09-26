package main

import "testing"

var Str = "D:/code/go/study/02-stdlib/testing/testcase/js.json"

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "李超",
		Age:   88,
		Skill: "打瓦",
	}
	res := monster.Store(Str)
	if !res {
		t.Fatalf("monster.Store() err,期望：%v,实际：%v", true, res)
	}
	t.Logf("执行正确")
}

func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore(Str)
	if !res {
		t.Fatalf("monster.ReStore() err,期望：%v,实际：%v", true, res)
	}

	if monster.Name != "李超" {
		t.Fatalf("monster.ReStore() err,期望：%v,实际：%v", "李超", monster.Name)
	}

	t.Logf("执行正确")
}
