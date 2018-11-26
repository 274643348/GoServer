package main

import (
	"./engine"
	"./zhenai/parse"
	"./scheduler"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parse.ParseCityList,
	//})

	//测试
	//engine.Run(engine.Request{
	//	Url:"http://album.zhenai.com/u/1903652003",
	//	ParseFunc:parse.PraseProfile,
	//})

	e :=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc:parse.ParseCityList,
	})

}
