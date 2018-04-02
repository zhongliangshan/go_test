package main

import (
	"reflect"
	"runtime"
	"fmt"
	"math"
	"time"
)

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}


func main(){
	fmt.Println("pow(10, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 10, 4))

	//var arr [5]int
	//arr2  :=  [3]int{1,2,3}
	//
	//fmt.Println(arr)
	//fmt.Println(arr2)

	timestamp := time.Now().Unix()
	fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02 15:04:05"))


	fmt.Println("2" != "2" && "2" != "1")
}
