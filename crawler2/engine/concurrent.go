package engine

type ConcurrentScheduler struct {
	// 定义调度器
	Scheduler   Scheduler
	WorkerCount int
	// interface{} 表示任何数据都是可以的
	ItemChan chan interface{}
}

type Scheduler interface {
	Submit(request Request)
	// 会改变 Request 里面的内容 所以最好需要用 指针的变量
	ConfigureMasterRequest(chan Request)
	RequestChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentScheduler) Run(seeds ...Request) {
	out := make(chan ParserResult)

	//
	e.Scheduler.Run()
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for i := 0; i < e.WorkerCount; i++ {
		in := e.Scheduler.RequestChan()
		CreateWorker(in, out, e.Scheduler)
	}

	// 循环遍历所有的输出out
	for {
		// 得到 输出的request
		result := <-out
		// 打印items

		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()

		}
		// 将Requests 加入调度器中
		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

var cacheUrl = make(map[string]int)

func isDuplicate(url string) bool {

	if cacheUrl[url] == 1 {
		return true
	}
	cacheUrl[url] = 1

	return false

}

func CreateWorker(in chan Request, out chan ParserResult, s Scheduler) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parserResult, err := Worker(request)
			if err != nil {
				continue
			}
			out <- parserResult
		}

	}()
}
