package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	max := 0
	temp :=0
	m := make(map[int]int,len(G))
	for _,v :=  range G{
		m[v] =0
	}
	for p :=head; p != nil; p =p.Next{
		if _,ok := m[p.Val]; ok{
			temp++
		}else{
			if temp != 0{
				max++
			}
			temp =0
		}
	}
	if temp != 0{
		max++
	}
	return max
}
func numComponents2(head *ListNode, G []int) int {
	j,t:=false,false
	num:=0
	for head!=nil{
		for i:=0; i<len(G);i++  {
			if head.Val==G[i] {
				head=head.Next
				G=append(G[:i], G[i+1:]...)
				j,t=true,true
				break
			}
		}
		if j==true&&head!=nil{
			j=false
		}else if t==true{
			t=false
			num++
		}else if head!=nil {
			head=head.Next
		}
	}
	return num
}
func main() {
	//ListNode6 :=&ListNode{5,nil}
	ListNode5:=&ListNode{1,nil}
	ListNode4:=&ListNode{2,ListNode5}
	ListNode3:=&ListNode{0,ListNode4}
	ListNode2:=&ListNode{4,ListNode3}
	ListNode:=&ListNode{3,ListNode2}
	G:=[]int{0,3,2,4,1}
	fmt.Println(numComponents2(ListNode,G))
	//arr:=[]int{2,5,1,4,6}
	//sort.Ints(arr)
	//for _,v:=range arr{
	//	println(v)
	//}
}
