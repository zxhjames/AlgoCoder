package main

import (
	"fmt"
	"sort"
)

//code720最长的单词
func main(){
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}


func longestWord(words []string) string {
	m:=make(map[uint8]int)
	//排序
	sort.Strings(words)
	for k:=0;k<len(words);k++ {
		for p:=0;p<len(words[k]);p++{
			m[words[k][p]]++
		}
	}
	return words[0]
}