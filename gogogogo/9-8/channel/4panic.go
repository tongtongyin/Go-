package main

import (
	"fmt"
	"time"
)

/*
	goroutine中使用recover，解决协程中出现panic，导致崩溃问题
	如下面例子：其中test()方法写错了；
			在main执行过程中，不希望因为go test()协程出现错误
			而导致其他协程sayHello或者main中断
*/

// 每隔一秒打印一次hello
func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello")
	}
}

func test()  {
	// 这里可以使用defer + recover
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	// 定义个map
	var myMap map[int]string
	myMap[0] = "golang"	// error

}
func main() {

	go sayHello()
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("main->ok=", i)
		time.Sleep(time.Second)
	}
}
/*
 说明：如果起了一个协程，但是这个协程出现了panic，如果没有捕获这个panic
	就会造成整个程序崩溃，这时可以在goroutine中使用recover来捕获panic，
	这样及时这个协程发生问题，但是主线程仍然不受影响，可以继续执行。
*/
