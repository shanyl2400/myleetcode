package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	cnt := make(map[byte]int)
	for _, i := range tasks {
		cnt[i]++
	}
	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)

}

func main() {
	tasks := []byte{
		// 'A', 'A', 'A', 'B', 'B', 'B',
		'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G',
	}
	n := 2

	fmt.Println(leastInterval(tasks, n))
}
