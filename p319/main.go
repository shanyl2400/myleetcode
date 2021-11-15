package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}

func main() {
	fmt.Println(bulbSwitch(100000))
}
