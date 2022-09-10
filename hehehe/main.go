package main

import (
	"fmt"
)

func main() {
	//islands := numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}})
	//islands := numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}})
	//fmt.Println(islands)

	//order := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	//order := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})
	//order := spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})
	//fmt.Println(order)

	//ints := []int{1, 2, 3}
	//ints := []int{5, 4, 7, 5, 3, 2}
	//ints := []int{1, 2, 0, 3, 0, 1, 2, 4}
	//ints := []int{4, 5, 2, 6, 3, 1}
	//ints := []int{3, 2, 1}
	//ints := []int{1, 1, 5}
	//ints := []int{2, 3, 1}
	//ints := []int{2, 3, 1, 0}
	//ints := []int{1, 3, 2}
	ints := []int{5, 1, 1}
	nextPermutation(ints)

	fmt.Println(ints)
}

//问面试官能不能 修改元素组 如果不能就用map 如果能就避免走回头路 直接修改元数组
func numIslands1(grid [][]byte) int {
	m := make(map[[2]int]bool)
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				m[[2]int{i, j}] = true
			}
		}
	}

	for ints := range m {
		queue := make([][2]int, 0, 0)
		queue = append(queue, ints)
		delete(m, ints)
		for len(queue) != 0 {
			l := len(queue)
			for i := 0; i < l; i++ {
				if m[[2]int{queue[0][0] - 1, queue[0][1]}] && queue[0][0]-1 >= 0 && grid[queue[0][0]-1][queue[0][1]] == '1' {
					queue = append(queue, [2]int{queue[0][0] - 1, queue[0][1]})
					delete(m, [2]int{queue[0][0] - 1, queue[0][1]})
				}
				if m[[2]int{queue[0][0] + 1, queue[0][1]}] && queue[0][0]+1 < len(grid) && grid[queue[0][0]+1][queue[0][1]] == '1' {
					queue = append(queue, [2]int{queue[0][0] + 1, queue[0][1]})
					delete(m, [2]int{queue[0][0] + 1, queue[0][1]})
				}
				if m[[2]int{queue[0][0], queue[0][1] - 1}] && queue[0][1]-1 >= 0 && grid[queue[0][0]][queue[0][1]-1] == '1' {
					queue = append(queue, [2]int{queue[0][0], queue[0][1] - 1})
					delete(m, [2]int{queue[0][0], queue[0][1] - 1})
				}
				if m[[2]int{queue[0][0], queue[0][1] + 1}] && queue[0][1]+1 < len(grid[0]) && grid[queue[0][0]][queue[0][1]+1] == '1' {
					queue = append(queue, [2]int{queue[0][0], queue[0][1] + 1})
					delete(m, [2]int{queue[0][0], queue[0][1] + 1})
				}
				delete(m, queue[0])
				queue = queue[1:]
			}

		}
		res++
	}

	return res
}

//优化 直接修改元数组 不用map  直接扫描 不用预处理
func numIslands2(grid [][]byte) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				queue := make([][2]int, 0, 0)
				queue = append(queue, [2]int{i, j})

				//修改元素组  避免回头路
				grid[i][j] = '0'

				for len(queue) != 0 {
					l := len(queue)
					for i := 0; i < l; i++ {
						if queue[0][0]-1 >= 0 && grid[queue[0][0]-1][queue[0][1]] == '1' {
							queue = append(queue, [2]int{queue[0][0] - 1, queue[0][1]})
							grid[queue[0][0]-1][queue[0][1]] = '0'
						}
						if queue[0][0]+1 < len(grid) && grid[queue[0][0]+1][queue[0][1]] == '1' {
							queue = append(queue, [2]int{queue[0][0] + 1, queue[0][1]})
							grid[queue[0][0]+1][queue[0][1]] = '0'

						}
						if queue[0][1]-1 >= 0 && grid[queue[0][0]][queue[0][1]-1] == '1' {
							queue = append(queue, [2]int{queue[0][0], queue[0][1] - 1})
							grid[queue[0][0]][queue[0][1]-1] = '0'

						}
						if queue[0][1]+1 < len(grid[0]) && grid[queue[0][0]][queue[0][1]+1] == '1' {
							queue = append(queue, [2]int{queue[0][0], queue[0][1] + 1})
							grid[queue[0][0]][queue[0][1]+1] = '0'
						}
						queue = queue[1:]
					}

				}
				res++
			}
		}
	}

	return res
}

//问面试官能不能 修改元素组 如果不能就用map 如果能就避免走回头路 直接修改元数组
func numIslands3(grid [][]byte) int {
	m := make(map[[2]int]bool)
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				m[[2]int{i, j}] = true
			}
		}
	}

	var travel func(x, y int)
	travel = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
			return
		}

		if m[[2]int{x, y}] && grid[x][y] == '1' {
			delete(m, [2]int{x, y})
			travel(x-1, y)
			travel(x+1, y)
			travel(x, y-1)
			travel(x, y+1)
		}
	}

	for ints := range m {
		travel(ints[0], ints[1])
		res++
	}
	return res
}

//问面试官能不能 修改元素组 如果不能就用map 如果能就避免走回头路 直接修改元数组
func numIslands(grid [][]byte) int {
	res := 0

	var travel func(x, y int)
	travel = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) {
			return
		}

		if grid[x][y] == '1' {
			grid[x][y] = '0'
			travel(x-1, y)
			travel(x+1, y)
			travel(x, y-1)
			travel(x, y+1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				travel(i, j)
				res++
			}
		}
	}

	return res
}

func reverse() {

}

//模拟法 空间复杂度稍高
func spiralOrder1(matrix [][]int) []int {
	m := make(map[[2]int]bool)
	res := make([]int, 0, 0)

	direction := 1
	row := 0
	col := 0
	res = append(res, matrix[0][0])
	m[[2]int{0, 0}] = true
	for len(res) != len(matrix)*len(matrix[0]) {
		if direction == 1 {
			col++
			if !m[[2]int{row, col}] && row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) {
				res = append(res, matrix[row][col])
				m[[2]int{row, col}] = true

			} else {
				col--
				direction = 2
			}
		} else if direction == 2 {
			row++
			if !m[[2]int{row, col}] && row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) {
				res = append(res, matrix[row][col])
				m[[2]int{row, col}] = true

			} else {
				row--
				direction = 3
			}
		} else if direction == 3 {
			col--
			if !m[[2]int{row, col}] && row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) {
				res = append(res, matrix[row][col])
				m[[2]int{row, col}] = true

			} else {
				col++
				direction = 4
			}
		} else if direction == 4 {
			row--
			if !m[[2]int{row, col}] && row >= 0 && row < len(matrix) && col >= 0 && col < len(matrix[0]) {
				res = append(res, matrix[row][col])
				m[[2]int{row, col}] = true

			} else {
				row++
				direction = 1
			}
		}
	}
	return res
}

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0, 0)
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		for i := right - 1; i >= left && bottom != top; i-- {
			res = append(res, matrix[bottom][i])
		}
		for i := bottom - 1; i >= top+1 && left != right; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

//下一个排列
func nextPermutation(nums []int) {
	var smaller int
	var larger int

	var travel func(nums []int)
	travel = func(nums []int) {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] >= nums[i+1] {
			continue
		} else {
			smaller = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > nums[smaller] {
			larger = i
			break
		}
	}

	if smaller == larger {
		travel(nums)
		return
	}

	nums[smaller], nums[larger] = nums[larger], nums[smaller]

	travel(nums[smaller+1:])
}
