package scheduler

import "github.com/zhongliangshan/test/crawler2/engine"

<<<<<<< HEAD
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


=======
// 通过队列实现并发
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) RequestChan() chan engine.Request {
	return s.requestChan
}

func (*QueuedScheduler) ConfigureMasterRequest(chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

// 需要一个总控的 GOROUTINES
func (s *QueuedScheduler) Run() {
	// 在每一个go 线程中 创建自己的 request 和 worker channel
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var RequestQueued []engine.Request
		var WorkerQueued []chan engine.Request

		for {
			var activerRequest engine.Request
			var activerWorker chan engine.Request
			// 当 request队列和worker队列斗准备好了有值的时候，就将request 发送给 worker
			if len(RequestQueued) > 0 && len(WorkerQueued) > 0 {
				activerRequest = RequestQueued[0]
				activerWorker = WorkerQueued[0]
				// 不能再这里直接发送 不然会卡死的
				// 需要在select 进行调度
			}
			// 使用select 进行调度
			select {
			case r := <-s.requestChan:
				RequestQueued = append(RequestQueued, r)
			case w := <-s.workerChan:
				WorkerQueued = append(WorkerQueued, w)
			case activerWorker <- activerRequest:
				RequestQueued = RequestQueued[1:]
				WorkerQueued = WorkerQueued[1:]
			}
		}
	}()
}
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
