package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int

func (il Intervals) Len() int {
	return len(il)
}
func (il Intervals) Less(i, j int) bool {
	return il[i][0] < il[j][0]
}

func (il Intervals) Swap(i, j int) {
	temp := il[i]
	il[i] = il[j]
	il[j] = temp
}
func merge(intervals [][]int) [][]int {
	//排序
	if len(intervals) < 1 {
		return nil
	}
	sort.Sort(Intervals(intervals))
	temp := make([]int, 2)
	temp[0] = intervals[0][0]
	temp[1] = intervals[0][1]

	output := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			//合并
			if temp[1] < intervals[i][1] {
				temp[1] = intervals[i][1]
			}
		} else {
			//下一个
			output = append(output, []int{temp[0], temp[1]})
			temp[0] = intervals[i][0]
			temp[1] = intervals[i][1]
		}
	}
	output = append(output, []int{temp[0], temp[1]})

	return output

}

func main() {
	intervals := [][]int{
		{8, 10}, {15, 18}, {1, 3}, {2, 6},
		// {1, 4}, {4, 5},
		// {2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10},
	}
	fmt.Println(merge(intervals))
}
