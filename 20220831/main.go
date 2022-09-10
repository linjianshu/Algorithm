package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)

	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}
	fmt.Scanln()

	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		result = append(result, num)
	}

	mmm := make(map[int]bool)
	res := 0
	count := 0
	j := len(result) - 1
	for i := len(result) - 1; i >= 0; i-- {
		if nums[i] != result[j] {
			mmm[nums[i]] = true
			count++
			res++
		} else {
			j--
			for j >= 0 && mmm[result[j]] {
				j--
			}
		}
	}

	fmt.Println(res)
	//D := 0
	//fmt.Scanf("%d\n", &D)
	//n := 0
	//fmt.Scanf("%d\n", &n)
	//nums := make([][2]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	x, y := 0, 0
	//	fmt.Scanf("%d %d\n", &x, &y)
	//	nums = append(nums, [2]int{x, y})
	//}
	//
	//res := 0
	//shen := 1000
	//pre := 0
	//for i := 0; i < len(nums); i++ {
	//
	//}
	//
	//fmt.Println(res)

	//n, m := 0, 0
	//fmt.Scanf("%d %d\n", &n, &m)
	//nums := make([][]int, n, n)
	//
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = make([]int, m, m)
	//}
	//
	//zhadan := false
	//xianjin := false
	//begin := [2]int{}
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		num := 0
	//		fmt.Scanf("%d", &num)
	//		nums[i][j] = num
	//		if num == 2 {
	//			begin = [2]int{i, j}
	//		}
	//
	//	}
	//	fmt.Scanln()
	//}
	//mm := make(map[[2]int]bool)
	//res := 1000000000
	//var travel func(i, j int, cur int)
	//travel = func(i, j int, cur int) {
	//	if i < 0 || i >= n || j < 0 || j >= m {
	//		return
	//	}
	//	if mm[[2]int{i, j}] == true {
	//		return
	//	}
	//	if nums[i][j] == 3 {
	//		if cur < res {
	//			res = cur
	//		}
	//		return
	//	}
	//
	//	mm[[2]int{i, j}] = true
	//	if nums[i][j] == 4 && !xianjin {
	//		xianjin = true
	//		if i-1 >= 0 {
	//			travel(i-1, j, cur+3)
	//		}
	//		if i+1 < n {
	//			travel(i+1, j, cur+3)
	//		}
	//
	//		if j-1 >= 0 {
	//			travel(i, j-1, cur+3)
	//		}
	//
	//		if j+1 < m {
	//			travel(i, j+1, cur+3)
	//		}
	//		xianjin = false
	//	} else if nums[i][j] == 6 && !zhadan {
	//		zhadan = true
	//		if i-1 >= 0 && nums[i-1][j] == 1 {
	//			nums[i-1][j] = 0
	//			travel(i-1, j, cur+1)
	//			nums[i-1][j] = 1
	//		} else {
	//			if i-1 >= 0 {
	//				travel(i-1, j, cur+1)
	//			}
	//		}
	//
	//		if i+1 < n && nums[i+1][j] == 1 {
	//			nums[i+1][j] = 0
	//			travel(i+1, j, cur+1)
	//			nums[i+1][j] = 1
	//		} else {
	//			if i+1 < n {
	//				travel(i+1, j, cur+1)
	//			}
	//		}
	//
	//		if j-1 >= 0 && nums[i][j-1] == 1 {
	//			nums[i][j-1] = 0
	//			travel(i, j-1, cur+1)
	//			nums[i][j-1] = 1
	//		} else {
	//			if j-1 >= 0 {
	//				travel(i, j-1, cur+1)
	//			}
	//		}
	//
	//		if j+1 < m && nums[i][j+1] == 1 {
	//			nums[i][j+1] = 0
	//			travel(i, j+1, cur+1)
	//			nums[i][j+1] = 1
	//		} else {
	//			if j+1 < m {
	//				travel(i, j+1, cur+1)
	//			}
	//		}
	//		zhadan = false
	//	} else if nums[i][j] == 1 {
	//		mm[[2]int{i, j}] = false
	//		return
	//	} else {
	//		if i-1 >= 0 {
	//			travel(i-1, j, cur+1)
	//		}
	//
	//		if i+1 < n {
	//			travel(i+1, j, cur+1)
	//		}
	//
	//		if j-1 >= 0 {
	//			travel(i, j-1, cur+1)
	//		}
	//
	//		if j+1 < m {
	//			travel(i, j+1, cur+1)
	//		}
	//	}
	//
	//	mm[[2]int{i, j}] = false
	//}
	//travel(begin[0], begin[1], 0)
	//fmt.Println(res)

	//fmt.Println(nums)
	//s := ""
	//m := make(map[string]int)
	//dic := ""
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//s = scanner.Text()
	//scanner.Scan()
	//dic = scanner.Text()
	//
	//split := strings.Split(dic, " ")
	//for i := 0; i < len(split); i++ {
	//	m[strings.ToUpper(split[i])] = i
	//}
	//
	//ans := make([]string, 0, 0)
	//sss := strings.Split(s, " ")
	//b := false
	//for i := 0; i < len(sss); i++ {
	//	if sss[i] == `"` && !b {
	//		b = true
	//	} else if sss[i] == `"` && b {
	//		b = false
	//	}
	//	temp := strings.ToUpper(sss[i])
	//	if v, ok := m[temp]; ok && !b {
	//		itoa := strconv.Itoa(v)
	//		ans = append(ans, itoa)
	//	} else {
	//		ans = append(ans, sss[i])
	//	}
	//}
	//
	//join := strings.Join(ans, " ")
	//fmt.Println(join)
}
