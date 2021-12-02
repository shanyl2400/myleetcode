package main

import "fmt"

type unionFind struct {
	parents []int
}

func (u *unionFind) find(x int) int {
	if x != u.parents[x] {
		u.parents[x] = u.find(u.parents[x])
	}
	return u.parents[x]
}
func (u *unionFind) union(x, y int) {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return
	}
	u.parents[rootX] = rootY
}
func (u *unionFind) isConnected(x, y int) bool {
	rootX := u.find(x)
	rootY := u.find(y)
	if rootX == rootY {
		return true
	}
	return false
}

func newUnionFind(n int) *unionFind {
	parents := make([]int, n)
	equals := make([]bool, n)
	for i := range parents {
		parents[i] = i
		equals[i] = true
	}
	return &unionFind{
		parents: parents,
	}
}

func equationsPossible(equations []string) bool {
	unionFind := newUnionFind(26)
	for i := range equations {
		if equations[i][1:3] == "==" {
			//equal
			id1 := int(equations[i][0] - 'a')
			id2 := int(equations[i][3] - 'a')

			unionFind.union(id1, id2)
		}
	}
	for i := range equations {
		if equations[i][1:3] == "!=" {
			val1 := equations[i][0]
			val2 := equations[i][3]
			if val1 == val2 {
				return false
			}
			id1 := int(equations[i][0] - 'a')
			id2 := int(equations[i][3] - 'a')

			if unionFind.isConnected(id1, id2) {
				return false
			}
		}

	}
	return true
}

func main() {
	eq0 := []string{
		"a==b",
		"b!=a",
	}

	eq1 := []string{
		"a==b",
		"b==a",
	}

	eq2 := []string{
		"a==b",
		"b==c",
		"a==c",
	}

	eq3 := []string{
		"a==b",
		"b!=c",
		"c==a",
	}

	eq4 := []string{
		"c==c",
		"b==d",
		"x!=z",
	}

	eq5 := []string{
		"a!=a",
	}

	eq6 := []string{
		"b!=c", "a==b", "e!=d", "b!=f", "a!=b",
	}
	fmt.Println(equationsPossible(eq0)) //false
	fmt.Println(equationsPossible(eq1)) //true
	fmt.Println(equationsPossible(eq2)) //true
	fmt.Println(equationsPossible(eq3)) //false
	fmt.Println(equationsPossible(eq4)) //true
	fmt.Println(equationsPossible(eq5)) //false
	fmt.Println(equationsPossible(eq6)) //false
}
