package main

import "fmt"

func trap(height []int) int {
	first := 0
	last := 1

	sum := 0
	for first < len(height)-1 {
		//遇到阻挡
		curSize := 0
		isFinish := false
		for height[first] > height[last] {
			curSize += height[first] - height[last]
			last++

			//查到结尾还是没有高点
			if last >= len(height) {
				isFinish = true
				break
			}
		}
		//last已经到结尾，还是没有找到高点
		if isFinish {
			if height[first] > 0 {
				height[first]--
				last = first + 1
			} else {
				first++
				last = first + 1
			}
			continue
		}

		sum += curSize
		// fmt.Println("first:", first, "last:", last, "sum:", sum)

		first = last
		last++
	}
	return sum
}

func main() {
	fmt.Println(trap([]int{
		// 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
		4, 2, 0, 3, 2, 5,
		// 4, 2, 3,
	}))
}
