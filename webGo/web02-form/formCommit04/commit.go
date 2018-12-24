package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"
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
	if r.Method == "GET" {
		t,err := template.ParseFiles("./web02-form/formCommit04//login.gtpl")
		if err != nil {
			panic(err)
		}

		//生成md5码，确保表单的唯一性
		crutim := time.Now().UnixNano()//当前所在时区的时间戳（纳秒）
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(crutim,10))
		token := fmt.Sprintf("%x",h.Sum(nil))

		fmt.Println("token:",token ,"crutime:",crutim)
		t.Execute(w,token)
	}else{
		token := r.Form.Get("token")
		if token != "" {
			//验证tocken的合法性

			fmt.Println("token:",token)
			fmt.Fprintln(w,token)
		}else
		{
			panic("find token error")
		}
	}

}


