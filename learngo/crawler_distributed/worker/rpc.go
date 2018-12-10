package worker

import "learngo/GoServer/learngo/crawler/engine"

////////创建npc服务器
type CrawlService struct {

}

func (CrawlService) Process(req Request,result *ParseResult) error{
	engineReq,err:= DeserializeRequest(req)

	if err != nil {
		return  err
	}

	engineResult,err :=engine.Worker(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(engineResult)
	return nil
}
