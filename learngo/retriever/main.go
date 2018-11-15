package main

import "fmt"
import "./xhxm"
import "./real"
type Retriever interface {
	Get(url string)string
}

func download(r Retriever)string{
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever

	r = xhxm.Retriever{"this is fake xhxm.com"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))
}
