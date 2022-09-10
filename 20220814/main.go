package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ss := make([]int, 0, 0)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text := sc.Text()
		split := strings.Split(text, ",")
		for _, value := range split {
			atoi, _ := strconv.Atoi(value)
			ss = append(ss, atoi)
		}
	}

	var n string
	fmt.Scanf("%s\n", &n)

	m := make(map[string]bool)
	mf := make(map[string]bool)
	var travel func(s string)
	travel = func(s string) {
		if len(s) > len(n) {
			mf[s] = true
			return
		}

		if len(s) == len(n) {
			if s[0] == '0' {
				mf[s] = true
				return
			}

			if m[s] {
				return
			}

			ss, err1 := strconv.Atoi(s)
			nn, err2 := strconv.Atoi(n)
			if err1 != nil || err2 != nil {
				mf[s] = true
				return
			} else {
				if ss <= nn {
					m[s] = true
					return
				}
				mf[s] = true
				return
			}
		}

		if s[0] == '0' {
			mf[s] = true
		} else {
			m[s] = true
		}

		for i := 0; i <= 9; i++ {
			if !m[string('0'+i)+s] && !mf[string('0'+i)+s] {
				travel(string('0'+i) + s)
			}

			if !m[s+string('0'+i)] && !mf[s+string('0'+i)] {
				travel(s + string('0'+i))
			}
		}
	}
	travel("25")

	fmt.Println(len(m))

	//for s := range m {
	//	fmt.Println(s)
	//
	//}

	//var s string
	//fmt.Scanf("%s\n", &s)
	//s = s[2:]
	//s = s[:len(s)-2]
	//split := strings.Split(s, "],[")
	//ints := make([][]int, len(split), cap(split))
	//for i := 0; i < cap(ints); i++ {
	//	s2 := split[i]
	//	temp := strings.Split(s2, ",")
	//	for _, value := range temp {
	//		atoi, err := strconv.Atoi(value)
	//		if err != nil {
	//
	//		} else {
	//			ints[i] = append(ints[i], atoi)
	//
	//		}
	//	}
	//}
	//
	//fmt.Println(canvisit(ints))

}

func canvisit(rooms [][]int) bool {
	n := len(rooms)
	vis := make([]bool, n)
	num := 0

	var travel func([][]int, int)
	travel = func(rooms [][]int, x int) {
		vis[x] = true
		num++
		for _, it := range rooms[x] {
			if !vis[it] {
				travel(rooms, it)
			}
		}
	}

	travel(rooms, 0)
	return num == n
}
