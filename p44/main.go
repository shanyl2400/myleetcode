package main

import "fmt"

func isMatch(s string, p string) bool {
	state := make([][]bool, len(p)+1)
	for i := range state {
		state[i] = make([]bool, len(s)+1)
	}
	state[0][0] = true
	//特殊处理开头连续的*
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			state[i+1][0] = true
		} else {
			break
		}
	}

	for i := 1; i <= len(p); i++ { //p[i]
		for j := 1; j <= len(s); j++ { //s[j]
			pw := p[i-1]
			sw := s[j-1]
			if pw == '*' {
				state[i][j] = state[i-1][j] || state[i][j-1]
			} else if pw == '?' || pw == sw {
				state[i][j] = state[i-1][j-1]
			}
		}
	}
	// for i := range state {
	// 	fmt.Println(state[i])
	// }

	return state[len(p)][len(s)]
}

func main() {
	// s := "aaabzbz"
	// p := "aa?b*zbz"
	// s := ""
	// p := "*****"
	// s := "abcabczzzde"
	// p := "*abc???de*"
	// s := "mississippi"
	// p := "m??*ss*?i*pi"
	s := "adceb"
	p := "*a*b"
	fmt.Println(isMatch(s, p))
}
