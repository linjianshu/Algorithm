package main

import (
	"fmt"
)

func main() {
	//index := pivotIndex([]int{-1, -1, -1, -1, -1, 0})
	//fmt.Println(index)

	//insert := searchInsert([]int{1, 3, 5, 6}, 7)
	//fmt.Println(insert)

	//rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})

	//setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})

	palindrome := longestPalindrome("cbbd")
	fmt.Println(palindrome)
}

//前缀和 + 逐项遍历
func pivotIndex1(nums []int) int {
	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = nums[i] + preSum[i]
	}

	for i := 0; i < len(preSum)-1; i++ {
		var left int
		if i > 0 {
			left = preSum[i]
		} else {
			left = 0
		}
		var right int
		if i == len(preSum)-1 {
			right = 0
		} else {
			right = preSum[len(preSum)-1] - preSum[i+1]
		}

		if left == right {
			return i
		}
	}

	return -1
}

//前缀和 + 遍历 简洁一点
func pivotIndex2(nums []int) int {
	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = nums[i] + preSum[i]
	}

	for i := 1; i < len(preSum); i++ {
		left := preSum[i-1]
		right := preSum[len(preSum)-1] - preSum[i]

		if left == right {
			return i - 1
		}
	}

	return -1
}

func searchInsert(nums []int, target int) int {
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
	return left
}

//属实不会 二维数组翻转的话 根据对角线翻转 然后再横向翻转  多试一试
func rotate1(matrix [][]int) {
	//正对角线翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//水平横向翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-1-j] = matrix[i][len(matrix[i])-1-j], matrix[i][j]
		}
	}

	return
}

//使用数组作为桥梁 存一下
func rotate(matrix [][]int) {
	res := make([]int, 0, len(matrix)*len(matrix))
	//代表列
	for i := 0; i < len(matrix[0]); i++ {
		//代表行
		for j := len(matrix) - 1; j >= 0; j-- {
			res = append(res, matrix[j][i])
		}
	}

	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = res[index]
			index++
		}
	}
	return
}

func setZeroes1(matrix [][]int) {
	m := make(map[[2]int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				m[[2]int{i, j}] = true
			}
		}
	}

	for ints := range m {
		for i := 0; i < len(matrix[ints[0]]); i++ {
			matrix[ints[0]][i] = 0
		}
		for i := 0; i < len(matrix); i++ {
			matrix[i][ints[1]] = 0
		}
	}
}

func setZeroes(matrix [][]int) {
	mRow := make(map[int]bool)
	mCol := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				mRow[i] = true
				mCol[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if mRow[i] || mCol[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func findDiagonalOrder(mat [][]int) []int {
	return []int{}
}

//中心拓展 好丑好丑
func longestPalindrome1(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		temp := s[i : i+1]
		for j := 1; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] != s[i+j] {
				break
			}
			//待优化的点 符合的时候一遍一遍赋值 很消耗时间
			temp = s[i-j : i+j+1]
		}

		if i+1 < len(s) && s[i] == s[i+1] {
			temp1 := s[i : i+2]
			for j := 1; i-j >= 0 && i+j+1 < len(s); j++ {
				if s[i-j] != s[i+j+1] {
					break
				}
				temp1 = s[i-j : i+j+1+1]
			}

			if len(temp) < len(temp1) {
				temp = temp1
			}
		}

		if len(temp) > len(res) {
			res = temp
		}
	}
	return res
}

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		temp := s[i : i+1]
		left, right := i-1, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		temp = s[left+1 : right]

		if i+1 < len(s) && s[i] == s[i+1] {
			temp1 := s[i : i+2]
			left, right = i-1, i+2
			for left >= 0 && right < len(s) && s[left] == s[right] {
				left--
				right++
			}
			temp1 = s[left+1 : right]

			if len(temp1) > len(temp) {
				temp = temp1
			}
		}

		if len(temp) > len(res) {
			res = temp
		}
	}
	return res
}
