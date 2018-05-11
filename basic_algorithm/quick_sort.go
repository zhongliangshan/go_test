package main

import (
	"fmt"
)

func findPos(arr []int , start , end int)int {
	val := arr[start]
	pos := start
	start ++
	for start < end {
		if arr[start] < val {
			arr[start], arr[pos] = arr[pos], arr[start]
			pos++
		}
		start++

	}
	return pos
}


func quickSort2(arr []int , start , end int)([]int) {
	if start < end {
		pos := findPos(arr , start , end)
		if pos != start || pos == end {
			quickSort2(arr , start , pos)
			quickSort2(arr , pos , end)
		}


	}


	return arr
}

func main() {
	var s  = []int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	//for i:=0;i<10;i++ {
	//	s[i] = rand.Intn(10)
	//}
	fmt.Println(s)

	fmt.Println(quickSort2(s ,0 , len(s)))
}
