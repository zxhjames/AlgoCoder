package main

import "strconv"

//数字翻译成字符串
func main(){
	//a:=12345
	//r:=strconv.Itoa(a)
	//fmt.Println(strings.Contains(a,"2"))
	//fmt.Println(r)
	//fmt.Println(r[0])
}

func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	str:=strconv.Itoa(num)
	return dfs7(str,"25",0,1) + 1
}


func dfs7(str,limit string,i,j int) int {
	if i == len(str)-1 && j == len(str)-1 {
		return 0
	}
	if j == len(str) - 1 {
		return 1
	}

}