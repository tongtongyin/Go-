package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 1.定义结构体 Name:"红孩儿" Age:10 Skill:"吐火"
type Monster struct {
	Name string
	Age int
	Skill string
}
// 2.给monster绑定方法Store，可以将一个Monster变量(对象)，序列化后保存到文件中
func (M *Monster) Store () bool {
	data, err := json.Marshal(M)
	if err != nil {
		fmt.Println("Marshal err=", err)
		return false
	}
	fileName := "d:/monster.ser"
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
		return false
	}
	return true
}
// 3.给Monster绑定方法ReStore，可以将一个序列化的Monster从文件中读取
func (M *Monster) ReStore() bool {
	// 首先从文件中读取序列化的内容
	fileName := "d:/monster.ser"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("ReadFile err=", err)
		return false
	}
	// 进行反序列化
	err = json.Unmarshal(data, M)
	if err != nil {
		fmt.Println("Unmarshal err=", err)
	}
	return true
}
// 4.并反序列化为Monster对象，检查反序列化，名字正确
