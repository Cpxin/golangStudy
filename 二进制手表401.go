package main

import "fmt"

func readBinaryWatch(num int) []string {
	read:=[]string{}
	s:=""
	for h:=0;h<12;h++ {
		for m:=0;m<60 ;m++  {
			addNum:=bNum(h,m)
			if addNum==num&&m<10{
				s=fmt.Sprintf("%d:0%d",h,m)
				read=append(read,s)
			}else if addNum==num&&m>=10 {
				s=fmt.Sprintf("%d:%d",h,m)
				read=append(read,s)
			}
		}
	}
	return read
}

func bNum(h int,m int)int  {
	numH,numM:=0,0
	var i uint8=0
	for i=0;i<4;i++{
		if h&(1<<i)!=0 {
			numH++
		}
	}
	for i=0;i<6;i++{
		if m&(1<<i)!=0 {
			numM++
		}
	}
	return numH+numM
}

func main() {
	for _,v:=range readBinaryWatch(1){
		println(v)
	}
	
}
