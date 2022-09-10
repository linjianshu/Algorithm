package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	l := make([]int, 0, 0)
	r := make([]int, 0, 0)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		l = append(l, num)
	}

	fmt.Scanln()

	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		r = append(r, num)
	}
	fmt.Scanln()
	beauti := make([]int, 0, 0)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scanf("%d", &num)
		beauti = append(beauti, num)
	}

	ans := make([]int, 0, 0)
	for i := 0; i < n; i++ {
		count := 0
		l1, r1 := l[i], r[i]
		for l1 <= r1 {
			temp := l1
			cal := 0
			wu := [10]int{}
			for temp != 0 {
				cc := temp % 10
				wu[cc]++
				//cal ^= cc
				temp /= 10
			}
			for k := 0; k < len(wu); k++ {
				if wu[k]%2 == 0 {
					continue
				} else {
					cal ^= k
				}
			}
			if cal == beauti[i] {
				count++
			}
			l1++
		}
		ans = append(ans, count)
	}

	res := make([]string, 0, 0)
	for i := 0; i < len(ans); i++ {
		itoa := strconv.Itoa(ans[i])
		res = append(res, itoa)
	}

	fmt.Println(strings.Join(res, " "))

	//n, k := 0, 0
	//fmt.Scanf("%d %d\n", &n, &k)
	//nums := make([]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	nums = append(nums, num)
	//}
	//
	//sort.Ints(nums)
	//sum := 0
	//for i := 0; i < len(nums); i++ {
	//	sum += nums[i]
	//}
	//avg := float64(sum) / float64(len(nums))
	//for float64(nums[len(nums)-1]) > avg*float64(k) {
	//	sub1 := avg - float64(nums[0])
	//	sub2 := float64(nums[len(nums)-1]) - avg
	//	if sub1 > sub2 {
	//		sum -= nums[0]
	//		nums = nums[1:]
	//		avg = float64(sum) / float64(len(nums))
	//	} else {
	//		sum -= nums[len(nums)-1]
	//		nums = nums[:len(nums)-1]
	//		avg = float64(sum) / float64(len(nums))
	//	}
	//
	//}
	//fmt.Println(len(nums))

	//w, n := 0, 0
	//nums := make([]int, 0, 0)
	//fmt.Scanf("%d %d\n", &w, &n)
	//for i := 0; i < n; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	nums = append(nums, num)
	//}
	//
	//success := 0
	//m := make(map[int]int)
	//var travel func(i int, rest []int, color []int)
	//travel = func(i int, rest []int, color []int) {
	//
	//	if len(color) == 0 {
	//		success++
	//		return
	//	}
	//	if len(rest) == 0 {
	//		success++
	//		return
	//	}
	//	if color[0] > len(rest) {
	//		return
	//	}
	//
	//	paint := color[0]
	//	for k := 0; k < paint; k++ {
	//		m[i+k]++
	//	}
	//
	//	if len(color) == 1 {
	//		success++
	//	}
	//	for j := 0; j < len(rest)-paint; j++ {
	//		if paint+1+j < len(rest) {
	//			travel(i+paint+1+j, rest[paint+1+j:], color[1:])
	//		}
	//	}
	//}
	//
	//rest := make([]int, w, w)
	//
	//travel(0, rest, nums)
	//max := 0
	//for _, v := range m {
	//	if v > max {
	//		max = v
	//	}
	//}
	//
	//count := 0
	//res := make([]int, 0, 0)
	//for key, value := range m {
	//	if value == success {
	//		count++
	//		res = append(res, key+1)
	//	}
	//}
	//
	//sort.Ints(res)
	//sss := make([]string, 0, 0)
	//for i := 0; i < len(res); i++ {
	//	itoa := strconv.Itoa(res[i])
	//	sss = append(sss, itoa)
	//}
	//fmt.Println(count)
	////fmt.Println(success)
	//fmt.Println(strings.Join(sss, " "))
}
