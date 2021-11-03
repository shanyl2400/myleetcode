package main

import "fmt"

type item struct {
	val byte
	idx int
}
type stack struct {
	mem []*item
	top int
}

func (s *stack) push(b *item) {
	s.mem[s.top] = b
	s.top++
}
func (s *stack) pop() *item {
	s.top--
	return s.mem[s.top]
}
func (s *stack) reset() {
	s.top = 0
}
func (s *stack) peek() *item {
	return s.mem[s.top-1]
}
func (s *stack) print() {
	for i := s.top - 1; i >= 0; i-- {
		fmt.Print(s.mem[i].val, ":", s.mem[i].idx, "   ")
	}
	fmt.Println()
}

func newStack(size int) *stack {
	return &stack{
		mem: make([]*item, size),
		top: 0,
	}
}
func maximalRectangle(matrix [][]byte) int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = matrix[i][j] - '0'
			if i > 0 && matrix[i][j] > 0 {
				matrix[i][j] = matrix[i][j] + matrix[i-1][j]
			}
		}
	}
	if len(matrix) < 1 {
		return 0
	}
	st := newStack(len(matrix[0]) + 2)
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		st.reset()
		st.push(&item{val: 0, idx: 0})

		matrix[i] = append(matrix[i], 0)
		for j := range matrix[i] {
			val := matrix[i][j]
			if matrix[i][j] >= st.peek().val {
				st.push(&item{
					val: val,
					idx: j + 1,
				})
			} else {
				min := j + 1
				for st.peek().val > val {
					it := st.pop()
					area := (j + 1 - it.idx) * int(it.val)
					if maxArea < area {
						maxArea = area
					}
					if it.idx < min {
						min = it.idx
					}
				}
				st.push(&item{val: val, idx: min})
			}
		}
	}

	return maxArea
}
func main() {
	matrix := [][]byte{
		// {'1', '0', '1', '0', '0'},
		// {'1', '0', '1', '1', '1'},
		// {'1', '1', '1', '1', '1'},
		// {'1', '0', '0', '1', '0'},
	}

	// matrix := [][]byte{
	// 	{'1'},
	// }

	fmt.Println(maximalRectangle(matrix))
}
