package diffName

import "fmt"

func DiffName(){
	fmt.Print("包名可以跟文件名不同，但是一个文件下只能有一个包名（多个go文件，但package不一样）")
}
