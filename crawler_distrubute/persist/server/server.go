package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"github.com/zhongliangshan/test/crawler_distrubute/persist"
	"github.com/zhongliangshan/test/crawler_distrubute/rpctools"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	log.Fatal("", ItemSaverServer(":1234", "data_profile"))
}

func ItemSaverServer(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	rpcParams := rpctools.RpcParams{
		NetWork: "tcp",
		Host:    host,
	}
	serverParams := rpctools.RpcServerParams{
		RpcParams: rpcParams,
		Service: persist.ItemService{
			Client: client,
			Index:  index,
		},
	}

	err = serverParams.ServerRpc()
	if err != nil {
		return err
	}

	return nil
}
