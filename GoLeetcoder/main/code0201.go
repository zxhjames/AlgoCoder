package main


func main() {

}



type ListNode struct {
	Val int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m:=make(map[int]int)
	tmp := head
	for tmp!=nil {
		m[tmp.Val]++
		tmp = tmp.Next
	}
	new := &ListNode{}
	tmp = new
	for k,_:=range m {
		tmp.Val = k
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return new
}
