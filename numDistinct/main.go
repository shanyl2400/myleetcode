package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	records := make([][]int, len(t))
	for i := range t {
		records[i] = make([]int, len(s))
	}

	for i := range t {
		for j := range s {
			history := 0
			grand := 0
			if j-1 >= 0 {
				history = records[i][j-1]
				if i-1 >= 0 {
					grand = records[i-1][j-1]
				}
			}

			val := 0
			if s[j] == t[i] {
				if i == 0 {
					val = history + 1
				} else {
					val = history + grand + val
				}
			} else {
				val = history
			}

			records[i][j] = val
		}
	}

	return records[len(t)-1][len(s)-1]
}

func main() {
	// fmt.Println(numDistinct("adbdadeecadeadeccaeaabdabdbcdabddddabcaaadbabaaedeeddeaeebcdeabcaaaeeaeeabcddcebddebeebedaecccbdcbcedbdaeaedcdebeecdaaedaacadbdccabddaddacdddc", "bcddceeeebecbc"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}
