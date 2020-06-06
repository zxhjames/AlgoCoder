package main

func main(){

}

type MyStack struct {
	size int
	stack []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		size:0,
		stack:[]int{},
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.size++
	this.stack = append(this.stack,x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	this.size--
	l:=len(this.stack)
	k:=this.stack[l-1]
	this.stack = this.stack[0:len(this.stack)-1]
	return k
}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */