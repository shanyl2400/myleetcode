package main

import "fmt"

type rect struct {
	height int
	start  int
}

func largestRectangleArea(heights []int) int {
	stack := make([]*rect, len(heights)+2)
	top := 0
	maxArea := 0
	minStart := 0

	heights = append(heights, 0)
	stack[top] = &rect{height: 0, start: 0}
	top++

	for i := range heights {
		minStart = i
		if heights[i] >= stack[top-1].height {
			stack[top] = &rect{height: heights[i], start: i + 1}
			top++
		} else {
			for heights[i] < stack[top-1].height {
				//pop
				hr := stack[top-1]
				top--

				//calculate area
				area := (i - hr.start + 1) * hr.height
				fmt.Println("height:", hr.height, "start:", hr.start, "i:", i, "area:", area)
				if maxArea < area {
					maxArea = area
				}
				minStart = hr.start
			}
			//push stack
			stack[top] = &rect{height: heights[i], start: minStart}
			top++
		}
	}

	return maxArea
}

func main() {
	heights := []int{
		// 2, 1, 5, 6, 2, 3,
		2, 4,
	}
	fmt.Println(largestRectangleArea(heights))
}
