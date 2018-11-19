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
		err:=handler(writer,request)
		if err!=nil {

			log.Printf("Error occurred handling request: %s",err.Error())
			code:=http.StatusOK
			switch {
				case os.IsNotExist(err):{
					code =http.StatusNotFound
				}
				case os.IsPermission(err):{
					code =http.StatusForbidden
				}
				default:{
					code = http.StatusInternalServerError
				}
			}
			http.Error(writer,http.StatusText(code),code)
			
		}
	}
}



//将业务逻辑提出来
func filelistingserver2(){

	http.HandleFunc("/liu/", errWrapper(filelisting.HandlerFileList))
	err :=http.ListenAndServe(":8888",nil)
	if  err != nil{
		panic(err)
	}
}
func main() {
	filelistingserver2()
}
