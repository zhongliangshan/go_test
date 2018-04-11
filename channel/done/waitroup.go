//通过 go 的 sync.WaitGroup
package main

import (
	"fmt"
	"sync"
)

func doWork3(id int, w worker3) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done()
	}
}

func createChan3(id int , wg *sync.WaitGroup) worker3 {
	w := worker3{
		in : make(chan int),
		done : func() {
			wg.Done()
		},
	}
	go doWork3(id , w)
	return w
}

type worker3 struct {
	in chan int
	done func()
}

func chanDemo3() {
	var channels [10]worker3
	var wg sync.WaitGroup
	for i := 0; i<10;i++ {
		channels[i] = createChan3(i , &wg)
	}

	wg.Add(20)
	for i , worker := range channels{
		worker.in <- 'a' + i
	}

	for i , worker := range channels{
		worker.in <- 'A' + i
	}

	wg.Wait()

}


func main() {
	chanDemo3()
}


