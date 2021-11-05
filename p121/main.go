package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0
	for i := range prices {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > max {
			max = prices[i] - minPrice
		}
	}
	return max
}

func main() {
	nums := []int{
		// 7, 1, 5, 3, 6, 4,
		// 7, 6, 5, 4, 2, 1,
		// 7, 1, 6, 3, 5, 4,
		// 7, 2, 6, 1, 5, 4,
		// 3, 2, 6, 5, 0, 3,
		// 2, 1, 4, 5, 2, 9, 7,
		1, 3, 5, 4, 3, 7, 6, 9, 2, 4,
		// 2, 1, 2, 0, 1,
	}
	fmt.Println(maxProfit(nums))
}
