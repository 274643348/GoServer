package main

import (

	"./fetcher"
	"fmt"
	"regexp"
)

func main() {
	all,_ :=fetcher.Fetcher("http://www.zhenai.com/zhenghun")
	printCityList(all)

}



//通过正则表达式获取城市和url列表
func printCityList(contents []byte){
	re :=regexp.MustCompile(`href=("http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+")[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(contents,-1)
	for _,m :=range matchs{
			fmt.Printf("City: %s,URL: %s\n",m[2],m[1])
	}

	fmt.Println(len(matchs))
}