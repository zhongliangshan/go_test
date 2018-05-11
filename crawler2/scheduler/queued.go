package scheduler

import "github.com/zhongliangshan/test/crawler2/engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WordkerChan chan chan engine.Request
}

func (q *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	q.WordkerChan <- worker
}

func (q *QueuedScheduler) WorderChan() chan engine.Request {
	return make(chan engine.Request)
}

// 创建两个消息队列 request 和 worker 通过消息队列进行传递
func (q *QueuedScheduler) Run() {
	// 创建两个消息的管道
	q.WordkerChan = make(chan chan engine.Request)
	q.RequestChan =make(chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue [] chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) >0 && len(workerQueue) >0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case request := <-q.RequestChan:
				requestQueue = append(requestQueue , request)
			case worker := <-q.WordkerChan:
				workerQueue = append(workerQueue , worker)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}

	}()
}

func (q *QueuedScheduler) Submit(request engine.Request) {
	q.RequestChan <-request
}


