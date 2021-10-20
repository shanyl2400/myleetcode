package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, len(s))
	top := 0

	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[top] = s[i]
			top++
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if top-1 < 0 {
				return false
			}
			if stack[top-1] != s[i]-1 && stack[top-1] != s[i]-2 {
				return false
			}
			top--
		} else {
			return false
		}
	}
	if top > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{"))
}
