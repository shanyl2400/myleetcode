package main

import (
	"fmt"
	"strings"
)

var ans [][]string

func backtrace(s string, idx int, dict map[string]bool, arr []string) {
	if idx >= len(s) {
		ans = append(ans, arr)
		return
	}
	for i := 1; i <= len(s)-idx; i++ {
		str := s[idx : idx+i]
		if dict[str] {
			//append & backtrace
			newArr := make([]string, 0, len(arr)+1)
			newArr = append(newArr, arr...)
			newArr = append(newArr, str)
			backtrace(s, idx+i, dict, newArr)
			//delete arr
		}
	}

}

func wordBreak(s string, wordDict []string) []string {
	if s == "" || len(wordDict) == 0 {
		return nil
	}

	ans = make([][]string, 0)
	dict := make(map[string]bool, len(wordDict))
	for i := range wordDict {
		dict[wordDict[i]] = true
	}

	backtrace(s, 0, dict, make([]string, 0))
	ret := make([]string, 0, len(ans))
	for i := range ans {
		ret = append(ret, strings.Join(ans[i], " "))
	}
	return ret
}

func main() {
	// s := "catsanddog"
	// dict := []string{
	// 	"cat", "cats", "and", "sand", "dog",
	// }
	// s := "pineapplepenapple"
	// dict := []string{
	// 	"apple", "pen", "applepen", "pine", "pineapple",
	// }
	s := "catsandog"
	dict := []string{
		"cats", "dog", "sand", "and", "cat",
	}
	res := wordBreak(s, dict)
	for i := range res {
		fmt.Println(res[i])
	}
}
