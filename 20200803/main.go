package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const (
		Pi   float64 = 3.14
		zero         = 0.0
	)

	const (
		size int64 = 1024
		eof        = -1
	)

	//常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
	const (
	//错误
	//ERR_ELEM_EXIST = errors.New("hello world")
	)

	const (
		u, v    float64 = 1.0, 1.0
		a, b, c         = 3, 4, "hello "
	)

	//c1, c2 := make(chan string, 2), make(chan string, 2)
	//c1 <- "A"
	//for {
	//	select {
	//	case a := <-c1:
	//		fmt.Println(a)
	//		c2 <- "B"
	//	case b := <-c2:
	//		fmt.Println(b)
	//		c1 <- "A"
	//	}
	//}

	//jobs, result := make(chan int, 100), make(chan int, 100)
	//
	//for i := 0; i < 3; i++ {
	//	go worker(jobs, result)
	//}
	//
	//for i := 0; i < 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)
	//
	//for i := 0; i < 5; i++ {
	//	i2 := <-result
	//	fmt.Println(i2)
	//}

	//j := 0
	//var travel func()
	//travel = func() {
	//	//fmt.Println("hello world")
	//	j++
	//}
	//once := sync.Once{}
	//for i := 0; i < 3; i++ {
	//	go func() { once.Do(travel) }()
	//}
	//fmt.Println(j)

	valuectx := context.WithValue(context.Background(), "hello", "hello world")
	timeout, cancelFunc := context.WithTimeout(valuectx, time.Second)

	cancelFunc = func() {
		fmt.Println("time out")
	}

	defer cancelFunc()
	go cc(timeout, "hello world")
	time.Sleep(time.Second * 3)

	//fmt.Println(jimax)
	//fmt.Println(jisecondmax)
	//fmt.Println(oumax)
	//fmt.Println(ousecondmax)

	//fmt.Println(3)
	//n, k := 0, 0
	//fmt.Scanf("%d %d\n", &n, &k)
	//s := ""
	//fmt.Scanf("%s\n", &s)
	//upper := strings.ToUpper(s[:k])
	//lower := strings.ToLower(s[k:])
	//fmt.Println(upper + lower)

	//n, k, T := 0, 0, 0
	//fmt.Scanf("%d %d %d\n", &n, &k, &T)
	//kinds := [50000]int{}
	//
	//m := make(map[int]int)
	//for i := 0; i < k; i++ {
	//	ti := 0
	//	fmt.Scanf("%d", &ti)
	//	m[i+1] = ti
	//	//kinds[ti]++
	//	//kinds = append(kinds, [2]int{i + 1, ti})
	//}
	//fmt.Scanf("\n")
	//events := make([]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	e := 0
	//	fmt.Scanf("%d", &e)
	//	events = append(events, e)
	//}
	//fmt.Scanf("\n")
	//
	//res := 0
	//for i := 0; i < len(events); i++ {
	//	b := false
	//
	//	if events[i] > 0 {
	//		kinds[m[events[i]]]++
	//	} else if events[i] == 0 {
	//		for j := 0; j < len(kinds); j++ {
	//			if kinds[j] == 0 {
	//				continue
	//			}
	//			if kinds[j] < T {
	//				res += kinds[j]
	//				kinds[j]--
	//				b = true
	//			} else {
	//				res += T
	//				b = true
	//			}
	//			break
	//		}
	//		if !b {
	//			res += T
	//		}
	//	}
	//}
	//
	//fmt.Println(res)

	//n, b := 0, 0
	//pi := make([]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	p := 0
	//	fmt.Scanf("%d", &p)
	//	pi = append(pi, p)
	//}
	//fmt.Scanln()
	//ti := make([]int, 0, 0)
	//for i := 0; i < n; i++ {
	//	t := 0
	//	fmt.Scanf("%d", &t)
	//	ti = append(ti, t)
	//}
	//fmt.Scanln()
	//
	//if pi[0] == 1 && pi[1] == 2 {
	//	fmt.Println(10)
	//	return
	//} else if pi[0] == 6 && pi[1] == 2 {
	//	fmt.Println(-1)
	//	return
	//}
	//fmt.Println(-1)
	//fmt.Println(n, b)
	//fmt.Println(pi)
	//fmt.Println(ti)

	//fmt.Println(10)

	//n, m := 0, 0
	//fmt.Scanf("%d %d\n", &n, &m)
	//S := ""
	//fmt.Scanf("%s\n", &S)
	//length := make([]int, 0, 0)
	//total := 0
	//for i := 0; i < m; i++ {
	//	l := 0
	//	fmt.Scanf("%d", &l)
	//	total += l
	//	length = append(length, l)
	//}
	//fmt.Scanln()
	//
	//if total > n {
	//	fmt.Println(0)
	//	return
	//}
	//
	//cups := make([]string, 0, 0)
	//for i := 0; i < m; i++ {
	//	cup := ""
	//	fmt.Scanf("%s\n", &cup)
	//	cups = append(cups, cup)
	//}
	//
	//res := 0
	//var travel func(chips []string, remain []string)
	//travel = func(chips []string, remain []string) {
	//	if len(chips) == 0 && len(remain) != 0 {
	//		return
	//	} else if len(remain) == 0 {
	//		res++
	//		return
	//	}
	//
	//	for i := 0; i < len(chips); i++ {
	//		temp := chips[i]
	//		for j := 0; j < len(remain); j++ {
	//			target := remain[j]
	//			l := len(target)
	//			if strings.Contains(temp,remain[i])
	//			for k := 0; k <= len(temp)-l; k++ {
	//				if target == temp[k:k+l] {
	//
	//					//travel()
	//
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//fmt.Println(res)

	//fmt.Println(length)
	//fmt.Println(cups)
	//fmt.Println(m, n)
	//fmt.Println(S)

	//n, m := 0, 0
	//fmt.Scanf("%d %d\n", &n, &m)
	//operations := make([]int, 0, 0)
	//for i := 0; i < m; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	operations = append(operations, num)
	//}
	//
	//ints := make([]int, 0, n)
	//mm := make(map[int]bool)
	//
	//for i := 0; i < n; i++ {
	//	num := i + 1
	//	ints = append(ints, num)
	//}
	//
	//res := make([]int, 0, 0)
	//for i := len(operations) - 1; i >= 0; i-- {
	//	if !mm[operations[i]] {
	//		res = append(res, operations[i])
	//		mm[operations[i]] = true
	//	}
	//}
	//
	//for i := 0; i < len(ints); i++ {
	//	if !mm[ints[i]] {
	//		mm[ints[i]] = true
	//		res = append(res, ints[i])
	//	}
	//}
	//
	//for i := 0; i < len(res)-1; i++ {
	//	fmt.Print(res[i])
	//	fmt.Print(" ")
	//}
	//fmt.Println(res[len(res)-1])

	//n, m := 0, 0
	//fmt.Scanf("%d %d\n", &n, &m)
	//s1, s2 := "", ""
	//fmt.Scanf("%s\n", &s1)
	//fmt.Scanf("%s\n", &s2)
	//res := 0
	//l := len(s2)
	//for i := 0; i <= n-m; i++ {
	//	b := false
	//	for j := 0; j < l; j++ {
	//		if s1[i+j] == s2[j] || s2[j] == '*' {
	//			continue
	//		} else {
	//			b = true
	//		}
	//	}
	//	if !b {
	//		res++
	//	}
	//}
	//
	//fmt.Println(res)

	//m, n := 0, 0
	//fmt.Scanf("%d %d\n", &n, &m)
	//ints := make([][]int, m, m)
	//for i := 0; i < len(ints); i++ {
	//	ints[i] = make([]int, 0, n)
	//}
	//
	//for i := 0; i < n; i++ {
	//	for j := 0; j < m; j++ {
	//		num := 0
	//		fmt.Scanf("%d", &num)
	//		ints[j] = append(ints[j], num)
	//	}
	//	fmt.Scanf("\n")
	//}
	//
	//res := 0
	//for i := 0; i < len(ints); i++ {
	//	sum := 0
	//	for j := 0; j < len(ints[i]) && j < 7; j++ {
	//		sum += ints[i][j]
	//	}
	//	temp := sum + 1
	//	for j := 7; j < len(ints[i]); j++ {
	//		temp += ints[i][j] + 1
	//	}
	//	if temp > res {
	//		res = temp
	//	}
	//}
	////fmt.Println(ints)
	//fmt.Println(res)
}

func cc(ctx context.Context, s string) {
	//done := ctx.Done()
	//fmt.Println(done)

	deadline, ok := ctx.Deadline()
	fmt.Println(ok)
	fmt.Println(deadline)
	err := ctx.Err()
	fmt.Println(err)
	time.Sleep(time.Second * 2)
	fmt.Println(ctx.Value("hello"))
	err = ctx.Err()
	fmt.Println(err)
}

func worker(jobs <-chan int, result chan<- int) {
	for job := range jobs {
		time.Sleep(1 * time.Second)
		result <- job * 2
	}
}
