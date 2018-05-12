package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func main() {
	pat := `/.|,|"|'|\\|;|:/`
	fmt.Println(pat)
	re, _ := regexp.Compile(pat)
	var mark string
	mark = `\`
	mark = re.ReplaceAllString(mark, "")
	fmt.Println(mark)
}
