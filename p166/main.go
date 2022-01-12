package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	//整除
	remainder := numerator % denominator
	if remainder == 0 {
		return fmt.Sprintf("%d", numerator/denominator)
	}

	s := []byte{}
	if numerator < 0 != (denominator < 0) {
		s = append(s, '-')
	}
	numerator = abs(numerator)
	denominator = abs(denominator)

	integerPart := numerator / denominator
	s = append(s, strconv.Itoa(integerPart)...)
	s = append(s, '.')

	remainderMap := make(map[int]int)
	remainder = numerator % denominator
	for remainder != 0 && remainderMap[remainder] == 0 {
		remainderMap[remainder] = len(s)
		remainder *= 10
		s = append(s, '0'+byte(remainder/denominator))
		remainder %= denominator
	}

	if remainder > 0 {
		insertIndex := remainderMap[remainder]
		s = append(s[:insertIndex], append([]byte{'('}, s[insertIndex:]...)...)
		s = append(s, ')')

	}

	return string(s)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	fmt.Println(fractionToDecimal(-50, 8))
}
