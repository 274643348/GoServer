package main

/*
通过time.sleep来确保协程中将channel传递的信息输出,很牵强,怎么才能等待goroutine结束呢?
1:channel的输出是阻塞的,
 */

import (
	"fmt"
	"sync"
)

//type worker struct{
//	in chan int
//	done chan bool
//}
//func createWorker(id int )worker{
//	w :=worker{
//		in:make(chan int),
//		done:make(chan bool),
//	}
//	go doworker(id,w)
//	return w
//}


////发一个任务接受一个任务，顺序执行，没有并行，没有意义
//func doworker(id int ,w worker){
//	for n:= range w.in {
//		fmt.Printf("Worker %d receive %c\n",id,n)
//		//通过新的goroutine，并行发送；
//		w.done<-true
//	}
//}
//func chanDemo1(){
//	 channels :=[10]worker{}
//
//	for i:=0;i<10;i++{
//		channels[i] = createWorker(i)
//	}
//
//	for i:=0;i<10;i++{
//		channels[i].in<- 'a'+i
//		//发送一个任务接受一个接受一个任务再继续，但是是顺序执行，没有并行的运行
//		<-channels[i].done
//	}
//
//	for i:=0;i<10;i++{
//		channels[i].in<- 'A'+i
//		//发送一个接受一个
//		<-channels[i].done
//	}
//	//time.Sleep(time.Millisecond)
//}

//等待全部结束的问题，和方案一
//func doworker(id int ,w worker){
//	for n:= range w.in {
//		fmt.Printf("Worker %d receive %c\n",id,n)
//
//		//报错，同一个协程中channel是阻塞的
//		//w.done<-true
//
//
//		//解决方案一：通过新的goroutine，并行发送；
//		go  func(){
//			w.done<-true
//		}()
//	}
//}
//
//func chanDemo1(){
//	channels :=[10]worker{}
//
//	for i:=0;i<10;i++{
//		channels[i] = createWorker(i)
//	}
//
//	for i:=0;i<10;i++{
//		channels[i].in<- 'a'+i
//	}
//
//	for i:=0;i<10;i++{
//		//小写字母会全部输出，这里将会报错，因为channel的发是阻塞式的；
//		//例如：上一个字母a发给了channel[0]后，channel[0].done<-true,发送了一个true，
//		//等待接受由于接受在下方的循环中，所以再次通过大写的A触发channel[0].done时发生阻塞
//		//报错deadlock
// 		channels[i].in<- 'A'+i
//	}
//
//	for _,worker :=range channels{
//		<-worker.done
//		<-worker.done
//	}
//}

//方案二：
//func doworker(id int ,w worker){
//	for n:= range w.in {
//		fmt.Printf("Worker %d receive %c\n",id,n)
//		w.done<-true
//	}
//}
//
//
//func chanDemo1(){
//	channels :=[10]worker{}
//
//	for i:=0;i<10;i++{
//		channels[i] = createWorker(i)
//	}
//
//	for i:=0;i<10;i++{
//		channels[i].in<- 'a'+i
//	}
//
//	for _,worker :=range channels{
//		<-worker.done
//	}
//	for i:=0;i<10;i++{
//		channels[i].in<- 'A'+i
//	}
//
//	for _,worker :=range channels{
//		<-worker.done
//	}
//}




//waitgroup---

//写法1：指针
//type worker_wait struct{
//	in chan int
//	wg *sync.WaitGroup
//}
//func createWorker_wait(id int ,wg *sync.WaitGroup)worker_wait{
//	w :=worker_wait{
//		in:make(chan int),
//		wg:wg,
//	}
//	go doworker_wait(id,w)
//	return w
//}
//
//func doworker_wait(id int ,w worker_wait){
//	for n:= range w.in {
//		fmt.Printf("Worker %d receive %c\n",id,n)
//		w.wg.Done()
//	}
//}
//写法2：函数是编程
type worker_wait struct{
	in chan int
	Done func()
}
func createWorker_wait(id int ,wg *sync.WaitGroup)worker_wait{
	w :=worker_wait{
		in:make(chan int),
		Done: func() {
			wg.Done()
		},
	}
	go doworker_wait(id,w)
	return w
}

func doworker_wait(id int ,w worker_wait){
	for n:= range w.in {
		fmt.Printf("Worker %d receive %c\n",id,n)
		w.Done()
	}
}

func chanDemo1(){
	wg :=sync.WaitGroup{}
	channels :=[10]worker_wait{}
	//20个任务
	wg.Add(20)
	for i:=0;i<10;i++{
		channels[i] = createWorker_wait(i,&wg)
	}

	for i:=0;i<10;i++{
		channels[i].in<- 'a'+i
	}

	for i:=0;i<10;i++{
		channels[i].in<- 'A'+i
	}

	//等待任务做完
	wg.Wait()
}

func main() {
	chanDemo1()
}
