package main

import "fmt"

func doMyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := doMyPow(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return doMyPow(x, n)
	}
	return 1 / doMyPow(x, -n)
}
func main() {
	// fmt.Println(myPow(-2, 2))
	// fmt.Println(myPow(0.01, 2147483647))
	fmt.Println(myPow(3, 8)) //8192
}
