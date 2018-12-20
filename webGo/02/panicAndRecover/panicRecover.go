package main

import "fmt"

//go没有像java那样的异常机制，他不能抛出异常，
// 而是用panic和recover机制

//panic是一个内建函数，可以中断原有的控制流程，
// 函数F调用panic，函数F的执行被中断，但是F的延迟函数会正常执行
//然后F返回到调用他的地方，在调用的地方，F的行为就像调用了panic这个过程继续向上
//直到发生panic的goroutine中所有调用的函数返回，此时程序退出


//Recover
//是一个内建函数，可以让一个进入panic的goroutine恢复过来。
//recover仅在延迟函数（defer）中有效。
//正常的执行过程中，调用recover会返回nil，并且没有其他任何效果。
//如果goroutine陷入恐慌，调用recover可以捕获到panic的输入，并恢复正常运行
func main() {
	fmt.Println("main start")
	panicTest()
	fmt.Println("main end")
}

func panicTest(){
	defer func() {
		panicData :=recover()
		fmt.Println("panicTest-----catch panic ----",panicData)
	}()
	fmt.Println("panicTest start")
	callPanic();
	fmt.Println("panicTest end")
}

func callPanic(){

	defer func() {
		panicData :=recover()
		fmt.Println("callPanic------catch panic and again----",panicData)
		panic(panicData)
	}()
	fmt.Println("callPanic start")
	panic("panic error")
	fmt.Println("callPanic end")
}
