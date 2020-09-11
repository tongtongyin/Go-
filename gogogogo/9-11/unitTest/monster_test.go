package main

import "testing"

func TestMonster_Store(t *testing.T) {
	// 创建monster
	monster := Monster {
		Name: "红孩儿",
		Age: 10,
		Skill: "吐火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()函数错误，希望为=%v 实际为%v",
			true, res)
	}
	t.Logf("monster.Store() 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()函数错误，希望为=%v 实际为%v",
			true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Name函错误，希望为=%v 实际为%v",
			"红孩儿", monster.Name)
	}
	t.Logf("monster.ReStore() 测试成功")
}
