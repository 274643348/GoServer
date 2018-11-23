package main

import (
	"fmt"
	"time"
)

/*
经测试发现select是阻塞的
0：首先会执行所有的case XXX：中的XXX，等待看那个XXX县过来，
1：如果case中的channel都都没有输入，则deadlock（可以通过default来防止卡死）
2：如果有多个channel输入，随机任意一个
3：如果完成一个channel后重写for循环，再次select

select和time.after
1:在每一次select中都会运行time.after创建新的channel time
2:case <-time.after 和其他case相比，看谁谁快，如果同时，就select就随机选一个执行；
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
