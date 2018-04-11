package main

import "fmt"

func main() {
	var c1 , c2  chan int
	for {
		select {
			case n := <-c1:
				fmt.Println("received from c1 : %d" , n)
			case n:= <-c2:
				fmt.Println("received from c2 : %d" , n)
			default:
				fmt.Println("received from default")
		}
	}

}
