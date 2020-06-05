package main

import (
	"fmt"
	"sort"
)

//重叠区间 AC
func main () {
	//var arr  = [][]int{
	//	{1,3},
	//	{2,6},
	//	{8,10},
	//	{15,18},
	//}
	var arr1 = [][]int {
		{2,3},
		{2,2},
		{3,3},
		{1,3},
		{5,7},
		{2,2},
		{4,6},
	}
	fmt.Println(merge(arr1))
}

type IntSlice [][]int
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool {
	if s[i][0] < s[j][0] {
		return true
	}else if s[i][0] == s[j][0] {
		return s[i][1] < s[j][1]
	}
	return false
}

func merge(intervals [][]int) [][]int {
	var temp = [][]int{}
	var len int = len(intervals)
	if len == 0 {
		return temp
	}
	//按照数组第一位大小进行排序
	sort.Sort(IntSlice(intervals))
	temp = append(temp,intervals[0])
	var k int = 0
	//要考虑三种情况
	for i:=1;i<len;i++ {
		if intervals[i][0] > temp[k][1] {
			temp = append(temp,intervals[i])
			k++
			continue
		}
		if intervals[i][0] <= temp[k][1] && intervals[i][1] >= temp[k][1] {
			temp[k][1] = intervals[i][1]
		}
	}
	return temp
}