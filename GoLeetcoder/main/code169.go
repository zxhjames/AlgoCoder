package main

/**
多数元素
 */

func main() {

}


//多数元素
func majorityElement(nums []int) int {
	var edge int = len(nums) / 2
	m:=make(map[int]int)
	var n int
	var f int
	for _,v:=range nums {
		m[v]++
	}
	for k,v:=range m {
		if v > edge && v > n{
			n = v
			f = k
		}
	}
	return f
}
