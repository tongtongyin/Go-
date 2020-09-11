package main

import (
	"encoding/json"
	"fmt"
)

/*
 将结构体、map、切片进行序列化
*/

type Monster struct {
	Name string // `json:"m_name"`	// 反射机制
	Age int // `json:"age"`
	Birthday string
	Sal float64
	Skill string
}
// 将结构体序列化
func testStruct()  {
	monster := Monster{
		Name: "牛魔王",
		Age: 500,
		Birthday: "2011-11-11",
		Sal: 8000.1,
		Skill: "牛魔拳",
	}
	// 将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}
// 将map序列化
func testMap()  {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("a map序列化后=%v\n", string(data))
}
// 对切片进行序列化
func testSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "beijing"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"Mexico","Hawaii"}
	slice = append(slice, m2)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("slice序列化后=%v\n", string(data))
}

// 对基本数据类型序列化
func testFloat64()  {
	var num float64 = 2345.67
	// 对num进行序列化
	data, err := json.Marshal(num)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	// 输出序列化后的结果
	fmt.Printf("num序列化后=%v\n", string(data))
}
func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
