package main

import (
	"fmt"
	"learngo/GoServer/webGo/01/say"
)

var (
	x = 1;
	q = 2;
)

func main() {
	diffName.DiffName();

	fmt.Println(x,q);
	//变量使用1
	var a int;
	a = 10;
	fmt.Println(a);

	//变量使用2
	var b, c, d int;
	b = 10;
	c = 10;
	d = 10;
	fmt.Println(b, c, d);

	//变量使用3
	var e, f, g int = 10, 10, 10;
	fmt.Println(e, f, g);

	//简化版赋值(这种形式只能在函数中使用,在函数外部是不行)
	h, i, j := 10, 10, 10;
	fmt.Println(h, i, j);

	//神奇的"_"
	_, k := 10, 10;
	fmt.Println("第一个10将会被丢弃", k);

	//声明的变量必须被使用,否则编译报错
}
