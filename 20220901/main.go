package main

import (
	"fmt"
)

func main() {
	contrib := getWakeLockContrib([][]int{{1234, 1235}, {1236, 1238}})
	fmt.Println(contrib)
	//str := findStr("Hello world Hello world", "l**")
	//fmt.Println(str)
	//s := ""
	//fmt.Scanf("%s\n", &s)
	//split := strings.Split(s, "|")
	//m := make(map[string][][2]string)
	//for _, v := range split {
	//	value := strings.Split(v, " ")
	//	for i := 0; i < len(value); i++ {
	//		m[value[0]] = append(m[value[0]], [2]string{value[1], value[2]})
	//	}
	//}
	//
	//n := len(m)
	//var travel func(start string, mm map[string]int, result int) int
	//travel = func(start string, mm map[string]int, result int) int {
	//	if len(mm) == n {
	//		return result
	//	}
	//
	//	mm[start] = 0
	//	for i := 0; i < len(m[start]); i++ {
	//		i2 := m[start][i]
	//		atoi, err := strconv.Atoi(i2[1])
	//		if err != nil {
	//			fmt.Println(err.Error())
	//			return 0
	//		}
	//		if mm[i2[0]] < atoi {
	//			result -= mm[i2[0]]
	//			mm[i2[0]] = atoi
	//			result += atoi
	//		}
	//
	//		travel()
	//	}
	//	for _, v := range m[start] {
	//
	//	}
	//	return -0
	//}
	//
	//res := 1000000
	//for key := range m {
	//	ans := travel(key, map[string]int{}, 0)
	//	if ans < res {
	//		res = ans
	//	}
	//}
	//
	//fmt.Println(res)
}

func getWakeLockContrib(wakeLock [][]int) []int {

	return []int{1, 2}
}

func findStr(str string, mod string) int {
	// write code here
	rev := ""
	ans := -1

	for i := len(mod) - 1; i >= 0; i-- {
		rev += string(mod[i])
	}
	var travel func(string, string) bool
	travel = func(target string, res string) bool {
		if len(target) == 0 && len(res) == 0 {
			ans++
			return true
		}
		b := false
		for i := 0; i < len(target); i++ {
			if target[i] == res[i] || target[i] == '*' || res[i] == '*' {

			} else {
				return false
			}
		}

		if !b {
			return true
		}
		return false
	}

	for i := len(str) - 1; i >= len(mod)-1; i-- {
		b := travel(str[i-len(mod)+1:i+1], rev)
		if b {
			ans = i
		}
	}

	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode) *ListNode {
	// write code here
	var preThree *ListNode
	var afterSix *ListNode
	var Three *ListNode
	var Six *ListNode
	var pre *ListNode
	cur := head
	for cur != nil {
		if cur.Val == 3 {
			Three = cur
			preThree = pre
		}

		if cur.Val == 6 {
			Six = cur
			afterSix = cur.Next
		}
		pre = cur
		cur = cur.Next
	}

	var prepre *ListNode
	var current *ListNode
	current = Three
	for current != afterSix {
		temp := current.Next
		current.Next = prepre
		prepre = current
		current = temp

		//current = current.Next
	}

	Three.Next = afterSix
	preThree.Next = Six

	return head
}

func ncov_defect(grid [][]int) int {
	// write code here
	res := 0

	var travel func(i, j int)
	travel = func(i, j int) {

		if i-1 >= 0 && grid[i-1][j] == 0 {
			res++
		}

		if i+1 < len(grid) && grid[i+1][j] == 0 {
			res++
		}

		if j-1 >= 0 && grid[i][j-1] == 0 {
			res++
		}

		if j+1 < len(grid[i]) && grid[i][j+1] == 0 {
			res++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				travel(i, j)
			}
		}
	}

	return res
}

func nucleicAcidTestWay(n int) int {
	// write code here
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	dp := make([]int, n+1, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
