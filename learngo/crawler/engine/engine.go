package engine

import "log"
import "../fetcher"
func Run(seeds ...Request){
	var requests []Request
	for _,r:=range seeds{
		requests =append(requests, r)
	}

	for len(requests) > 0{
		//获取第一个request
		r:=requests[0]
		requests = requests[1:]

		//获取目标中Url的body
		log.Printf("Fetching %s",r.Url)
		body,err :=fetcher.Fetcher(r.Url)
		if err != nil {
			log.Printf("Fetch: error " +
				"fetching  url %s : %v",r.Url,err)
			continue
		}

		//
		parseResult := r.ParseFunc(body)
		requests = append(requests,parseResult.Requests...)

		for _,item := range parseResult.Items{
			log.Printf("Get item is %v",item)
		}
		log.Printf("--------%d-------\n\n\n",len(parseResult.Items))


	}
}