// 这样的话 会有数据冲突
package main

import (
	"time"
	"fmt"
)

type atomicInt int

func (a *atomicInt) increment(){
	*a++
}

func main() {
	var a atomicInt
	a = 0
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a)

}
