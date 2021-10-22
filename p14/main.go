package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) < 2 {
		return ""
	}
	index := 0
out:
	for i := range strs[0] {
		for _, v := range strs[1:] {
			if len(v) <= i {
				break out
			}
			if strs[0][i] != v[i] {
				break out
			}
		}
		index++
	}
	return strs[0][:index]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"ab",
		"a",
	}))
}
