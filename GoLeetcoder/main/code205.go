package main

import "fmt"

//code 205
func main(){
	fmt.Println(1 ^ 1,1 | 1,1 & 1)
}

//同构字符串
func isIsomorphic(s string, t string) bool {
	l1:=len(s)
	l2:=len(t)
	if l1!=l2 {
		return false
	}
	m1:=make(map[uint8]int)
	m2:=make(map[uint8]int)
	for k:=0;k<l1;k++{
		if m1[s[k]]!=m2[t[k]] {
			return false
		}
		m1[s[k]] += k+1
		m2[t[k]] += k+1
	}
	return true
}
