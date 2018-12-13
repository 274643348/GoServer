package Client

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"learngo/GoServer/learngo/crawler_distributed/rpcsupport"
	"log"
)



import (
"context"
"github.com/pkg/errors"

)

func ItemSaver(host string)(chan engine.Item,error){
	client,err := rpcsupport.NewClient(host)
	if err != nil {
		return nil,err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for true {
			item :=<-out
			fmt.Printf("Got item #%d %v\n",itemCount,item)
			itemCount ++
			result:=""
			err := client.Call(config.ItemSaveRpc,item,&result)
			if err != nil {
				log.Printf("Item saver: error Save item %v:%v",item,err)
			}
		}
	}()

	return  out,nil

}

func Save(client *elastic.Client,Index string,item engine.Item) error{

	if item.Type == "" {
		return  errors.New("itemsaver error: item.Type must have")
	}

	indexService := client.Index().
		Index(Index).
		Type(item.Type).BodyJson(item)

	if item.Id != ""{
		indexService.Id(item.Id)
	}

	_ ,err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}