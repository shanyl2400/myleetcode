package main

import "fmt"

func isConnected(s1, s2 string) bool {
	cnt := 0
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	graph := make([][]int, len(wordList))
	target := -1
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			target = i
		}
		graph[i] = make([]int, len(wordList))
	}
	if target == -1 {
		return 0
	}

	for i := 0; i < len(wordList); i++ {
		for j := i + 1; j < len(wordList); j++ {
			if isConnected(wordList[i], wordList[j]) {
				graph[i][j] = 1
				graph[j][i] = 1
			}
		}
	}
	visited := make([]bool, len(wordList))
	return dfs(graph, visited, len(wordList)-1, target, 1)
}

func dfs(graph [][]int, visited []bool, start, target int, idx int) int {
	if start == target {
		return idx
	}
	minCnt := -1
	for i := range graph[start] {
		if graph[start][i] == 1 && !visited[i] {
			//mark as visited
			visited[i] = true
			//connected
			cnt := dfs(graph, visited, i, target, idx+1)
			if cnt > 0 && (minCnt == -1 || minCnt > cnt) {
				//find
				minCnt = cnt
			}
			//backtracking
			visited[i] = false
		}
	}
	return minCnt
}

func main() {
	begin := "hit"
	end := "coz"
	wordList := []string{
		"hot", "dot", "dog", "lot", "log", "cog", "coz",
		// "hot", "dot", "dog", "lot", "log",
	}
	fmt.Println(ladderLength(begin, end, wordList))
}
