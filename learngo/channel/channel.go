package main

import (
	"fmt"
	"time"
)

//func chanDemo1(){
//	var c chan int//c =nil  暂时没法用，后期select会用到
//}


//func chanDemo2(){
//	//创造出一个chan
//	c :=make(chan int)
//
//	//chan<-发送数据给chan
//	c<-1
//	c<-2
//
//	//<-chan接受数据给chan
//	n := <-c
//	fmt.Println(n)
//
//	//运行报错fatal error: all goroutines are asleep - deadlock!
//	//在c<-1时会死锁，因为没有goroutine接收，后期通过buffered channel来处理
//}

//创造chan实现接受
//func chanDemo3(){
//		//创造出一个chan
//		c :=make(chan int)
//
//		go func(){
//			for {
//				n:= <-c
//				fmt.Println(n)
//			}
//
//			c<-4
//		}()
//		//chan<-发送数据给chan
//		c<-1
//		c<-2
//
//		//防止main结束，程序被杀掉
//		time.Sleep(time.Millisecond)
//	}

//多个chan的使用
//func work(id int ,c chan int){
//		for {
//			fmt.Printf("worker %d received %c\n",id,<-c)
//		}
//}
//
//func chanDemo4(){
//	//创造出多个chan
//	c :=[10]chan int{}
//
//	for i:=0;i<10 ;i++  {
//		c[i] = make(chan int)
//		go work(i,c[i])
//	}
//
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'a'+i
//	}
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'A'+i
//	}
//
//	//防止main结束，程序被杀掉
//	time.Sleep(time.Millisecond)
//
//	//最终的输出是乱序的，但都会打出来
//}



//chan作为返回值
//func creatWork(id int) chan int{
//	c := make(chan int)
//
//	go func(){
//		for {
//			fmt.Printf("worker %d received %c\n",id,<-c)
//		}
//	}()
//	return c
//}
//
//func chanDemo5(){
//	//创造出多个chan
//	c :=[10]chan int{}
//
//	for i:=0;i<10 ;i++  {
//		c[i] = creatWork(i)
//	}
//
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'a'+i
//	}
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'A'+i
//	}
//
//	//防止main结束，程序被杀掉
//	time.Sleep(time.Millisecond)
//
//	//最终的输出是乱序的，但都会打出来
//}

////chan返回为<-外部只能写入
//func creatWork(id int) chan<- int{
//	c := make(chan int)
//
//	go func(){
//		for {
//			fmt.Printf("worker %d received %c\n",id,<-c)
//		}
//	}()
//	return c
//}
//
//func chanDemo6(){
//	//创造出多个chan
//	c :=[10]chan<- int{}
//
//	for i:=0;i<10 ;i++  {
//		c[i] = creatWork(i)
//	}
//
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'a'+i
//		n := <-c[i]
//	}
//	for i:=0;i<10 ;i++  {
//		c[i]<- 'A'+i
//	}
//
//	//防止main结束，程序被杀掉
//	time.Sleep(time.Millisecond)
//
//	//最终的输出是乱序的，但都会打出来
//}

//第二个参数为bufferedChannle，缓冲允许传入三次
//func bufferedChannle(){
//	c := make(chan int,3)
//
//	c<-1
//	c<-2
//	c<-3
//	c<-4
//}

//bufferedChannle，不会影响chan的传递
//func work(id int ,c chan int)  {
//	for {
//		fmt.Printf("worker %d received %c\n",id ,<-c)
//	}
//}
//func bufferedChannle2(){
//	c := make(chan int,3)
//	go work(0,c)
//	c<- 'a'
//	c<- 'b'
//	c<- 'c'
//	c<- 'd'
//	time.Sleep(time.Millisecond)
//}


//关闭chan，chan还是能收到的0
func work(id int ,c chan int) {

	//关闭后，还是会受到<-c,int--0 char--长方体 string--""
	for {
		//通过ok,来判断是否c关闭
		fmt.Printf("worker %d received %d\n", id, <-c)
	}

	//方法一：判断chan是否close
	//for {
	//	//通过ok,来判断是否c关闭
	//	n,ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %d\n", id, n)
	//}

	//方法二：判断chan是否close
	//for n := range c{
	//	//通过ok,来判断是否c关闭
	//	fmt.Printf("worker %d received %d\n", id, n)
	//}
}
func channelClose(){
	c := make(chan int)
	go work(0,c)
	c<-'a'
	c<-'b'
	c<-0
	c<-'d'
	close(c)
	time.Sleep(time.Millisecond)
}


func main() {
	channelClose()
}
