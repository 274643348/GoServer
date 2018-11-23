package main

import (
	"fmt"
	"time"
)

/*
经测试发现select是阻塞的
1：如果case中的channel都都没有输入，则deadlock
2：如果有多个channel输入，随机任意一个
3：如果完成一个channel后重写for循环，再次select
 */

func sleep(seconds float32, endSignal chan<- bool) {
	time.Sleep(time.Duration(seconds) * time.Second)
	endSignal <- true
}

func main() {
	endSignal := make(chan bool, 1)
	go sleep(0.0001, endSignal)

	defaultnum,timenum:=0,0

	for{
		defaultnum++
		select {
		case <-endSignal:
			fmt.Println("The end!")
			//fmt.Println("default-num" ,defaultnum)
			//fmt.Println("time-after-num" ,timenum)
		case <-time.After(2 * time.Second):
			timenum++
			fmt.Println("time-0!" )
			//fmt.Println("default-num" ,defaultnum)
			//fmt.Println("time-after-num" ,timenum)
		case <-time.After(2 * time.Second):
			timenum++
			fmt.Println("time-1!" )
			//fmt.Println("default-num" ,defaultnum)
			//fmt.Println("time-after-num" ,timenum)
		case <-time.After(2 * time.Second):
			timenum++
			fmt.Println("time-2!" )
			//fmt.Println("default-num" ,defaultnum)
			//fmt.Println("time-after-num" ,timenum)
		case <-time.After(3 * time.Second):
			timenum++
			fmt.Println("time-3!" )
			//fmt.Println("default-num" ,defaultnum)
			//fmt.Println("time-after-num" ,timenum)
		//default :
		//	timenum++
		//	fmt.Println("default-num" ,defaultnum)
		}
	}
}
