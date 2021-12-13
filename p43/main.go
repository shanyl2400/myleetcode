package main

import (
	"fmt"
	"strconv"
)

func add(num1 string, num2 string) string {
	diff := len(num1) - len(num2)
	if diff < 0 {
		diff = -diff
	}
	//padding zeros
	if diff > 0 {
		padding := ""
		for i := 0; i < diff; i++ {
			padding = padding + "0"
		}
		if len(num1) < len(num2) {
			num1 = padding + num1
		} else {
			num2 = padding + num2
		}
	}

	c := byte(0)
	res := ""
	for i := len(num1) - 1; i >= 0; i-- {
		d1 := num1[i] - '0'
		d2 := num2[i] - '0'
		s := d1 + d2 + c
		c = s / byte(10)
		res = string(byte(s%10)+'0') + res
	}
	if c > 0 {
		res = string(c+'0') + res
	}
	return res
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	c := byte(0)
	s := byte(0)
	mid := make([]string, len(num1))
	zeros := ""
	for i := len(num1) - 1; i >= 0; i-- {
		c = 0
		d1 := num1[i] - '0'
		mid[i] = ""
		for j := len(num2) - 1; j >= 0; j-- {
			d2 := num2[j] - '0'
			s = d1*d2 + c
			c = s / 10
			mid[i] = string((s%10)+'0') + mid[i]
		}
		if c > 0 {
			mid[i] = string((c%10)+'0') + mid[i]
		}
		mid[i] = mid[i] + zeros
		zeros = zeros + "0"
	}

	//combine
	prev := mid[0]
	for i := 1; i < len(mid); i++ {
		prev = add(prev, mid[i])
	}
	return prev
}

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 200; j++ {
			a1 := strconv.Itoa(i)
			a2 := strconv.Itoa(j)
			res := multiply(a1, a2)
			r, _ := strconv.Atoi(res)
			if r != i*j {
				fmt.Println("wrong, ", i, j)
			} else {
				fmt.Println("success, ", i, j, res)
			}
		}
	}
	// fmt.Println(multiply("2", "3"))
}
