package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root==nil{
		return 0
	}
	if root.Val<L {
		return rangeSumBST(root.Right,L,R)
	}else if root.Val>R {
		return rangeSumBST(root.Left,L,R)
	}else {
		return root.Val+rangeSumBST(root.Left,L,R)+rangeSumBST(root.Right,L,R)
	}
}
func main() {
	TreeNode1:=&TreeNode{5,nil,nil}
	TreeNode0:=&TreeNode{0,nil,nil}
	TreeNode2:=&TreeNode{3,TreeNode0,TreeNode1}
	TreeNode3:=&TreeNode{18,nil,nil}
	TreeNode4:=&TreeNode{7,TreeNode2,TreeNode3}
	TreeNode5:=&TreeNode{15,nil,nil}
	TreeNode6:=&TreeNode{10,TreeNode4,TreeNode5}
	fmt.Println(rangeSumBST(TreeNode6,7,15))
}
