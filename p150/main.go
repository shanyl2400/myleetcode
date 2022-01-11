package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	vals []int
	top  int
}

func (s *stack) push(val int) {
	if len(s.vals) <= s.top {
		s.vals = append(s.vals, val)
	} else {
		s.vals[s.top] = val
	}
	s.top++
}
func (s *stack) pop() int {
	s.top--
	return s.vals[s.top]
}
func (s *stack) peek() int {
	return s.vals[s.top-1]
}
func newStack() *stack {
	return &stack{
		top:  0,
		vals: make([]int, 0),
	}
}

func evalRPN(tokens []string) int {
	s := newStack()
	for _, token := range tokens {
		switch token {
		case "+":
			val1 := s.pop()
			val2 := s.pop()
			s.push(val1 + val2)
		case "-":
			val1 := s.pop()
			val2 := s.pop()
			s.push(val2 - val1)
		case "*":
			val1 := s.pop()
			val2 := s.pop()
			s.push(val1 * val2)
		case "/":
			val1 := s.pop()
			val2 := s.pop()
			s.push(val2 / val1)
		default:
			val, _ := strconv.Atoi(token)
			s.push(val)
		}
	}

	return s.peek()
}

func main() {
	tokens := []string{
		// "2", "1", "+", "3", "*",
		// "4", "13", "5", "/", "+",
		"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
	}
	fmt.Println(evalRPN(tokens))
}
