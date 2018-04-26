// 城市列表解析器
package main

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"github.com/zhongliangshan/test/crawler2/zhenai/parser"
	"github.com/zhongliangshan/test/crawler2/scheduler"
)

func main() {
	e := engine.ConcurrentScheduler{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParserCityList,
	})
}
