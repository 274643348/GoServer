package scheduler

import (
	"learngo/GoServer/learngo/crawler/engine"
)
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (e *SimpleScheduler) WorkerChan() chan engine.Request {
	return e.workerChan
}

func (e *SimpleScheduler) WorkerReady(r chan engine.Request) {
}

func (e *SimpleScheduler) Run() {
	e.workerChan = make(chan  engine.Request)
}

func ( e *SimpleScheduler) Submit(r engine.Request) {
	//存在无限等待的可能，当受到out的输入时，不断的向scheduler中submit，导致worker被占满，
	//然后继续向scheduler中submit，等待worker中in的接受，但是所有的worker都在工作，没有人接受，
	//于是最初的out接受数据的逻辑就等待在了submit
	//当任意一个worker中的逻辑做完后，调用out<-result时，由于上一个<-out在等待无法接受，所以out<-result也在等待，就这样爱你爱你。。。

	// e.workerChan<-r

	//方案一：将e.workerChan<-r丢到一个新的goroutine中，防止submit一直等待倒是双向等待；
	go func() {

		e.workerChan<-r
	}()
}

//配置输入chan到sheduler中
func (e *SimpleScheduler) ConfigureMustWorkChan(r chan engine.Request){
	e.workerChan = r
}
