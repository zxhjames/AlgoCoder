package main

import (
	"math"
	"fmt"
)

//最小栈

type MinStack struct {
	stack []int
	minElem int
}


/** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{stack:nil}
//}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack,x)
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	var n = math.MaxInt64
	for _,v := range this.stack {
		if v < n {
			n = v
		}
	}
	return n
}

func main(){
	obj:=Constructor()
	obj.Push(10)
	obj.Push(2)
	obj.Push(11)

	fmt.Println(obj)
	obj.Pop()
	fmt.Println(obj)
}
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
