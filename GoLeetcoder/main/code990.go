package main

import "fmt"

func main(){
	fmt.Println(equationsPossible([]string{"a==b","b!=a"}))
}

// @Description
// 1 扫描所有 a == b 加入并查集
// 2 扫描所有 a != b 确认是否合规
func equationsPossible(equations []string) bool {
	cMap := make([]int, 26)
	for i := 0; i < len(cMap); i++ {
		cMap[i] = i
	}
	// scan equations
	for i := 0; i < len(equations); i++ {
		//相等的等式
		if equations[i][1] == '=' {
			Merge(cMap, int(equations[i][0]-'a'), int(equations[i][3]-'a'))
			//fmt.Println("merge", int(equations[i][0]-'a'), int(equations[i][3]-'a'), "to", cMap)
		}
	}
	// scan inequations
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '!' {
			fa := find(cMap, int(equations[i][0]-'a'))
			fb := find(cMap, int(equations[i][3]-'a'))
			if fa == fb || fb == cMap[fa] {
				return false
			}
		}
	}
	//fmt.Println(cMap)
	return true
}

func Merge(cmap []int, i int, j int) {
	fa := find(cmap, i)
	fb := find(cmap, j)
	if fa != fb {
		cmap[fa] = fb
	}
}

func find(cmap []int, i int) int {
	j := i
	for cmap[j] != i {
		j, cmap[j] = cmap[j], i
	}
	return cmap[i]
}

