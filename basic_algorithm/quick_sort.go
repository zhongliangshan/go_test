package main

import (
	"fmt"
	//"math/rand"
	//"strconv"
	//"regexp"
	//"syscall"
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


func quick_sort2(arr []int , start , end int)([]int) {
	if start +1 < end {
		pos := findPos(arr , start , end)
		quick_sort2(arr , start , pos)
		quick_sort2(arr , pos+1 , end)
	}

	return arr
}

func main() {
	var s  = []int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	//for i:=0;i<10;i++ {
	//	s[i] = rand.Intn(10)
	//}
	fmt.Println(s)

	fmt.Println(quick_sort2(s ,0 , len(s)))


	//pat := "/.|,|\"|;|:/"
	//re, _ := regexp.Compile(pat)
	//
	//mark := "asdasdsd"
	////将匹配到的部分替换为"##.#"
	//mark = re.ReplaceAllString(mark, "")
	//fmt.Println(mark)

}
