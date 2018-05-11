package scheduler

import "github.com/zhongliangshan/test/crawler/engine"

<<<<<<< HEAD
type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

// 将输入的chan 赋值到woker chan 中 然后 再将 request 添加到chan中
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.WorkerChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {s.WorkerChan <-r}()
}

=======
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
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
