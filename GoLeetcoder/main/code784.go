package main

import "fmt"

//字母全排列
func main(){
	//var s string = "a1b2"
	fmt.Println('a','z','A','Z')
	//fmt.Println(letterCasePermutation(s))

}

//func letterCasePermutation(S string) []string {
//	//初始化结果集数组
//	var ans= []string{S}
//	dfs6(ans,S,0)
//	return ans
//}
//
//
//func dfs6(ans []string,s string,k int)(){
//	if k == len(s) {
//		ans = append(ans,s)
//		return
//	}
//	r:=[]rune(s)
//	if !isBigWord(r[k]) && !isLittleWord(r[k]){
//		dfs6(ans,string(r),k+1)
//	}else{
//		r[k] = r[k] + 32
//		dfs6(ans,string(r),k+1)
//		r[k] = r[k] - 32
//		dfs6(ans,string(r),k+1)
//	}
//}
//
//
//func isBigWord(c int32) bool {
//	if c >=65 && c <=90 {
//		return true
//	}
//	return false
//}
//
//
//func isLittleWord(c int32) bool {
//	if c>=97 && c<=122 {
//		return true
//	}
//	return false
//}