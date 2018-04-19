package main

import (
	"fmt"

	real2 "retriver/real"
	"retriver/sham"
)

type Retriver interface {
	Get(url string) string
}

func download(r Retriver) string {
	return r.Get("http://www.baidu.com");
}

func main() {
	var r Retriver
	r = sham.Retriver{"this is test"}
	fmt.Println(download(r))

	r = real2.Retriver{}
	fmt.Println(download(r))
}
