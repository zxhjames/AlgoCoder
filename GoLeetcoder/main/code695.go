package main

import (
	"fmt"
)

//岛屿最大面积
func main(){
	var grid = [][]int {
		{1,1,1,1,0},
		{1,1,0,1,0},
		{1,1,0,0,0},
		{0,0,0,0,1},
	}
	fmt.Println(maxAreaOfIsland(grid))
}


func maxAreaOfIsland(grid [][]int) int {
	vertical:=len(grid) //x
	horizontal:=len(grid[0]) //y
	count:=0
	for i:=0;i<vertical;i++ {
		for j:=0;j<horizontal;j++ {
			if grid[i][j]==1{
				t:= _dfs(grid,i,j,vertical,horizontal)
				if t > count{
					count = t
				}
			}
		}
	}
	return count
}

func _dfs(g [][]int,i,j,v,h int) int{
	if i<0 || i==v || j<0 || j==h || g[i][j]!=1 {
		return 0
	}
	g[i][j]=0
	count:=1
	count+=_dfs(g,i+1,j,v,h)
	count+=_dfs(g,i-1,j,v,h)
	count+=_dfs(g,i,j+1,v,h)
	count+=_dfs(g,i,j-1,v,h)
	return count
}