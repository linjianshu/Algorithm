package main

import (
	"sort"
)

func main() {
	//boat := rowTheBoat(8, 2, 13)
	//fmt.Println(boat)

	//sort := shell_sort([]int{5, 6, 1, 4}, []int{2, 1})
	//sort := shell_sort([]int{1, 4, 3, 7, 5}, []int{3, 1})
	//fmt.Println(sort)

	//inMap := checkStrInMap([][]string{{"a", "b", "c"}, {"f", "e", "d"}, {"g", "n", "m"}}, []string{"fed", "fam", "bca"})
	//fmt.Println(inMap)
}

func rowTheBoat(x float32, y float32, m float32) int {
	// write code here
	if x > m {
		return 1
	}

	count := 0
	for m > 0 {

		if x < y && x < m {
			return -1
		}

		m -= x
		if m <= 0 {
			return count + 1
		}
		m += y
		x = 0.8 * x
		count++
	}
	return count
}
func rowTheBoat1(x float32, y float32, m float32) int {
	// write code here
	count := 0
	for m > 0 {
		if x >= m {
			return count + 1
		}

		if x <= y && x < m {
			return -1
		}

		m -= x
		m += y
		x = 0.8 * x
		count++
	}
	return count
}

type pair struct {
	x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exitst(board [][]string, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, w)
	}
	//for i := range vis {
	//	vis[i] = make([]bool, w)
	//}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j][0] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}

		vis[i][j] = true
		defer func() { vis[i][j] = false }()

		for _, dir := range directions {
			if newi, newj := i+dir.x, j+dir.y; 0 <= newi && newi < h && 0 <= newj && newj < w && !vis[newi][newj] {
				if check(newi, newj, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
func checkStrInMap(m [][]string, strArr []string) []bool {
	// write code here
	res := make([]bool, 0, 0)
	for i := 0; i < len(strArr); i++ {
		res = append(res, exitst(m, strArr[i]))
	}
	return res
}

func shell_sort(numlist []int, gaplist []int) int {
	// write code here
	sort.Slice(gaplist, func(i, j int) bool {
		if gaplist[i] > gaplist[j] {
			return true
		}

		return false
	})
	res := 0
	last := 1000000
	for i := 0; i < len(gaplist); i++ {
		gap := gaplist[i]
		if gap < 1 {
			break
		}
		if gap >= last {
			continue
		}
		for j := 0; j < len(numlist)-gap; j++ {
			if numlist[j] >= numlist[j+gap] {
				numlist[j], numlist[j+gap] = numlist[j+gap], numlist[j]
				res++
			}
		}

		if gap == 1 {
			break
		}
		last = gap
	}

	return res
}
