package main

import "fmt"

func trailingZeroes(n int) int {
	cnt := 0
	for i := 5; i <= n; i += 5 {
		tmp := i
		for tmp%5 == 0 {
			cnt++
			tmp /= 5
		}

	}
	return cnt
}
func main() {
	fmt.Println(trailingZeroes(30))
}
