package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func forceChange() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)

}

func enums() {
	const (
		php = iota
		java
		python
		golang
	)

	const (
		a = 1 << (10 * iota)
		b
		c
		d
	)

	fmt.Println(php, java, python, golang)
	fmt.Println(a, b, c, d)
}

func main() {
	euler()
	forceChange()
	enums()
}
