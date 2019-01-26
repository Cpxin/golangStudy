package main

import "math"

func binaryGap(N int) int {
	var binary=[]int{}
	m:=0
	for  N!=0{
		m=N%2
		N=N/2
		binary=append(binary,m)
	}
	num :=make(map[int]int)
	zero:=0
	for _,bin:=range binary {
		_,ok :=num[0]
		if bin==0&&ok {
			zero++
		}else if bin==1{
			num[1]++
			if !ok {
				num[0]=zero
			}else if num[0]<=zero {
				num[0]=zero
			}
			zero=0
		}
	}
	if num[1]<=1 {
		return 0
	}else {
		return num[0]+1
	}
}
func binaryGap2(N int) int {
	last := -1
	ans := 0
	un := uint(N)
	var i uint
	for i = 0 ; i < 32 ; i++ {
		if(( un>>i ) & 1 ) > 0  {
			if last >= 0 {
				ans = int(math.Max(float64(ans), float64(int(i) - last)))
			}
			last = int(i)
		}
	}
	return ans
}

func main() {
	println(binaryGap(5))
	//println(uint(5))
}
