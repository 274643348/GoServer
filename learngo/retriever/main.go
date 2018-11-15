package main

import (
	"./xhxm"
	"fmt"
	"time"
)
import "./real"
type Retriever interface {
	Get(url string)string
}

func download(r Retriever)string{
	return r.Get("http://www.baidu.com")
}


func main() {
	var r Retriever
	//retriever :=xhxm.XhxmRetriever{"this is fake xhxm.com"}
	//r = &retriever
	//fmt.Println(download(r))
	//
	//fmt.Printf("%T-----%v\n",r,r)
	//
	//retriever.Contents= "xhxm222"
	//
	//fmt.Printf("%T----%v\n",r,r)

	retriever :=xhxm.XhxmRetriever{"this is fake xhxm.com"}
	r = retriever
	fmt.Println(download(r))

	fmt.Printf("%T-----%v\n",r,r)

	retriever.Contents= "xhxm222"

	fmt.Printf("%T----%v\n",r,r)

	//r.ShowName2()//接口类型只能调用接口中定义的方法，无法调用引用这的自己的方法
	retriever.ShowName()





	realRetriever :=real.Retriever2{"xhxm",time.Minute}
	r = &realRetriever

	fmt.Printf("%T----%v\n",r,r)

	realRetriever.UserAgent = "xhxm222"

	fmt.Printf("%T----%v\n",r,r)

	//fmt.Println(download(r))

	//通过.(type)获取真实对象；
	switch v:=r.(type) {
	case *real.Retriever2:
		fmt.Printf("type---%T\n",r)
		v.ShowName2()

		//指针接管
		citeR :=v
		citeR.UserAgent = "xhxm333"
		fmt.Printf("type---%v\n",r)

		//指针接管
		copyR :=*v
		copyR.UserAgent = "xhxm444"
		fmt.Printf("type---%v\n",r)


	case xhxm.XhxmRetriever:
		fmt.Printf("type---%T",r)
		v.ShowName()

	}

	if citeR,ok :=r.(*real.Retriever2);ok  {
		fmt.Printf("*rel.Retrever2---%T---%s\n",citeR,citeR.UserAgent)
	}else
	{
		fmt.Println("not a *real.Retriever2")
	}

	if citeR,ok :=r.(xhxm.XhxmRetriever);ok  {
		fmt.Printf("xhxm.XhxmRetriever---%T---%s\n",citeR,citeR.Contents)
	}else
	{
		fmt.Println("not a xhxm.XhxmRetriever")
	}
}
