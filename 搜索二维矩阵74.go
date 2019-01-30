package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for _,v:=range matrix{
		if len(v)!=0&&v[0]<=target&& v[len(v)-1]>=target {
			for k,tar:=range v {
				if tar==target {
					return true
				}else if k==len(v)-1 {
					return false
				}
			}
		}
	}
	return false


}
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix[0])==0 {
		return false
	}
	min:=0
	max:=len(matrix)*len(matrix[0])-1
	for min<=max  {
		mid:=min+(max-min)/2
		if matrix[mid/len(matrix)][mid%len(matrix)]>target {
			max=mid-1
		}else if matrix[mid/len(matrix)][mid%len(matrix)]<target {
			min=mid+1
		}else {
			return true
		}
	}
	return false

}

func main() {
	//matrix:=[][]int{
	//{1,   3,  5,  7},
	//{10, 11, 16, 20},
	//{23, 30, 34, 50},
	//}
	matrix:=[][]int{
		{1, 1},
	}
	//fmt.Print(len(matrix))
	//matrix:=[][]int{}
	fmt.Print(searchMatrix(matrix,2))
	//_,ok := matrix[0][0]
	//for k,v:=range matrix {
	//	fmt.Printf("k is %d,v is %d",k,v)
	//}
}
