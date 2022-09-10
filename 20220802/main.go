package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//sum := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	//fmt.Println(sum)

	//matrix := updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	//matrix := updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}})
	//matrix := updateMatrix([][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 1, 1, 1}})
	//fmt.Println(matrix)

	//str := reverseStr("abcd", 4)
	//str := reverseStr("abcdefg", 2)
	//str := reverseStr("abcdefg", 8)
	//fmt.Println(str)

	//newnew()

	//count := findEwordCount("Nice to meet you")
	//fmt.Println(count)

	//	area := getArea(&Point{
	//		X: 1,
	//		Y: 1,
	//	}, &Point{
	//		X: 2,
	//		Y: 2,
	//	}, &Point{
	//		X: 1,
	//		Y: 3,
	//	}, &Point{
	//		X: 0,
	//		Y: 2,
	//	})
	//	fmt.Println(area)

	//input := bufio.NewScanner(os.Stdin)
	//r, c := 0, 0
	//fmt.Scanf("%d %d\n", &r, &c)
	//s := make([][]string, r, r)
	//for i := 0; i < r; i++ {
	//	s[i] = make([]string, 0, c)
	//	input.Scan()
	//	text := input.Text()
	//	split := strings.Split(text, " ")
	//	for _, v := range split[:c] {
	//		s[i] = append(s[i], v)
	//	}
	//}
	//fmt.Println(s)
	//i := dfs(s)
	//fmt.Println(i)
	//for j := 0; j < len(i); j++ {
	//	fmt.Printf("(%v,%v)\n", i[j][0], i[j][1])
	//}

}
func numIslands1(grid [][]byte) int {
	res := 0
	if len(grid) == 0 {
		return 0
	}

	var travel func(x, y int)
	travel = func(x, y int) {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return
		}

		if grid[x][y] == '0' {
			return
		}

		if grid[x][y] == '1' {
			grid[x][y] = '0'
			travel(x+1, y)
			travel(x-1, y)
			travel(x, y+1)
			travel(x, y-1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				travel(i, j)
			}
		}

	}

	return res
}

func maxAreaOfIsland1(grid [][]int) int {
	var travel func(x, y int) int
	travel = func(x, y int) int {
		if grid[x][y] == 0 {
			return 0
		}
		grid[x][y] = 0
		var a, b, c, d int
		if x-1 >= 0 && grid[x-1][y] == 1 {
			a = travel(x-1, y)
		}
		if x+1 < len(grid) && grid[x+1][y] == 1 {
			b = travel(x+1, y)
		}

		if y-1 >= 0 && grid[x][y-1] == 1 {
			c = travel(x, y-1)
		}

		if y+1 < len(grid[x]) && grid[x][y+1] == 1 {
			d = travel(x, y+1)
		}
		return a + b + c + d + 1
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				temp := travel(i, j)
				if temp > res {
					res = temp
				}
			}
		}
	}

	return res
}

func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				queue := make([][2]int, 0, 0)
				queue = append(queue, [2]int{i, j})
				temp := 0
				for len(queue) != 0 {
					q := queue[0]
					x := q[0]
					y := q[1]
					grid[x][y] = 0
					if x-1 >= 0 && grid[x-1][y] == 1 {
						queue = append(queue, [2]int{x - 1, y})
						grid[x-1][y] = 0
					}
					if x+1 < len(grid) && grid[x+1][y] == 1 {
						queue = append(queue, [2]int{x + 1, y})
						grid[x+1][y] = 0
					}
					if y-1 >= 0 && grid[x][y-1] == 1 {
						queue = append(queue, [2]int{x, y - 1})
						grid[x][y-1] = 0
					}
					if y+1 < len(grid[x]) && grid[x][y+1] == 1 {
						queue = append(queue, [2]int{x, y + 1})
						grid[x][y+1] = 0
					}
					queue = queue[1:]
					temp++
				}
				if temp > res {
					res = temp
				}
			}
		}

	}

	return res
}

func numIslands(grid [][]byte) int {
	res := 0
	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				queue := make([][2]int, 0, 0)
				queue = append(queue, [2]int{i, j})
				for len(queue) != 0 {
					q := queue[0]
					x := q[0]
					y := q[1]
					grid[x][y] = '0'
					if x-1 >= 0 && grid[x-1][y] == '1' {
						queue = append(queue, [2]int{x - 1, y})
						grid[x-1][y] = '0'
					}
					if x+1 < len(grid) && grid[x+1][y] == '1' {
						queue = append(queue, [2]int{x + 1, y})
						grid[x+1][y] = '0'
					}
					if y-1 >= 0 && grid[x][y-1] == '1' {
						queue = append(queue, [2]int{x, y - 1})
						grid[x][y-1] = '0'
					}
					if y+1 < len(grid[x]) && grid[x][y+1] == '1' {
						queue = append(queue, [2]int{x, y + 1})
						grid[x][y+1] = '0'
					}
					queue = queue[1:]
				}
			}
		}
	}

	return res
}

