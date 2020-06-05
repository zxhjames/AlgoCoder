package main

import (
	"fmt"
	"strings"
)

/**
罗马数字转换成整数 AC
 */
func main() {
	var str string = "MCMXCIV"

	//fmt.Println(strings.Index(str,"X")) //获取第一个子字符的起始位置
	//fmt.Println(strings.Count(str,"X")) //统计一个字符串中某个字符出现个数
	//fmt.Println(strings.Compare(str,"VY")) //比较字符串
	//fmt.Println(strings.Contains(str,"VY")) //判断是否包含子串
	//fmt.Println(strings.ContainsAny(str,"VY"))
	//fmt.Println(strings.Fields(str)) //转换成字符串数组
	//fmt.Println(strings.HasPrefix(str,"XX")) //判断后者是否是前者的前缀
	//fmt.Println(strings.HasSuffix(str,"YX")) //判断后者是否是前者的后缀
	//fmt.Println(strings.Replace(str,"IXX","AAA",-1)) //使用new 代替 str 中的old ，一共替换n个
	//var s = strings.Split(str,"")
	//fmt.Println(s,len(s))
	fmt.Println(romanToInt(str))
}
//
func romanToInt(s string) int {
	//生成map数组
	m:=map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
	var arr = strings.Split(s,"") //首先转换成数组
	var sum int = 0
	var len int = len(arr)
	if len == 0 {
		return 0
	}
	sum = sum + m[arr[0]]
	for i:=1;i<len;i++ {
		var res int = m[arr[i]] / m[arr[i-1]]
		var temp string = arr[i-1]
		if (res == 10 || res == 5)  &&  (temp == "I" || temp == "X" || temp == "C") {
			sum = sum + m[arr[i]] - m[arr[i-1]] - m[arr[i-1]]
		}else{
			sum = sum + m[arr[i]]
		}
	}
	return sum
}
