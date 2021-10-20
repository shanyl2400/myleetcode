package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	dip := make([]int, len(s))
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] == ')' {
			if s[i-1] == '(' {
				dip[i] = 2
				if i >= 2 {
					dip[i] = dip[i] + dip[i-2]
				}
			} else if i-dip[i-1]-1 >= 0 {
				if s[i-dip[i-1]-1] == '(' {
					dip[i] = 2
					if i-dip[i-1]-2 >= 0 {
						dip[i] = dip[i] + dip[i-dip[i-1]-2]
					}
					if i >= 1 {
						dip[i] = dip[i] + dip[i-1]
					}
				}

			}
		}
	}

	max := 0
	for i := range dip {
		if dip[i] > max {
			max = dip[i]
		}
	}
	return max
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses("())"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("(()())"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses(")(((((()())()()))()(()))("))
	fmt.Println(longestValidParentheses(")(())(()()))("))
	fmt.Println(longestValidParentheses(")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())"))
}
