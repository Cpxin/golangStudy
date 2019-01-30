package main

func canConstruct(ransomNote string, magazine string) bool {
	str1:=make(map[string]int)
	str2:=make(map[string]int)
	for i:=0; i<len(ransomNote);i++  {
		str1[ransomNote[i:i+1]]++
	}
	for j:=0;j<len(magazine) ;j++  {
		str2[magazine[j:j+1]]++
	}
	num:=0
	for k,v:=range str1{
		if str2[k]>=v {
			num++
		}
	}
	if num==len(str1) {
		return true
	}else {
		return false
	}
}

func main() {
	ransomNote:="abc"
	//println(ransomNote[0:1])
	magazine:="aaabbbccc"
	println(canConstruct(ransomNote,magazine))
}
