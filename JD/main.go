package main

import (
	"fmt"
	"sort"
)

func removeAnagrams(words []string) []string {
	res := make([]string, 0, 0)

	pre := [26]int{}
	for i := 0; i < len(words); i++ {
		temp := [26]int{}
		for j := 0; j < len(words[i]); j++ {
			u := words[i][j] - 'a'
			temp[u]++
		}

		if temp != pre {
			res = append(res, words[i])
			pre = temp
		} else {
			pre = temp
		}
	}
	return res
}

//func maxConsecutive(bottom int, top int, special []int) int {
//	sort.Ints(special)
//	index := 0
//	res := math.MinInt
//	for index < len(special) {
//		pre := bottom
//		for i := bottom; i <= top; i++ {
//			if i == special[index] {
//				sub := i - pre
//				if sub > res {
//					res = sub
//				}
//				pre = i
//				index++
//			}
//		}
//	}
//	return res
//}

func largestCombination(candidates []int) int {
	res := 1
	for i := 0; i < 63; i++ {
		res = res | 1<<i
	}
	fmt.Printf("%b", res)
	for i := 0; i < len(candidates); i++ {
		res = res & candidates[i]
	}
	return res
}

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	res := 0

	for i := 0; i < len(special)-1; i++ {
		sub := special[i+1] - special[i] - 1
		if sub == 1 {

		} else if sub > res {
			res = sub
		}
	}
	i := special[0] - bottom
	if i == 1 {

	} else if i > res {
		res = i
	}

	i2 := top - special[len(special)-1]
	if i2 == 1 {

	} else if i2 > res {
		res = i2
	}
	return res
}

func main() {
	//anagrams := removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})
	//anagrams := removeAnagrams([]string{"a", "b", "a"})
	//anagrams := removeAnagrams([]string{"a", "b", "c", "d", "e"})
	//fmt.Println(anagrams)

	//consecutive := maxConsecutive(2, 9, []int{4, 6})
	//consecutive := maxConsecutive(3, 15, []int{7, 9, 13})
	//consecutive := maxConsecutive(6, 8, []int{7, 6, 8})
	//fmt.Println(consecutive)

	combination := largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
	fmt.Println(combination)

	//a("yesno", "yyseno")
	//a("pop", "oppppipi")
	//a("abcd", "gfbadcfgg")
	//a("faq", "qaaf")
	//a("abc", "abc")
	//n := 0
	//fmt.Scanf("%d\n", &n)
	//sl1 := make([]string, 0, n)
	//sl2 := make([]string, 0, n)
	//for i := 0; i < n; i++ {
	//	s1 := ""
	//	s2 := ""
	//	fmt.Scanf("%v\n", &s1)
	//	fmt.Scanf("%v\n", &s2)
	//	sl1 = append(sl1, s1)
	//	sl2 = append(sl2, s2)
	//}
	//
	//for i := 0; i < n; i++ {
	//	a(sl1[i], sl2[i])
	//}

}

func a(target string, s2 string) {
	m := make(map[int]int)
	for i := 0; i < len(target); i++ {
		m[int(target[i]-'a')]++
	}

	left, right := 0, 0
	for right < len(s2) {
		for right-left != len(target) {
			if v, ok := m[int(s2[right]-'a')]; ok {
				if v == 1 {
					delete(m, int(s2[right]-'a'))
				} else {
					m[int(s2[right])-'a']--
				}
			} else {
				m[int(s2[right]-'a')]--
			}
			right++
		}
		if len(m) == 0 {
			fmt.Println("YES")
			return
		}
		m[int(s2[left]-'a')]++
		if m[int(s2[left]-'a')] == 0 {
			delete(m, int(s2[left]-'a'))
		}
		left++
	}
	fmt.Println("NO")
}
