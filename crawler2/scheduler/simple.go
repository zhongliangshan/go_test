package scheduler

import "github.com/zhongliangshan/test/crawler2/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterRequest(r chan engine.Request) {
	s.WorkerChan = r
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.WorkerChan <- r}()
}

