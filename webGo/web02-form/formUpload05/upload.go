package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)


//通过上面的实例我们可以看到我们上传文件主要三步处理：
//1：表单中增加enctype="multipart/form-data"
//2：服务端调用r.ParseMultipartForm,把上传的文件存储在内存和临时文件中
//3：使用r.FormFile获取文件句柄，然后对文件进行存储等处理。

func main() {
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		panic(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		t,err := template.ParseFiles("./web02-form/formUpload05/upload.gtpl")
		if err != nil {
			panic(err)
		}

		err = t.Execute(w,nil)
		if err != nil {
			panic(err)
		}

	}else{
		r.ParseMultipartForm(32<<20)
		file,handler,err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer  file.Close()

		fmt.Fprintln(w,"%v",handler.Header)

		PahtExists("./web02-form/formUpload05/uploadFile")

		f,err := os.OpenFile("./web02-form/formUpload05/uploadFile/"+handler.Filename,os.O_WRONLY|os.O_CREATE,0666)
		if err != nil {
			fmt.Println("ljy error :" ,err)
			return
		}

		defer  f.Close()
		io.Copy(f,file)
	}
}


func PahtExists(path string){
	_,err := os.Stat(path);
	if err != nil {
		err := os.Mkdir(path,os.ModePerm)
		if err != nil {
			fmt.Println("ljy error :creat " + path +"error")
		}
	}
}
