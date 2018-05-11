package scheduler

import "github.com/zhongliangshan/test/crawler2/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
	panic("implement me")
}

func (s *SimpleScheduler) WorderChan() chan engine.Request {
	return  s.WorkerChan
}


func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.WorkerChan <- r}()
}

