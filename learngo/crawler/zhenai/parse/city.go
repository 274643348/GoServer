package parse
import (
	"../../engine"
	"regexp"
)

//`<a href="http://album.zhenai.com/u/1445042275" target="_blank">筱静</a>`

var(
	profileRe =regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`)
	cityRe =regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	)

func ParseCity(contents []byte) engine.ParseRusult{
	matchs := profileRe.FindAllSubmatch(contents,-1)
	result := engine.ParseRusult{}
	for _,m :=range matchs{
		//result.Items = append(result.Items,"User : "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:PraseProfile,
		})
		//fmt.Printf("City: %s,URL: %s\n",m[2],m[1])
	}

	matchs = cityRe.FindAllSubmatch(contents,-1)
	for _,m:=range matchs  {
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:ParseCity,
		})
	}
	//fmt.Println(len(matchs))
	return result
}
