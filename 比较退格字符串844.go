package main

func backspaceCompare(S string, T string) bool {
	for i:=0;i<len(S); {
		if string(S[i])=="#" {
			if i==0 {
				S=S[1:]
				continue
			}else {
				S=S[:i-1]+S[i+1:]
				i=0
				continue
			}
		}
		i++
	}
	for j:=0;j<len(T) ; {
		if string(T[j])=="#" {
			if j==0 {
				T=T[1:]
				continue
			}else {
				T=T[:j-1]+T[j+1:]
				j=0
				continue
			}
		}
		j++
	}
	if S==T {
		return true
	}else {
		return false
	}
}

func main() {
	S:="#j##xfix"
	T:="j###xfix"
	//for i:=0;i<len(S);{
	//	if string(S[i])=="#" {
	//		if i==0 {
	//			S=S[1:]
	//			continue
	//		}else {
	//			S=S[:i-1]+S[i+1:]
	//			i=0
	//			continue
	//		}
	//	}
	//	i++
	//}
	//for j:=0;j<len(T) ; {
	//	if string(T[j])=="#" {
	//		if j==0 {
	//			T=T[1:]
	//			continue
	//		}else {
	//			T=T[:j-1]+T[j+1:]
	//			j=0
	//			continue
	//		}
	//	}
	//	j++
	//}
	//println(S)

	println(backspaceCompare(S,T))
}
