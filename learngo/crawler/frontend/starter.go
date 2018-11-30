package main


import (
	"net/http"

	"./controller"
)

func main() {

	//http.Handle("/", http.FileServer(http.Dir("frontend/view")))

	http.Handle("/search", controller.SearchResultHandler{})

	err := http.ListenAndServe(":9999", nil)

	if err != nil {
		panic(err)
	}
}
