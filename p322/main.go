package main

import "fmt"

func min(x, y int) int {
	if x == 0 {
		return y
	} else if y == 0 {
		return x
	}
	if x > y {
		x = y
	}
	return x
}

func coinChange(coins []int, amount int) int {
	state := make([]int, amount+1)
	for i := range state {
		state[i] = -1
	}
	state[0] = 0

	for i := 1; i < len(state); i++ {
		for j := range coins {
			if i >= coins[j] && state[i-coins[j]] >= 0 {
				if state[i] == -1 {
					state[i] = state[i-coins[j]] + 1
				} else {
					state[i] = min(state[i], state[i-coins[j]]+1)
				}
			}
		}
	}

	return state[amount]
}

func main() {
	coins := []int{
		1, 2, 5,
		// 2,
		// 1,
		// 2, 5, 10, 1,
	}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
