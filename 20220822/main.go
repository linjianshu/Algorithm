package main

import (
	"container/heap"
	"fmt"
)

func main() {
	window := maxSlidingWindow([]int{1, -1}, 1)
	fmt.Println(window)
}

type myheap struct {
	Idex []int
	Nums []int
}

func (m *myheap) Len() int {
	//TODO implement me
	//panic("implement me")
	return len(m.Idex)
}

func (m *myheap) Less(i, j int) bool {
	//TODO implement me
	//panic("implement me")
	return m.Nums[m.Idex[i]] > m.Nums[m.Idex[j]]
}

func (m *myheap) Swap(i, j int) {
	//TODO implement me
	//panic("implement me")
	m.Idex[i], m.Idex[j] = m.Idex[j], m.Idex[i]
}

func (m *myheap) Push(x interface{}) {
	//TODO implement me
	//panic("implement me")
	m.Idex = append(m.Idex, x.(int))
}

func (m *myheap) Pop() interface{} {
	//TODO implement me
	//panic("implement me")
	v := m.Nums[m.Idex[len(m.Idex)-1]]
	m.Idex = m.Idex[:len(m.Idex)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, 0)
	h := myheap{}
	h.Idex = make([]int, 0, 0)
	h.Nums = nums
	for i := 0; i < len(nums); i++ {
		heap.Push(&h, i)

		if i < k-1 {
			continue
		}

		for {
			i2 := h.Idex[0]
			if i2 <= i-k {
				heap.Pop(&h)
			} else {
				res = append(res, nums[i2])
				break
			}
		}
	}

	return res
}
