package main

import "fmt"

type num struct {
	val int
	idx int
}
type heap struct {
	size int
	mem  []num
}

func newHeap() *heap {
	return &heap{
		size: 0,
		mem:  make([]num, 1),
	}
}
func (h *heap) push(val int, idx int) {
	h.size = h.size + 1
	i := h.size
	if h.size >= len(h.mem) {
		h.mem = append(h.mem, num{})
	}

	for ; i > 1 && h.mem[i/2].val < val; i /= 2 {
		h.mem[i].idx = h.mem[i/2].idx
		h.mem[i].val = h.mem[i/2].val
	}
	h.mem[i] = num{
		val: val,
		idx: idx,
	}
}

func (h *heap) pop() num {
	if h.size < 1 {
		return num{idx: -1}
	}
	i, child := 1, 0
	val := h.mem[1]
	latest := h.mem[h.size]
	h.size--

	for i = 1; i*2 <= h.size; i = child {
		child = i * 2
		if child != h.size && h.mem[child+1].val > h.mem[child].val {
			child++
		}

		if latest.val < h.mem[child].val {
			h.mem[i] = h.mem[child]
		} else {
			break
		}
	}
	h.mem[i] = latest
	return val
}

func maxSlidingWindow(nums []int, k int) []int {
	h := newHeap()
	for i := 0; i < k; i++ {
		h.push(nums[i], i)
	}
	ret := []int{
		h.mem[1].val,
	}
	for i := k; i < len(nums); i++ {
		h.push(nums[i], i)
		for h.mem[1].idx <= i-k {
			h.pop()
		}
		ret = append(ret, h.mem[1].val)
	}

	return ret
}
func maxSlidingWindow3(nums []int, k int) []int {
	left := 0
	right := k - 1
	max := 0
	ret := make([]int, 0)
	for right < len(nums) {
		if left == 0 {
			max = findMax(nums, left, right, nums[0], false)
		} else {
			if nums[left-1] == max {
				max = findMax(nums, left, right, max, true)
			}
			if nums[right] > max {
				max = nums[right]
			}
		}

		ret = append(ret, max)

		left++
		right++
	}

	return ret
}
func findMax(nums []int, start, end, preMax int, checkPre bool) int {
	max := nums[start]
	for i := start; i <= end; i++ {
		if checkPre && nums[i] == preMax {
			return preMax
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{
		// 1, 3, -1, -3, 5, 3, 6, 7,
		// 1, -1,
		// 1,
		// 9, 11,
		// 4, -2,
		9, 10, 9, -7, -4, -8, 2, -6,
	}
	// for i := 10000; i >= 1; i-- {
	// 	for j := 0; j < 10; j++ {
	// 		nums = append(nums, i)
	// 	}
	// }
	// k := 50000
	// k := 2
	// k := 2
	// k := 3
	k := 5
	fmt.Println(maxSlidingWindow(nums, k))

}
