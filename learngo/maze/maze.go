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

type point struct{
	i,j int
}

var dirs = [4]point{
	{-1,0},
	{0,-1},
	{1,0},
	{0,1},

}

func(p point)add(r point)point{
	return point{p.i+r.i,p.j+r.j}
}

//二维数组中某个左边的值
func(p point)at(grid [][]int)(int,bool){
	if(p.i<0 || p.i >=len(grid)) {
		return 0,false
	}
	if(p.j<0 || p.j >=len(grid[0])) {
		return 0,false
	}
	return grid[p.i][p.j],true
}

func walk(maze [][]int ,start,end point)[][]int{

	//创建steps二维数组存放，走过的路径
	steps := make([][]int, len(maze))
	for i:=range steps{
		steps[i] = make([]int ,len(maze[i]))
	}

	//存放探索数据的队列
	Q := []point{start}

	//开始探索队列
	for len(Q)>0 {
		cur := Q[0]
		Q = Q[1:]

		//发现终点-退出
		if cur == end{
			break
		}

		//判断周围的点是否满足加入队列的要求
		for _, dir := range dirs {
			next := cur.add(dir)

			//一：next的值为0；
			//二：steps的值也为0
			//三：next不为起点start
			//满足的话才能走
			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}

			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			if next == start {
				continue

			}

			//next的steps = 当前探索点的值 + 1
			curSteps,_ := cur.at(steps)
			steps[next.i][next.j] = curSteps +1

			//追加到探索队列
			Q = append(Q,next)

			}
		}

	return steps

}
func main() {
	maze := readMze("maze/maze.in")
	//for _,row :=range maze {
	//	for _,val :=range row{
	//		fmt.Printf("%3d",val)
	//	}
	//	fmt.Println()
	//}

	//走迷宫
	steps := walk(maze,point{0,0},point{len(maze) - 1,len(maze[0]) - 1})

	for _,row :=range  steps{
		for _,val := range row  {
			fmt.Printf("%3d",val)
		}
		fmt.Println()
	}
}
