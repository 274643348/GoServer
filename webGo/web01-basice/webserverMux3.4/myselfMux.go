package main

import (
	"fmt"
	"net/http"
)
//自定义多级路由

type MyMux struct {

}
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHellowName(w,r)
		return
	}
	http.NotFound(w,r)
	return

}

func sayHellowName(w http.ResponseWriter, r *http.Request){
	r.ParseForm();


	fmt.Fprint(w,"hello myroute","\n")
	fmt.Fprint(w, "from :",r.Form,"\n")
	fmt.Fprint(w, "url :",r.URL,"\n")
	fmt.Fprint(w, "path :",r.URL.Path,"\n")
	fmt.Fprint(w, "scheme :",r.URL.Scheme,"\n")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090",mux)
}
