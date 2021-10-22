package main

import "fmt"

func reverse(x int) int {
	ret := 0
	flag := false
	for x != 0 {
		v := x % 10
		x = x / 10
		if v == 0 && !flag {
			continue
		}
		ret = ret*10 + v
		flag = true
	}
	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return ret
}
func main() {
	fmt.Println(reverse(123))
}
