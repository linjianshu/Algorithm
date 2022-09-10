package main

import (
	"fmt"
)

func main() {
	//s := ""
	//fmt.Scanf("%s\n", &s)
	//s = s[1:]
	//s = s[:len(s)-1]
	//split := strings.Split(s, ",")
	//ints := make([]int, 0, len(split))
	//
	//for i := 0; i < len(split); i++ {
	//	atoi, err := strconv.Atoi(split[i])
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	ints = append(ints, atoi)
	//}
	//
	//res := 0
	//l, r := 0, len(ints)-1
	//for l < r {
	//	length := r - l
	//	width := int(math.Min(float64(ints[l]), float64(ints[r])))
	//	temp := length * width
	//	if temp > res {
	//		res = temp
	//	}
	//	if math.Min(float64(ints[l+1]), float64(ints[r])) <= math.Min(float64(ints[l]), float64(ints[r-1])) {
	//		r -= 1
	//	} else {
	//		l += 1
	//	}
	//}
	//fmt.Println(res)

	//ints := []int{1, 2, 3, 5, 6, 4, 3, 2, 1}
	//fmt.Println(findMax(ints))

	fmt.Println(buyFruit())
}

func findMax(ints []int) int {
	idx := 0
	for i := 1; i < len(ints); i++ {
		if ints[i] >= ints[idx] {
			idx = i
		} else {
			break
		}
	}
	return ints[idx]
}

func a(a, b int) (q int, int, float64) {
	return 0, 0, 0.0
}
func buyFruit() int {
	res := 0
	for i := 0; i <= 50/7; i++ {
		for j := 0; j <= (50-7*i)/3; j++ {
			for k := 0; k <= (50-7*i-3*j)*3; k += 3 {
				if 7*i+3*j+k/3 == 50 && i+j+k == 50 {
					res++
				}
			}
		}
	}

	return res
}
