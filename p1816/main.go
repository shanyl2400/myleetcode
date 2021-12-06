package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	return strings.Join(strings.Split(s, " ")[:k], " ")
}

func main() {
	s := "Hello how are you Contestant"
	fmt.Println(truncateSentence(s, 4))
}
