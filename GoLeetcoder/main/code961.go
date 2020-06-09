package main


func main(){

}


func repeatedNTimes(A []int) int {
	m:=make(map[int]int)
	for _,value:=range A {
		m[value]++
	}
	for k,v:=range m {
		if v == len(A)/2 {
			return k
		}
	}
	return 0
}