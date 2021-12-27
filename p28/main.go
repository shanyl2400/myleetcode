package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	k := -1
	pi := computePrefix(needle)
	for i := range haystack {
		for k >= 0 && haystack[i] != needle[k+1] {
			k = pi[k]
		}
		if haystack[i] == needle[k+1] {
			k++
		}
		if k == len(needle)-1 {
			return i - len(needle) + 1
		}
	}
	return -1
}
func computePrefix(p string) []int {
	pi := make([]int, len(p))
	pi[0] = -1
	k := -1

	for i := 1; i < len(p); i++ {
		for k >= 0 && p[i] != p[k+1] {
			k = pi[k]
		}
		if p[k+1] == p[i] {
			k++
		}
		pi[i] = k
	}
	return pi
}

func main() {
	haystack := "hello"
	needle := "ll"
	// haystack := "aaaaa"
	// needle := "bba"
	// haystack := "aaa"
	// needle := "aaaa"
	fmt.Println(strStr(haystack, needle))
}
