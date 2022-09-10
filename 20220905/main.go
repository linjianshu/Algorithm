package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
)

func init() {
	fmt.Println(20220906)
}
func main() {
	ints := guibingpaixu([]int{1, 5, 4, 3, 2, 6})
	fmt.Println(ints)

	i := maopaopaixu([]int{1, 5, 4, 3, 2, 6})
	fmt.Println(i)

	i2 := selectpaixu([]int{1, 5, 4, 3, 2, 6})
	fmt.Println(i2)

	marshal, err := json.Marshal(&Student{
		Sno:   "2020170281",
		Sname: "林健树",
		Ssex:  1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))

	stu := Student{}
	err = json.Unmarshal([]byte(`{"学号":"2020170281","姓名":"林健树","性别":1}`), &stu)
	if err != nil {
		panic(err)
	}

	fmt.Println(stu)

	m["2020170281"] = "林健树"
	m["2020170282"] = "江文涛"
	m["2020170283"] = "李继恩"

	bytes, err := json.Marshal(&m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	mm := make(map[string]string)
	err = json.Unmarshal(bytes, &mm)
	if err != nil {
		panic(err)
	}
	fmt.Println(mm)

	nums := []int{1, 4, 3, 6, 2}
	mh := myheap{}
	for j := 0; j < len(nums); j++ {
		heap.Push(&mh, nums[j])
	}

	i5, err := json.Marshal(&mh)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(i5))

	var unmarshalMyheap myheap
	err = json.Unmarshal(i5, &unmarshalMyheap)
	if err != nil {
		panic(err)
	}
	fmt.Println(unmarshalMyheap)

	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 2)
	//ch1 <- 1
	//i3, err := json.Marshal(&ch1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(i3))
	//
	//err = json.Unmarshal(i3, &ch2)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(ch2)

	i4, err := json.Marshal(guibingpaixu)
	if err != nil {
		panic(err)
	}
	fmt.Println(i4)

	var travel func(ints []int) []int
	err = json.Unmarshal(i4, &travel)
	if err != nil {
		panic(err)
	}
	fmt.Println(travel)
}

type myheap []int

func (m *myheap) Len() int {
	return len(*m)
}

func (m *myheap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *myheap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap) Push(i interface{}) {
	*m = append(*m, i.(int))
}

func (m *myheap) Pop() interface{} {
	x := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return x
}

type Dog struct {
}

func (d Dog) Run(a int) int {
	return a * 2
}

type Runner interface {
	Run(int) int
}

var m = map[string]string{}

type Student struct {
	Sno   string `json:"学号"`
	Sname string `json:"姓名"`
	Ssex  int    `json:"性别"`
}

func guibingpaixu(ints []int) []int {
	var merge func(a []int, b []int) []int
	merge = func(a []int, b []int) []int {
		if len(a) == 0 {
			return b
		} else if len(b) == 0 {
			return a
		}
		ans := make([]int, 0, 0)
		ai, bi := 0, 0
		for ai < len(a) || bi < len(b) {
			if ai == len(a) {
				ans = append(ans, b[bi])
				bi++
			} else if bi == len(b) {
				ans = append(ans, a[ai])
				ai++
			} else {
				if a[ai] < b[bi] {
					ans = append(ans, a[ai])
					ai++
				} else {
					ans = append(ans, b[bi])
					bi++
				}
			}
		}
		return ans
	}

	var travel func(ints []int) []int
	travel = func(ints []int) []int {
		if len(ints) == 0 || len(ints) == 1 {
			return ints
		}

		mid := len(ints) / 2
		a := travel(ints[:mid])
		b := travel(ints[mid:])
		return merge(a, b)
	}

	return travel(ints)
}

func maopaopaixu(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		for j := 0; j < len(ints)-1-i; j++ {
			if ints[j] > ints[j+1] {
				ints[j], ints[j+1] = ints[j+1], ints[j]
			}
		}
	}
	return ints
}

func selectpaixu(ints []int) []int {
	for i := 0; i < len(ints); i++ {
		temp := i
		for j := i + 1; j < len(ints); j++ {
			if ints[j] < ints[temp] {
				temp = j
			}
		}
		ints[i], ints[temp] = ints[temp], ints[i]
	}

	return ints
}

func insertpaixu(ints []int) []int {
	res := make([]int, 0, 0)

	return res
}
