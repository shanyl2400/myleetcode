package main

import "fmt"

func longestPalindrome(s string) string {
	max := string(s[0])
	for i := range s[:len(s)-1] {
		if s[i] == s[i+1] {
			//回文检查
			start, end := getPalindrome(i, i+1, s)
			if end-start+1 > len(max) {
				max = s[start : end+1]
			}
		}

		if i <= len(s)-3 {
			if s[i] == s[i+2] {
				//回文检查
				start, end := getPalindrome(i, i+2, s)
				if end-start+1 > len(max) {
					max = s[start : end+1]
				}
			}
		}
	}
	return max
}
func getPalindrome(start, end int, s string) (int, int) {
	for start >= 0 && end < len(s) {
		if s[start] != s[end] {
			break
		}
		start--
		end++
	}
	start++
	end--
	return start, end
}

func main() {
	fmt.Println(longestPalindrome("ac"))

}
