package main

import "fmt"

type temperature struct {
	temp int
	idx  int
}
type stack struct {
	temperatures []*temperature
	top          int
}

func (s *stack) push(temperature *temperature) {
	s.top++
	s.temperatures[s.top] = temperature
}
func (s *stack) pop() *temperature {
	ret := s.temperatures[s.top]
	s.top--
	return ret
}
func (s *stack) size() int {
	return s.top + 1
}
func (s *stack) peek() *temperature {
	return s.temperatures[s.top]
}
func (s *stack) print() {
	for i := range s.temperatures {
		if s.temperatures[i] != nil {
			fmt.Print(s.temperatures[i])
		}
	}
	fmt.Println()
}

func newStack(max int) *stack {
	return &stack{
		top:          -1,
		temperatures: make([]*temperature, max),
	}
}

func dailyTemperatures(temperatures []int) []int {
	stack := newStack(len(temperatures))
	ret := make([]int, len(temperatures))
	for i := range temperatures {
		for stack.size() > 0 && stack.peek().temp < temperatures[i] {
			d := stack.pop()
			ret[d.idx] = i - d.idx
		}
		stack.push(&temperature{
			idx:  i,
			temp: temperatures[i],
		})
	}
	return ret
}

func main() {
	// temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	// temperatures := []int{30, 40, 50, 60}
	temperatures := []int{30, 60, 90}
	fmt.Println(dailyTemperatures(temperatures))
}
