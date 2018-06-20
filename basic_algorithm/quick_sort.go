package main

import (
	"github.com/zhongliangshan/test/basic_algorithm/helper"
	"fmt"
)

func findPos(arr []interface{} , start , end int)int {
	val := arr[start]
	pos := start
	start ++
	for start < end {
		if b , _ :=helper.CheckVal(arr[start] , val) ; b {
			arr[start], arr[pos] = arr[pos], arr[start]
			pos++
		}
		start++

	}
	return pos
}


func quickSort2(arr []interface{} , start , end int)([]interface{}) {
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
	helper.GlobalArray.GenaratNearlySortedmArray(10000 , 1000)
	//var s  = []int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	//for i:=0;i<10;i++ {
	//	s[i] = rand.Intn(10)
	//}
	//fmt.Println(helper.GlobalArray.Array)

	helper.GlobalArray.Array = quickSort2(helper.GlobalArray.Array ,0 , len(helper.GlobalArray.Array))
	if !helper.GlobalArray.CheckSorted() {
		fmt.Println("sorted error ")
	}

}
