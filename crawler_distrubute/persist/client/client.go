package client

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler_distrubute/rpctools"
	"log"
)

func ItemSaver(
	host string) (chan engine.Item, error) {
	// 建立客户端
	clientParams := rpctools.RpcParams{
		NetWork: "tcp",
		Host:    host,
	}
	client, err := clientParams.ClientRpc()

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item "+
				"#%d: %v", itemCount, item)
			itemCount++
			result := ""
			err := client.Call("ItemService.Saver", item, &result)
			if err != nil {
				log.Printf("Item Saver: error "+
					"saving item %v: %v",
					item, err)
			}
		}
	}()

	return out, nil
}
