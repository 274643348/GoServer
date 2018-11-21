package main

import (
	"fmt"
	"runtime"
	"time"
)

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


