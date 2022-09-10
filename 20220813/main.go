package main

import (
	"fmt"
	"strconv"
)

func main() {

	nums := 0
	fmt.Scanf("%d\n", &nums)
	ints := make([]int, 0, 0)
	for i := 0; i < nums; i++ {
		s := ""
		fmt.Scanf("%s", &s)
		atoi, _ := strconv.Atoi(s)
		ints = append(ints, atoi)
	}

	temp := make([]int, len(ints), cap(ints))
	copy(temp, ints)

	mm := make(map[int]bool)
	i := 0
	var travel func()
	travel = func() {
		if len(mm) == len(temp)-1 {
			return
		}
		for mm[ints[0]] == true {
			tt := ints[0]
			ints = append(ints[1:], tt)
		}

		two := ints[:2]
		ints = append(ints[2:], two...)
		for mm[ints[0]] == true {
			tt := ints[0]
			ints = append(ints[1:], tt)
		}
		m := make(map[int]int)

		for j := 0; j < len(ints); j++ {
			m[ints[j]] = j
		}

		for mm[ints[0]] == true {
			tt := ints[0]
			ints = append(ints[1:], tt)
		}

		if ints[0] != temp[i] {
			ints[0], ints[m[temp[i]]] = ints[m[temp[i]]], ints[0]
			i++
			zero := ints[0]
			mm[zero] = true
			ints = append(ints[1:], zero)
			travel()
		} else {
			i++
			zero := ints[0]
			mm[zero] = true
			ints = append(ints[1:], zero)
			travel()
		}
	}
	travel()
	fmt.Println(ints)

	//var n, m, k int
	//fmt.Scanf("%d %d %d\n", &n, &m, &k)
	//
	//var temp string
	//fmt.Scanf("%s", &temp)
	//temp = temp[:k]
	//
	//mapp := make(map[[2]int]bool)
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		mapp[[2]int{i, j}] = false
	//	}
	//}
	//mapp[[2]int{0, 0}] = true
	//delete(mapp, [2]int{0, 0})
	//
	//ini := [2]int{0, 0}
	//for i := 0; i < len(temp); i++ {
	//	switch temp[i] {
	//	case 'W':
	//		ini = [2]int{ini[0] - 1, ini[1]}
	//	case 'A':
	//		ini = [2]int{ini[0], ini[1] - 1}
	//	case 'S':
	//		ini = [2]int{ini[0] + 1, ini[1]}
	//	case 'D':
	//		ini = [2]int{ini[0], ini[1] + 1}
	//	}
	//
	//	if v, ok := mapp[ini]; ok {
	//		if v == false {
	//			mapp[ini] = true
	//			delete(mapp, ini)
	//		}
	//	}
	//
	//	if len(mapp) == 0 {
	//		fmt.Println("Yes")
	//		fmt.Println(i + 1)
	//		return
	//	}
	//}
	//
	//fmt.Println("No")
	//res := 0
	//for _, b := range mapp {
	//	if !b {
	//		res++
	//	}
	//}
	//fmt.Println(res)

	//res := 0
	//var n, t int
	//fmt.Scanf("%d %d\n", &n, &t)
	//ints := make([]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	var temp int
	//	fmt.Scanf("%d", &temp)
	//	ints = append(ints, temp)
	//}
	//
	//nums := make([]int, len(ints), cap(ints))
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = (i + 1) * t
	//}
	//
	//sort.Ints(ints)
	//x := 0
	//y := 0
	//
	//for x < len(nums) && y < len(ints) {
	//	if ints[y] < nums[x] {
	//		res++
	//		y++
	//	} else if ints[y] >= nums[x] {
	//		x++
	//		y++
	//	}
	//}
	//fmt.Println(res)
}
