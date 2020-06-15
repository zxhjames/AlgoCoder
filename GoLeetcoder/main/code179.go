package main

import (
	"sort"
	"strconv"
	"fmt"
	"strings"
)

//code 最大数
func largestNumber(nums []int) string {
	ss := make([]string, len(nums))
	for i, num := range nums {
		ss[i] = strconv.Itoa(num)
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i]+ss[j] >= ss[j]+ss[i]
	})
	o := strings.Join(ss, "")
	if o[0] == '0' {
		return "0"
	}
	return o
}


func getIndex(n int) (int) {
	for n>=10 {
		n/=10
	}
	return n
}



func main(){
	fmt.Println(largestNumber([]int{3,30,34,5,9}))
}
