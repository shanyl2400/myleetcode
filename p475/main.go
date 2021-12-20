package main

import "fmt"

func findRadius(houses []int, heaters []int) int {
	if len(houses) == 0 || len(heaters) == 0 {
		return 0
	}
	radius := 0
	for i := range houses {
		minDiff := 1000000001
		for j := range heaters {
			diff := heaters[j] - houses[i]
			if diff < 0 {
				diff = -diff
			}
			if minDiff > diff {
				minDiff = diff
			}
		}
		if radius < minDiff {
			radius = minDiff
		}
	}
	return radius
}

func main() {
	houses := []int{
		// 1, 2, 3, 4,
		// 1, 2, 3,
		1, 5,
	}
	heaters := []int{
		// 1, 4,
		// 2,
		10,
	}

	fmt.Println(findRadius(houses, heaters))
}
