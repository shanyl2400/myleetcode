package main

import "fmt"

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	ans := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(countPrimes(10))
}
