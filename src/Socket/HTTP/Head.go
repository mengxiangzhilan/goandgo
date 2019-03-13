package HTTP

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func Httpmain() {
	url := "https://www.baidu.com/"
	//response, err := http.Head(url)
	response, err := http.Get(url) //get方法获取资源内容，响应内容为 response 的Body属性。它是一个io.ReadCloser类型。
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(2)
	}
	fmt.Println(response.Status)        //响应状态
	for k, v := range response.Header { //Header属性对应HTTP 响应的header域
		fmt.Println(k+":", v)
	}
	fmt.Println("******************分割线*********************8")

	if response.Status != "200 OK" {
		fmt.Println(response.Status)
	}
	b, _ := httputil.DumpResponse(response, false)
	fmt.Println(string(b))
	contentTypes := response.Header["Content-Type"]
	if !acceptableCharset(contentTypes) {
		fmt.Println("Cannont handle", contentTypes)
	}
	var buf [512]byte
	reader := response.Body
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}
}
func acceptableCharset(contentTypes []string) bool {
	// each type is like [text/html; charset=UTF-8]
	// we want the UTF-8 only
	/*服务器提供内容时使用的字
	符集编码，甚至传输编码，通常是用户代理和服务器之间协商的那结果，但我们使用的Get
	的命令很简单，它不包括用户代理的内容协商组件。因此，服务器可以自行决定使用什么字
	符编码。*/
	for _, cType := range contentTypes {
		if strings.Index(cType, "UTF-8") != -1 {
			return true
		}
	}
	return false
}
