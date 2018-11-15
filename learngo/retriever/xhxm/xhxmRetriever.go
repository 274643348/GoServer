package xhxm

import "fmt"

type XhxmRetriever struct {
	Contents string
}

func (r XhxmRetriever) Get(url string) string{
	return r.Contents
}

func (r XhxmRetriever) ShowName(){
	fmt.Println("xhxmRetriever-----showname")
}