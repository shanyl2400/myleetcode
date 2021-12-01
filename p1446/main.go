package main

import "fmt"

func maxPower(s string) int {
	if len(s) < 1 {
		return 0
	}
	max := 1
	curMax := 1
	prev := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == prev {
			curMax++
			if curMax > max {
				max = curMax
			}
		} else {
			curMax = 1
			prev = s[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxPower("tourist"))
}
