package main

import "fmt"

func main() {

	fmt.Println("======================基本赋值======================")
	//assignment to entry in nil map
	//var numbers map[string]int
	//numbers["one"] = 1
	//numbers["two"] = 2
	//numbers["three"] = 3

	//var numbers map[string]int = make(map[string]int)
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3

	for key,value :=range numbers{
		fmt.Println(key + "====" + fmt.Sprint(value))
	}

	fmt.Println("===================增删改查=========================")



	//直接初始化赋值
	numbers2:= map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}
	numbers2["four"] = 4
	numbers2["two"] = 0

	for key,value :=range numbers2{
		fmt.Println(key + "====" + fmt.Sprint(value))
	}

	fmt.Println("------------------delete----------------")
	delete(numbers2,"one")
	delete(numbers2,"one")
	delete(numbers2,"null")
	for key,value :=range numbers2{
		fmt.Println(key + "====" + fmt.Sprint(value))
	}

	fmt.Println("---------find--------")
	one ,ok := numbers2["one"]
	if ok == false {
		fmt.Println("----one not find----")
	}else {
		fmt.Println("----one : "+ fmt.Sprint(one) +"----")
	}

	fmt.Println(`======================和slice一样是"引用类型"======================`)

	copyNumber2 := numbers2
	fmt.Println(copyNumber2["two"])
	copyNumber2["two"] = 2
	fmt.Println(numbers2["two"])


}
