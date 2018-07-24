package main

import (
	"github.com/robfig/cron"
	"time"
	"fmt"
)

func main() {
	c := cron.New()

	c.AddFunc("* * * * * *" , func() {
		b := "2018-06-27 15:58:58"
		t, err := time.Parse("2006-01-02 15:04:05"  , b)
		fmt.Println(t)
		fmt.Println(err)
		a := t.Unix()
		fmt.Println(a)
	})

	c.Start()
	select {

	}
}
