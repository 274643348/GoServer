package main

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler_distributed/persist"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
)

//包装itemsaverservice的服务；
func main() {
	err := serverRpc(":1234","dating_profile")
	if err != nil {
		panic(err)
	}

	//log.Fatal(serverRpc(":1234","dating_profile"))
}

func serverRpc(host,index string)error{
	client ,err:= elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host,&persist.ItemSaverService{
		Client:client,
		Index:index,
	})

}
