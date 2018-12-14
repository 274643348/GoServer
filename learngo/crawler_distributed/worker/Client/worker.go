package Client

import (
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"learngo/GoServer/learngo/crawler_distributed/worker"
	"net/rpc"
)

func CreaterWorkerProcess(clientChan *chan * rpc.Client)(engine.Processor,error){
	//转为通道来处理；
	//client,err := rpcsupport.NewClient(fmt.Sprintf(":%d",config.WorkerPort0))
	//
	//if err != nil {
	//	return  nil,err
	//}

	return func(req engine.Request) (engine.ParseRusult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		c := <-*clientChan
		err := c.Call(config.CrawlServiceRpc,sReq,&sResult)

		if err != nil {
			return engine.ParseRusult{},err
		}

		return worker.DeserializeParseResult(sResult),nil


	},nil
}