// QuickSort
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func partition(aris []int, begin int, end int) int {
	pvalue := aris[begin]
	i := begin
	j := begin + 1
	for j < end {
		if aris[j] < pvalue {
			i++
			aris[i], aris[j] = swap(aris[i], aris[j])
			fmt.Println(i,j)
		}
		j++
	}
	aris[i], aris[begin] = swap(aris[i], aris[begin])
	return i
}

func quick_sort(aris []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(aris, begin, end)
		quick_sort(aris, begin, mid)
		quick_sort(aris, mid+1, end)
	}
}

func rand_array(aris []int, lent int) {
	for i := 0; i < lent; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		aris[i] = r.Intn(1000)
	}
}

func main() {
	intas := []int{20, 1, 21, 7, 20, 4, 77, 1, 22, 0}
	//rand_array(intas, 10)
	fmt.Println(intas)
	quick_sort(intas, 0, 10)
	fmt.Println(intas)
}