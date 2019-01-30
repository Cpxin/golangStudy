package main

func intersect(nums1 []int, nums2 []int) []int {
	nums1map:=make(map[int]int)
	for _,num1:=range nums1 {
		nums1map[num1]++
	}
	agg:=[]int{}
	for _,num2:=range nums2 {
		//_,ok :=nums1map[num2]
		//if ok&&nums1map[num2]!=0 {
		if nums1map[num2]>0 {
			nums1map[num2]--
			agg=append(agg,num2)
		}
	}
	return agg
}

func main() {
	nums1:=[]int{1,2,2,1}
	nums2:=[]int{2,2}
	println(intersect(nums1,nums2))
	//var nums1map map[int]int
	//nums1map:=make(map[int]int)
	//nums1map[1]++
	//println(nums1map[1])
}
