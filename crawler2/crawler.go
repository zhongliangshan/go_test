// 城市列表解析器
package main

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler2/scheduler"
<<<<<<< HEAD
	"github.com/zhongliangshan/test/crawler2/persist"
=======
	"github.com/zhongliangshan/test/crawler2/zhenai/parser"
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
)

func main() {
	itemChan, err := persist.ItemSaver("data_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentScheduler{
<<<<<<< HEAD
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemChan,
=======
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
