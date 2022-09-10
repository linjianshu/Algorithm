package main

import "fmt"

func main() {
	//insert := searchInsert([]int{1, 3, 5, 6}, 2)
	//fmt.Println(insert)

	fill := floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2)
	fmt.Println(fill)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	left, right := 0, n
	for left <= right {
		mid := left + (right-left)/2
		if !isBadVersion(mid) {
			left = mid + 1
		} else if isBadVersion(mid) {
			right = mid - 1
		}
	}
	return left
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}

	if left == len(nums) {
		return left
	}
	if left < 0 {
		return 0
	}
	return left
}

//广度优先搜索 很重要的一点在于 不走回头路
func floodFill1(image [][]int, sr int, sc int, newColor int) [][]int {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{sr, sc})
	old := image[sr][sc]

	//如果新旧相同 不走回头路会失效 因此要额外特殊判断
	if old == newColor {
		return image
	}

	for len(queue) != 0 {
		x := queue[0][0]
		y := queue[0][1]
		if !(x >= 0 && x < len(image) && y >= 0 && y < len(image[x])) {
			queue = queue[1:]
			continue
		}
		if image[x][y] != old {
			queue = queue[1:]
			continue
		}

		//可以优化的思路是 不直接加入到队列中 而是先判断在不在范围内再加入到队列中
		image[x][y] = newColor
		queue = append(queue, [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}...)
		queue = queue[1:]
	}
	return image
}

//深度优先搜索
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	old := image[sr][sc]
	offSet := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	//如果新旧相同 不走回头路会失效 因此要额外特殊判断
	if old == newColor {
		return image
	}

	var travel func(point [2]int)
	travel = func(point [2]int) {
		x := point[0]
		y := point[1]

		if !(x >= 0 && x < len(image) && y >= 0 && y < len(image[x])) {
			return
		}

		if image[x][y] != old {
			return
		}
		image[x][y] = newColor
		for i := 0; i < 4; i++ {
			travel([2]int{x + offSet[i][0], y + offSet[i][1]})
		}
	}
	travel([2]int{sr, sc})

	return image
}

//深度优先搜索
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	var travel func(x, y int, temp int) int
	travel = func(x, y int, temp int) int {
		if !(x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x])) {
			return temp
		}
		if grid[x][y] == 0 {
			return temp
		}
		grid[x][y] = 0
		temp++
		temp = travel(x+1, y, temp)
		temp = travel(x-1, y, temp)
		temp = travel(x, y+1, temp)
		temp = travel(x, y-1, temp)
		return temp
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			temp := travel(i, j, 0)
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
