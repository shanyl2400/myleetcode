package main

import "fmt"

func doCountSubstrings(s string, mid int) int {
	l := mid - 1
	r := mid + 1
	count := 0
	if r < len(s) && s[r] == s[mid] {
		count++
		r++
		for l >= 0 && r < len(s) {
			if s[l] == s[r] {
				count++
			} else {
				break
			}
			l--
			r++
		}
	}
	l = mid - 1
	r = mid + 1
	for l >= 0 && r < len(s) {
		if s[l] == s[r] {
			count++
		} else {
			break
		}
		l--
		r++
	}

	return count
}

func countSubstrings(s string) int {
	ret := 0
	for i := range s {
		ret++
		ret += doCountSubstrings(s, i)
	}

	return ret
}

func main() {
	// s := "aaa"
	// s := "abcdedc"
	s := "aaaaa"
	fmt.Println(countSubstrings(s))
}
