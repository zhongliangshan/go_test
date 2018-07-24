package main

import (
	"net/http"
	"fmt"
	"strings"
)

func router(resp http.ResponseWriter , req *http.Request) {
	req.ParseForm() // 解析参数

	fmt.Println(req.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(resp, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {

	http.HandleFunc("/" , router)

	err := http.ListenAndServe(":10000", nil)

	if err != nil {
		panic(err)
	}
}
