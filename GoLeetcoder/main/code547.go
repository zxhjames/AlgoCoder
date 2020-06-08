package main

import "fmt"

//547朋友圈
func main(){
	var cmap = [][]int{
		{1,0,0,1},
		{0,1,1,0},
		{0,1,1,1},
		{1,0,1,1},
	}
	fmt.Println(findCircleNum(cmap))
}

//func findCircleNum(M [][]int) int {
//	res:=make([]int,len(M))
//	//初始化序列数组
//	for k:=0;k<len(M);k++{
//		res[k] = k
//	}
//	fmt.Println(res)
//	//一次探索
//	for i:=0;i<len(M);i++{
//		for j:=i+1;j<len(M[0]);j++{
//			//满足题设条件
//			if M[i][j] == 1{
//				MergeElem(res,i,j)
//				fmt.Println(res)
//			}
//		}
//	}
//	for k:=0;k<len(M);k++{
//		res[k] = FindRoot(res,res[k])
//	}
//	count:=0
//	tmp:=-1
//	for _,v:=range res {
//		if tmp!=v{
//			count++
//			tmp = v
//		}
//	}
//	return count
//}
//
//func MergeElem(res []int,i,j int){
//	a:=FindRoot(res,i)
//	b:=FindRoot(res,j)
//	if a == b {
//		return
//	}
//	res[a] = b
//}
//
//func FindRoot(res []int,k int) int{
//	if k == res[k] {
//		return k
//	}
//	return FindRoot(res,res[k])
//}


func findCircleNum(M [][]int) int {
	//初始化一个布尔数组与计数器
	tmp,res:=make([]bool,len(M)),0
	for i:=0;i<len(M);i++{
		if !tmp[i]{
			dfs4(M,i,tmp)
			res++
		}
	}
	return res
}

func dfs4(M[][]int,i int,tmp []bool){
	for j:=0;j<len(M);j++{
		if !tmp[j]&&M[i][j]==1{
			tmp[j] = true
			dfs4(M,j,tmp)
		}
	}
}