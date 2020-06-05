package main

import "fmt"

func main(){
	var k int = rangeBitwiseAnd(5,7,)
	fmt.Println(k)
}

func rangeBitwiseAnd(m int, n int) int {
	var k int
	for k = n;k>m;{
		k = k & (k-1)
	}
	return k
}
