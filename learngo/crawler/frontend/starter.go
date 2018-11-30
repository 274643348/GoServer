package main


import (
	"learngo/GoServer/learngo/crawler/frontend/controller"
	"net/http"

)

func main() {

	//Age:(<30)
	//Age:([20 TO 30])

	//防止找不到css文件
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./crawler/frontend/view"))))

	http.Handle("/search", controller.CreateSearchResultHandler("./crawler/frontend/view/template.html"))

	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		panic(err)
	}
}
