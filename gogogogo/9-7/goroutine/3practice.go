/*
 练习：
 1.启动1个协程，将1-2000的数放到一个channel中，比如numChan
 2.启动8个协程，从numChan取出数(比如n)，并计算1+...+n的值，存放到resChan
 3.最后8个协程协同完成工作后，再遍历resChan，显示结果，如res[1] = 1;res[10]=55
 注意：考虑resChan chan int 是否合适？
*/
package main

func main() {

}
