package main

import "fmt"

func romanToInt(s string) int {
	valMap := map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 {
			s0 := s[i : i+2]
			val, exist := valMap[s0]
			if exist {
				ret = ret + val
				i = i + 1
				continue
			}
		}

		val := valMap[string(s[i])]
		ret = ret + val
	}

	return ret
}

func main() {
	fmt.Println(romanToInt(("MCMXCIV")))
}
