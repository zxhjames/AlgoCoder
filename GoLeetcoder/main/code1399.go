package main

func main(){

}

func countLargestGroup(n int) int {
	m1,m2:=make(map[int][]int),make(map[int]int)
	for i:=1;i<=n;i++{
		tmp:=getKey(i)
		m1[tmp] = append(m1[tmp], i)
	}
	for _,v:=range m1 {
		m2[len(v)]++
	}
	K:=-1
	for _,v:=range m2 {
		if K < v {
			K = v
		}
	}
	return K
}


func getKey(n int) (int){
	sum:=0
	for n!=0 {
		sum+=n%10
		n = n/10
	}
	return sum
}