func dfs(s [][]string) [][2]int {
	res := make([][2]int, 0, 0)
	m := make(map[[2]int]bool)
	b := false
	var travel func(x, y int)
	travel = func(x, y int) {
		if m[[2]int{x, y}] {
			return
		}
		if x < 0 || x >= len(s) || y < 0 || y >= len(s[0]) {
			return
		}
		if x == len(s)-1 && y == len(s[0])-1 {
			m[[2]int{x, y}] = true
			res = append(res, [2]int{x, y})
			b = true
			return
		}
		if s[x][y] == "1" {
			return
		}

		res = append(res, [2]int{x, y})
		m[[2]int{x, y}] = true
		travel(x-1, y)
		if b {
			return
		}
		travel(x+1, y)
		if b {
			return
		}
		travel(x, y-1)
		if b {
			return
		}
		travel(x, y+1)
		if b {
			return
		} else {
			res = res[:len(res)-1]
		}
	}

	travel(0, 0)

	return res
}

func findEwordCount(s string) int {
	// write code here
	res := 0
	for i := 0; i < len(s); {
		if s[i] == 'e' {
			res++
			for i < len(s) && s[i] != ' ' {
				i++
			}
		} else {
		}
		i++
	}
	return res
}

type Point struct {
	X int
	Y int
}

func getArea(p1 *Point, p2 *Point, p3 *Point, p4 *Point) int64 {
	// write code here
	p12 := math.Sqrt(float64(math.Abs(float64((p1.X-p2.X)*(p1.X-p2.X)))) + math.Abs(float64((p1.Y-p2.Y)*(p1.Y-p2.Y))))
	p23 := math.Sqrt(float64(math.Abs(float64((p2.X-p3.X)*(p2.X-p3.X)))) + math.Abs(float64((p2.Y-p3.Y)*(p2.Y-p3.Y))))
	p34 := math.Sqrt(float64(math.Abs(float64((p3.X-p4.X)*(p3.X-p4.X)))) + math.Abs(float64((p3.Y-p4.Y)*(p3.Y-p4.Y))))
	p41 := math.Sqrt(float64(math.Abs(float64((p4.X-p1.X)*(p4.X-p1.X)))) + math.Abs(float64((p4.Y-p1.Y)*(p4.Y-p1.Y))))
	p24 := math.Sqrt(float64(math.Abs(float64((p2.X-p4.X)*(p2.X-p4.X)))) + math.Abs(float64((p2.Y-p4.Y)*(p2.Y-p4.Y))))
	k1 := (p12 + p41 + p24) / 2
	k2 := (p23 + p34 + p24) / 2
	temp1 := math.Sqrt(k1*(k1-p12)*(k1-p41)*(k1-p24)) + math.Sqrt(k2*(k2-p23)*(k2-p34)*(k2-p24))
	return int64(math.Round(temp1))
}

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]), len(grid[0]))
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 {
				if j == 0 {
					dp[i][j] = grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
				}
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func updateMatrix1(mat [][]int) [][]int {
	queue := make([][2]int, 0, 0)
	m := make(map[[2]int]bool)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	var push func(x, y int)
	push = func(x, y int) {
		if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[0]) && !m[[2]int{x, y}] {
			queue = append(queue, [2]int{x, y})
		}
	}

	res := make([][]int, len(mat), len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(mat[0]), len(mat[0]))
	}

	index := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			ints := queue[0]
			x := ints[0]
			y := ints[1]

			if _, ok := m[ints]; ok {
				queue = queue[1:]
				continue
			}
			m[ints] = true
			res[x][y] = index
			queue = queue[1:]

			push(x+1, y)
			push(x-1, y)
			push(x, y-1)
			push(x, y+1)
		}
		index++
	}

	return res
}

func updateMatrix(mat [][]int) [][]int {
	dp := make([][]int, len(mat), len(mat))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(mat[0]), len(mat[0]))
		for j := 0; j < len(dp[i]); j++ {
			if mat[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = math.MaxInt >> 2
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if i >= 1 {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j])
			}
			if j >= 1 {
				dp[i][j] = min(dp[i][j-1]+1, dp[i][j])
			}
		}
	}

	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[i]) - 1; j >= 0; j-- {
			if i < len(dp)-1 {
				dp[i][j] = min(dp[i+1][j]+1, dp[i][j])
			}
			if j < len(dp[i])-1 {
				dp[i][j] = min(dp[i][j+1]+1, dp[i][j])
			}
		}
	}

	return dp
}

func reverseStr(s string, k int) string {

	var travel func(s []byte) string
	travel = func(s []byte) string {
		if len(s) == 0 {
			return string(s)
		} else if len(s) < k {
			for i := 0; i < len(s)/2; i++ {
				s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
			}
			return string(s)
		}

		for i := 0; i < k/2; i++ {
			s[i], s[k-i-1] = s[k-i-1], s[i]
		}

		if 2*k >= len(s) {
			return string(s)
		}
		return string(s[:2*k]) + travel(s[2*k:])
	}

	return travel([]byte(s))
}

func newnew() {
	i := new(struct{ a int })
	(*i).a = 1
	fmt.Println(i)
}

func Solution1(N int) int {
	// write your code in Go 1.4
	itoa := strconv.Itoa(N)
	b := false
	if itoa[0] == '-' {
		itoa = itoa[1:]
		b = true
	}
	res := math.MinInt
	for i := 0; i < len(itoa); i++ {
		temp := itoa[:i] + "5" + itoa[i:]
		if b {
			temp = "-" + temp
		}
		atoi, _ := strconv.Atoi(temp)
		if atoi > res {
			res = atoi
		}
	}

	return res
}
