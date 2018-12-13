package main

import (
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/scheduler"
	"learngo/GoServer/learngo/crawler/zhenai/parse"
	"learngo/GoServer/learngo/crawler_distributed/persist/Client"
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

	itemsaver ,err:=Client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	e :=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemsaver,
	}
	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parse.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/zhengzhou",
		Parse:engine.NewFuncParser(parse.ParseCity,"ParseCity"),
	})

}
