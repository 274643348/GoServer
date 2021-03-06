package parse

import (
	"fmt"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/moder"
	"learngo/GoServer/learngo/crawler_distributed/config"
	"regexp"
	"strconv"
)

const Name =`<span class="nickName" [^>]+>([^<]+)</span>`
const age =`<div class="m-btn purple" [^>]+>([^<]+)岁</div>`
const Gender =`<a href="http://www.zhenai.com/zhenghun[^>]+>[^<]+(男士|女士)征婚</a>`
//<div class="id" data-v-35c72236="">ID：1768325386</div>
const id = `<div class="id" [^>]+>ID：([^<]+)</div>`

const house = `<div class="m-btn pink" [^>]+>([^<]+房)</div>`
const car = `<div class="m-btn pink" [^>]+>([^<]+车)</div>`
const stature = `<div class="m-btn pink" [^>]+>(体型[^<]+)</div>`
//const other =`<div class="purple-btns" data-v-ff544c08><div class="m-btn purple" data-v-ff544c08>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`.+`+
//	`[^>]+>[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`+
//	`[^>]+>([^<]+)</div>`




func parseProfile(contents []byte,url string ,name string) engine.ParseRusult{
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("PraseProfile error :%v\n",err)
		}

	}()

	profile := moder.Profile{}

	profile.Name = extract(contents,Name)
	age ,err := strconv.Atoi(extract(contents,age))
	if err != nil {
		age = 0
	}

	profile.Age = age
	profile.Gender =extract(contents,Gender)
	profile.House =extract(contents,house)
	profile.Car =extract(contents,car)
	profile.Stature = extract(contents,stature)
	var Id =extract(contents,id)





	//re =regexp.MustCompile(other)
	//matchs = re.FindSubmatch(contents)
	////fmt.Printf("-----------------------wwwww%s\n",matchs)
	//
	//if matchs != nil{
	//
	//	age:=string(matchs[2])
	//	profile.Age = age
	//	profile.Height = string(matchs[4])
	//	profile.Weight =string(matchs[5])
	//	profile.Income =string(matchs[7])
	//	profile.Marriage =string(matchs[14])
	//	profile.Occupation = string(matchs[15])
	//	//profile.
	//}

	result := engine.ParseRusult{

		Items:[]engine.Item{
			{
				Url:"http://album.zhenai.com/u/"+Id,
				Id:Id,
				Type:"zhenai",
				Payload:profile,
			},
		},
	}


	//fmt.Println(len(matchs))
	return result
}

func extract(b []byte,text string) string {
	re :=regexp.MustCompile(text)
	matchs := re.FindSubmatch(b)
	//fmt.Printf("wwwww%s\n",matchs)
	if matchs != nil {
		return string(matchs[1])
	}
	return `"无"`
}



//func NewProfileParser(name string)engine.ParserFunc{
//	return func(contents []byte, url string) engine.ParseRusult {
//		return ParseProfile(contents,url,name)
//	}
//}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parser(contents []byte, url string) engine.ParseRusult {
	return parseProfile(contents,url,p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile ,p.userName
}

func NewProfileParser(name string) *ProfileParser{
	 return &ProfileParser{userName:name}
}
