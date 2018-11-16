package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	Iotil("liu.txt")
	//fmt.Printf("file operation")
	////_,err1 := os.Create("liu.txt")
	////if err1!=nil {
	////	fmt.Printf("err---%s",err1)
	////}
	//
	//file,err := os.Open("liu.txt")
	//if err!=nil {
	//	fmt.Printf("err---%s",err)
	//}
	//scanner :=bufio.NewScanner(file)
	//
	//for  scanner.Scan()  {
	//	println(scanner.Text())
	//}



}

func Iotil(name string){
	if contents,err:=ioutil.ReadFile(name);err ==nil {
		result :=strings.Replace(string(contents[:]),"\n","",1)
		fmt.Printf(result)
	}
}
