package main

import "fmt"

func lower(s byte) byte {
	if (s >= 'a' && s <= 'z') || (s >= '0' && s <= '9') {
		return s
	}
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return 0
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	right := len(s) - 1
	left := 0
	diff := 0
	for left < right {
		s1 := lower(s[left])
		s2 := lower(s[right])
		if s1 == 0 {
			left++
			continue
		}
		if s2 == 0 {
			right--
			continue
		}
		diff = int(s1 - s2)

		if diff != 0 {
			return false
		}

		left++
		right--
	}
	return true
}

func main() {
	// s := "A man, a plan, a canal: Panama"
	s := "race a car"
	// s := "0P"
	fmt.Println(isPalindrome(s))
}
