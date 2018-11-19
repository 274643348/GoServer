package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"./filelisting"
)


//第一版文件web服务器
func filelistingserver1(){
	http.HandleFunc("/liu/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/liu/"):]
		file,err:=os.Open(path)
		if err!=nil {
			////服务器没有宕机,直接异常，客户端很难看
			//panic(err)

			//显示错误信息给客户端(用户看到错误信息，不好)
			http.Error(writer,err.Error(),http.StatusInternalServerError)
			return
		}
		defer file.Close()

		all,err:=ioutil.ReadAll(file)
		if err !=nil {
			panic(err)
		}

		writer.Write(all)

	})

	err :=http.ListenAndServe(":8888",nil)
	if err!=nil {
		panic(err)
	}
}


//第二版剥离业务逻辑到其他文件
type appHandler func(writer http.ResponseWriter, request *http.Request)error

func errWrapper(handler appHandler)func(writer http.ResponseWriter, request *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		//自己处理panic
		defer func(){
			r:=recover()
			if r==nil {
				return
			}
			log.Printf("Panic:%v",r)
			http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)

		}()


		err:=handler(writer,request)
		if err!=nil {

			log.Printf("Error occurred handling request: %s",err.Error())

			if userError,ok:=err.(userError);ok {
				http.Error(writer,userError.Message(),http.StatusBadRequest)
				return
			}

			code:=http.StatusOK
			switch {
				case os.IsNotExist(err):{
					code =http.StatusNotFound
				}
				case os.IsPermission(err):{
					code =http.StatusForbidden
				}
				default:{
					//有些是可以给用户看的怎么处理
					code = http.StatusInternalServerError
				}
			}
			http.Error(writer,http.StatusText(code),code)
			
		}
	}
}



//将业务逻辑提出来
func filelistingserver2(){

	//模拟接受所有用户---浏览器访问http://localhost:8888/abc---abc小于"/liu/"所以内部会报错
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err :=http.ListenAndServe(":8888",nil)
	if  err != nil{
		panic(err)
	}
}

//定义一个自己的error
type userError interface {
	error
	Message()string
}
func main() {
	filelistingserver2()
}
