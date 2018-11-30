package view

import (
	"html/template"
	"io"
	"learngo/GoServer/learngo/crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

//获取filename的template模版
func CreateSearchResultView(filename string) SearchResultView {

	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}

}

//template执行data数据到w中
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {

	return s.template.Execute(w, data)

}
