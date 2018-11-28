package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"learngo/GoServer/learngo/crawler/moder"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := moder.Profile{
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
	}
	id,err :=save(expected)
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
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v",resp)
	fmt.Printf("%s",*resp.Source)

	//反序列化
	var actual moder.Profile
	json.Unmarshal(*resp.Source,&actual)
	if expected != actual{
		 t.Errorf("expected %v/n,actual %v",expected,actual)
	}
}