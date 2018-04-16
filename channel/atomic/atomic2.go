// 这样的话 会有数据冲突
package main

import (
	"time"
	"fmt"
	"sync"
)

type atomicInt2 struct {
	num int
	lock sync.Mutex
}

func (atomic *atomicInt2) increment2(){
	atomic.lock.Lock()
	defer atomic.lock.Unlock()
	atomic.num++
}

func (atomic *atomicInt2) get() int {
	atomic.lock.Lock()
	defer atomic.lock.Unlock()
	return atomic.num
}

func main() {
	var atomic atomicInt2
	atomic.num = 0
	atomic.increment2()

	go func() {
		atomic.increment2()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(atomic.get())

}
