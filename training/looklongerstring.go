package main

import "fmt"

func looklongerstring (str string) (int , string) {
	start := 0
	max := 0
	min := 0
	occuredStr := make(map[rune]int)

	for i , ch := range []rune(str) {
		if lastId ,ok := occuredStr[ch];ok && lastId >= start {
			start = lastId + 1
		}

		if i - start + 1 > max {
			min = i
			max = i - start + 1
		}

		occuredStr[ch] = i
	}

	return max , str[min:min+max]
}

func main() {
	num , str := looklongerstring("aadasdaa")
	fmt.Println(num , str)
}
