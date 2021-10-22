package main

import "fmt"

func myAtoi(s string) int {
	ret := 0
	//去空格
	skip := 0
	for i := range s {
		if s[i] == ' ' {
			skip++
		} else {
			break
		}
	}
	if skip >= len(s) {
		return 0
	}

	touchNum := false
	for _, v := range s[skip:] {
		x := getNum(v, touchNum)
		if x >= 0 {
			prev := ret
			ret = ret * 10
			ret = ret + x

			if ret < prev {
				//越界
				ret = 2147483648
				break
			}

			touchNum = true
		} else if x == -2 {
			break
		} else {
			touchNum = true
		}
	}

	if s[skip] == '-' {
		ret = -ret
	}
	if ret < -2147483648 {
		return -2147483648
	}
	if ret > 2147483647 {
		return 2147483647
	}
	return ret
}
func getNum(s rune, touchNum bool) int {
	if s >= '0' && s <= '9' {
		return int(s - '0')
	}
	if (s == '-' || s == '+') && !touchNum {
		return -1
	} else {
		return -2
	}
}
func main() {
	fmt.Println(myAtoi("18446744073709551617"))
}
