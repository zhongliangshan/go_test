package main

import "fmt"

func split(arr [] int)[]int {
	if (len(arr) < 2) {
		return arr
	}
	middle := len(arr) / 2
	left := arr[0:middle]
	right := arr[middle:]

	return merge_sort(split(left) , split(right))
}

func merge_sort(left , right []int)[]int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res , left[0])
			left = left[1:]
		} else {
			res = append(res , right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		res = append(res, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		res = append(res, right[0])
		right = right[1:]
	}

	return res
}

func main() {
	var s =[]int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	fmt.Println(s)
	fmt.Println(split(s))

}
