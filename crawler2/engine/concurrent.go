package engine

import (
	"log"
)

type ConcurrentScheduler struct {
	// 定义调度器
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	// 会改变 Request 里面的内容 所以最好需要用 指针的变量
	ConfigureMasterRequest(chan Request)
}

func (e *ConcurrentScheduler) Run (seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureMasterRequest(in)
	for _ ,r := range seeds {
		e.Scheduler.Submit(r)
	}

	for i :=0;i<e.WorkerCount;i++ {
		CreateWorker(in , out)
	}

	itemCount := 0
	// 循环遍历所有的输出out
	for {
		// 得到 输出的request
		result := <-out
		// 打印items

		for _ , item :=range result.Items {
			log.Printf("Got Item : %d  :%s" ,itemCount , item)
			itemCount++
		}
		// 将Requests 加入调度器中
		for _ , r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func CreateWorker(in chan Request , out chan ParserResult) {
	go func() {
		for {
			request := <-in
			parserResult, err := Worker(request)
			if err != nil {
				continue
			}
			out <- parserResult
		}

	}()
}
