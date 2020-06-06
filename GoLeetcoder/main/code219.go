package main

import "fmt"

//存在重复元素
func main(){
	var arr = []int{1}
	var k = 1
	fmt.Println(containsNearbyDuplicate(arr,k))
}

//我的垃圾解法
func containsNearbyDuplicate(nums []int, k int) bool {
	m:=make(map[int][]int)
	l:=len(nums)
	for p:=0;p<l;p++ {
		m[nums[p]] = append(m[nums[p]],p)
		ml:=len(m[nums[p]])
		if ml > 1 {
			m[nums[p]][ml-2] = m[nums[p]][ml-1] - m[nums[p]][ml-2]
		}
	}
	for _,v:=range m {
		for p:=0;p<len(v);p++ {
			if v[p] <= k {
				return true
			}
		}
	}
	return false
}

//高效解法
func containsNearbyDuplicate1(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok && (i-m[v] <= k) {
			return true
		}
		m[v] = i
	}
	return false
}
