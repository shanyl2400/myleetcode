package main

import "fmt"

func minWindow(s string, t string) string {
	left := 0
	right := 0

	dist := 0
	begin := 0
	size := len(s) + 1

	tFreq := make([]int, 128)
	//统计t中的单词
	for i := range t {
		tFreq[t[i]]++
	}
	winFreq := make([]int, 128)
	for right < len(s) {
		cur := s[right]
		//字母不在s中，向后移动
		if tFreq[cur] == 0 {
			right++
			continue
		}
		//window中缺少cur，移动right
		if winFreq[cur] < tFreq[cur] {
			dist++
		}
		winFreq[s[right]]++
		right++

		for dist == len(t) {
			if right-left < size {
				size = right - left
				begin = left
			}

			rest := s[left]
			//不需要的单词
			if tFreq[rest] == 0 {
				left++
				continue
			}

			//dest数量刚好
			if tFreq[rest] == winFreq[rest] {
				dist--
			}
			//从window中去掉rest
			left++
			winFreq[rest]--
		}
	}
	if size > len(s) {
		return ""
	}
	return s[begin : begin+size]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}
