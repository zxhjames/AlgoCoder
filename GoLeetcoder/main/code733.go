package main

import "fmt"

//图像渲染
func main(){
	var image = [][]int{
		{1,1,1},
		{1,1,0},
		{1,0,1},
	}
	fmt.Println(floodFill(image,1,1,2))
}
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	return dfs3(image,sr,sc,newColor,image[sr][sc])
}
func dfs3(image [][]int, sr int, sc int, newColor,oldColor int) [][]int {
	if sr < 0 || sr == len(image) || sc < 0 || sc == len(image[0]) || image[sr][sc] !=oldColor || image[sr][sc] == newColor {
		return image
	}
	image[sr][sc] = newColor
	dfs3(image,sr+1,sc,newColor,oldColor)
	dfs3(image,sr-1,sc,newColor,oldColor)
	dfs3(image,sr,sc+1,newColor,oldColor)
	dfs3(image,sr,sc-1,newColor,oldColor)
	return image
}

