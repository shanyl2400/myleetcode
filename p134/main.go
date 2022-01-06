package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	sumGas := 0
	sumCost := 0
	for i := range gas {
		sumGas += gas[i]
		sumCost += cost[i]
	}
	if sumGas < sumCost {
		return -1
	}

	start := 0
	curGas := 0
	for i := 0; i < len(gas); i++ {
		curGas = curGas - cost[i] + gas[i]
		if curGas < 0 {
			curGas = 0
			start = i + 1
		}
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
