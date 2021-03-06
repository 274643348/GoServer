package persist

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/engine"
	"log"
)

func ItemSaver(Index string)(chan engine.Item,error){
	out := make(chan engine.Item)

	client,err := elastic.NewClient(
		//默认寻找服务器
		//elastic.SetURL()

		//客户端维护集群状态的，但是集群不跑在本机上而是跑在docker中，所以必须false
		elastic.SetSniff(false))

	if err != nil{
		return nil,err
	}

	go func() {
		itemCount := 0
		for true {
			item :=<-out
			fmt.Printf("Got item #%d %v\n",itemCount,item)
			itemCount ++
			err := Save(client,Index,item)
			if err != nil {
				log.Printf("Item saver: error Save item %v:%v",item,err);
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