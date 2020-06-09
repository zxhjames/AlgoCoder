package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println(commonChars([]string{"bella","label","roller"}))
}

//code1002,寻找常用字符
func commonChars(A []string) []string {
	var str []string
	var tmp string
	l:=len(A)
	m:=map[uint8]int{}
	for k:=0;k<l;k++{
		var b= []byte(A[k])

	}
	for k,v:=range m {
		if v%l == 0{
			for i:=0;i<v/l;i++{
				t:=fmt.Sprintf("%c",k)
				str = append(str,t)
			}
		}
	}
	return str
}