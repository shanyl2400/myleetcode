package main

import "fmt"

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	for i := 0; i < 32; i++ {
		ret <<= 1
		ret |= num & 1
		num >>= 1
	}
	return ret
}
func main() {
	n := uint32(0b0)
	fmt.Println(reverseBits(n))
}
