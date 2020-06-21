package main

import (
	"strings"
	"fmt"
)

/**
二叉树序列化与反序列化
 */


//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }
func main(){
	var str string = "[1,2,3,null,null,4,5]"
	var tmp = strings.Split(str,",")
	for i:=0;i<len(tmp);i++ {
		fmt.Printf("%s\n",tmp[i])
	}
	s:="[1]"
	tmp = strings.Split(s,",")
	fmt.Printf("%s",tmp)

}

//type Codec struct {
//	quene []*TreeNode
//}
//
//func Constructor() Codec {
//
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	//二叉树序列化成字符串
//	//采用层序遍历
//	begin,end:="[","]"
//	if root == nil {
//		return begin+end
//	}
//	begin+=strconv.Itoa(root.Val)
//	this.quene = append(this.quene,root)
//	for len(this.quene)!=0 {
//		top:=this.quene[0]
//		if top.Left == nil {
//			begin+=",null"
//		}else if top.Left!=nil {
//			begin+=","+strconv.Itoa(top.Left.Val)
//			this.quene = append(this.quene,top.Left)
//		}else if top.Right == nil {
//			begin+=",null"
//		}else if top.Right!=nil {
//			begin+=","+strconv.Itoa(top.Right.Val)
//			this.quene = append(this.quene,top.Right)
//		}
//		if len(this.quene) == 1 {
//			break
//		}
//		this.quene = this.quene[1:len(this.quene)]
//	}
//	return begin+end
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	if data == "[]"{
//		return nil
//	}
//	var tmp = strings.Split(data,",")
//}


/**
* Your Codec object will be instantiated and called as such:
* obj := Constructor();
* data := obj.serialize(root);
* ans := obj.deserialize(data);
*/
