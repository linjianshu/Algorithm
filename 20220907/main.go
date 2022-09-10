package main

import (
	"container/heap"
	"fmt"
	_ "src/JianZhiOffer/model"
)

func main() {
	window := maxSlidingWindow([]int{1, -1}, 1)
	fmt.Println(window)
}

type myheap struct {
	nums []int
	idx  []int
}

func (m *myheap) Len() int {
	return len((*m).idx)
}

func (m *myheap) Swap(i, j int) {
	(*m).idx[i], (*m).idx[j] = (*m).idx[j], (*m).idx[i]
}

func (m *myheap) Less(i, j int) bool {
	return m.nums[(*m).idx[i]] > m.nums[(*m).idx[j]]
}

func (m *myheap) Push(x interface{}) {
	(*m).idx = append((*m).idx, x.(int))
}

func (m *myheap) Pop() interface{} {
	temp := (*m).idx[m.Len()-1]
	(*m).idx = (*m).idx[:m.Len()-1]
	return temp
}

func maxSlidingWindow(nums []int, k int) []int {
	mh := myheap{nums: nums}
	res := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		heap.Push(&mh, i)
		if i < k-1 {

		} else {
			for mh.idx[0] <= i-k {
				heap.Pop(&mh)
			}
			res = append(res, nums[mh.idx[0]])
		}
	}
	return res
}
