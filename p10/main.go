package main

import "fmt"

func isMatch(s string, p string) bool {
	return doMatch(s, p, 0, 0)
}

func doMatch(s string, p string, si int, pi int) bool {
	//字符串结束
	if si >= len(s) && pi >= len(p) {
		return true
	}
	if pi+1 < len(p) && p[pi+1] == '*' {
		//*的情况
		//省略*
		ret := doMatch(s, p, si, pi+2)
		if ret {
			return true
		}
		//匹配*
		for idx := 0; idx < len(s)-si; idx++ {
			if matchChar(s[si+idx], p[pi]) {
				ret := doMatch(s, p, si+1+idx, pi+2)
				if ret {
					return true
				}
			} else {
				return false
			}
		}
	} else if si >= len(s) || pi >= len(p) || !matchChar(s[si], p[pi]) {
		return false
	}
	return doMatch(s, p, si+1, pi+1)
}

func matchChar(s byte, p byte) bool {
	if p == '.' || p == s {
		return true
	}
	return false
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
}
