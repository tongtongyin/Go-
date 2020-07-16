// 获取URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// 习题1.8 如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀
		// 判断URL是否有http://前缀，如果没有则加上
		if strings.HasPrefix(url, "http://") == false{
			url = "http://" + url
			//fmt.Println(url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)


	}
}