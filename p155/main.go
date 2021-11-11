package main

import "fmt"

type MinStack struct {
	mem []int
	ass []int
	top int
	cap int
}

func Constructor() MinStack {
	return MinStack{
		cap: 10,
		mem: make([]int, 10),
		ass: make([]int, 10),
		top: -1,
	}
}

func (this *MinStack) Push(val int) {
	min := val
	if this.top >= 0 {
		if this.ass[this.top] < min {
			min = this.ass[this.top]
		}
	}

	this.top++
	if this.top >= this.cap {
		this.mem = append(this.mem, val)
		this.ass = append(this.ass, min)
		this.cap++
	} else {
		this.mem[this.top] = val
		this.ass[this.top] = min
	}
}

func (this *MinStack) Pop() {
	this.top--
}

func (this *MinStack) Top() int {
	return this.mem[this.top]
}

func (this *MinStack) GetMin() int {
	if this.top < 0 {
		return 0
	}
	return this.ass[this.top]
}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}
