package engine

import "log"

<<<<<<< HEAD
// 定义调度器和线程个数
type ConcurrentScheduler struct {
	Scheduler Scheduler
=======
type ConcurrentScheduler struct {
	Scheduler   Scheduler
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	WorkerCount int
}

type Scheduler interface {
<<<<<<< HEAD
	Submit(r Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (s *ConcurrentScheduler) Run(seeds ...Request) {
	// 创建输入和输出的管道
	in := make(chan Request)
	out := make(chan ParserResult)
	s.Scheduler.ConfigureMasterWorkerChan(in)
	for _ ,r := range seeds {
		s.Scheduler.Submit(r)
	}

	for i :=0 ; i< s.WorkerCount ; i++ {
		CreateWorker(in , out)
	}

	for {
		result := <-out
		for _ , item := range result.Items {
			log.Printf("got item %s" , item)
		}

		for _ , request := range result.Requests {
			s.Scheduler.Submit(request)
=======
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
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
		}

	}
}

<<<<<<< HEAD
func CreateWorker(in chan Request , out chan ParserResult) {
	go func() {
		for {
			request := <- in
=======
func (e ConcurrentScheduler) CreateWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
			parserResult, err := Worker(request)
			if err != nil {
				continue
			}
<<<<<<< HEAD

			out <- parserResult
		}
	}()
}
=======
			out <- parserResult
		}
	}()
}
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
