package main
//岛屿的周长
func main(){

}


func islandPerimeter(grid [][]int) int {
	vertical:=len(grid) //x
	horizontal:=len(grid[0]) //y
	for i:=0;i<vertical;i++ {
		for j:=0;j<horizontal;j++ {
			if grid[i][j]==1{
				return dfs2(grid,i,j,vertical,horizontal)
			}
		}
	}
	return 0
}


func dfs2(grid [][]int,i,j,v,h int) int{
	if i<0 || i==v || j<0 || j==h || grid[i][j]==0 {
		return 1
	}
	grid[i][j] = 2
	return dfs2(grid,i,j-1,v,h)+dfs2(grid,i-1,j,v,h)+dfs2(grid,i,j+1,v,h)+dfs2(grid,i+1,j,v,h)
}
