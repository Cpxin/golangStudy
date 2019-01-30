package main

func circularArrayLoop(nums []int) bool {
	for i:=0;i<len(nums);i++ {
		if nums[i]>0 {
			if pos(i,nums){     //向前
				return true
			}
		}else {
			if neg(i,nums) { //向后
				return true
			}
		}
	}
	return false
}
func pos(i int,nums []int)bool  {   //i:起点索引 j:起点下一个索引
	m:=make(map[int]int)
	m[i]=1
	mid:=i
	for k:=i;k<len(nums) ;  {
		mid=k
		k=k+nums[k]
		if k>=len(nums) {
			k=k%len(nums)
		}
		if nums[k]>0 {
			m[k]++
			if mid!=k&&m[k]==2 {
				return true
			}else if mid==k{
				return false
			}
		}else {
			return false
		}
	}
	return false
}
func neg(i int,nums []int)bool  {
	m:=make(map[int]int)
	m[i]=1
	mid:=i
	for k:=i;k<len(nums) ;  {    //
		mid=k
		k=k+nums[k]            //v:k的下一个索引
		if k<0 {
			k=len(nums)-(-k%len(nums))
		}
		if nums[k]<0 {
			m[k]++
			if m[k]==2&&mid!=k {
				return true
			}else if mid==k {
				return false
			}
		}else {
			return false
		}
	}
	return false
}
func main() {
	nums:=[]int{ -1, -2,-3,-4,-5}
	println(circularArrayLoop(nums))
}
