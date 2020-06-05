package main

import (
	"strconv"
	"fmt"
)

/**
回文数
 */

func main() {
	fmt.Println(isPalindrome(101))
}

func isPalindrome(x int) bool {
	//强转换
	var str string = strconv.Itoa(x)
	if len(str) == 1 {
		return true
	}
	i:=0
	j:=len(str)-1
	for ;i<=j;{
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
