package persist

import "log"

func ItemSaver2() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d :%v", itemCount, item)
			itemCount++
		}
	}()

	return out
}
