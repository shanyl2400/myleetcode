package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	max := 0
	cur := 0
	begin := 0
	for i := 0; i < len(s); i++ {
		offset := isDuplicate(begin, s[:i], s[i])
		if offset >= 0 {
			begin = offset + 1
			cur = i - begin + 1
		} else {
			cur = cur + 1
		}
		if max < cur {
			max = cur
		}
	}
	return max
}

func isDuplicate(begin int, s string, c byte) int {
	ret := -1
	for i := begin; i < len(s); i++ {
		if s[i] == c {
			ret = i
		}
	}
	return ret
}

func main() {
	// blqsearxxxbiwqa
	// pwwkew
	// aab
	// bbtablud
	fmt.Println(lengthOfLongestSubstring("blqsearxxxbiwqa"))
}
