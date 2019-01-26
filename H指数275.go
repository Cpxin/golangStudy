package main

import "fmt"

func hIndex(citations []int) int {
	max:=len(citations)
	if max==0 {
		return 0
	}
	if citations[max-1]==0 {
		return 0
	}
	for i:=0;i<max;i++{
		if i+1<max&&i+1<citations[max-i-1] {
			continue
		}else if i+1==max {
			return i
		}else if i+1>citations[max-i-1] {
			return i
		}else{
			return i+1
		}
	}
	return 0
}
func hIndex2(citations []int) int {
	max:=len(citations)-1
	it:=0
	for i,_:=range citations {
		if i+1<citations[max-i]&&it<=i+1 {
			it=i+1
		}
	}
	return it
}

func main() {
	citations:=[]int{0,1,3,5,6}
	fmt.Print(hIndex2(citations))
}
