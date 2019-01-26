package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
   Right *TreeNode
}

func findMode1(root *TreeNode) []int {
	arr1:=[]int{}
	arr2:=[]int{}
	mode:=make(map[int]int)
	if root.Left !=nil{
		arr1=append(findMode1(root.Left),root.Val)
	}else if root.Right!=nil {
		arr2=append(findMode1(root.Right),root.Val)
	}else {
		arr1=append(arr1,arr2...)
		for _,v:=range arr1 {
			mode[v]++
		}
		max:=root.Val
		end:=[]int{}
		for _,v:=range mode {
			if  v>max{
				max=v
			}
		}
		end=append(end,max)
		return end
	}
	return arr1
}
func findMode(root *TreeNode) []int {
	res:=[]int{}
	if root==nil {
		return res
	}
	pre:=new(TreeNode)

	cur:=1
	max:=0
	findModeCore(root,&pre,&cur,&max,&res)
	return res
}
func findModeCore(root *TreeNode,pre **TreeNode,cur *int,max *int,res *[]int)  {
	if root==nil {
		return
	}
	findModeCore(root.Left, pre, cur, max,res)
	if pre!=nil&&root.Val==*pre.Val{
		cur++
	}else {
		cur=1
	}
	if (cur >= max) {
		if (cur == max) {
			res = append(res, root.Val)
		} else {
			res = nil
			res = append(res, root.Val)
			max = cur
		}
	}
	pre=root
	findModeCore(root.Right, pre, cur, max,res)
}
func main() {
	//TreeNode1:=&TreeNode{2,nil,nil}
	//TreeNode0:=&TreeNode{2,TreeNode1,nil}
	//TreeNode2:=&TreeNode{1,nil,TreeNode0}

	//TreeNode3:=&TreeNode{18,nil,nil}
	//TreeNode4:=&TreeNode{7,TreeNode2,TreeNode3}
	//TreeNode5:=&TreeNode{15,nil,nil}
	//TreeNode6:=&TreeNode{10,TreeNode4,TreeNode5}
	//for _,v:=range findMode(TreeNode2){
	//	println(v)
	//}
	pre:=new(TreeNode)
	fmt.Printf("a is %d",pre)

}
