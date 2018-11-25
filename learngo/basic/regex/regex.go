package main

import (
	"fmt"
	"regexp"
)

const text = "my email is 274643348@qq.com"


//** ***@***.***
func regexp1() {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`.+@.+\..+`)
	match := re.FindString(text)
	fmt.Println(match)

}

//** ***@***.***---aA0@aA0.aA0
func regexp2(text string) {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)

}

//如果存在多个匹配项，FindAllString
const text2 = "my email is 274643348@qq.com" +
	"email is 1893916095t@qq.com" +
	"email is 274643348@gmail.com"

func regexp3(test string) {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(test,-1)
	fmt.Println(match)

}


const text3 = "my email is 274643348@qq.com" +
	"email is 1893916095t@qq.com" +
	"email is 274643348@gmail.com.cn"
//aA0@aA0.aA0,@和最后一个.aA0之间出现.可以
func regexp4(test string) {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(test,-1)
	fmt.Println(match)

}


func regexp_submatch(test string)[][]string {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(test,-1)
	fmt.Println(match)
	return match

}

func main() {
	//regexp1()
	//regexp2(text2)
	//regexp3(text2)

	//regexp3(text3)
	//regexp4(text3)

	match :=regexp_submatch(text3)
	for _,m:=range match{
		fmt.Printf("%s  %s  %s\n",m[1],m[2],m[3])
	}
}
