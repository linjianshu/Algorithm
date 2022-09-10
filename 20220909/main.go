package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)
	nums := make([][]int, 0, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, 0, 0)
	}
	for i := 0; i < n; i++ {
		ints := make([]int, 2, 2)
		m1, m2 := 0, 0
		fmt.Scanf("%d %d\n", &m1, &m2)
		ints[0] = m1
		ints[1] = m2
		nums = append(nums, ints)
	}

	res := 100000
	var travel func(i int, color int, ans int)
	travel = func(i int, color int, ans int) {
		if i == m {
			if res > ans {
				res = ans
			}
			return
		}

		if color == nums[i][0] {
			travel(i+1, color, ans)
		} else {
			travel(i+1, nums[i][0], ans+color*nums[i][0])
		}

		if color == nums[i][1] {
			travel(i+1, color, ans)
		} else {
			travel(i+1, nums[i][1], ans+color*nums[i][1])

		}
	}
	travel(0, 1, 0)
	fmt.Println(res)
}
