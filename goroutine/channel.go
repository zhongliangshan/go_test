package main

import (
	"fmt"
	"time"
)

type client chan string
var entering = make(chan client)

func main() {
	ch := make(chan string)
	go func() {
		ch <- "那你很棒棒哦😯 "
		time.Sleep(1 * time.Millisecond)
		}()
	go func() {entering <- ch}()
	o := <-entering
	time.Sleep(1* time.Second)
	fmt.Println(<-o)
}