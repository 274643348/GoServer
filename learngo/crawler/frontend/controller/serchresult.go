package controller

import (
	"golang.org/x/net/context"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/frontend/model"
	"learngo/GoServer/learngo/crawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//创建resultHandler，完成view的对象，和elastic的对象；
func CreateSearchResultHandler(template string) SearchResultHandler {

	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}

}


func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}


	////将格式化字符串写入w的io.write
	//fmt.Fprintf(w,"q=%s,from=%d",q,from)

	var page model.SearchResult
	page, err = h.getSearchResult(q, from)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}


//通过search中的q和from获取到model数据
func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {

	var result model.SearchResult
	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())


	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	if result.PrevFrom <=0 {
		result.PrevFrom=0
	}

	result.NextFrom = result.Start + len(result.Items)

	return result, nil

}
