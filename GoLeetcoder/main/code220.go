package main

import "math"

//存在重复元素2

func main() {

}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 10000 {
		return false
	}
	l:=len(nums)
	for i:=0;i<l;i++ {
		for j:=i+1;j<l;j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= float64(t) && math.Abs(float64(i-j))<=math.Abs(float64(k)){
				return true
			}
		}
	}
	return false
}

func Solution2(nums []int,k int,t int) bool {

}