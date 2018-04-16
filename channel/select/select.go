package main

import (
	"fmt"
	"time"
	"math/rand"
)

func genChan() chan int{
	c := make(chan int)
	go func() {
		i := 0
		for  {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			c <-i
			i++
		}
	}()
	return c
}

func doWork(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createChan(id int) chan<- int {
	c := make(chan int)
	go doWork(id ,c)
	return c
}

func main() {
	var c1 , c2 = genChan() , genChan()
	w := createChan(0)

	var values []int
	var tm = time.After(time.Second * 10)
	var ticker = time.Tick(time.Millisecond * 1500)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = w
			activeValue = values[0]
		}

		n := 0
		select {
		case n = <-c1:
			values = append(values , n)
		case n = <-c2:
			values = append(values , n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-ticker:
			fmt.Println("len values:" , len(values))
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}

}
