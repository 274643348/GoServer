package main

import (
	"fmt"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/scheduler"
	"learngo/GoServer/learngo/crawler/zhenai/parse"
	"learngo/GoServer/learngo/crawler_distributed/config"
	itemSever "learngo/GoServer/learngo/crawler_distributed/persist/Client"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	worker "learngo/GoServer/learngo/crawler_distributed/worker/Client"
	"net/rpc"
	"strings"
)


//var(
//	itemSaverHost = flag.String("itemsaver_host","","itemsaver host")
//
//	workerHosts = flag.String("worker_hosts","","worker hosts host")
//)
//终端输入：
//--***=":1234"  --***=":9000,:9001"
func main() {
	//flag.Parse()
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParseFunc:parse.ParseCityList,
	//})

	//测试
	//engine.Run(engine.Request{
	//	Url:"http://album.zhenai.com/u/1903652003",
	//	ParseFunc:parse.PraseProfile,
	//})

	itemsaver ,err:=itemSever.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}


	pool :=createClientPool(strings.Split(":9000,:9001",","))

	processor,err := worker.CreaterWorkerProcess(&pool)
	if err != nil {
		panic(err)
	}
	e :=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan:itemsaver,
		Request:processor,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		Parse:engine.NewFuncParser(parse.ParseCityList,config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun/zhengzhou",
	//	Parse:engine.NewFuncParser(parse.ParseCity,config.ParseCity),
	//})

}


func createClientPool(host []string)chan *rpc.Client{
	var clients []*rpc.Client
	fmt.Println(host);
	for _,h := range host{
		client ,err:= rpcsupport.NewClient(h)
		if err != nil {
			fmt.Printf("error connecting to %s:%v", h, err)
		}else
		{
			clients = append(clients,client)
		}
	}

	out := make(chan * rpc.Client)
	go func(){
		for {
			for _,client := range  clients{
				out <-client
			}
		}
	}()
	return out
}
