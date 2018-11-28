package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/moder"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1768325386",
		Type:"zhenai",
		Id:"1768325386",
		Payload:moder.Profile{
			Name :"小红",
			Gender :"女",
			Age : "11",
			Height :"11",
			Weight :"11",
			Income :"11",
			Marrige :"11",
			Education :"11",
			Occupation :"11",
			Hukou :"11",
			Xingzuo :"11",
			House :"11",
			Car :"11",
		},

	}
	err :=save(expected)
	if err != nil {
		panic(err)
	}

	client,err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp,err := client.Get().
		Index("dating_profile").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v",resp)
	fmt.Printf("%s",*resp.Source)

	//反序列化
	var actual engine.Item
	json.Unmarshal(*resp.Source,&actual)

	actualProfile ,_ := moder.FromJsonObj(actual.Payload)

	actual.Payload = actualProfile

	if expected != actual{
		 t.Errorf("expected %v/n,actual %v",expected,actual)
	}
}