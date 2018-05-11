// 城市列表解析器
package main

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler2/persist"
	"github.com/zhongliangshan/test/crawler2/scheduler"
	"github.com/zhongliangshan/test/crawler2/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentScheduler{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
