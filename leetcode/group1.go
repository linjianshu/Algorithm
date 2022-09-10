package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//使用哈希表也可以
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//对10取余就行了
func reverse(x int) int {
	flag := false
	s := strconv.Itoa(x)
	if strings.Contains(s, "-") {
		s = s[1:]
		flag = true
	}

	slice := []byte(s[:])
	fmt.Println(slice)
	length := len(slice)
	for i := 0; i < length/2; i++ {
		temp := slice[i]
		slice[i] = s[length-1-i]
		slice[length-1-i] = temp
	}

	newstring := string(slice)
	newint, err := strconv.Atoi(newstring)
	var result int
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if flag {
		result = 0 - newint
	} else {
		result = newint
	}
	if result > 1<<31-1 || result <= -1<<31 {
		return 0
	} else {
		return result
	}
}

//直接使用字符串的reverse函数也可以
func isPalindrome(x int) bool {
	y := x
	if x < 0 {
		return false
	}
	var newString string
	for x != 0 {
		temp := x % 10
		newString += strconv.Itoa(temp)
		x = (x - temp) / 10
	}

	atoi, _ := strconv.Atoi(newString)
	if atoi == y {
		return true
	} else {
		return false
	}
}

//可以使用读两位判断前后大小,然后移动两位再判断,直到移动到末尾
func romanToInt(s string) int {
	sum := 0
	var index int
	if strings.Contains(s, "IV") {
		index = strings.Index(s, "IV")
		s = s[:index] + s[index+2:]
		sum += 4
	}
	if strings.Contains(s, "IX") {
		index = strings.Index(s, "IX")
		s = s[:index] + s[index+2:]
		sum += 9
	}

	if strings.Contains(s, "XL") {
		index = strings.Index(s, "XL")
		s = s[:index] + s[index+2:]
		sum += 40
	}

	if strings.Contains(s, "XC") {
		index = strings.Index(s, "XC")
		s = s[:index] + s[index+2:]
		sum += 90
	}

	if strings.Contains(s, "CD") {
		index = strings.Index(s, "CD")
		s = s[:index] + s[index+2:]
		sum += 400
	}

	if strings.Contains(s, "CM") {
		index = strings.Index(s, "CM")
		s = s[:index] + s[index+2:]
		sum += 900
	}
	for _, v := range s {
		switch {
		case strings.Contains(string(v), "I"):
			sum += 1
		case strings.Contains(string(v), "V"):
			sum += 5
		case strings.Contains(string(v), "X"):
			sum += 10
		case strings.Contains(string(v), "L"):
			sum += 50
		case strings.Contains(string(v), "C"):
			sum += 100
		case strings.Contains(string(v), "D"):
			sum += 500
		case strings.Contains(string(v), "M"):
			sum += 1000
		}
	}

	return sum
}

func climbStairs(n int) int {
	var sum int
	if n == 1 {
		return 1
	}
	i2 := 2
	if n == 2 {
		return 2
	}
	i3 := 3
	if n == 3 {
		return i3
	} else {
		x := i3
		y := i2
		for i := n - 3; i > 0; i-- {
			sum = x + y
			y = x
			x = sum
		}
	}
	return sum
}

//可以取每一个单词的同一位置的字母，看是否相同。  反思:边界没考虑到 , 例如只有一个string进来,或者进来的都是相同的,思路不够缜密
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	s := strs[0]
	final := false
	for i := 0; i < len(s); i++ {
		b := true
		for j := 1; j < len(strs); j++ {
			b = strings.HasPrefix(strs[j], s[0:i+1])
			if !b {
				final = true
				break
			}
		}

		if final {
			return s[0:i]
		}
	}
	if !final {
		return strs[0]
	}

	return ""
}

//堆栈最好!!!把这题记住!如果想要用递归的话就得同时使得两个小问题成立,也就是 return isvalid(1)&&isvalid(2)
func isValid(s string) bool {
	m := map[string]string{"(": ")", "[": "]", "{": "}"}
	newStr := make([]string, 0, 10)
	for len(s) != 0 {
		if len(newStr) == 0 {
			newStr = append(newStr, string(s[0]))
		} else if m[newStr[len(newStr)-1]] == string(s[0]) {
			newStr = newStr[:len(newStr)-1]
		} else {
			newStr = append(newStr, string(s[0]))
		}
		s = s[1:]
	}
	if len(newStr) > 0 {
		return false
	} else {
		return true
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//可以用递归,递归属实看不懂
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	a := make([]int, 0, 10)
	for l1.Next != nil {
		a = append(a, l1.Val)
		l1 = l1.Next
	}
	a = append(a, l1.Val)
	for l2.Next != nil {
		a = append(a, l2.Val)
		l2 = l2.Next
	}
	a = append(a, l2.Val)
	sort.Ints(a)
	//定义一个要返回的listnode
	var listNode ListNode
	//定义一个指针类型的head
	var head *ListNode
	//赋初值 初指针指向listnode
	head = &listNode
	for i := 0; i < len(a)-1; i++ {
		head.Val = a[i]
		head.Next = &ListNode{
			Val:  a[i+1],
			Next: nil,
		}
		//next的指针赋给head
		head = head.Next
	}
	return &listNode
}

//双指针更好
func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			temp := nums[:i]
			temp = append(temp, nums[i+1:]...)
			nums = temp
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}

//对[]进行forrange 会遍历它的cap 而不是 len  可以选择不使用数组修改的方式 使用覆盖的方式
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			temp := nums[:i]
			if i != len(nums)-1 {
				temp = append(temp, nums[i+1:]...)
			}
			nums = temp
			//如果判断到了 就要-1 因为数组的结构变了
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}

//直接调用api太无聊了
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	index := strings.Index(haystack, needle)
	return index
}

func strStr1(haystack string, needle string) int {
	length := len(needle)
	lenstack := len(haystack)
	if length == 0 {
		return 0
	}
	if lenstack < length {
		return -1
	}
	for i := 0; i < lenstack-length+1; i++ {
		if (haystack[i : i+length]) == needle {
			return i
		}
	}
	return -1
}

//投机取巧了 看了源码
func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !func(i int) bool {
			return nums[i] >= target
		}(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func permute(nums []int) [][]int {
	final := make([][]int, 0)
	m := make(map[int]bool)

	var travel func(nums []int, ans []int)
	travel = func(nums []int, ans []int) {
		//res := make([]int, 0, len(nums))
		for _, num := range nums {
			if m[num] {
				continue
			}
			m[num] = true
			c := append(ans, num)
			if len(c) == len(nums) {
				m[num] = false
				final = append(final, c)
				continue
			}
			travel(nums, c)
			m[num] = false
		}
	}
	travel(nums, []int{})
	return final
}
