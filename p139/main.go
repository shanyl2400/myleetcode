package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	state := make([]bool, len(s)+1)
	for i := range s {
		if i == 0 || state[i] {
			words := findWords(s[i:], wordDict)
			for j := range words {
				state[i+len(words[j])] = true
			}
		}
	}
	return state[len(s)]
}
func findWords(s string, wordDict []string) []string {
	ret := make([]string, 0)
	for i := range wordDict {
		if len(s) < len(wordDict[i]) {
			continue
		}
		if s[:len(wordDict[i])] == wordDict[i] {
			ret = append(ret, wordDict[i])
		}
	}
	return ret
}

func main() {
	// s := "leetcode"
	// words := []string{
	// 	"leet", "code",
	// }
	// s := "applepenapple"
	// words := []string{
	// 	"apple", "pen",
	// }
	// s := "catsandog"
	// words := []string{
	// 	"cats", "dog", "sand", "and", "cat",
	// }

	// s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	// words := []string{
	// 	"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa",
	// }
	// s := "bb"
	// words := []string{
	// 	"a", "b", "bbb", "bbbb",
	// }
	// s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	// words := []string{
	// 	"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba",
	// }
	s := "aaabbb"
	words := []string{}
	fmt.Println(wordBreak(s, words))
}
