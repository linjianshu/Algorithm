package main

import "fmt"

func main() {
	//window := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//fmt.Println(window)

	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	param_1 := obj.Get(1)
	fmt.Println(param_1)
	obj.Put(3, 3)
	get := obj.Get(2)
	fmt.Println(get)
	obj.Put(4, 4)
	i := obj.Get(1)
	fmt.Println(i)
	i2 := obj.Get(3)
	fmt.Println(i2)
	i3 := obj.Get(4)
	fmt.Println(i3)

	/**
	 * Your LRUCache object will be instantiated and called as such:
	 */
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, 0)
	queue := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		for len(queue) != 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i < k-1 {

		} else {
			for i-k+1 > queue[0] {
				queue = queue[1:]
			}
			ans := queue[0]
			res = append(res, nums[ans])
		}
	}

	return res
}

type LRUCache struct {
	L        int
	Capacity int
	Head     *LinkNode
	Tail     *LinkNode
	M        map[int]*LinkNode
}

type LinkNode struct {
	Key   int
	Value int
	Pre   *LinkNode
	Next  *LinkNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		L:        0,
		Capacity: capacity,
		Head: &LinkNode{
			Key:   0,
			Value: 0,
			Pre:   nil,
			Next:  nil,
		},
		Tail: &LinkNode{
			Key:   0,
			Value: 0,
			Pre:   nil,
			Next:  nil,
		},
		M: map[int]*LinkNode{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Pre = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	res := this.M[key]
	if res == nil {
		return -1
	}
	temp := res.Value
	//如果存在 就先删掉再添加
	delete(this.M, key)
	res.Next.Pre = res.Pre
	res.Pre.Next = res.Next
	this.L--
	this.Put(key, temp)
	return temp
}

func (this *LRUCache) Put(key int, value int) {
	res := this.M[key]
	//如果存在 就先删掉再添加
	if res != nil {
		delete(this.M, key)
		this.L--
		this.Put(key, value)
	}

	newLinkNode := &LinkNode{
		Key:   key,
		Value: value,
	}

	this.M[key] = newLinkNode
	if this.L == this.Capacity {
		//如果溢出 就删除尾巴
		delete(this.M, this.Tail.Pre.Key)
		this.L--
		this.Tail.Pre.Pre.Next = this.Tail
		this.Tail.Pre = this.Tail.Pre.Pre
	}
	this.Head.Next.Pre = newLinkNode
	newLinkNode.Next = this.Head.Next
	this.Head.Next = newLinkNode
	newLinkNode.Pre = this.Head
	this.L++
}
