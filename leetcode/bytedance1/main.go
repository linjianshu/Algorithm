package main

import (
	"fmt"
	"sort"
)

func main() {

	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
	//numbers := addTwoNumbers(&ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//	},
	//}, &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//})
	//
	//fmt.Println(numbers)

	//substring := lengthOfLongestSubstring("abc")
	//fmt.Println(substring)

	palindrome := longestPalindrome("cbbd")
	fmt.Println(palindrome)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//要考虑特殊情况 以及最后一位的进位 以及进位标识符的重置
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := &ListNode{}
	cur := res
	isCarry := 0

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 == nil {
			sum = l2.Val + isCarry
			l2 = l2.Next
		} else if l2 == nil {
			sum = l1.Val + isCarry
			l1 = l1.Next
		} else {
			sum = l1.Val + l2.Val + isCarry
			l1 = l1.Next
			l2 = l2.Next
		}
		if sum >= 10 {
			isCarry = 1
			sum -= 10
		} else {
			isCarry = 0
		}
		cur.Next = &ListNode{
			Val:  sum,
			Next: nil,
		}
		cur = cur.Next
	}

	if isCarry == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return res.Next
}

//要注意 滑动窗口的扩大和缩小 优化还不太会 使用while移动左指针优化重复的循环
func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[uint8]int)
	left, right := 0, 0
	for left < len(s) {
		for right < len(s) {
			if m[s[right]] != 0 {
				break
			} else {
				//扩大窗口
				m[s[right]]++
				right++
			}
		}
		temp := right - left
		if temp > res {
			res = temp
		}
		//缩小窗口
		m[s[left]]--
		left++
	}
	return res
}

//中心拓展 好像不是特别熟练 差点忘记了艹
func longestPalindrome(s string) string {
	res := ""
	distance := 0
	for i := 0; i < len(s); i++ {
		temp := 1
		left, right := i-1, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			temp += 2
			left--
			right++
		}
		if temp > distance {
			distance = temp
			res = s[left+1 : right]
		}
	}

	for i := 0; i < len(s); i++ {
		temp := 0
		left, right := i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			temp += 2
			left--
			right++
		}
		if temp > distance {
			distance = temp
			res = s[left+1 : right]
		}
	}
	return res
}
