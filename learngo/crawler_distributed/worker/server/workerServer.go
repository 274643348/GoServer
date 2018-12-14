package main

import (
	"flag"
	"fmt"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"learngo/GoServer/learngo/crawler_distributed/worker"
)

//接受命令行参数
//go run workerServer.go  --port=9000
var port = flag.Int("port",0,"the port for me to listen on")

func main(){
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	fmt.Printf("workerServe port :%d",*port)
	rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{})
}

