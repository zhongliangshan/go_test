package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0;i<100;i++ {
		go func(i int) {
			fmt.Println("hello goroutine %d" , i)
		}(i)

	}

	time.Sleep(time.Millisecond)
}
