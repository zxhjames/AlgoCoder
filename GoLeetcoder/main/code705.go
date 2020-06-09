package main

func main(){

}

type MyHashSet struct {
	data []int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	var data []int
	for i:=0;i<=1000001;i++{
		data[i] = 0
	}
	return MyHashSet{
		data:data,
	}
}


func (this *MyHashSet) Add(key int)  {
	this.data[key] = key
}


func (this *MyHashSet) Remove(key int)  {
	this.data[key] = 0
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.data[key] == key {
		return true
	}
	return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
