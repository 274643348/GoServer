package scheduler

import "../engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (e *QueuedScheduler) Submit(r engine.Request) {
	e.requestChan<-r
}

func (e *QueuedScheduler) WorkerReady(w chan engine.Request){
	e.workerChan <-w
}

func (e *QueuedScheduler) WorkerChan( )chan engine.Request {
	return make(chan engine.Request)
}

func (e *QueuedScheduler) Run(){

	//scheduler中定义自己的request和worker的chan，用于接受外界的数据
	//在自己的goroutine中接受上面的两个chan数据，放在队列中
	//当两个队列都存在时，赋值后，非nil就会在select中运行，然后将两个数据移除队列

	//创建chan
	e.workerChan = make(chan chan engine.Request)
	e.requestChan= make(chan engine.Request)

	go func() {
		//存放request和worker
		var requestQ []engine.Request
		var workerQ [] chan engine.Request

		for  {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			//满足条件，即可在select中向选中的activeWorker发送activeRequest
			if len(requestQ) >=1 && len(workerQ) >=1 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {

				case r :=<-e.requestChan:{
					requestQ = append(requestQ, r)
				}
				case w :=<-e.workerChan:{

					workerQ = append(workerQ, w)
				}
				case activeWorker <- activeRequest:
					requestQ = requestQ[1:]
					workerQ = workerQ[1:]

			}
		}
	}()
}


