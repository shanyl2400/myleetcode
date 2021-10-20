package main

import "fmt"

func letterCombinations(digits string) []string {
	return doLetterCombinations(digits, 0)
}
func doLetterCombinations(digits string, index int) []string {
	if index >= len(digits) {
		return nil
	}
	subLetters := doLetterCombinations(digits, index+1)

	letters := numToLetters(digits[index])
	ret := make([]string, 0)
	for i := range letters {
		if len(subLetters) < 1 {
			ret = append(ret, string(letters[i]))
		} else {
			for j := range subLetters {
				ret = append(ret, string(letters[i])+subLetters[j])
			}
		}
	}
	return ret
}
func numToLetters(digit byte) []byte {
	index := digit - '0'
	if index == 7 {
		return []byte{
			'p', 'q', 'r', 's',
		}
	}
	if index == 8 {
		return []byte{
			't', 'u', 'v',
		}
	}
	if index == 9 {
		return []byte{
			'w', 'x', 'y', 'z',
		}
	}
	index = index - 2
	ret := make([]byte, 3)
	for i := byte(0); i < 3; i++ {
		x := 'a' + i + index*byte(3)
		ret[i] = x
	}

	return ret
}

func main() {
	fmt.Println(letterCombinations(""))
}
