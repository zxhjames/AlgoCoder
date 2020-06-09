package main

import "fmt"

func main(){
	fmt.Println(uniqueOccurrences([]int{-3,0,1,-3,1,1,1,-3,10,0}))
}

func uniqueOccurrences(arr []int) bool {
	hm1,hm2:=make(map[int]int),make(map[int]int)
	for _,v:=range arr {
		hm1[v]++
	}
	for _,v:=range hm1 {
		hm2[v]++
	}
	fmt.Println(len(hm1),len(hm2))
	if len(hm2)!=len(hm1){
		return false
	}
	return true
}
