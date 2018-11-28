package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver()chan interface{}{
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for true {
			item :=<-out
			fmt.Printf("Got item #%d %v\n",itemCount,item)
			itemCount ++
			_,err :=save(item)
			if err != nil {
				log.Printf("Item saver: error save item %v:%v",item,err);
			}
		}
	}()

	return  out

}

func save(item interface{})(id string ,err error){
	client,err := elastic.NewClient(
		//默认寻找服务器
		//elastic.SetURL()

		//客户端维护集群状态的，但是集群不跑在本机上而是跑在docker中，所以必须false
		elastic.SetSniff(false))

	if err != nil{
		return "",err
	}

	resp,err := client.Index().
		Index("dating_profile").
		Type("zhenai").BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "",err
	}

	return resp.Id,nil
}