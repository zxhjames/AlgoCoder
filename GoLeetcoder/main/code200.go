package main

import "fmt"

//岛屿数量
func main(){
	var grid = [][]byte {
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	vertical:=len(grid) //x
	horizontal:=len(grid[0]) //y
	count:=0
	for i:=0;i<vertical;i++ {
		for j:=0;j<horizontal;j++ {
			if grid[i][j]=='1'{
				DFS(grid,i,j,vertical,horizontal)
				count++
			}
		}
	}
	return count
}

func DFS(grid [][]byte,i,j,v,h int) (){
	//朝向四周遍历
	//var dir = [][]int{
	//	//上 下 左 右
	//	{-1,0},{1,0},{0,-1},{0,1},
	//}
	//for k:=0;k<4;k++{
	//	toX:=i+dir[k][0]
	//	toY:=j+dir[k][1]
	//	grid[i][j] = '#'
	//	if  toX<0 || toX==v || toY<0 || toY==h || grid[toX][toY] == '0' || grid[toX][toY] == '#'{
	//		continue
	//	}else{
	//		DFS(grid,toX,toY,v,h)
	//	}
	//}
	//return
	if i<0 || i==v || j<0 || j==h || grid[i][j]=='0' || grid[i][j]=='#' {
		return
	}
	grid[i][j]='#'
	DFS(grid,i+1,j,v,h)
	DFS(grid,i-1,j,v,h)
	DFS(grid,i,j+1,v,h)
	DFS(grid,i,j-1,v,h)

}