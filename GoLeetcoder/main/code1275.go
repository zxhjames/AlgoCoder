package main

//井字棋游戏

func main(){

}

func tictactoe(moves [][]int) string {
	isEnd:=0 //用来记录是否结束
	var m = [][]int{
		{0,0,0},
		{0,0,0},
		{0,0,0},
	}
	//下棋过程
	for i:=0;i<len(moves);i++ {
		if i%2 == 0 {
			m[moves[i][0]][moves[i][1]] = 1
			if isWin(m,1) {
				return "A"
			}
		}else{
			m[moves[i][0]][moves[i][1]] = -1
			if isWin(m,-1){
				return "B"
			}
		}
		isEnd++
	}
	if isEnd == 9 {
		return "Draw"
	}
	return "Pending"

}

func isWin(ans [][]int,player int) bool {
	//分为三个类别
	if ans[0][0] == ans[0][1] && ans[0][1] == ans[0][2] && ans[0][2] == player ||
		ans[1][0] == ans[1][1] && ans[1][1] == ans[1][2] && ans[1][2] == player ||
		ans[2][0] == ans[2][1] && ans[2][1] == ans[2][2] && ans[2][2] == player ||
		ans[0][0] == ans[1][0] && ans[1][0] == ans[2][0] && ans[2][0] == player ||
		ans[0][1] == ans[1][1] && ans[1][1] == ans[2][1] && ans[2][1] == player ||
		ans[0][2] == ans[1][2] && ans[1][2] == ans[2][2] && ans[2][2] == player ||
		ans[0][0] == ans[1][1] && ans[1][1] == ans[2][2] && ans[2][2] == player ||
		ans[0][2] == ans[1][1] && ans[1][1] == ans[2][0] && ans[2][0] == player {
			return true
		}
	return false
}