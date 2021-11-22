package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	maxNum := int(math.Sqrt(float64(n)))

	state := make([]int, n+1)
	for i := range state {
		for j := 1; j <= maxNum; j++ {
			prev := i - j*j
			if prev >= 0 && (state[prev]+1 < state[i] || state[i] == 0) {
				state[i] = state[prev] + 1
			}
		}
	}

	return state[n]
}

func main() {
	fmt.Println(numSquares(12))
}
