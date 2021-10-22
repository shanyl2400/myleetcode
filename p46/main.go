package main

import "fmt"

func permute(nums []int) [][]int {
	switch len(nums) {
	case 0:
		return nil
	case 1:
		return [][]int{{nums[0]}}
	case 2:
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	ret := make([][]int, 0)
	for i := range nums {
		res := make([]int, 0, len(nums))
		for j := range nums {
			if j != i {
				res = append(res, nums[j])
			}
		}
		sub := permute(res)
		for p := range sub {
			ret = append(ret, append(sub[p], nums[i]))
		}
	}
	return ret
}

func permute2(nums []int) [][]int {
	res := make([][]int, 0)
	used := make([]bool, len(nums))
	path := make([]int, 0)

	dfs(nums, 0, path, used, &res)
	return res
}

func dfs(nums []int, depth int, path []int, used []bool, res *[][]int) {
	if depth == len(nums) {
		*res = append(*res, path)
		return
	}

	for i := range nums {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		dfs(nums, depth+1, path, used, res)
		used[i] = false
		newPath := make([]int, len(path)-1)
		copy(newPath, path[:len(path)-1])
		path = newPath
	}
}

func main() {
	nums := []int{1, 2, 3}
	ret := permute2(nums)
	fmt.Println(ret)
}
