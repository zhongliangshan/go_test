package main

import (
	//"errors"
	"fmt"
)

// panic recover 需要在defer中使用
func tryRecover() {
	defer func() {
		r := recover()
		if err , ok := r.(error);ok{
			fmt.Println("Error occurred:", err)
		} else {
			panic(err)
		}
	}()
	//panic(errors.New("this is an error"))
	b := 0
	fmt.Println(1/b)
}


func main() {
	tryRecover()
}
