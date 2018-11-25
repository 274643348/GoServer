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
func regexp2() {
	//re :=regexp.MustCompile(".+@.+\\..+")
	re :=regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)

}

func main() {
	//regexp1()
	regexp2()
}
