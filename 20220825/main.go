package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s := ""
	fmt.Scanf("%s\n", &s)
	s = s[1:]
	s = s[:len(s)-1]
	split := strings.Split(s, ",")
	ints := make([][2]int, len(split), len(split))


	for i := 0; i < len(split); i++ {
		atoi, err := strconv.Atoi(split[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		ints[i] = [2]int{i, atoi}
	}
	res := 0
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			x := ints[j][0] - ints[i][0]
			y := int(math.Min(float64(ints[i][1]), float64(ints[i][1])))
			temp := x * y
			if res < temp {
				res = temp
			}
		}
	}
	fmt.Println(res)
	//fmt.Println(ints)
	//fmt.Println(DagPathNum([][]int{{1, 2, 3}, {3}, {3}, {4}, {}}))
}

func DagPathNum(nodes [][]int) int {
	// write code here
	res := 0
	var travel func(i int)
	travel = func(i int) {
		if i == len(nodes)-1 {
			res++
			return
		}

		roads := nodes[i]
		for j := 0; j < len(roads); j++ {
			travel(roads[j])
		}
	}

	travel(0)

	return res
}
