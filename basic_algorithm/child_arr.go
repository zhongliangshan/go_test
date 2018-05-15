package main

import "fmt"

func childArr(arr , arr2 []int)bool {
	i ,j := 0,0
	flag := false
	for i < len(arr) && j < len(arr2) {
		if arr[i] == arr2[j] {
			i++
			j++
			flag = true
			continue
		}
		i++

		if flag == true {
			return false
		}
	}

	if i == len(arr) && j == len(arr2) {
		return true
	}

	return false
}

func main() {
	arr:=[]int{1,2,3,4,5,6}
	arr2 :=[]int{3,4,5,6,7,8}
	fmt.Println(childArr(arr , arr2))
}
