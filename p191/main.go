package main

import "fmt"

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num = num & (num - 1)
		cnt++
	}

	return cnt
}

func main() {
	// num := uint32(0b00000000000000000000000000001011)
	// num := uint32(0b00000000000000000000000010000000)
	num := uint32(0b11111111111111111111111111111101)
	fmt.Println(hammingWeight(num))
}
