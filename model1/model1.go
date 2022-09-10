package model1

import "fmt"

func init() {
	fmt.Println("model1")
	s := [2]int{}
	i := cap(s)
	fmt.Println(i)
	ss := make([]int, 0, 0)
	i2 := cap(ss)
	fmt.Println(i2)
	c := make(chan int, 1)
	i3 := cap(c)
	fmt.Println(i3)
}
