package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
)


func main() {


	 c := make(chan os.Signal ,1)
	 signal.Notify(c , syscall.SIGINT , syscall.SIGTERM)

	 for {
		 select{
		 case <-c:
			 os.Exit(0)
			 break
		 case <-time.After(time.Duration(1) * time.Second):
			 fmt.Println(time.Now().Unix())
		 }
	 }

}
