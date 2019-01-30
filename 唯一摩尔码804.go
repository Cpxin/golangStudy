package main

import "fmt"

var dict = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

func uniqueMorseRepresentations(words []string) int {
	visit := make(map[string]struct{})
	for _, word := range words {
		visit[transform(word)] = struct{}{}
	}
	return len(visit)
}

func transform(word string) (result string) {
	for _, c := range word {
		result = result + dict[int(c-'a')]
	}
	return result
}

func main() {
	words:=[]string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
