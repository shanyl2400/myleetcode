package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type node struct {
	end    int
	height int
}

type nodeList []node

func (n nodeList) Len() int {
	return len(n)
}
func (n nodeList) Less(i, j int) bool {
	return n[i].height > n[j].height
}
func (n nodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n *nodeList) Push(v interface{}) {
	*n = append(*n, v.(node))
}
func (n *nodeList) Pop() interface{} {
	a := *n
	v := a[len(a)-1]
	*n = a[:len(a)-1]
	return v
}

func getSkyline(buildings [][]int) [][]int {
	size := len(buildings)
	boundaries := make([]int, 0, size*2)
	for i := range buildings {
		boundaries = append(boundaries, buildings[i][0], buildings[i][1])
	}
	sort.Ints(boundaries)

	idx := 0
	h := nodeList{}
	ans := make([][]int, 0)
	for _, boundary := range boundaries {
		for idx < size && buildings[idx][0] <= boundary {
			//building is in boundary
			heap.Push(&h, node{
				end:    buildings[idx][1],
				height: buildings[idx][2],
			})
			idx++
		}

		//pop all out of range
		for len(h) > 0 && h[0].end <= boundary {
			heap.Pop(&h)
		}
		max := 0
		if len(h) > 0 {
			max = h[0].height
		}
		if len(ans) == 0 || max != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, max})
		}
	}
	return ans
}

func main() {
	building := [][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}
	// building := [][]int{
	// 	{0, 2, 3},
	// 	{2, 5, 3},
	// }
	fmt.Println(getSkyline(building))
}
