package main

import "fmt"

func intToRoman(num int) string {
	val := []int{
		1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
	}
	str := []string{
		"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
	}

	rest := num
	ret := ""
	for rest != 0 {
		for i := range val {
			if rest >= val[i] {
				rest = rest - val[i]
				ret = ret + str[i]
				break
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(intToRoman((1994)))
}
