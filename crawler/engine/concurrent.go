package engine

import "log"

type ConcurrentScheduler struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(request Request)
	//ConfigureMasterRequest(chan Request)
	Run()
	WorkReady(request Request)
}

func (e ConcurrentScheduler) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	//e.Scheduler.ConfigureMasterRequest(in)

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(in, out)
	}
	for {
		parserResult := <-out

		for _, item := range parserResult.Items {
			log.Printf("Got item %s", item)
		}

		for _, request := range parserResult.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func (e ConcurrentScheduler) CreateWorker(in chan Request, out chan ParserResult) {
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
