package main

import (
	"math"
	"fmt"
)

func main() {
	r5:=TreeNode{
		0,
		nil,
		nil,
	}
	r4:=TreeNode{
		Val:0,
		Left:nil,
		Right:nil,
	}
	r3:=TreeNode{
		0,
		&r4,
		&r5,
	}
	r2:=TreeNode{
		Val:0,
		Left:nil,
		Right:nil,
	}
	r1:=TreeNode{
		0,
		&r2,
		&r3,
	}
	fmt.Println(dfs(&r1))
}

type TreeNode struct {
	    Val int
	    Left *TreeNode
	    Right *TreeNode
}
//平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if math.Abs(dfs(root.Left) - dfs(root.Right)) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}


func dfs(root *TreeNode) (float64) {
	if root != nil {
		return math.Max(dfs(root.Left), dfs(root.Right)) + 1
	}
	return 0
}
