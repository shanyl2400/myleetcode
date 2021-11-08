package main

import "fmt"

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	recordS := make([]int, 10)
	recordG := make([]int, 10)
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			recordS[secret[i]-'0']++
			recordG[guess[i]-'0']++
		}
	}
	for i := range recordS {
		if recordG[i] >= recordS[i] {
			recordG[i] = recordS[i]
		}
		cows += recordG[i]
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func main() {
	secret := "1807"
	guess := "7810"
	fmt.Println(getHint(secret, guess))
}
