// 获取URL
package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// 请求命令行输入的URL，返回resp
		resp, err := http.Get(url)
		// 请求出错
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//读取body字段到标准错误输出的文件描述符
		b, err := io.Copy(os.Stderr, resp.Body)
		// 关闭resp的Body流，防止资源泄露
		resp.Body.Close()
		// 请求出错
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// 请求成功输出b
		fmt.Printf("%s", b)
	}
}
