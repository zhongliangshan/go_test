package main

import "fmt"

func sum( n int , c chan int) {
	total := 0

	for i:=0;i<n;i++ {
		total+=i
	}

	c <- total
}


func main() {
	//n:= 10
	//
	//c := make(chan int)
	//
	//go sum(n , c)
	//go sum(n/2 ,c)
	//
	//x , y := <-c , <-c
	//
	//fmt.Println(x ,y , x+y)

	// 设置channel有缓存的读
	n := 10

	c := make(chan int , 1)
	go sum(n , c)
	go sum(n/2 ,c)
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y , x+y)
}
