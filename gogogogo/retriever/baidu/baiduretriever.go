package baidu
// 定义一个结构体Retriever
type Retriever struct {
	Contents string	// 定义一个字符串变量Contents
}

// 定义Get方法，参数为url字符串，返回字符串类型
func (r Retriever) Get(url string) string {
	return r.Contents
}

