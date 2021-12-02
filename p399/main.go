package main

import "fmt"

type unionFind struct {
	parents []int
	weights []float64
}

func (u *unionFind) find(x int) int {
	//路径压缩
	if x != u.parents[x] {
		origin := u.parents[x]
		u.parents[x] = u.find(u.parents[x])
		u.weights[x] = u.weights[x] * u.weights[origin]
	}
	return u.parents[x]
}
func (u *unionFind) union(x, y int, val float64) {
	rootX := u.find(x)
	rootY := u.find(y)

	if rootX == rootY {
		return
	}
	u.parents[rootX] = rootY
	u.weights[rootX] = u.weights[y] * val / u.weights[x]
}
func (u *unionFind) isConnected(x, y int) float64 {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX != rootY {
		return -1
	}
	return u.weights[x] / u.weights[y]
}

func newUnionFind(n int) *unionFind {
	parents := make([]int, n)
	weights := make([]float64, n)
	for i := range parents {
		parents[i] = i
		weights[i] = 1
	}
	return &unionFind{
		parents: parents,
		weights: weights,
	}
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	nodeMap := make(map[string]int, len(equations)*2)
	unionFind := newUnionFind(len(equations) * 2)
	id := 0
	//初始化
	for i := range equations {
		val1 := equations[i][0]
		val2 := equations[i][1]
		id1, exists1 := nodeMap[val1]
		if !exists1 {
			nodeMap[val1] = id
			id1 = id
			id++
		}
		id2, exists2 := nodeMap[val2]
		if !exists2 {
			nodeMap[val2] = id
			id2 = id
			id++
		}
		unionFind.union(id1, id2, values[i])
	}
	//运算
	res := make([]float64, len(queries))
	for i := range queries {
		val1 := queries[i][0]
		val2 := queries[i][1]

		id1, exists1 := nodeMap[val1]
		id2, exists2 := nodeMap[val2]
		if !exists1 || !exists2 {
			res[i] = -1
		} else {
			res[i] = unionFind.isConnected(id1, id2)
		}
	}
	return res
}

func main() {
	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{
		2.0, 3.0,
	}
	queries := [][]string{
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}
	fmt.Println(calcEquation(equations, values, queries))
}
