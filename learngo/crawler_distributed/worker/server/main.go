package server

import (
	"fmt"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"learngo/GoServer/learngo/crawler_distributed/worker"
)

func main(){
	rpcsupport.ServeRpc(fmt.Sprintf(":%d",config.WorkerPort0),worker.CrawlService{})
}

