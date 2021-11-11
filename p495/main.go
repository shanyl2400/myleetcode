package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	ret := 0
	for i := 1; i < len(timeSeries); i++ {
		diff := timeSeries[i] - timeSeries[i-1]
		if diff >= duration {
			ret = ret + duration
		} else {
			ret = ret + diff
		}
	}
	ret = ret + duration
	return ret
}

func main() {
	series := []int{
		// 1, 2, 3, 4, 5, 6, 7, 8, 9,
		1, 2,
	}
	fmt.Println(findPoisonedDuration(series, 2))
}
