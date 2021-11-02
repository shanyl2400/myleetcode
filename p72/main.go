package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}
	word1 = word1 + "0"
	word2 = word2 + "0"

	state := make([][]int, len(word1))
	for i := range word1 {
		state[i] = make([]int, len(word2))
	}

	for i := range word1 {
		for j := range word2 {
			if i == 0 && j == 0 {
				continue
			}
			dist1 := 100000
			dist2 := 100000
			dist3 := 100000
			if i > 0 {
				dist1 = state[i-1][j] + 1
			}
			if j > 0 {
				dist2 = state[i][j-1] + 1
			}
			if i > 0 && j > 0 {
				dist3 = state[i-1][j-1]
				if word1[i-1] != word2[j-1] {
					dist3 = dist3 + 1
				}
			}
			state[i][j] = min(dist1, dist2, dist3)
		}
	}

	// for i := 0; i < len(word1); i++ {
	// 	for j := 0; j < len(word2); j++ {
	// 		fmt.Print(state[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	return state[len(word1)-1][len(word2)-1]
}
func min(x, y, z int) int {
	if y < x {
		x = y
	}
	if z < x {
		x = z
	}
	return x
}

func main() {
	word1 := "horse"
	word2 := "ros"
	// word1 := "intention"
	// word2 := "execution"
	// word1 := "a"
	// word2 := "ab"
	// word1 := "plasma"
	// word2 := "altruism"
	// word1 := "a"
	// word2 := "b"
	fmt.Println(minDistance(word1, word2))
}
