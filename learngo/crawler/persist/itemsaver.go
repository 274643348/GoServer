package persist

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/engine"
	"log"
)

func ItemSaver()(chan engine.Item,error){
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
			err :=save(client,item)
			if err != nil {
				log.Printf("Item saver: error save item %v:%v",item,err);
			}
		}
	}()

	return  out,nil

}

func save(client *elastic.Client,item engine.Item) error{

	if item.Type == "" {
		return  errors.New("itemsaver error: item.Type must have")
	}

	indexService := client.Index().
		Index("dating_profile").
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