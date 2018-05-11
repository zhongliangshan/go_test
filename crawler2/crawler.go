// 城市列表解析器
package main

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler2/persist"
	"github.com/zhongliangshan/test/crawler2/scheduler"
	"github.com/zhongliangshan/test/crawler2/zhenai/parser"
)

func main() {
	e := engine.ConcurrentScheduler{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
