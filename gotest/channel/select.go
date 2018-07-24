package main

import (
	"time"
	"fmt"
)

func main() {
	for {
		ticker := time.Tick(time.Duration(100) * time.Millisecond)

		select {
		case <-ticker:

			fmt.Println("定时器")

		}
	}

}
