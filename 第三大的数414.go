package main

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	ret:=[]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		ret = append(ret, nums[i])
	}
	if len(ret)<=2 {
		return ret[len(ret)-1]
	}else {
		return ret[len(ret)-3]
	}
}

func main() {
	nums:=[]int{-2,3,6,6}
	println(thirdMax(nums))
}
