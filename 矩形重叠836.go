package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec1[0]>=rec2[2] {
		return false
	}else if rec1[2]<=rec2[0] {
		return false
	}else if rec1[3]<=rec2[1] {
		return false
	}else if rec1[1]>=rec2[3] {
		return false
	}else {
		return true
	}
}

func main() {
	rec1:=[]int{0,0,1,1}
	rec2:=[]int{1,0,2,1}
	println(isRectangleOverlap(rec1,rec2))
}
