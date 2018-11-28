package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface{
	Submit(Request)
	WorkerChan()chan Request
	WorkerReady(chan Request)
	Run()
}


func (e *ConcurrentEngine)Run(seeds ...Request){



	////所有的worker共用一个输入输出
	//in := make(chan Request)
	//out := make(chan ParseRusult)
	//
	////配置输出chan到shceduler
	//e.Scheduler.ConfigureMustWorkChan(in)

	out := make(chan ParseRusult)
	e.Scheduler.Run()

	for i:=0; i<e.WorkerCount;i++  {
		createWorker(e.Scheduler.WorkerChan(),out,&e.Scheduler)
	}

	//将request注入scheduler
	for _,r :=range seeds  {
		e.Scheduler.Submit(r)
	}

	//接受out的数据
	for  {
		result := <-out


		for _,item :=range result.Items  {
			go func() {
				e.ItemChan <- item
			}()
		}

		for _,request := range result.Requests  {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}



}

func createWorker(in chan Request,out chan ParseRusult,s *Scheduler){

	//每一个worker都有自己的chan，用于针对自己的chan接受
	go func() {
		for{

			(*s).WorkerReady(in)
			request := <-in
			result,err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

var visitedUrls  = make(map[string]bool)

func isDuplicate(url string) bool{
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
