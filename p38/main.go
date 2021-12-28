package main

import "fmt"

func countAndSay(n int) string {
	say := "1"
	for i := 2; i <= n; i++ {
		res := ""
		cnt := 0
		for idx := range say {
			if idx == 0 || say[idx-1] == say[idx] {
				cnt++
			} else {
				res = fmt.Sprintf("%v%v%v", res, cnt, say[idx-1]-'0')
				cnt = 1
			}
		}
		if cnt > 0 {
			res = fmt.Sprintf("%v%v%v", res, cnt, say[len(say)-1]-'0')
		}
		say = res
	}

	return say
}

func main() {
	fmt.Println(countAndSay(4))

}
