package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 1 && s[0] >= '1' && s[0] <= '9' {
		return true
	} else if (len(s) == 2) && (s[0] == '1' && s[1] >= '0' && s[1] <= '9') ||
		(s[0] == '2' && s[1] >= '0' && s[1] <= '6') {
		return true
	}
	return false
}
func numDecodings(s string) int {
	if s == "" {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[len(s)] = 1
	if isValid(s[len(s)-1:]) {
		dp[len(s)-1] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		if isValid(s[i : i+1]) {
			dp[i] = dp[i] + dp[i+1]
		}

		if isValid(s[i : i+2]) {
			dp[i] = dp[i] + dp[i+2]
		}

	}

	return dp[0]
}

func main() {
	fmt.Println(numDecodings("123123"))
}
