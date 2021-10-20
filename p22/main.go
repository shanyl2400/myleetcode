package main

import "fmt"

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	for i := 0; i < n; i++ {
		//build i
		nums := n - i
		start := ""
		end := ""
		for j := 0; j < nums; j++ {
			start = start + "("
			end = end + ")"
		}
		commas := start + end

		if i > 0 {
			subRes := generateParenthesis(i)
			fmt.Println(subRes)
			for x := range subRes {
				ret = append(ret, commas+subRes[x])
			}
		} else {
			ret = append(ret, commas)
		}
	}

	return ret
}

func main() {
	fmt.Println(generateParenthesis(3))
}
