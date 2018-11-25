package parse
import (
	"../../engine"
	"regexp"
)

//`<a href="http://album.zhenai.com/u/1445042275" target="_blank">筱静</a>`

const cityre =`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseRusult{
	re :=regexp.MustCompile(cityre)
	matchs := re.FindAllSubmatch(contents,-1)
	result := engine.ParseRusult{}
	for _,m :=range matchs{
		result.Items = append(result.Items,"User : "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParseFunc:engine.NilParser,
		})
		//fmt.Printf("City: %s,URL: %s\n",m[2],m[1])
	}

	//fmt.Println(len(matchs))
	return result
}
