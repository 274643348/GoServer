package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface{
	Submit(Request)
	ConfigureMustWorkChan(chan Request)
}


func (e *ConcurrentEngine)Run(seeds ...Request){



	//所有的worker共用一个输入输出
	in := make(chan Request)
	out := make(chan ParseRusult)

	//配置输出chan到shceduler
	e.Scheduler.ConfigureMustWorkChan(in)

	for i:=0; i<e.WorkerCount;i++  {
		createWorker(in,out)
	}

	//将request注入scheduler
	for _,r :=range seeds  {
		e.Scheduler.Submit(r)
	}

	//接受out的数据
	for  {
		result := <-out
		for _,item :=range result.Items  {
			fmt.Printf("Get item :%v\n",item)
		}

		for _,request := range result.Requests  {
			e.Scheduler.Submit(request)
		}
	}



}

func createWorker(in chan Request,out chan ParseRusult){
	go func() {
		for{
			request := <-in
			result,err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
