package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/zhongliangshan/test/basic/func/fib/fib"
)

// defer 先进后出
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)
}

func writeFile(filename string) {
	// 先创建一个文件
	//file , err :=os.Create(filename)
	file , err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//fmt.Println("error : %s" , err)
		if pathError , ok := err.(*os.PathError);!ok {
				panic(err)
		} else {
			fmt.Println(pathError.Op , pathError.Path , pathError.Err)
		}
		return
	}
	// 先定义函数退出 就关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i:= 0; i< 20;i++ {
		fmt.Fprintln(writer , f())
	}
}

func main() {
	tryDefer()

	writeFile("fib.txt")
}
