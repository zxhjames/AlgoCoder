package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(findBestValue([]int{2,3,5},10))
}


func findBestValue(arr []int, target int) int {
	//首先求一下和
	sum,l:=0,len(arr)
	sort.Ints(arr)
	for i:=0;i<l;i++ {
		sum+=arr[i]
	}
	if sum <= target{
		return arr[l-1]
	}
	tmp:=target/l
	for i:=0;i<l-1;i++{
		if arr[i] <= tmp && tmp <= arr[i+1] {
			return tmp
		}
	}
	return 0
}
