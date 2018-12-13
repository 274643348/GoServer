package Client

import (
	"fmt"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"learngo/GoServer/learngo/crawler_distributed/worker"
)

func CreaterWorkerProcess()(engine.Processor,error){
	client,err := rpcsupport.NewClient(fmt.Sprintf(":%d",config.WorkerPort0))

	if err != nil {
		return  nil,err
	}

	return func(req engine.Request) (engine.ParseRusult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		err := client.Call(config.CrawlServiceRpc,sReq,&sResult)

		if err != nil {
			return engine.ParseRusult{},err
		}

		return worker.DeserializeParseResult(sResult),nil


	},nil
}