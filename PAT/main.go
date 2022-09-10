package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	format := ""
	//reader要放在外面啊啊啊啊!!!
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		bytes, _ := reader.ReadBytes('\n')
		split := strings.Split(strings.Trim(string(bytes), "\r\n"), " ")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		c, _ := strconv.Atoi(split[2])
		format += fmt.Sprintf("Case #%d: %v\n", i+1, a+(b-c) > 0)
	}
	fmt.Println(format[:len(format)-1])
	//reader := bufio.NewReader(os.Stdin)
	//bytes, _ := reader.ReadBytes('\n')
	//bytes = bytes[:len(bytes)-1]
	//split := strings.Split(string(bytes), " ")
	//res := make([]int, 0, 0)
	//for i := 0; i < len(split); i += 2 {
	//	a, _ := strconv.Atoi(split[i])
	//	b, _ := strconv.Atoi(split[i+1])
	//
	//	temp := a * b
	//	temp1 := b - 1
	//	if b == 0 {
	//		continue
	//	} else if temp == 0 && temp1 == 0 {
	//		res = append(res, 0, 0)
	//	} else if temp != 0 {
	//		res = append(res, temp, temp1)
	//	}
	//}
	//ans := make([]string, 0, len(res))
	//for i := 0; i < len(res); i++ {
	//	ans = append(ans, fmt.Sprintf("%d", res[i]))
	//}
	//join := strings.Join(ans, " ")
	//fmt.Print(join)

	//var num string
	//fmt.Scanf("%s", &num)
	//thisNum := writeThisNum1(num)
	//fmt.Println(thisNum)

	//correct := answerCorrect("PAT")
	//correct := answerCorrect("PAAT")
	//correct := answerCorrect("AAPATAA")
	//correct := answerCorrect("AAPAATAAAA")
	//correct := answerCorrect("xPATx")
	//correct := answerCorrect("PT")
	//correct := answerCorrect("Whatever")
	//correct := answerCorrect("APAAATAA")
	//correct := answerCorrect("APT")
	//correct := answerCorrect("APATTAA")
	//fmt.Println(correct)

	//var nums int
	//fmt.Scanf("%d\n", &nums)
	//for i := 0; i < nums; i++ {
	//	var s string
	//	fmt.Scanf("%s\n", &s)
	//	fmt.Println(answerCorrect(s))
	//}

	//threeNAddOne()

	//var n int
	//fmt.Scanf("%d\n", &n)
	//i := susudui(n)
	//fmt.Println(i)

	//var length int
	//var n int
	//fmt.Scanf("%d %d\n", &length, &n)
	//ints := make([]int, length, length)
	//for i := 0; i < length; i++ {
	//	var num int
	//	fmt.Scanf("%d", &num)
	//	ints[i] = num
	//}
	//
	//youyi(ints, n)

	//reader := bufio.NewReader(os.Stdin)
	//readByte, _ := reader.ReadBytes('\n')
	//readByte = readByte[:len(readByte)-1]
	//split := strings.Split(string(readByte), " ")
	//for i := 0; i < len(split)/2; i++ {
	//	split[i], split[len(split)-1-i] = split[len(split)-1-i], split[i]
	//}
	//join := strings.Join(split, " ")
	//fmt.Println(join)
}

//数组元素循环右移问题  艹!!!太久没做忘记了!!!!!! 数组倒序问题 三次
func youyi(ints []int, n int) {
	var travel func(ints []int)
	travel = func(ints []int) {
		for i := 0; i < len(ints)/2; i++ {
			ints[i], ints[len(ints)-1-i] = ints[len(ints)-1-i], ints[i]
		}
	}
	n = n % len(ints)
	travel(ints[:len(ints)-n])
	travel(ints[len(ints)-n:])
	travel(ints)

	res := make([]string, len(ints), len(ints))
	for i := 0; i < len(ints); i++ {
		res[i] = fmt.Sprintf("%d ", ints[i])
	}
	join := strings.Join(res, " ")
	fmt.Println(join)
}

