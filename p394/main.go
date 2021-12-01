package main

import (
	"fmt"
	"strconv"
)

func decodeParts(s string) (string, int) {
	ret := ""
	idx := 0
	for idx < len(s) && s[idx] != ']' {
		if s[idx] >= '0' && s[idx] <= '9' {
			//read num
			numStr := string(s[idx])
			offset := 1
			for s[idx+offset] >= '0' && s[idx+offset] <= '9' {
				numStr = numStr + string(s[idx+offset])
				offset++
			}
			num, _ := strconv.Atoi(numStr)
			idx += offset + 1 //skip the '['
			str, skip := decodeParts(s[idx:])
			for t := 0; t < num; t++ {
				ret = ret + str
			}
			idx += skip
		} else {
			ret = ret + string(s[idx])
			idx++
		}
	}

	return ret, idx + 1
}

func decodeString(s string) string {
	ret, _ := decodeParts(s)
	return ret
}

func main() {
	// s := "zz3[a2[tt]]xx"
	// s := "3[a2[c]]"
	// s := "3[a]2[bc]"
	// s := "2[abc]3[cd]ef"
	s := "abc3[cd]xyz"
	fmt.Println(decodeString(s))
}
