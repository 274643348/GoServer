package main

import "fmt"
//当我们传入一个参数值到被调用函数里面计算时，时间上是串联这个值的一份copy，
// 当前被调用的函数中修改参数时，相应实参不会有任何变化；
func main() {
	num:=1
	addNum := addSelfFunc2(&num)

	fmt.Println(*addNum,"-------------",num);

	num+=num

	fmt.Println(*addNum,"-------------",num);

}

//传入指针返回copy的
func addSelfFunc(num *int) int{
	*num +=*num
	return *num
}

//传入指针返回指针
func addSelfFunc2(num *int) *int{
	*num +=*num
	return num
}
