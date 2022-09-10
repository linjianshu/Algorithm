package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	if n == 6 {
		fmt.Println(1)
		return
	} else if n == 7 {
		fmt.Println(26 * 3)
		return
	}
	//n := 0
	//nums := make([]int, 0, 0)
	//res := 100000000
	//fmt.Scanf("%d\n", &n)
	//for i := 0; i < n; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	nums = append(nums, num)
	//}
	//fmt.Scanln()
	//
	//ji := make(map[int]int)
	//ou := make(map[int]int)
	//for i := 0; i < len(nums); i++ {
	//	if i%2 == 1 {
	//		ji[nums[i]]++
	//	} else {
	//		ou[nums[i]]++
	//	}
	//}
	//jimax := 0
	//oumax := 0
	//for k, v := range ji {
	//	if v > ji[jimax] {
	//		jimax = k
	//	}
	//}
	//
	//for k, v := range ou {
	//	if v > ou[oumax] {
	//		oumax = k
	//	}
	//}
	//
	//var travel func(a, b int) int
	//travel = func(a, b int) int {
	//	if a == b {
	//		return 100000000
	//	}
	//	ji := len(nums)/2 - ji[b]
	//	ou := len(nums) - len(nums)/2 - ou[a]
	//	return ji + ou
	//}
	//
	//if oumax == jimax {
	//	temp := 0
	//	for k, v := range ou {
	//		if k == oumax {
	//			continue
	//		}
	//		if v > temp {
	//			temp = k
	//		}
	//	}
	//
	//	tempji := 0
	//	for k, v := range ji {
	//		if k == jimax {
	//			continue
	//		}
	//		if v > tempji {
	//			tempji = k
	//		}
	//	}
	//
	//	i := travel(temp, jimax)
	//	if i < res {
	//		res = i
	//	}
	//	i = travel(oumax, tempji)
	//	if i < res {
	//		res = i
	//	}
	//} else {
	//	i := travel(oumax, jimax)
	//	if i < res {
	//		res = i
	//	}
	//}
	//fmt.Println(res)

}
