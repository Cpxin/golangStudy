package main

import (
	"fmt"
	"time"
)

var Nums =[][]int{}
var Munms=make(map[string]int)

func permuteUnique(nums []int) [][]int {
	t1 := time.Now()
	end:=[]int{}
	Nums=nil
	Munms=make(map[string]int)
	if len(nums)==0||len(nums)==1 {
		Nums=append(Nums,nums)
		return Nums
	}
	first(nums,end)
	elapsed := time.Since(t1)
	fmt.Println(elapsed)
	return Nums

}
func first(nums []int,end []int)  {
	if len(nums)!=0 {
		Tnums:=nums
		for i:=0;i<len(Tnums) ;i++  {
			Tend:=append(end[:len(end):len(end)],Tnums[i])
			nums=append(Tnums[:i:i],Tnums[i+1:]...)
			first(nums,Tend)
			Tend=Tend[:len(Tend)-1]
		}
	}else {
		temp:=""
		for _,v:=range end{
			temp+=fmt.Sprintf("%d",v)
		}
		Munms[temp]++
		if Munms[temp]==1 {
			Nums=append(Nums,end)
		}
	}
}

func main() {

	nums:=[]int{1,2,3,4}

	fmt.Println(permuteUnique(nums))




}
