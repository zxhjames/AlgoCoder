package main

import (
	"strconv"
	"strings"
	"fmt"
)

/**
38.外观数列
 */

func main () {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	var str []string = []string{"1"}
	var tmp []string = []string{}
	var l int = 0
	var count int = 1
	for i:=1;i<n;i++ {
		l = len(str)
		for j:=0;j<l;j++ {
			if j+1 < l && str[j+1] == str[j] {
				count++
				continue
			}else {
				tmp = append(tmp,strconv.Itoa(count))
				tmp = append(tmp,str[j])
				count = 1
			}
		}
		if count != 1 {
			tmp = append(tmp,strconv.Itoa(count))
			tmp = append(tmp,str[l-1])
		}
		str = tmp
		tmp = []string{}
	}
	return strings.Join(str,"")
}
