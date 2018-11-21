package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	goroutine相当于Coroutine协程
	1：任何函数只要加来go，就交给调度器来执行；
	2：不需要在函数定义时确认是否为异步；
	3：调度器在合适的点进行切换；
	4：通过-race检测数据冲突；

	切换点(只是参考，不能保证切换，也不保证在其他地方不切换)：
	1：I/O，printf
	2：channel
	3：等待锁
	4：函数调用（有时）
	5：runtime.Gosched()


	Coroutine协程：
	1：轻量级"线程"；（线程开到百个就很不错了，协程直接上千）
	2：非抢占式多任务处理，由协程主动交出控制权；
	3：编译器/解释器/虚拟机层面的多任务；（编译器级别的多任务，go语言吧go func解释为线程，通过调度器来控制切换）
	4：多个协程可以在一个或都个线程上运行；（由go语言的调度器来控制）
 */


func normaFunc(){
	fmt.Printf("goroutine start \n")

	for i := 0; i < 10; i++ {
		func(){
			for{

				fmt.Printf("hellow from toroutine %d\n",i)
			}
		}()
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)
}

func goFunc(){
	fmt.Printf("goroutine start \n")

	for i := 0; i < 10; i++ {
		go func(){
			for{

				fmt.Printf("hellow from toroutine %d\n",i)
			}
		}()
	}
	fmt.Printf("goroutine end  \n")

}

func goFunc1(){
	fmt.Printf("goroutine start \n")

	for i := 0; i < 10; i++ {
		go func(){
			for{

				fmt.Printf("hellow from toroutine %d\n",i)
			}
		}()
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)
}

func goFunc2(){
	fmt.Printf("goroutine start \n")

	for i := 0; i < 10; i++ {
		go func(i int){
			for{

				fmt.Printf("hellow from toroutine %d\n",i)
			}
		}(i)
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)
}

func goCoroutine(){
	fmt.Printf("goroutine start \n")
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int){
			for{
				a[i]++
			}
		}(i)
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)
	fmt.Printf("goroutine end2  \n")

	fmt.Println(a)
}


func goCoroutine2(){
	fmt.Printf("goroutine start \n")
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(i int){
			for{
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)
	fmt.Printf("goroutine end2  \n")

	fmt.Println(a)
}


func goCoroutineErr(){
	fmt.Printf("goroutine start \n")
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(){
			for{
				a[i]++
				runtime.Gosched()
			}
		}()
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)

	fmt.Println(a)
}

func goCoroutineErr2(){
	fmt.Printf("goroutine start \n")
	a := [10]int{}
	for i := 0; i < 10; i++ {
		go func(ii int ){
			for{
				a[ii]++
				runtime.Gosched()
			}
		}(i)
	}
	fmt.Printf("goroutine end  \n")

	time.Sleep(time.Millisecond)

	fmt.Println(a)
}

func main() {
	//出现死循环
	//normaFunc()

	//main太快结束，只输出了start和end，所以没有输出goroutine中的输出
	//goFunc()

	//没有死循环，先输出end，在输出"hellow from。。。。。"，闭包i导致最终为10
	//goFunc1()

	//通过值传递保留当前的值，还是先输出end
	//goFunc2()

//-----------------非抢占式，只有协程主动交出

	//不交出控制权尝试,会在sleep这卡住，无法输出end2
	//goCoroutine()


	//通过runtime.Gosched(),交出控制权，在最后输出数组
	//goCoroutine2()


	//runtime error: index out of range,因为i闭包，最后为10，所以越界
	//goCoroutineErr()

	//通过go run XX.go 命令可以运行，在终端查看错误
	//通过go run -race XX.go 命令查看数据冲突
	//如下run没有问题，但是通过-race的话机会出现DATA RACE数据冲突，最后printr的时候read，但协程汇总存在write
	goCoroutineErr2()


}


