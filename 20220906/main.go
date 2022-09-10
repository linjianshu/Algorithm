package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	m, n := 0, 0
	fmt.Scanf("%d,%d\n", &m, &n)
	nums := make([][]int, m, m)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, n, n)
	}
	for k := 0; k < m; k++ {
		s := ""
		fmt.Scanf("%s\n", &s)
		split := strings.Split(s, ",")
		for i := 0; i < len(split); i++ {
			atoi, _ := strconv.Atoi(split[i])
			nums[k][i] = atoi
		}
	}

	//for i := 0; i < m; i++ {
	//	for j := 0; j < n; j++ {
	//		num := 0
	//		fmt.Scanf("%d", &num)
	//		nums[i][j] = num
	//	}
	//	fmt.Scanln()
	//}

	res := 0
	mm := make(map[[2]int]bool)

	var travel func(i, j int) int
	travel = func(i, j int) int {
		if mm[[2]int{i, j}] {
			return 0
		}
		if !(i >= 0 && j >= 0 && i < m && j < n) {
			return 0
		}
		mm[[2]int{i, j}] = true
		if nums[i][j] == 0 {
			return 0
		}
		return travel(i+1, j) + travel(i-1, j) + travel(i, j-1) + travel(i, j+1) + 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == 1 && !mm[[2]int{i, j}] {
				temp := travel(i, j)
				if res < temp {
					res = temp
				}
			}
		}
	}
	fmt.Println(res)
}
