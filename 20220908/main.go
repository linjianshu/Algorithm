package main

import (
	"fmt"
)

func main() {
	m, n := 0, 0
	fmt.Scanf("%d %d\n", &m, &n)
	nums := make([][]int, m, m)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, n, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := 0
			fmt.Scanf("%c", &num)
			nums[i][j] = num
		}
		fmt.Scanln()
	}

	mm := make(map[[2]int]bool)
	dir := []int{1, 2, 3, 4}
	res := 1000000000
	var travel func(i, j int, direction int, steps int)
	travel = func(i, j int, direction int, steps int) {
		if mm[[2]int{i, j}] {
			return
		}
		if !(i >= 0 && j >= 0 && i < m && j < n) {
			return
		}

		if nums[i][j] == 'X' {
			return
		}

		if nums[i][j] == 'E' {
			if res > steps {
				res = steps
				return
			}
		}

		//nums[i][j] = 'X'
		mm[[2]int{i, j}] = true

		if direction == 0 {
			travel(i+1, j, dir[0], steps+1)
			travel(i-1, j, dir[1], steps+1)
			travel(i, j-1, dir[2], steps+1)
			travel(i, j+1, dir[3], steps+1)
		} else {
			if direction == dir[0] {
				travel(i+1, j, dir[0], steps+1)
			} else {
				travel(i+1, j, dir[0], steps+2)
			}

			if direction == dir[1] {
				travel(i-1, j, dir[1], steps+1)
			} else {
				travel(i-1, j, dir[1], steps+2)
			}

			if direction == dir[2] {
				travel(i, j-1, dir[2], steps+1)
			} else {
				travel(i, j-1, dir[2], steps+2)
			}

			if direction == dir[3] {
				travel(i, j+1, dir[3], steps+1)
			} else {
				travel(i, j+1, dir[3], steps+2)
			}
		}
	}

	travel(0, 0, 0, 0)
	if res == 1000000000 {
		res = -1
	}
	fmt.Println(res)
}

//
//func maxScore(energy int, actions [][]int) int {
//	// write code here
//	dp := make([][]int, energy+1, energy+1)
//	for i := 0; i < len(dp); i++ {
//		dp[i] = make([]int, len(actions)+1, len(actions)+1)
//	}
//
//	for i := 0; i < len(dp[0]); i++ {
//		dp[0][i] = 0
//	}
//
//	for i := 0; i < len(dp); i++ {
//		dp[i][0] = 0
//	}
//
//	var max func(a, b int) int
//	max = func(a, b int) int {
//		if a > b {
//			return a
//		}
//
//		return b
//	}
//	for i := 1; i < len(dp); i++ {
//		for j := 1; j < len(dp[i]); j++ {
//			if i < actions[j-1][0] {
//				dp[i][j] = dp[i][j-1]
//			} else {
//				dp[i][j] = max(dp[i][j-1], dp[i-actions[j-1][0]][j-1]+actions[j-1][1])
//			}
//		}
//	}
//
//	return dp[len(dp)-1][len(dp[0])-1]
//}
//
//func Decrypt(encryptedNumber int, decryption int, number int) int {
//	// write code here
//	var travel func(x, y int, m int) int
//	travel = func(x, y int, m int) int {
//		if y == 0 {
//			return 1 % m
//		}
//
//		if y == 1 {
//			return x % m
//		}
//
//		a := travel(x, y-1, m)
//		b := travel(x, 1, m)
//		return (a * b) % m
//	}
//
//	return travel(encryptedNumber, decryption, number)
//}
//
//func showDown(inHand string) string {
//	split := strings.Split(inHand, " ")
//	value := [13]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
//	//asldkfjaslkdfj A  K  Q  J  10 9  8  7  6  5  4  3  2
//	colors := [4]int{0, 0, 0, 0}
//	for i := 0; i < len(split); i++ {
//		switch split[i][1:] {
//		case "A":
//			value[0]++
//		case "K":
//			value[1]++
//		case "Q":
//			value[2]++
//		case "J":
//			value[3]++
//		case "10":
//			value[4]++
//		case "9":
//			value[5]++
//		case "8":
//			value[6]++
//		case "7":
//			value[7]++
//		case "6":
//			value[8]++
//		case "5":
//			value[9]++
//		case "4":
//			value[10]++
//		case "3":
//			value[11]++
//		case "2":
//			value[12]++
//		}
//
//		if split[i][0] == 'S' {
//			colors[0]++
//		} else if split[i][0] == 'H' {
//			colors[1]++
//		} else if split[i][0] == 'C' {
//			colors[2]++
//		} else if split[i][0] == 'D' {
//			colors[3]++
//		}
//	}
//	res := 4
//
//	max := 0
//	for i := 0; i < len(colors); i++ {
//		if max < colors[i] {
//			max = colors[i]
//		}
//	}
//
//	if max == 5 {
//		return "HuangJiaTongHuaShun"
//	}
//
//	same := 0
//	for i := 0; i < len(value); i++ {
//		if same < value[i] {
//			same = value[i]
//		}
//	}
//
//	distinct := make([]int, 0, 13)
//	for i := 0; i < len(value); i++ {
//		distinct = append(distinct, value[i])
//	}
//	sort.Ints(distinct)
//
//	if same == 4 {
//		return "SiTiao"
//	}
//
//	if distinct[len(distinct)-1] >= 3 && distinct[len(distinct)-2] >= 2 {
//		return "HuLu"
//	} else if distinct[len(distinct)-1] == 2 && distinct[len(distinct)-2] == 2 {
//		return "LiangDui"
//	}
//	if same == 3 {
//		return "SanTiao"
//	}
//	if same == 2 {
//		return "YiDui"
//	}
//
//	ans := ""
//	switch res {
//	case 1:
//		ans = "HuangJiaTongHuaShun"
//	case 2:
//		ans = "TongHuaShun"
//	case 3:
//		ans = "SiTiao"
//	case 4:
//		ans = "HuLu"
//	case 5:
//		ans = "TongHua"
//	case 6:
//		ans = "ShunZi"
//	case 7:
//		ans = "SanTiao"
//	case 8:
//		ans = "LiangDui"
//	case 9:
//		ans = "YiDui"
//	case 10:
//		ans = "GaoPai"
//	}
//	return ans
//}
