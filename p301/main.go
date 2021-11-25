package main

import (
	"fmt"
)

func isValid(s string) bool {
	count := 0
	for i := range s {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--

			if count < 0 {
				return false
			}
		}
	}
	return true
}

func doRemove(ans *[]string, lbrackets, rbrackets int, s string, start int) {
	if lbrackets == 0 && rbrackets == 0 {
		//过滤不合法的结果
		if isValid(s) {
			*ans = append(*ans, s)
		}
		return
	}

	for i := start; i < len(s); i++ {
		//过滤重复项
		if i > start && s[i] == s[i-1] {
			continue
		}
		//
		if lbrackets+rbrackets > len(s)-i {
			return
		}
		if s[i] == '(' && lbrackets > 0 {
			//尝试删除左括号
			doRemove(ans, lbrackets-1, rbrackets, s[:i]+s[i+1:], i)
		}
		if s[i] == ')' && rbrackets > 0 {
			//尝试删除右括号
			doRemove(ans, lbrackets, rbrackets-1, s[:i]+s[i+1:], i)
		}
	}
}

func removeInvalidParentheses(s string) []string {
	lbrackets := 0
	rbrackets := 0
	for i := range s {
		if s[i] == '(' {
			lbrackets++
		} else if s[i] == ')' {
			if lbrackets == 0 {
				rbrackets++
			} else {
				lbrackets--
			}
		}
	}
	ans := make([]string, 0)
	doRemove(&ans, lbrackets, rbrackets, s, 0)
	return ans
}

func main() {
	//s := "()())()"
	//s := "(a)())()"
	//s := "(a()b()("
	//s := ")("
	//s := "))n((i()"
	s := "(()()()))(("
	fmt.Println(removeInvalidParentheses(s))
}
