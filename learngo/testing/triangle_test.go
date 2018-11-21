package main

import (
	"fmt"
	"testing"
	"./testfunc"
)

/*
单元测试：
1：方法名TextXXX,xxx大写，
2：文件名必须以_test结尾

命令行go test  .
运行目录下所有单元测试:go test .(显示详情的话 -v)
测试制定文件中的指定单元测试go test -v -run TestA select_test.go
 */

//func main() {
//
//}

//前缀答谢TestX,X大写，参数是*texting.T
func TestTriangle(t *testing.T){
	tests :=[]struct{a,b,c int}{
		{3,4,5},
		{0,12,13},
		{0,15,17},
		{3000,4000,5000},
	}
	for _,tt:=range  tests{
		if actual:=testfunc.CalcTriangle(tt.a,tt.b) ;actual != tt.c{
			fmt.Printf("calcTringle(%d,%d) got %d; expected %d\n",tt.a,tt.b,actual,tt.c)
		}
	}
}

func TestTriangleB(t *testing.T){
	tests :=[]struct{a,b,c int}{
		{3,4,5},
		{0,12,13},
		{0,15,17},
		{3000,4000,5000},
	}
	for _,tt:=range  tests{
		if actual:=testfunc.CalcTriangle(tt.a,tt.b) ;actual != tt.c{
			fmt.Printf("calcTringle(%d,%d) got %d; expected %d\n",tt.a,tt.b,actual,tt.c)
		}
	}
}