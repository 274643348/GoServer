package persist

import (
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler_distributed/persist/Client"
)

/*
	针对不同的程序，提供通用的client和server接口使用
 */

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func(s *ItemSaverService)Save(item engine.Item,result *string)error{
	err := Client.Save(s.Client,s.Index,item)
	fmt.Printf("Item %v saved.\n",item)

	if err == nil {
		*result = "OK"
		fmt.Printf("Item error : %v\n",err)
	}
	return err
}
