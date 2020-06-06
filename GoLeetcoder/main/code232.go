package main

func main(){

}

type MyQueue struct {
	size int
	quene []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		size:0,
		quene:[]int{},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.quene = append(this.quene, x)
	this.size++
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var k = this.quene[0]
	this.quene = this.quene[1:]
	this.size--
	return k
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.quene[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

