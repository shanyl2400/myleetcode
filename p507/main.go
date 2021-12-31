package main

import "fmt"

func checkPerfectNumber(num int) bool {
	return num == 6 || num == 28 || num == 496 || num == 8128 || num == 33550336
}

func main() {
	fmt.Println(checkPerfectNumber(1))
}
