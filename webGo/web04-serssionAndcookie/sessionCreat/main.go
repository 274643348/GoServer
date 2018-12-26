package main

import (
	"fmt"
	"html/template"
	"learngo/GoServer/webGo/web04-serssionAndcookie/sessionCreat/sessionManager"
	"net/http"
)


func main() {
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("main main")
}
func login(w http.ResponseWriter, r *http.Request) {
	sess := sessionManager.GlobalSessions.SessionStart(w, r)
	_ = r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		//w.Header().Set("Content-Type", "text/html")
		_ = t.Execute(w, sess.Get("username"))
	} else {
		_ = sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
func init(){
	fmt.Println("main init")
}
func init(){
	fmt.Println("main init2")
}