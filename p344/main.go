package main

import "fmt"

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}

func main() {
	s := []byte{
		'h', 'e', 'l', 'l', 'o',
	}
	reverseString(s)
	fmt.Println(string(s))
}
