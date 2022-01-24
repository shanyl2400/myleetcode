package main

import "fmt"

func calculate(s string) int {
	digit := 0
	preSign := '+'
	stack := make([]int, 0)
	for i, ch := range s {
		isDigit := false
		if ch >= '0' && ch <= '9' {
			digit = digit*10 + int(ch-'0')
			isDigit = true
		}
		if (!isDigit && ch != ' ') || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, digit)
			case '-':
				stack = append(stack, -digit)
			case '*':
				stack[len(stack)-1] *= digit
			default:
				stack[len(stack)-1] /= digit
			}
			preSign = ch
			digit = 0
		}
	}
	ans := 0
	for _, s := range stack {
		ans += s
	}
	return ans
}

func main() {
	s := "3+5 / 2 "
	fmt.Println(calculate(s))
}
