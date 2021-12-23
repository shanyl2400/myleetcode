package main

import (
	"fmt"
	"strings"
)

func reverse(s []byte) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return string(s)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	ret := ""
	for i := range words {
		if ret != "" {
			ret = ret + " "
		}
		ret = ret + reverse([]byte(words[i]))
	}
	return ret
}

func main() {
	str := "Let's take LeetCode contest"
	fmt.Println(reverseWords(str))
}
