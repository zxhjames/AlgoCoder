package main

import "fmt"

func main(){

	fmt.Println(toLowerCase("Hello"))
	fmt.Println('A','Z')
}

func toLowerCase(str string) string {
	r:=[]rune(str)
	for i:=0;i<len(r);i++ {
		r[i] = r[i] + 32
	}
	return string(r)
}