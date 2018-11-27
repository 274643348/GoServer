package parse

import (
	"../../engine"
	"regexp"
)

const cityListre =`href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseRusult{
	re :=regexp.MustCompile(cityListre)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseRusult{}
	for _,m :=range matchs{
		//result.Items = append(result.Items,"City : " + string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:ParseCity,
		})
		//fmt.Printf("City: %s,URL: %s\n",m[2],m[1])
		//
		//break
	}

	//fmt.Println(len(matchs))
	return result
}