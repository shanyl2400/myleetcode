package main

import "fmt"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) [][]string {
	ret := make([][]string, 0)
	for i := 1; i < len(s); i++ {
		str := s[:i]
		if isPalindrome(str) {
			items := partition(s[i:])
			for j := range items {
				item := []string{
					str,
				}
				item = append(item, items[j]...)
				ret = append(ret, item)
			}
		}
	}
	if isPalindrome(s) {
		ret = append(ret, []string{s})
	}
	return ret
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
