package view

import (
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/frontend/model"
	"learngo/GoServer/learngo/crawler/moder"
	"os"
	"testing"
)

//测试template模版
//func TestTemplate(t *testing.T){
//
//	template := template.Must(template.ParseFiles("template.html"))
//
//	page :=model.SearchResult{}
//		page.Hits = 123
//		item := engine.Item{
//			Url:  "http://album.zhenai.com/u/1029982807",
//			Type: "zhenai",
//			Id:   "1029982807",
//			Payload: moder.Profile{
//				Name:       "Lucy",
//				Gender:     "女",
//				Age:        22,
//				Height:     "170",
//				Weight:     "49",
//				Income:     "8001-12000元",
//				Marriage:   "未婚",
//				Education:  "大学本科",
//				Stature: "财务/申计",
//				Hukou:      "上海浦东新区",
//				Xingzuo:     "狮子座",
//				House:      "和家人同住",
//				Car:        "未购车",
//			},
//		}
//
//	for i := 0; i < 10; i++ {
//		page.Items = append(page.Items,item)
//	}
//
//	out,err :=os.Create("template.test.html")
//	err =template.Execute(out,page)
//	if err != nil {
//		panic(err)
//	}
//}


//测试view接口
func TestSearchResultView_Render(t *testing.T) {

	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

		page :=model.SearchResult{}
			page.Hits = 123
			item := engine.Item{
				Url:  "http://album.zhenai.com/u/1029982807",
				Type: "zhenai",
				Id:   "1029982807",
				Payload: moder.Profile{
					Name:       "Lucy",
					Gender:     "男",
					Age:        22,
					Height:     "170",
					Weight:     "49",
					Income:     "8001-12000元",
					Marriage:   "未婚",
					Education:  "大学本科",
					Stature: "财务/申计",
					Hukou:      "上海浦东新区",
					Xingzuo:     "狮子座",
					House:      "和家人同住",
					Car:        "未购车",
				},
			}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)

	if err != nil {
		panic(err)
	}

}
