package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	nums := make([][]int, n, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, 0, 3)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {

			//if j == 1 {
			//	for k := 0; k < 2; k++ {
			//		num := 0
			//		fmt.Scanf("%d", &num)
			//		nums[i] = append(nums[i], num)
			//	}
			//} else {
			//	for k := 0; k < 3; k++ {
			//		num := 0
			//		fmt.Scanf("%d", &num)
			//		nums[i] = append(nums[i], num)
			//
			//	}
			//}
			//fmt.Scanln()

			scanner.Scan()
			ss := scanner.Text()
			split := strings.Split(ss, " ")
			for _, value := range split {
				atoi, err := strconv.Atoi(value)
				if err != nil {

				} else {
					nums[i] = append(nums[i], atoi)
				}
			}
		}
	}
	fmt.Println(nums)

	var travel func(ints []int) int
	travel = func(ints []int) int {
		ans := 0
		if ints[1]-ints[0] == ints[2]-ints[1] {
			ans++
		}

		if ints[6]-ints[5] == ints[7]-ints[6] {
			ans++
		}

		if ints[3]-ints[0] == ints[5]-ints[3] {
			ans++
		}

		if ints[4]-ints[2] == ints[7]-ints[4] {
			ans++
		}

		a := [2]int{ints[3], ints[4]}
		b := [2]int{ints[1], ints[6]}
		c := [2]int{ints[0], ints[7]}
		d := [2]int{ints[5], ints[2]}

		max := 1

		xx := 1
		i := (a[1] - a[0]) / 2
		temp := a[0] + i
		if temp-b[0] == b[1]-temp {
			xx++
		}
		if temp-c[0] == c[1]-temp {
			xx++
		}
		if temp-d[0] == d[1]-temp {
			xx++
		}
		if xx > max {
			max = xx
		}

		xx = 1
		i = (b[1] - b[0]) / 2
		temp = b[0] + i
		if temp-a[0] == a[1]-temp {
			xx++
		}
		if temp-c[0] == c[1]-temp {
			xx++
		}
		if temp-d[0] == d[1]-temp {
			xx++
		}
		if xx > max {
			max = xx
		}

		xx = 1
		i = (c[1] - c[0]) / 2
		temp = c[0] + i
		if temp-a[0] == a[1]-temp {
			xx++
		}
		if temp-b[0] == b[1]-temp {
			xx++
		}
		if temp-d[0] == d[1]-temp {
			xx++
		}
		if xx > max {
			max = xx
		}

		xx = 1
		i = (d[1] - d[0]) / 2
		temp = d[0] + i
		if temp-a[0] == a[1]-temp {
			xx++
		}
		if temp-b[0] == b[1]-temp {
			xx++
		}
		if temp-c[0] == c[1]-temp {
			xx++
		}
		if xx > max {
			max = xx
		}

		return ans + max
	}
	for i := 0; i < len(nums); i++ {
		res := travel(nums[i])
		sprintf := fmt.Sprintf("Case #%d: %d\n", i+1, res)
		fmt.Print(sprintf)
	}
}
