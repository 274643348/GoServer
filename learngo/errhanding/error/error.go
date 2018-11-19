package main

import (
	"../fibonacci"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//文件测试err
func tryErrorWriteFile(filename string){
	//file,err:=os.Create(filename)
	file,err:=os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//自己的err（针对不同的err做不同的处理）
	//err = errors.New("this is a custom error")

	if err != nil {
		if pathError,ok :=err.(*os.PathError); !ok {
			//自己的err
			panic(err)
		}else{
			//已知的*patherror单独处理
			fmt.Println(pathError.Op,pathError.Path,pathError.Err)
		}
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
func main() {
	tryErrorWriteFile("liu.txt")
}
