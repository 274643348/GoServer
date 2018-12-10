package worker

import (
	"fmt"
	"github.com/pkg/errors"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/zhenai/parse"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Aegs interface{}
}


////////////////////////////用于将方法改为可传递的字符串；
type Request struct {
	Url string
	Parser SerializedParser
}

type ParseResult struct {
	Items []engine.Item
	Request []Request
}



////////////////////////序列化///////////////////
func SerializeRequest(r engine.Request) Request{
	funcName,args :=r.Parse.Serialize()
	return Request{
		Url:r.Url,
		Parser:SerializedParser{
			Name:funcName,
			Aegs:args,
		},
	}
}

func SerializeParseResult(r engine.ParseRusult) ParseResult {
	result := ParseResult{
		Items:r.Items,
	}

	for _,req:= range r.Requests{
		result.Request = append(result.Request,SerializeRequest(req))
	}

	return result
}




////////////////////////反序列化///////////////////
func DeserializeRequest(r Request) (engine.Request,error){

	parser,err:=deserializeParser(r.Parser)

	if err != nil {
		return  engine.Request{},err
	}

	return engine.Request{
		Url:r.Url,
		Parse:parser,
	},nil
}

func DeserializeParseResult(r ParseResult) (engine.ParseRusult,error){
	result :=engine.ParseRusult{
		Items:r.Items,
	}

	for _,req:=range r.Request{

		enginReq,err :=DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing " + "request: %v",err)
		}
		result.Requests = append(result.Requests,enginReq)
	}
	return result,nil
}


func deserializeParser(p SerializedParser)(engine.Parse,error){
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parse.ParseCityList,config.ParseCityList),nil
	case config.ParseCity:
		return engine.NewFuncParser(parse.ParseCity,config.ParseCity),nil
	case config.NilParser:
		return engine.NilParser{},nil
	case config.ParseProfile:
		username,ok:=p.Aegs.(string)
		if ok {
			return parse.NewProfileParser(username),nil
		}else {
			return nil,fmt.Errorf("invalid " + "arg:%v",p.Aegs)
		}

	}
	return nil,errors.New("unknown parser name")
}

