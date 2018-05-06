package scheduler

import "github.com/zhongliangshan/test/crawler/engine"

type Scheduler struct {
	WorkerChan chan engine.Request
}

// 为了导致程序卡死这边需要异步去赋值
func (s *Scheduler) Submit(request engine.Request) {
	go func() { s.WorkerChan <- request }()
}

func (s *Scheduler) ConfigureMasterRequest(r chan engine.Request) {
	s.WorkerChan = r
}
