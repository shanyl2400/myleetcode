package main

import "fmt"

func isUnique(astr string) bool {
	strMap := make(map[rune]bool)
	for _, x := range astr {
		res := strMap[x]
		if res {
			return false
		}
		strMap[x] = true
		fmt.Println(x)
	}
	return true
}

func main() {
	fmt.Println(isUnique("leetcode"))
}
