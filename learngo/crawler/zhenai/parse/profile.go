package parse

import (
	"fmt"
	"learngo/GoServer/learngo/crawler/engine"
	"learngo/GoServer/learngo/crawler/moder"
	"regexp"
)

const name =`<span class="nickName" [^>]+>([^<]+)</span>`
const age =`<div class="m-btn purple" [^>]+>([^<]+岁)</div>`
const Gender =`<a href="http://www.zhenai.com/zhenghun[^>]+>[^<]+(男士|女士)征婚</a>`
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




func PraseProfile(contents []byte) engine.ParseRusult{
	defer func(){
		err := recover()
		if err != nil {
			fmt.Printf("PraseProfile error :%v\n",err)
		}

	}()

	profile := moder.Profile{}
	re :=regexp.MustCompile(name)
	matchs := re.FindSubmatch(contents)
	//fmt.Printf("wwwww%s\n",matchs)
	if matchs != nil {
		profile.Name = string(matchs[1])
	}

	re =regexp.MustCompile(age)
	matchs = re.FindSubmatch(contents)
	//fmt.Printf("wwwww%s\n",matchs)
	if matchs != nil {
		profile.Age = string(matchs[1])
	}

	re =regexp.MustCompile(Gender)
	matchs = re.FindSubmatch(contents)
	//fmt.Printf("wwwww%s\n",matchs)
	if matchs != nil {
		profile.Gender = string(matchs[1])
	}


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
	//	profile.Marrige =string(matchs[14])
	//	profile.Occupation = string(matchs[15])
	//	//profile.
	//}

	result := engine.ParseRusult{

		Items:[]interface{}{profile},
	}


	//fmt.Println(len(matchs))
	return result
}