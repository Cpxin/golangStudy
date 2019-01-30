package main

func findPeakElement(nums []int) int {
	max:=0
	if len(nums)>0 {
		max=nums[0]
	}else {
		return 0
	}
	top:=0
	for k,v:= range nums  {
		if max<v {
			max=v
			top=k
		}else if k!=0{
			top=k-1
			return top
		}
	}
	return top
}

func main() {
	nums:=[]int{1,2,1}
	println(findPeakElement(nums))
}
