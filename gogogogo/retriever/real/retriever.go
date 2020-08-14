package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)
// 自定义Retriever类型
type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

// Retriever实现Retriever接口的Get方法
// 参数：url 返回页面内容
func (r Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(
		resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}



