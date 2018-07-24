package main

import (
	"os"
	"syscall"
	"os/signal"
	"sync"
	"fmt"
)

func main() {
	sig1 := make(chan os.Signal , 1)
	sig2 := make(chan os.Signal , 1)

	sigArr  := []os.Signal{syscall.SIGINT , syscall.SIGQUIT}
	sigArr2  := []os.Signal{syscall.SIGINT}


	signal.Notify(sig1 , sigArr...)
	signal.Notify(sig2 , sigArr2...)


	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for sig := range sig1 {
			fmt.Println("1:" , sig)
		}
		wg.Done()
	}()

	go func() {
		for sig := range sig2 {
			fmt.Println("2:" , sig)
		}
		wg.Done()
	}()

	wg.Wait()
}
