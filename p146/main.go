package main

import "fmt"

type LRUCache struct {
	vals []int
	keys []int
	idx  int
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		idx:  -1,
		cap:  capacity,
		vals: make([]int, capacity),
		keys: make([]int, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if this.idx < 0 {
		return -1
	}
	found := this.find(key)
	if found < 0 {
		return -1
	}
	//move
	val := this.vals[found]
	this.removeIndex(found)
	this.vals[this.idx] = val
	this.keys[this.idx] = key

	return val
}

func (this *LRUCache) Put(key int, value int) {
	found := this.find(key)
	if found >= 0 {
		//修改
		this.vals[found] = value

		this.removeIndex(found)
		this.vals[this.idx] = value
		this.keys[this.idx] = key
		return
	}
	//插入
	if this.idx >= this.cap-1 {
		//remove
		this.removeIndex(0)
		this.idx--
	}
	this.idx++
	this.vals[this.idx] = value
	this.keys[this.idx] = key
}

func (this *LRUCache) find(key int) int {
	found := -1
	for i := range this.keys {
		if this.keys[i] == key {
			found = i
		}
	}
	return found
}
func (this *LRUCache) removeIndex(index int) {
	for i := index; i < len(this.keys)-1; i++ {
		this.keys[i] = this.keys[i+1]
		this.vals[i] = this.vals[i+1]
	}
	this.keys[len(this.keys)-1] = 0
	this.vals[len(this.vals)-1] = 0
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	// obj.Put(1, 1)
	// obj.Put(2, 2)
	// fmt.Println(obj.Get(1))
	// obj.Put(3, 3)
	// fmt.Println(obj.Get(2))
	// obj.Put(4, 4)
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(4))

	obj.Put(2, 1)
	obj.Put(1, 1)
	obj.Put(2, 3)
	obj.Put(4, 1)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
}
