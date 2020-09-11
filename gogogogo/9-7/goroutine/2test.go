package main
/*
	测试：同样的数据量用goroutine和channel的方式 和 普通单线程的方式
		在运行速度方面有何差别
*/
import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	for num := 1; num <=150000; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {

		}
	}
	end := time.Now().Unix()
	fmt.Println("普通的方法耗时=", end - start)
}
