package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b := true
	nums := make([]int, 0) // 用于记录每一行内的数字
	nums1 := make([]int, 0)
	var n int                         // 输入行数
	sca := bufio.NewScanner(os.Stdin) // 创建一个Scanner结构体实例对象
	fmt.Scanf("%d\n", &n)
	for i := 0; i < 3*n; i++ { //读一行
		if sca.Scan() {
			strSlice := strings.Split(sca.Text(), " ") //转成string，再按空格符切分
			if len(strSlice) == 2 {
				continue
			}
			for _, v := range strSlice {
				num, _ := strconv.Atoi(v) // string转int
				if b {
					nums = append(nums, num)
				} else {
					nums1 = append(nums1, num)
				}
			}
			b = !b

			if len(nums) != 0 && len(nums1) != 0 {
				lingxing(nums, nums1)
				nums = []int{}
				nums1 = []int{}
			}
		}
	}
}

func yunshu(nums []int, q int) int {
	dp := make([][][2]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, q+1, q+1)
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(dp[0]); i++ {

		dp[0][i] = [2]int{0, i}
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if nums[i-1] >= 0 {
				dp[i][j] = [2]int{dp[i-1][j][0] + 1, dp[i-1][j][1] + nums[i-1]}
			} else {
				if dp[i-1][j][1]+nums[i-1] >= 0 {
					a := 1
					b := dp[i-1][j][0] + 1
					if max(a, b) == b {
						dp[i][j] = [2]int{b, dp[i-1][j][1] + nums[i-1]}
					} else {
						dp[i][j] = [2]int{1, j - nums[i-1]}
					}
				} else if j+nums[i-1] >= 0 {
					dp[i][j] = [2]int{1, j - nums[i-1]}
				} else {
					dp[i][j] = [2]int{0, j}
				}
			}
		}
	}

	return dp[len(nums)][q][0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func lingxing(a, b []int) {
	if len(a)%4 != 0 || len(a) > len(b) {
		fmt.Println("No")
		return
	}
	m := make(map[int]*ListNode)
	cur := &ListNode{
		Val:  b[0],
		Next: nil,
	}
	root := &ListNode{
		Val:  a[0],
		Next: cur,
	}
	m[cur.Val] = cur
	m[root.Val] = root
	for i := 1; i < len(a); i++ {
		if v, ok := m[a[i]]; ok {
			newnode := &ListNode{
				Val:  b[i],
				Next: nil,
			}
			if v.Next != nil {
				fmt.Println("No")
				return
			}
			v.Next = newnode
			m[newnode.Val] = newnode
		}
	}

	fmt.Println("Yes")
}

func distance(s, t string) int {
	res := 0
	l1 := len(s)
	l2 := len(t)
	left := 0
	right := 0

	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a < b {
			return b - a
		}
		return a - b
	}
	for right < l2 {

		for right-left == l1-1 {
			i := 0
			for l := left; l <= right; l++ {
				res += abs(int(t[l]), int(s[i]))
				i++
			}
			left++
		}
		right++
	}

	return res
}
