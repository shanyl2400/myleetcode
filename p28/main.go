package main

import "fmt"

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	j := 0
	for i := range haystack {
		j = 0
		for j < len(needle) && i+j < len(haystack) && haystack[i+j] == needle[j] {
			j++
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

func main() {
	// haystack := "hello"
	// needle := "ll"
	haystack := "aaaaa"
	needle := "bba"
	// haystack := "aaa"
	// needle := "aaaa"
	fmt.Println(strStr(haystack, needle))
}