func susudui(n int) int {
	res := 0
	last := 1
	num := 2
	visited := make(map[int]bool)

	var travel func(x int) bool
	travel = func(x int) bool {
		visited[x] = true
		for i := 2; i <= int(math.Pow(float64(i), 0.5)); i++ {
			if x%i == 0 {
				for j := 2; j*x <= n; j++ {
					visited[j*x] = true
				}
				return false

			}
		}

		for j := 2; j*x <= n; j++ {
			visited[j*x] = true
		}
		return true
	}
	for num <= n {
		if visited[num] {
			num++
			continue
		}

		if travel(num) {
			if num-last == 2 {
				res++
			}
			last = num
		}
		num++
	}

	return res
}

func threeNAddOne() {
	res := make([]int, 0, 0)
	nums := 0
	fmt.Scanf("%d\n", &nums)
	ints := make([]int, 0, nums)
	for i := 0; i < nums; i++ {
		var n int
		fmt.Scanf("%d", &n)
		ints = append(ints, n)
	}

	sort.Ints(ints)
	m := make(map[int]bool)
	for i := 0; i < len(ints); i++ {
		m[ints[i]] = false
	}

	for i := 0; i < len(ints); i++ {
		if m[ints[i]] {
			continue
		}
		for ints[i] != 1 {
			if ints[i]%2 == 0 {
				ints[i] = ints[i] / 2
				if _, ok := m[ints[i]]; ok {
					m[ints[i]] = true
				}
			} else {
				ints[i] = (3*ints[i] + 1) / 2
				if _, ok := m[ints[i]]; ok {
					m[ints[i]] = true
				}
			}
		}
	}
	for key, v := range m {
		if v == false {
			res = append(res, key)
		}
	}

	sort.Ints(res)
	ans := make([]string, 0, 0)
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, fmt.Sprintf("%v", res[i]))
	}
	join := strings.Join(ans, " ")

	fmt.Println(join)

}

func scoreRange() {
	var n int
	fmt.Scanf("%d\n", &n)
	high := math.MinInt8
	highRes := ""
	low := math.MaxInt8
	lowRes := ""
	for i := 0; i < n; i++ {
		var s string
		var ss2 string
		var ss1 string
		fmt.Scanf("%s %s %s\n", &ss2, &ss1, &s)
		score, _ := strconv.Atoi(s)
		if score > high {
			high = score
			highRes = ss2 + " " + ss1
		}
		if score < low {
			low = score
			lowRes = ss2 + " " + ss1
		}
	}
	fmt.Println(highRes)
	fmt.Println(lowRes)
}
func answerCorrect(s string) string {
	m := make(map[byte]int)
	var a string
	var b string
	var c string
	var p int
	var t int
	for i := 0; i < len(s); i++ {
		if s[i] == 'P' {
			p = i
		} else if s[i] == 'T' {
			t = i
		}

		m[s[i]]++
	}

	if t <= p {
		return "NO"
	}

	a = s[:p]
	c = s[t+1:]
	b = s[p+1 : t]

	if len(m) != 3 || m['P'] != 1 || m['T'] != 1 || len(b) == 0 {
		return "NO"
	}

	if len(b) > 1 && len(c) < len(a) || len(b)*len(a) != len(c) {
		return "NO"
	}

	return "YES"
}

func find(num int) (res int) {
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = (num*3 + 1) / 2
		}
		res++
	}
	return
}

func writeThisNum(n int64) string {
	m := map[byte]string{'1': "yi", '2': "er", '3': "san", '4': "si", '5': "wu", '6': "liu", '7': "qi", '8': "ba", '9': "jiu", '0': "ling"}
	ans := make([]string, 0, 0)
	res := 0
	c := n % 10
	n = n / 10
	res += int(c)
	s := fmt.Sprintf("%v", res)
	for i := 0; i < len(s); i++ {
		ans = append(ans, m[s[i]])
	}
	return strings.Join(ans, " ")

}

func writeThisNum1(n string) string {
	m := map[byte]string{'1': "yi", '2': "er", '3': "san", '4': "si", '5': "wu", '6': "liu", '7': "qi", '8': "ba", '9': "jiu", '0': "ling"}
	ans := make([]string, 0, 0)
	res := 0
	for i := 0; i < len(n); i++ {
		atoi, _ := strconv.Atoi(string(n[i]))
		res += atoi
	}
	s := fmt.Sprintf("%v", res)
	for i := 0; i < len(s); i++ {
		ans = append(ans, m[s[i]])
	}
	return strings.Join(ans, " ")

}
