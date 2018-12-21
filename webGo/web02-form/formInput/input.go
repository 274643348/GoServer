package main

import (
	"fmt"
	"html/template"
	"net/http"
)


func main() {

	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		panic(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:",r.Method)
	if r.Method == "GET" {
		t,err:= template.ParseFiles("./web02-form/formInput//login.gtpl")////////注意文件路径
		if err != nil {
			panic(err)
			return
		}

		err =t.Execute(w,nil)
		if err != nil {
			panic(err)
			return
		}
	}else {
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
	}
}
