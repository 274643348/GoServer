package main

import "fmt"

//函数类型的使用
type funcName func(int ,int)int

func main() {
	addNum := doFunc(1,2,addFunc);
	fmt.Println("1 add 2 ",addNum)

	subNum := doFunc(1,2,subFunc);
	fmt.Println("1 sub 2 ",subNum)

}


func doFunc(num1 int ,num2 int,name funcName)int{
	return  name(num1,num2)
}

func addFunc(num1 int ,num2 int)int{
	return num1 + num2
}

func subFunc(num1 int ,num2 int)int{
	return num1 - num2
}
