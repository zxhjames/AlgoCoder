package main

import "fmt"

func main(){
	fmt.Println(equationsPossible([]string{"a==b","b!=a"}))
}

func equationsPossible(equations []string) bool {
	var str1 = equations[0]
	var str2 = equations[1]
}