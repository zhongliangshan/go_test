package scheduler

import "github.com/zhongliangshan/test/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

// 将输入的chan 赋值到woker chan 中 然后 再将 request 添加到chan中
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.WorkerChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.WorkerChan <- r }()
}
