package main

import (
	"fmt"
	"os"
)

//读取本地配置到内存

func readMze(filename string)[][]int{
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row,col int
	fmt.Fscanf(file,"%d %d",&row,&col)

	maze :=make([][]int, row)
	for i:= range maze {
		maze[i] = make([]int ,col)
		for j:=range maze[i] {
			fmt.Fscanf(file,"%d",&maze[i][j])
		}
	}
	return maze
}
func main() {
	maze := readMze("maze/maze.in")
	for _,row :=range maze {
		for _,val :=range row{
			fmt.Printf("%3d",val)
		}
		fmt.Println()
	}
}
