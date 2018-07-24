package main

import (
	"os"
	"io"
)

func main() {
	name := "E:/data/go/src/github.com/zhongliangshan/test/file/read/aaa.txt"

	c := make(chan string)

	go bufRead(name , c)


	<-c
}

// 缓存读
func bufRead(name string , c chan string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte,1024)

	for {
		_, err := file.Read(buffer)

		if err == io.EOF {
			break
		}

	}

	c<-""
}
