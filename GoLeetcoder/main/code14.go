package main

import "fmt"

/**
最长公共前缀
 */

func main() {
	//var s string = "1234"
	//for _,s:=range s{
	//	fmt.Printf("%c",s)
	//}
	//fmt.Println(len(s))
	//var str []string = []string{"aa"}
	//fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	m:=make(map[string]int)
	var p int = 0
	var t string
	var f bool = true
	var str string = ""
	for {
		if !f {
			break
		}
		for i:=0;i<len(strs);i++ {
			if p == len(strs[i]) {
				f = false
				return str
			}
			t = fmt.Sprintf("%c",strs[i][p])
			m[t]++
		}
		var cot int = m[t]%len(strs)
		if cot == 0 {
			str = str + t
		}else {
			break
		}
		m = make(map[string]int)
		p++
	}
	return str
}
