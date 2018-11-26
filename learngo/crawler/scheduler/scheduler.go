package scheduler

import (
	"../engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func ( e *SimpleScheduler) Submit(r engine.Request) {

	e.workerChan<-r
}

//配置输入chan到sheduler中
func (e *SimpleScheduler) ConfigureMustWorkChan(r chan engine.Request){
	e.workerChan = r
}
