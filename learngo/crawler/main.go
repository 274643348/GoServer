package main

import (
	"./engine"
	"./zhenai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parse.ParseCityList,
	})

	//测试
	//engine.Run(engine.Request{
	//	Url:"http://album.zhenai.com/u/1903652003",
	//	ParseFunc:parse.PraseProfile,
	//})

}
