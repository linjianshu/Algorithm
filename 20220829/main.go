package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//n := 0
	//fmt.Scanf("%d\n", &n)
	//nums := make([][2]int, 0)
	//for i := 0; i < n; i++ {
	//	x1, x2 := 0, 0
	//	fmt.Scanf("%d %d\n", &x1, &x2)
	//	if len(nums) == 0 {
	//		nums = append(nums, [2]int{x1, x2})
	//	} else {
	//		b := false
	//		for j := 0; j < len(nums); j++ {
	//			if nums[j][0] <= x1 && nums[j][1]  x2 {
	//				b = true
	//				break
	//			}
	//		}
	//		if !b {
	//			nums = append(nums, [2]int{x1, x2})
	//		}
	//	}
	//}
	//
	//fmt.Println(len(nums))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	s := ""
	s = scanner.Text()
	split := strings.Split(s, " ")
	nums := make([]int, 0, 0)
	for i := 0; i < len(split); i++ {
		atoi, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		nums = append(nums, atoi)
	}
	a := xxx(nums[1:])
	b := xxx(nums[:len(nums)-1])
	fmt.Println(max(a, b))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func xxx(nums []int) int {
	dp := make([]int, len(nums)+1, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(dp)-1]
}
