package main

import "fmt"

func max(x, y int) int {
	if x < y {
		x = y
	}
	return x
}

func maxProfit(prices []int) int {
	size := len(prices)

	prev1, prev2 := 0, -prices[0]
	cur1, cur2 := 0, 0

	for i := 1; i < size; i++ {
		cur1 = max(prev1, prev2+prices[i])
		cur2 = max(prev2, prev1-prices[i])

		prev1, prev2 = cur1, cur2
	}
	return max(cur1, cur2)
}

func main() {
	prices := []int{
		// 7, 1, 5, 3, 6, 4,
		1, 2, 3, 4, 5,
		// 7, 6, 4, 3, 1,
	}
	fmt.Println(maxProfit(prices))
}
