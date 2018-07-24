package main

import (
	"fmt"
	"os"
	"os/exec"
)

func adder() func(i int) int{
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 函数似的编程
type iAdder func(int) (int , iAdder)

func adder2(base int) iAdder {
	return func(v int) (int , iAdder) {
		return base + v , adder2(base +v)
	}
}

func main() {
	a := adder2(0)

	for i:=0;i<10;i++ {
		var s int
		s , a = a(i)
		fmt.Println(s)
	}

	fmt.Println(os.Getpid() , os.Getppid())

	cmd := exec.Command("echo", "-n", "golang")

	readCloser, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	out := make([]byte , 30)

	n, e := readCloser.Read(out)
	if e != nil {
		panic(e)
	}

	fmt.Println(out[:n])
}
