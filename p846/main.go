package main

import (
	"fmt"
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	handMap := make(map[int]int, len(hand))
	for i := range hand {
		handMap[hand[i]]++
	}
	handKey := make([]int, 0, len(handMap))
	for k := range handMap {
		handKey = append(handKey, k)
	}
	sort.Ints(handKey)

	for _, key := range handKey {
		num := handMap[key]
		if num <= 0 {
			continue
		}
		for i := 0; i < groupSize; i++ {
			handMap[key+i] -= num
			if handMap[key+i] < 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	hand := []int{
		1, 2, 3, 6, 2, 3, 4, 7, 8,
		// 1, 2, 3, 4, 5,
	}
	groupSize := 3
	// groupSize := 4

	fmt.Println(isNStraightHand(hand, groupSize))
}
