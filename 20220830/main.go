package main

import (
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	nums := make([]int, 0, 0)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		nums = append(nums, num)
	}

	sub := make([]int, 0, 0)
	for i := 1; i < len(nums); i++ {
		temp := math.Abs(float64(nums[i] - nums[i-1]))
		sub = append(sub, int(temp))
	}

	if len(sub) == 1 {
		fmt.Println(0)
		return
	}

	maxI, secMaxI := 0, 0

	for i := 0; i < len(sub); i++ {
		if sub[i] > sub[maxI] {
			maxI = i
		}
	}

	if maxI == secMaxI {
		secMaxI = 1
	}

	mTrue := sub[maxI]
	sub[maxI] = 0
	for i := 0; i < len(sub); i++ {
		if sub[i] > sub[secMaxI] {
			secMaxI = i
		}
	}

	sub[maxI] = mTrue

	var max func(a, b int) int
	max = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	a, b := 0, 0
	if maxI == 0 {
		fmt.Println(sub[secMaxI])
	} else if maxI == len(sub)-1 {
		fmt.Println(sub[secMaxI])
	} else {
		a = nums[maxI-1]
		b = nums[maxI+1]
		left1 := math.Abs(float64(a - (a+b)/2))
		right1 := math.Abs(float64(b - (a+b)/2))
		ans1 := max(int(left1), int(right1))
		fmt.Println(max(ans1, sub[secMaxI]))
	}

	//fmt.Println(maxI, secMaxI)
	//fmt.Println(sub)

	//var min func(a, b int) int
	//min = func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//
	//n := 0
	//fmt.Scanf("%d\n", &n)
	//for i := 0; i < n; i++ {
	//	y, o, u := 0, 0, 0
	//	fmt.Scanf("%d %d %d\n", &y, &o, &u)
	//
	//	m := min(min(y, o), u)
	//	if o-m >= 2 {
	//		fmt.Println(m*2 + o - m - 1)
	//		continue
	//	}
	//	fmt.Println(m * 2)
	//}

	//n := 0
	//fmt.Scanf("%d\n", &n)
	//nums := make([]int, 0, n)
	//for i := 0; i < n; i++ {
	//	num := 0
	//	fmt.Scanf("%d\n", &num)
	//	nums = append(nums, num)
	//}
	//
	//for i := 0; i < len(nums); i++ {
	//	if nums[i]%2 == 0 {
	//		fmt.Println(nums[i])
	//		continue
	//	}
	//	itoa := strconv.Itoa(nums[i])
	//	b := false
	//	for j := 0; j < len(itoa); j++ {
	//		if (itoa[j]-'0')%2 == 0 {
	//			res := itoa[:j] + itoa[j+1:]
	//			if res[0] == '0' {
	//				res = res[1:] + string('0')
	//			}
	//			res += string(itoa[j])
	//
	//			atoi, err := strconv.Atoi(res)
	//			if err != nil {
	//				fmt.Println(err)
	//				return
	//			}
	//			fmt.Println(atoi)
	//			b = true
	//			break
	//		}
	//	}
	//
	//	if !b {
	//		fmt.Println(-1)
	//	}
	//}

}
