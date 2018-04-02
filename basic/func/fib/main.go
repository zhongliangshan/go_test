package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
	"github.com/zhongliangshan/test/basic/func/fib/fib"
)


type intGen func() int


func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fib.Fibonacci()

	for i :=0 ;i<10; i++ {
		fmt.Println(f())
	}

	printFileContents(f)


	str := "23213123,"
	fmt.Println(strings.Trim(str , ","))

}