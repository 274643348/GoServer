package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever2 struct {
	UserAgent string
	Timeout time.Duration
}

func (r *Retriever2) Get(url string) string {
	resp,err:=http.Get(url)
	if err!=nil {
		panic(err)
	}

	result,err:=httputil.DumpResponse(resp,true)
	if err != nil {
		panic(err)
	}
	return  string(result)
}

func (r Retriever2) ShowName2(){
	fmt.Println("Retriever2-----showname")
}