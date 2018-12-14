package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "hellow";
	//s[0]='c';
	fmt.Println(s);

	//修改字符串
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c);
	fmt.Println(s2);

	//字符串链接
	x := "hellow"
	y := " world"
	xy := x + y

	fmt.Println(xy);

	//通过切片修改sss
	sss := "hellow"
	fmt.Println(sss);
	aaa := &sss
	sss = "c" + sss[1:]

	fmt.Println("切片的意义")
	fmt.Printf("原始的sss:%v，现在的sss:%v\n", *aaa, sss);

	//多行字符串
	ddd := `hello
world`
	fmt.Println(ddd)


	//错误类型
	 err := errors.New("this is an error")
	if err != nil {
		fmt.Println(err)
	}
}
