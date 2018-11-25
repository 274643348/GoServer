package parse

import (
	"../../moder"
	"fmt"
	"regexp"
)
import "../../engine"

const name =`<span class="nickName" [^>]+>([^<]+)</span>`
const other =`<div class="purple-btns" data-v-ff544c08><div class="m-btn purple" data-v-ff544c08>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`.+`+
	`[^>]+>[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`+
	`[^>]+>([^<]+)</div>`



//<div class="purple-btns" data-v-ff544c08=""><div class="m-btn purple" data-v-ff544c08="">离异</div><div class="m-btn purple" data-v-ff544c08="">45岁</div><div class="m-btn purple" data-v-ff544c08="">天蝎座(10.23-11.21)</div><div class="m-btn purple" data-v-ff544c08="">167cm</div><div class="m-btn purple" data-v-ff544c08="">69kg</div><div class="m-btn purple" data-v-ff544c08="">工作地:阿克苏阿克苏市</div><div class="m-btn purple" data-v-ff544c08="">月收入:5-8千</div><div class="m-btn purple" data-v-ff544c08="">服务业</div><div class="m-btn purple" data-v-ff544c08="">高中及以下</div></div>
//const age =`<div class="purple-btns" data-v-ff544c08="">
//<div class="m-btn purple" data-v-ff544c08="">离异</div>
//<div class="m-btn purple" data-v-ff544c08="">40岁</div>
//<div class="m-btn purple" data-v-ff544c08="">射手座(11.22-12.21)</div>
//<div class="m-btn purple" data-v-ff544c08="">165cm</div>
//<div class="m-btn purple" data-v-ff544c08="">67kg</div>
//<div class="m-btn purple" data-v-ff544c08="">工作地:阿克苏柯坪</div>
//<div class="m-btn purple" data-v-ff544c08="">月收入:5-8千</div>
//<div class="m-btn purple" data-v-ff544c08="">自由职业</div>
//<div class="m-btn purple" data-v-ff544c08="">大专</div>
//</div>`

func PraseProfile(contents []byte) engine.ParseRusult{
	defer func(){
		err := recover()
		fmt.Printf("PraseProfile error :%v\n",err)

	}()

	re :=regexp.MustCompile(name)
	matchs := re.FindSubmatch(contents)
	//fmt.Printf("wwwww%s",matchs)
	name:=string(matchs[1])

	re =regexp.MustCompile(other)
	matchs = re.FindSubmatch(contents)
	//fmt.Printf("wwwww%s",matchs)
	profile := moder.Profile{}
	if matchs != nil{

		age:=string(matchs[2])
		profile.Name = name
		profile.Age = age
		profile.Height = string(matchs[4])
		profile.Weight =string(matchs[5])
		profile.Income =string(matchs[7])
		profile.Marrige =string(matchs[14])
		profile.Occupation = string(matchs[15])
		//profile.
	}

	result := engine.ParseRusult{

		Items:[]interface{}{profile},
	}


	//fmt.Println(len(matchs))
	return result
}