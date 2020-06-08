package main

import "sort"

//code575

func main(){

}

func distributeCandies(candies []int) int {
	m,l:=make(map[int]int),len(candies)/2
	sort.Ints(candies)
	for _,v := range candies {
		m[v]++
	}
	ml:=len(m)
	if ml <= l {
		return ml
	}else if ml > l{
		return l
	}
	return 0
}