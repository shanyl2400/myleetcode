package main

import "fmt"

func findAnagrams(s string, p string) []int {
	wordMap := make(map[byte]int)
	targetMap := make(map[byte]int)
	ret := make([]int, 0)
	diff := 0
	left := 0
	right := 0
	for i := range p {
		wordMap[p[i]]++
		diff++
	}
	for right < len(s) {
		nw := s[right]
		// fmt.Println(left, right)

		if wordMap[nw] > 0 {
			if targetMap[nw] < wordMap[nw] {
				diff--
			}
			targetMap[nw]++
		}

		if right-left >= len(p) {
			ow := s[left]
			if wordMap[ow] > 0 {
				if targetMap[ow] <= wordMap[ow] {
					diff++
				}
				targetMap[ow]--
			}
			left++
		}
		if diff == 0 {
			ret = append(ret, left)
		}
		right++
	}

	return ret
}

func main() {
	fmt.Println(findAnagrams(
		// "cbaebabacd",
		// "abc",
		// "abab",
		// "ab",
		"dinitrophenylhydrazinetrinitrophenylmethylnitramine",
		"trinitrophenylmethylnitramine", //19 20 21 22
	))
}

// a:1 e:3 h:2 i:4 l:2 m:2 n:4 o:1 p:1 r:3 t:4 y:2
// a:1 e:3 h:3 i:3 l:1 n:3 o:2 p:2 r:4 t:3 y:2
