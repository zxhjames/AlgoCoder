package main

import (
	"strings"
	"fmt"
)

/**
最后一个单词
 */

func main() {
	var strs  string = ""
	var strarr []string = strings.Fields(strs)
	fmt.Println(strarr)
}


func lengthOfLastWord(s string) int {

	var strs []string = strings.Fields(s)
	if len(strs) == 0 {
		return 0
	}
	return len(strs[len(strs) - 1])
}
