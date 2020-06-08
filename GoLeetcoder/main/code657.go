package main

import "fmt"

//机器人能否回到原来的点

func main(){
	s:="UD"
	fmt.Println(judgeCircle(s))
}

func judgeCircle(moves string) bool {
	dir := [][]int{
		//up
		{0,0},{-1,0},{1,0},{0,-1},{0,1},
	}
	m:=map[int32]int{
		'U':1,
		'D':2,
		'L':3,
		'R':4,
	}
	for _,value:=range moves {
		dir[0][0]+=dir[m[value]][0]
		dir[0][1]+=dir[m[value]][1]
	}
	if dir[0][0] == 0 && dir[0][1] == 0 {
		return true
	}
	return false
}
