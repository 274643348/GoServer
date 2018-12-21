package main

import (
	"fmt"
	"net/http"
)
//http包建立web服务器
//go build httpServer.go 编译成可执行文件，httpserver
//然后就可以访问localhost:9000 进行http访问
//localhost:9000/?url_long=111&url_long=222 携带数据进行访问
func main() {
	http.HandleFunc("/",sayhellowName)/////////////////////设置访问路由
	err :=http.ListenAndServe(":9000",nil)///////////设置监听的端口
	if err != nil {
		panic(err)
	}
}

func sayhellowName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()///////////////////////////解析参数，默认是不会解析的
	fmt.Println(r.Form)/////////////////////打印表单数据
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v :=range r.Form{
		fmt.Println("key :",k)
		fmt.Println("val :",v)
	}

	fmt.Fprintln(w,"hello astaxie")//////写入w的是输出到客户端的
}
