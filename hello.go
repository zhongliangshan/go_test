package main

import (
	"math"
	"fmt"
)


func main() {
	a := math.NaN()

	if math.IsNaN(a) {
		fmt.Println("is NAN")
	} else {
		fmt.Println(a)
	}

	for i:=0 ; i < 2 ;i++ {
		res := make(map[int]int)

		res[i]=i

		fmt.Println(res)
	}
}