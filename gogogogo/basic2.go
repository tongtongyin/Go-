package main

import (
	"fmt"
	"io/ioutil"
)
func main() {
	const filename = "abc.txt"
	// if的条件里可以赋值， if的条件里赋值的变量作用域就在这个if语句里
	if contents, err := ioutil.ReadFile(filename); err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
}
