package persist

import (
	"fmt"
)

func ItemSaver()chan interface{}{
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for true {
			item :=<-out
			fmt.Printf("Got item #%d %v\n",itemCount,item)
			itemCount ++
		}
	}()

	return  out

}
