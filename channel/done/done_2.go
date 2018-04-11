// 利用channel
package main

import (
	"fmt"
)

func doWork2(id int, w worker2) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done<-true
	}
}

func createChan2(id int) worker2 {
	w := worker2{
		in : make(chan int),
		done : make(chan bool),
	}
	go doWork2(id , w)
	return w
}

type worker2 struct {
	in chan int
	done chan bool
}

func chanDemo2() {
	var channels [10]worker2
	for i := 0; i<10;i++ {
		channels[i] = createChan2(i)
	}
	for i , worker := range channels{
		worker.in <- 'a' + i
	}

	for _ , worker := range channels {
		<-worker.done
	}

	for i , worker := range channels{
		worker.in <- 'A' + i
	}


	for _ , worker := range channels {
		<-worker.done
	}
}


func main() {
	chanDemo2()
}


