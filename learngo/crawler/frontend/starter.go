package main


import (
	"learngo/GoServer/learngo/crawler/frontend/controller"
	"net/http"

)

func main() {

	//Age:(<30)
	//Age:([20 TO 30])

	//防止找不到css文件
	http.Handle("/", http.FileServer(http.Dir("/Users/liujingyan/go/src/learngo/GoServer/learngo/crawler/frontend/view")))

	http.Handle("/search", controller.CreateSearchResultHandler("/Users/liujingyan/go/src/learngo/GoServer/learngo/crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		panic(err)
	}
}
