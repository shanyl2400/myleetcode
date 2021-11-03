package main

import "fmt"

func numTrees(n int) int {
	ret := float64(1)
	for i := 1; i <= n; i++ {
		ret = float64(ret) * float64(n+i) / float64(i+1)
	}

	return int(ret)
}

func main() {
	fmt.Println(numTrees(19))
}
