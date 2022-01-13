package main

import "fmt"

func titleToNumber(columnTitle string) int {
	ans := 0
	for _, col := range columnTitle {
		val := int(col - 'A' + 1)
		ans = ans*26 + val
	}

	return ans
}

func main() {
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("FXSHRXW"))
	fmt.Println(titleToNumber(""))
	fmt.Println(titleToNumber("D"))
}
