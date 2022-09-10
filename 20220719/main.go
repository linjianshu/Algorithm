package main

import (
	"fmt"
	"math"
)

func main() {
	//b := judge([]int{1, 2, 5, 2, 3})
	b := judge([]int{1, 2, 4, 6, 5})
	fmt.Println(b)
}

func judge(nums []int) bool {
	var travel func(nums []int, min int, max int) bool
	travel = func(nums []int, min int, max int) bool {
		if len(nums) == 1 || len(nums) == 0 {
			return true
		}

		newNums := make([]int, len(nums)-1, len(nums)-1)
		copy(newNums, nums)
		num := nums[len(nums)-1]

		if num < min || num > max {
			return false
		}

		b := true
		index := 0
		for i := 0; i < len(newNums); i++ {
			if newNums[i] > num {
				index = i
				b = false
				break
			}
		}
		if b {
			return true
		}

		return travel(newNums[:index], min, num) && travel(newNums[index:], num, max)
	}
	return travel(nums, math.MinInt32, math.MaxInt32)
}
