package main

import "fmt"

func main() {
	var arr = [][]byte{
		{1,0,1,0,0},
		{1,0,1,1,1},
		{1,1,1,1,1},
		{1,0,0,1,0},
	}
	fmt.Println(maximalSquare(arr))
}

func maximalSquare(matrix [][]byte) int {
	var flag bool = true //循环标记
	var sqe int = 0 //记录最大面积
	x:=len(matrix) //x行
	if x == 0{
		return 0
	}
	y:=len(matrix[0]) //y列

	//依次遍历
	for i:=0;i<x;i++{
		for j:=0;j<y;j++{
			if matrix[i][j] !=0 {
				var step int = 1
				toRight:=j
				toDown:=i
				for {
					if !flag {
						break
					}else{
						if toRight+1 < x && toDown+1 < y{
							step++
							//遍历边缘的数据
							for p:=0;p<step;p++{
								if matrix[toDown+1][j+p] == 0 || matrix[toRight+1][i+p] == 0{
									break
								}
							}
							//否则面积加一，标记已经访问
							res:=step*step
							if res > sqe {
								sqe = res
							}
							//标记已经访问
							for p:=0;p<step;p++{
								matrix[toDown+1][j+p] = 0
								matrix[toRight+1][i+p] = 0
							}
							toRight++
							toDown++
						}else{
							res:=step*step
							if res > sqe {
								sqe = res
							}
							flag = false
						}

					}
				}
				flag = true
			}
		}
	}
	return sqe
}


