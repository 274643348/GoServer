package main

import "fmt"

//函数返回一个闭包函数
//闭包函数对于外层函数的sum具有存储功能；
func Adder()func (int)int{
	sum :=0
	return func (v int)int{
		sum +=v
		return sum
	}
}

func main() {
	//获取闭包中的func
	fmt.Printf("第一个闭包函数对象------------\n")
	functionClosure :=Adder()
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))
	fmt.Println(functionClosure(1))

	fmt.Printf("\n第二个闭包函数对象------------\n")
	functionClosure2:=Adder()
	fmt.Println(functionClosure2(1))

}
