package main

import "fmt"

//杨辉三角
func main(){
	fmt.Println(generate(5))
}

func generate(numRows int)[][] int{
	var matrix [][]int
	if numRows == 1 {
		return [][]int{{1}}
	}else if numRows == 2{
		return [][]int{{1},{1,1}}
	}else{
		matrix = [][]int{{1},{1,1}}
		for i:=3;i<=numRows;i++{
			var tmp  = []int{1}
			for j:=1;j<i-1;j++{
				tmp = append(tmp, matrix[i-2][j-1]+matrix[i-2][j])
			}
			tmp = append(tmp,1)
			matrix = append(matrix,tmp)
		}
	}
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 1)
	return matrix
}