package main

import "fmt"

var data = []int{
	1,
	3,
	9,
	27,
	81,
	243,
	729,
	2187,
	6561,
	19683,
	59049,
	177147,
	531441,
	1594323,
	4782969,
	14348907,
	43046721,
	129140163,
	387420489,
	1162261467,
}

func isPowerOfThree(n int) bool {
	for i := range data {
		if n == data[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPowerOfThree(27))
}
