package main

import (
	"bufio"
	"fmt"
	"os"
	"../fibonacci"
	"strconv"
)

//defer 相当于栈，先进吼出

//确保调用在函数结束时发生
//参数在defer语句时计算
//defer列表为后进先出


//何时调用defer
//Open/Close
//Lock/Unlock
//PritHeader/PrintFooter


func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func tyrDeferErr(){
	defer  fmt.Println(1)
	defer  fmt.Println(2)
	fmt.Println(3)
	panic("error occurresd")
	fmt.Println(4)
}


//文件测试defer
func tryDeferWriteFile(filename string){
	file,err:=os.Create(filename)
	if err != nil {
		panic(err)
	}
	//defer 关闭文件
	defer file.Close()

	writer := bufio.NewWriter(file)
	//defer 写入
	defer writer.Flush()

	f:=fibonacci.Fibonacci()
	for i:=0;i<20 ;i++  {
		fmt.Fprintln(writer,"aaaa-----"+strconv.Itoa(f()))
	}

}


//defer时计算
func tryDeferCount(){
	for i:=0;i<100 ;i++  {
		defer fmt.Println(i)
		if i==30 {
			panic("printed too many")
		}
	}
}

func main() {
	//tryDefer()
	//tyrDeferErr()
	//tryDeferWriteFile("liu.txt")
	tryDeferCount()
}

