package engine

type ConcurrentScheduler struct {
	// 定义调度器
	Scheduler   Scheduler
	WorkerCount int
	// interface 表示可以传递任意类型的值
	ItemChan chan Item
}

type Scheduler interface {
	WorderIntifery
	Submit(request Request)
<<<<<<< HEAD
	//// 会改变 Request 里面的内容 所以最好需要用 指针的变量
	//ConfigureMasterRequest(chan Request)

	WorderChan() chan Request
	Run()
}

type WorderIntifery interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentScheduler) Run (seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()
	in := e.Scheduler.WorderChan()
	for i :=0;i<e.WorkerCount;i++ {
		CreateWorker(in , out , e.Scheduler)
	}

	for _ ,r := range seeds {
		e.Scheduler.Submit(r)
=======
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
		e.Scheduler.Submit(r)
	}

	for i := 0; i < e.WorkerCount; i++ {
		in := e.Scheduler.RequestChan()
		CreateWorker(in, out, e.Scheduler)
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
	}

	// 循环遍历所有的输出out
	for {
		// 得到 输出的request
		result := <-out
		// 打印items

<<<<<<< HEAD
		for _ , item :=range result.Items {
			go func(i Item) {
				e.ItemChan <- i
			}(item)
=======
		for _, item := range result.Items {
			log.Printf("Got Item : %d  :%s", itemCount, item)
			itemCount++
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
		}
		// 将Requests 加入调度器中
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

<<<<<<< HEAD
func CreateWorker(in chan Request , out chan ParserResult , s WorderIntifery) {
=======
func CreateWorker(in chan Request, out chan ParserResult, s Scheduler) {
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
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
