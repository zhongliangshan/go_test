// 这一版本的控制跟并发就没有太大的关系了, 因为这样得到的就是顺序执行，管道的读取时阻塞的
package main

import (
	"fmt"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done<-true
	}
}

func createChan(id int) worker {
	w := worker{
		in : make(chan int),
		done : make(chan bool),
	}
	go doWork(id , w)
	return w
}

type worker struct {
	in chan int
	done chan bool
}

func chanDemo() {
	var channels [10]worker
	for i := 0; i<10;i++ {
		channels[i] = createChan(i)
	}
	for i := 0; i < 10; i++ {
		channels[i].in <- 'a' + i
		<-channels[i].done
	}

	for i := 0; i < 10; i++ {
		channels[i].in <- 'A' + i
		<-channels[i].done
	}
}


func main() {
	chanDemo()
}


