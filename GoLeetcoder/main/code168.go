package main

import "fmt"

func main(){
	fmt.Println(convertToTitle(52))
	var b = []byte{
		'a','b',
	}
}
func convertToTitle(n int) string {
	r:=[]rune{
		'.','A','B','C','D','E','F','G','H','I','J','K',
		'L','M','N','O','P','Q','R','S','T','U','V','W',
		'X','Y','Z',
	}
	runes:=[]rune{}
	for n>26 {
		tmp:=n%26
		n = n/26
		if tmp == 0 {
			runes = append(runes,'Z')
			for k:=0;k<n-1;k++{
				runes = append(runes,'A')
			}
			reverse(runes)
			return string(runes)
		}
		runes = append(runes,r[tmp])

	}
	runes = append(runes,r[n])
	//reverse
	reverse(runes)
	return string(runes)
}


func reverse(runes []rune){
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
}