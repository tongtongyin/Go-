// 获取URL
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		// fmt.Println(reflect.TypeOf(resp.Status))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", b)

		// 习题1.9 打印出HTTP协议的状态码
		fmt.Println(resp.Status)  
	}
}