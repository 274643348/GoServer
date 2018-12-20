package main

import (
	"fmt"
)

func main() {
	errorCatch();
}

func errorCatch(){
	defer func() {
		fmt.Println("defer")
	}()

	fmt.Println("do some thing")

	//后进先出
	for i := 0; i < 5; i++ {
		defer fmt.Println(i);
	}
}
