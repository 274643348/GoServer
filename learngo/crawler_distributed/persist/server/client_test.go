package main

import (
	"fmt"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/moder"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

/*
	用于测试itemSaverService的服务端serverRpc（及./itemsaveRpcServer.go）的正确性
 */

func TestItemServer(t *testing.T){
	const host =":1234"
	//start itemSaverServer
	go serverRpc(host,"test1")

	time.Sleep(time.Second)
	//start itemSaverClient
	client ,err :=rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//Call save
	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1768325386",
		Type:"zhenai",
		Id:"1768325386",
		Payload:moder.Profile{
			Name :       "小红",
			Gender :     "女",
			Age :        11,
			Height :     "11",
			Weight :     "11",
			Income :     "11",
			Marriage:    "11",
			Education :  "11",
			Stature : "11",
			Hukou :      "11",
			Xingzuo :    "11",
			House :      "11",
			Car :        "11",
		},

	}
	result:=""
	err = client.Call("ItemSaverService.Save",expected,&result)
	if err != nil && result != "OK" {
		fmt.Printf("result:%s  err :%v",result,err)
	}
}
