package main

import (
	"fmt"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"learngo/GoServer/learngo/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawService(t *testing.T){
	const host = ":9000"
	go rpcsupport.ServeRpc(
		host,worker.CrawlService{})
	time.Sleep(time.Second)

	client,err :=rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1927253334",
		Parser:worker.SerializedParser{
			Name:config.ParseProfile,
			Aegs:"安静的学",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc,req,&result)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println(result)
	}

}
