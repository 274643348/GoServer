package main

import (
	"fmt"
)

//panic
//停止当前函数执行
//一直想上返回，执行每一层的defer
//如果没遇到recover，程序退出

//recover
//仅在defer调用中使用
//获取panic的值
//如果无法处理，可充新panic

func tryRecover(){
	defer func(){
		r:=recover()
		if err,ok:=r.(error);ok {
			fmt.Println("Error occured:",err)
		}else{
			//继续往外抛
			panic(fmt.Sprintln("I don't know what to do",r))
		}
	}()

	//主动报错
	//panic(errors.New("this is an error"))

	//运行报错
	//b:=0
	//a :=5/b
	//fmt.Println(a)

	//非error类型错误
	panic(123)


}

func main() {
	tryRecover()
}
