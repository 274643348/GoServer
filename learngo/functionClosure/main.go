package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//函数返回一个闭包函数
//闭包函数对于外层函数的sum具有存储功能；
//shift + fn + F6 整体替换
//alt + command + M独立一块代码
func Adder2()func (int)int{
	sum :=0
	return func (v int)int{
		sum +=v
		return sum
	}
}


func fibonacci() intGen{
	a,b:=0,1
	return func() int{
		a,b=b,a+b
		return  a
	}
}



func main() {
	//Closure()

	//fmt.Println("闭包的应用----斐波那契数列")
	//fibonacciText()

	fmt.Println("斐波那契数列---实现reader接口实现输出")
	readerFibonacci()

}

func Closure() {
	//获取闭包中的func
	fmt.Printf("第一个闭包函数对象------------\n")
	functionClosure := Adder2()
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Printf("\n第二个闭包函数对象------------\n")
	functionClosure2 := Adder2()
	fmt.Println(functionClosure2(1))
}

func fibonacciText() {
	f := fibonacci()
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
	println(f())
}

func readerFibonacci(){
	f := fibonacci()
	printFileContents(f)
}


type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	 next :=g()
	if next > 10000 {
		return 0, io.EOF
	}
	 s := fmt.Sprintf("%d\n",next)
	 return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


