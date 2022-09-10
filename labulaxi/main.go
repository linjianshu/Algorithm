package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(len("376L3000403044521443622-01-051042400FE130BOSCH#"))
	//i := jump([]int{2, 1})
	//fmt.Println(i)

	//product := maxProduct([]string{"a", "aa", "aaa", "aaaa"})
	//fmt.Println(product)

	//sum := threeSum([]int{-1, 0, 1, 2, -1, -4})
	//fmt.Println(sum)

	//fmt.Println(math.MaxInt32)
	//i := divide(7, 3)
	//fmt.Println(i)

	//binary := addBinary("101111", "10")
	//fmt.Println(binary)

	//bits := countBits(8)
	//fmt.Println(bits)

	//arrayLen := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	//fmt.Println(arrayLen)

	//k := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	//fmt.Println(k)

	//length := findMaxLength([]int{0, 1, 1})
	//fmt.Println(length)

	//index := pivotIndex([]int{1, 2, 3})
	//index := pivotIndex([]int{1, 7, 3, 6, 5, 6})
	//fmt.Println(index)

	//characters := numberOfWeakCharacters([][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}})
	//fmt.Println(characters)

	//inclusion := checkInclusion("adc", "dcda")
	//fmt.Println(inclusion)

	//anagrams := findAnagrams("abab", "ab")
	//fmt.Println(anagrams)

	//substring := lengthOfLongestSubstring("hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//fmt.Println(substring)

	//node := ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//}

	//end := removeNthFromEnd(&node, 2)
	//fmt.Println(end)

	//fmt.Printf("%d", '[')
	//palindrome := isPalindrome("Marge, let's \"[went].\" I await {news} telegram.")
	//fmt.Println(palindrome)

	//palindrome := validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga")
	//fmt.Println(palindrome)

	//substrings := countSubstrings("aaa")
	//fmt.Println(substrings)

	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}

	l.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	//cycle := detectCycle(l)
	//fmt.Println(cycle)

	//node := getIntersectionNode(l5, l2)
	//fmt.Println(node)

	//list := reverseList(l)
	//fmt.Println(list)

	//l6 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//			},
	//		},
	//	},
	//}

	//l7 := &ListNode{
	//	Val: 0,
	//}
	//numbers := addTwoNumbers(l6, l7)
	//fmt.Println(numbers)

	//reorderList(l6)

	//n1 := &Node{
	//	Val:   1,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n2 := &Node{
	//	Val:   2,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n3 := &Node{
	//	Val:   3,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n4 := &Node{
	//	Val:   4,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n5 := &Node{
	//	Val:   5,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n6 := &Node{
	//	Val:   6,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n7 := &Node{
	//	Val:   7,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n8 := &Node{
	//	Val:   8,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	////n9 := &Node{
	////	Val:   9,
	////	Prev:  nil,
	////	Next:  nil,
	////	Child: nil,
	////}
	////n10 := &Node{
	////	Val:   10,
	////	Prev:  nil,
	////	Next:  nil,
	////	Child: nil,
	////}
	//n11 := &Node{
	//	Val:   11,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n12 := &Node{
	//	Val:   12,
	//	Prev:  nil,
	//	Next:  nil,
	//	Child: nil,
	//}
	//n1.Next = n2
	//n2.Prev = n1
	//n2.Next = n3
	//n3.Prev = n2
	//n3.Next = n4
	//n3.Child = n7
	//n7.Next = n8
	//n8.Prev = n7
	////n8.Next = n9
	//n4.Prev = n3
	//n4.Next = n5
	//n5.Prev = n4
	//n5.Next = n6
	//n6.Prev = n5
	////n8.Prev = n7
	////n8.Next = n9
	////n9.Prev = n8
	////n9.Next = n10
	////n10.Prev = n9
	//n8.Child = n11
	//n11.Next = n12
	//n12.Prev = n11
	//
	////n1.Next = nil
	////n2.Next = nil
	////n3.Next = nil
	////n1.Child = n2
	////n2.Child = n3
	//node := flatten(n1)
	//fmt.Println(node)

	ll := &ListNode{
		Val:  5,
		Next: nil,
	}
	ll2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	ll3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	ll.Next = ll2
	ll2.Next = ll3
	ll3.Next = ll
	//ll := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	//ll := &ListNode{
	//	Val:  1,
	//	Next: nil,
	//}
	node := insert(ll3, 6)
	fmt.Println(node)
}

//跳跃游戏Ⅱ 暴力遍历不可以 用map稍微好一点点 这里可以看出 在哪里判断map里存不存在很重要
//之前在for里判断 那么就得多执行一个for循环的线性遍历过程 现在在for外判断
func jump1(nums []int) int {
	m := make(map[int]int)

	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	var dp func(index int) int
	dp = func(index int) int {
		min := len(nums)

		if index >= len(nums)-1 {
			return 0
		}

		if nums[index] == 0 {
			return math.MaxInt32
		}

		if v, ok := m[index]; ok {
			return v
		}
		temp := 0
		for j := index; j <= index+nums[index]; j++ {
			if j == index {
				continue
			}
			temp = dp(j) + 1
			if min > temp {
				min = temp
			}
		}
		m[index] = min
		return min
	}

	return dp(0)
}

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	farthest := 0
	res := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {
		temp := i + nums[i]

		if temp > farthest {
			farthest = temp
		}
		if end == i {
			end = farthest
			res++
		}
	}
	return res
}

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for k := range m {
		if m[k] != 3 {
			return k
		}
	}
	return 0
}

func maxProduct(words []string) int {
	res := 0

	for i, word := range words {
		for j := i + 1; j < len(words); j++ {
			if !strings.ContainsAny(words[j], word) {
				temp := len(word) * len(words[j])
				if res < temp {
					res = temp
				}
			}
		}
	}
	return res
}

//数组 双指针
func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return [][]int{}
	}

	m := make(map[int]int, 6)
	m1 := make(map[[3]int]bool)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	res := make([][]int, 0)

	left, right := 0, 1
	for left <= len(nums)-3 {
		right = left + 1
		for right <= len(nums)-1 {
			i := m[0-nums[left]-nums[right]]
			if i <= right {
				right++
				continue
			}
			temp := []int{nums[left], nums[right], nums[i]}
			sort.Ints(temp)
			if _, ok := m1[[3]int{temp[0], temp[1], temp[2]}]; !ok {
				m1[[3]int{temp[0], temp[1], temp[2]}] = true
				res = append(res, temp)
			}
			right++
		}
		left++
	}
	return res
}

func divide(a int, b int) int {
	flag := 0
	if a < 0 && b < 0 {
		flag = 0
	} else if a < 0 || b < 0 {
		flag = 1
	}

	res := 0
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	oldb := b
	temp := 1

	for a >= b {
		//a>b 那么b翻倍 翻倍的倍数加到结果里
		for a >= b && b<<1 < a {
			b <<= 1
			temp <<= 1
		}
		res += temp
		//重置翻倍倍数
		temp = 1
		//扣掉被减掉的数字
		a -= b
		//重置被除数
		b = oldb
	}

	if flag == 1 {
		res = -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

//双指针 数组要考虑用双指针呀!!!! 我还在想什么位运算呢我搞什么飞机 strconv.itoa strconv.atoi 这类api耗时很长
//用char 和 char相减 也能得到 uint8
func addBinary(a string, b string) string {
	var flag uint8
	aPoint := len(a) - 1
	bPoint := len(b) - 1
	final := ""
	for aPoint >= 0 || bPoint >= 0 {

		var res uint8
		if aPoint >= 0 && bPoint >= 0 {
			res = a[aPoint] - '0' + b[bPoint] - '0' + flag
		} else if aPoint >= 0 {
			res = a[aPoint] - '0' + flag
		} else {
			res = b[bPoint] - '0' + flag
		}

		if res >= 2 {
			res -= 2
			final += strconv.Itoa(int(res))
			flag = 1
		} else {
			final += strconv.Itoa(int(res))
			flag = 0
		}
		aPoint--
		bPoint--
	}
	if flag == 1 {
		final += "1"
	}

	ret := ""
	for i := len(final) - 1; i >= 0; i-- {
		ret += string(final[i])
	}
	return ret
}

func countBits1(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = strings.Count(fmt.Sprintf("%b", i), "1")
	}
	return res
}

//b-k算法 一个数x的1的个数 是 x & x-1 直至结果为0的次数
func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 {
			res[i] = i
			continue
		} else {
			temp := 0
			j := i
			for j != 0 {
				j = j & (j - 1)
				temp++
			}
			res[i] = temp
		}
	}
	return res
}

//如果是2的幂 那么x&x(-1) 肯定=0
func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	high := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			res[i] = 1
			high = i
		} else {
			res[i] = res[i-high] + 1
		}
	}
	return res
}

//用了前缀和 注意 ： 线性for运算如果时间超时 但是如果线性分布的可以考虑二分法
func minSubArrayLen1(target int, nums []int) int {
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		preSum[i+1] = sum
	}
	min := math.MaxInt
	left, right := 0, 0
	for left <= len(preSum)-1 && right <= len(preSum)-1 {
		if preSum[right]-preSum[left] >= target {
			dis := right - left
			if dis < min {
				min = dis
			}
			left++
		} else if preSum[right]-preSum[left] < target {
			right++
		}
	}
	if min == math.MaxInt {
		return 0
	} else {
		return min
	}
}

//滑动窗口 注意窗口的滑动和窗口值sum的改变 左移就要减掉 右移就要加上 移动的时候注意别索引超了 模版不大对劲 要使用两个for比较好使一点
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	left, right := 0, 0
	sum := nums[0]
	for left <= right {
		temp := 0

		if sum < target {
			right++
			if right > len(nums)-1 {
				break
			}
			sum += nums[right]

		} else {

			temp = right - left + 1
			if temp < res {
				res = temp
			}
			sum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

//简单双指针
func numSubarrayProductLessThanK1(nums []int, k int) int {
	res := 0
	left := 0
	for left < len(nums) {
		right := left
		product := 1
		for right < len(nums) {
			product *= nums[right]
			if product < k {
				res++
			} else {
				break
			}
			right++
		}
		left++
	}
	return res
}

//滑动窗口 外层优先选用fori 内层用for吧
func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	left := 0
	product := 1
	for i := left; i < len(nums); i++ {
		product *= nums[i]

		for left <= i && product >= k {
			product /= nums[left]
			left++
		}

		res += i - left + 1

	}
	return res
}

//这样就超时了 呜呜 遇到数组的问题 要考虑尝试使用 双指针 动态规划 滑动窗口 二分法 前缀和思想知道了么
func findMaxLength1(nums []int) int {
	res := 0
	left := 0
	m := make(map[int]int, 2)
	for i := left; i < len(nums); i++ {
		right := i
		for right < len(nums) {
			m[nums[right]]++
			if m[0] == m[1] && m[0]*2 > res {
				res = m[0] * 2
			}
			right++
		}
		delete(m, 0)
		delete(m, 1)
	}
	return res
}

//实在搞不明白了 太牛了鸭
func findMaxLength(nums []int) int {
	res := 0
	m := make(map[int]int)
	m[0] = -1
	count := 0
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if v, ok := m[count]; ok {
			res = max(res, i-v)
		} else {
			m[count] = i
		}
	}
	return res
}

//前缀和思想
func pivotIndex(nums []int) int {
	l := len(nums) + 1
	preSum := make([]int, l)

	preSum[0] = 0
	for i := 0; i < l-1; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	for i := 0; i < l-1; i++ {
		if preSum[i]-preSum[0] == preSum[l-1]-preSum[i+1] {
			return i
		}
	}
	return -1
}

//超时
func numberOfWeakCharacters1(properties [][]int) int {
	res := 0
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] < properties[j][0]
		} else {
			return properties[i][1] > properties[j][1]
		}
	})
	for i := len(properties) - 1; i >= 0; i-- {
		j := 0
		for j <= i {
			if properties[i][0] > properties[j][0] && properties[i][1] > properties[j][1] {
				properties = append(properties[0:j], properties[j+1:]...)
				i--
				res++
				continue
			}
			j++
		}
	}
	return res
}

//超时
func numberOfWeakCharacters2(properties [][]int) int {
	res := 0
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] < properties[j][0]
		} else {
			return properties[i][1] > properties[j][1]
		}
	})
	m := make(map[[2]int]bool)
	mExist := make(map[[2]int]int)
	for _, property := range properties {
		mExist[[2]int{property[0], property[1]}]++
	}
	for i := len(properties) - 1; i >= 0; i-- {
		j := 0
		for j <= i {
			if properties[i][0] > properties[j][0] && properties[i][1] > properties[j][1] {
				if _, ok := m[[2]int{properties[j][0], properties[j][1]}]; !ok {
					res += mExist[[2]int{properties[j][0], properties[j][1]}]
					m[[2]int{properties[j][0], properties[j][1]}] = true
				}
			}
			j++
		}
	}
	return res
}

//利用排序的妙处在于 攻击力从小到大 防御力从大到小 这样遍历的时候已经默认攻击力已经排好序了 那么只要比较防御力的大小就可以 防御力大的把defense及时更新就好
func numberOfWeakCharacters(properties [][]int) int {
	res := 0
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] < properties[j][0]
		} else {
			return properties[i][1] > properties[j][1]
		}
	})

	defense := 0
	for i := len(properties) - 1; i >= 0; i-- {
		if defense <= properties[i][1] {
			defense = properties[i][1]
		} else {
			res++
		}
	}
	return res
}

type NumMatrix struct {
	lineSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	lineSum := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		preSum := make([]int, len(matrix[0])+1)
		preSum[0] = 0
		for j := 0; j < len(matrix[0]); j++ {
			preSum[j+1] = preSum[j] + matrix[i][j]
		}
		lineSum[i] = preSum
	}

	return NumMatrix{lineSum: lineSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.lineSum[i][col2+1] - this.lineSum[i][col1]
	}
	return sum
}

//用深比较是不是有点骚
func checkInclusion1(s1 string, s2 string) bool {
	m := make(map[uint8]int)
	l := len(s1)
	for i := 0; i < l; i++ {
		m[s1[i]]++
	}

	for i := 0; i < len(s2)-1-l+1; i++ {
		m1 := make(map[uint8]int)
		for j := i; j < i+l; j++ {
			m1[s2[j]]++
		}
		if reflect.DeepEqual(m1, m) {
			return true
		}
	}
	return false
}

//切片没有==运算符 但是数组有==运算符  滑动窗口 要巧用字符串中每个char和'a'的关系
//滑动窗口划走的时候 要把之前的恢复
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	cnt1 := [26]int{}
	cnt2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}

	if cnt1 == cnt2 {
		return true
	}

	s1Lth := len(s1)
	for i := len(s1); i < len(s2); i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-s1Lth]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}

	return false
}

//巧用数组的可比较性质 巧用int数组 字符串中的每个char-'a'之后得到的数值
//滑动窗口滑起来之后 记得把之前的恢复一下
func findAnagrams(s string, p string) []int {
	cntS, cntP := [26]int{}, [26]int{}
	if len(s) < len(p) {
		return []int{}
	}

	for i := 0; i < len(p); i++ {
		cntP[p[i]-'a']++
		cntS[s[i]-'a']++
	}

	res := make([]int, 0)

	if cntP == cntS {
		res = append(res, 0)
	}

	PLen := len(p)
	for i := PLen; i < len(s); i++ {
		cntS[s[i-PLen]-'a']--
		cntS[s[i]-'a']++
		if cntS == cntP {
			res = append(res, i-PLen+1)
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	cnt := [1000]int{}

	left, right := 0, 0
	for left < len(s) {
		for right < len(s) && cnt[s[right]-'a'] == 0 {
			cnt[s[right]-'a']++
			temp := right - left + 1
			if temp > res {
				res = temp
			}
			right++
		}
		cnt[s[left]-'a']--
		left++
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针涉及链表这块最好有个头结点  nil 指向原始的链表 好操作呜呜呜
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head == nil {
		return nil
	}
	l := &ListNode{}
	l.Next = head
	slow, fast := l, l
	count := 0
	for count != n {
		fast = fast.Next
		count++
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return l.Next
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	temp := make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] <= 'A'+25) || (s[i] >= 'a' && s[i] <= 'a'+25) || (s[i] >= '0' && s[i] <= '9') {
			temp = append(temp, s[i])
		}
	}

	l := len(temp)
	s2 := string(temp)

	fmt.Println(s2)
	for i := 0; i < len(temp)/2; i++ {
		if temp[i] != temp[l-1-i] {
			return false
		}
	}
	return true
}

//还没做完  用广度优先算法来做
func highestPeak(isWater [][]int) [][]int {
	queue := make([][2]int, 0)
	for i := 0; i < len(isWater); i++ {
		for j := 0; j < len(isWater[i]); j++ {
			if isWater[i][j] == 1 {
				queue = append(queue, [2]int{})
			}
		}

	}
	return [][]int{}
}

func validPalindrome(s string) bool {
	b := false

	left, right := 0, len(s)-1
	var travel func(left, right int) bool

	travel = func(left, right int) bool {
		for left <= right {
			if s[left] != s[right] {
				return false
			} else {
				left++
				right--
			}
		}
		return true
	}

	for left <= right {
		if s[left] != s[right] && !b {
			return travel(left+1, right) || travel(left, right-1)
		} else {
			left++
			right--
		}
	}
	return true
}

//双指针
func countSubstrings1(s string) int {
	res := 0
	left, right := 0, 0
	for left <= len(s)-1 {

		for right <= len(s)-1 {
			b := true
			for i := 0; i <= (right-left)/2; i++ {
				if s[left+i] != s[right-i] {
					b = false
					break
				}
			}
			if b {
				res++
			}
			right++
		}
		left++
		right = left

	}

	return res
}

//中心拓展
func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			res++
		}

		left, right = i, i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
			res++
		}
	}

	return res
}

//快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head.Next, head.Next.Next
	for slow != fast && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func detectCycle1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	m := make(map[*ListNode]bool, 0)
	temp := head
	for temp != nil && temp.Next != nil && !m[temp] {
		m[temp] = true
		temp = temp.Next
	}

	if temp == nil || temp.Next == nil {
		return nil
	}

	return temp
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	l1 := headA
	for l1 != nil {
		m[l1] = true
		l1 = l1.Next
	}

	l2 := headB
	for l2 != nil && !m[l2] {
		l2 = l2.Next
	}
	return l2
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			return nil
		} else if l1 == nil {
			l1 = headB
		} else if l2 == nil {
			l2 = headA
		}
	}
	return l1
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	cur := head
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	var res *ListNode
	var travel func(list *ListNode) *ListNode
	travel = func(list *ListNode) *ListNode {
		if list.Next == nil {
			res = list
			return list
		}
		node := travel(list.Next)
		list.Next = nil
		node.Next = list
		return list
	}
	travel(head)

	return res
}

//递归
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	var reverse func(l1 *ListNode) *ListNode
	reverse = func(l1 *ListNode) *ListNode {
		cur := l1
		var pre *ListNode
		for cur != nil {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}
		return pre
	}
	listNode := reverse(l1)
	listNode2 := reverse(l2)

	isCarry := 0

	var travel func(ll1 *ListNode, ll2 *ListNode) *ListNode
	travel = func(ll1 *ListNode, ll2 *ListNode) *ListNode {
		if ll1 == nil && ll2 == nil {
			return nil
		} else if ll1 == nil {
			sum := ll2.Val + isCarry
			if sum >= 10 {
				sum -= 10
				isCarry = 1
			} else {
				isCarry = 0
			}
			l := &ListNode{
				Val:  sum,
				Next: nil,
			}
			ll2 = ll2.Next
			l.Next = travel(ll1, ll2)
			return l
		} else if ll2 == nil {
			sum := ll1.Val + isCarry
			if sum >= 10 {
				sum -= 10
				isCarry = 1
			} else {
				isCarry = 0
			}
			l := &ListNode{
				Val:  sum,
				Next: nil,
			}
			ll1 = ll1.Next
			l.Next = travel(ll1, ll2)
			return l
		} else {
			sum := ll1.Val + ll2.Val + isCarry
			if sum >= 10 {
				sum -= 10
				isCarry = 1
			} else {
				isCarry = 0
			}
			ll1 = ll1.Next
			ll2 = ll2.Next
			l := &ListNode{
				Val:  sum,
				Next: nil,
			}
			l.Next = travel(ll1, ll2)
			return l
		}
	}
	node := travel(listNode, listNode2)
	res := reverse(node)
	if isCarry == 1 {
		res = &ListNode{Val: 1, Next: res}
	}
	return res

}

//用栈的特性来做对齐
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	queue1 := make([]int, 0)
	queue2 := make([]int, 0)
	for l1 != nil {
		queue1 = append(queue1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		queue2 = append(queue2, l2.Val)
		l2 = l2.Next
	}

	var pre *ListNode
	var cur *ListNode

	isCarry := 0
	for len(queue1) != 0 || len(queue2) != 0 {
		q1, q2 := 0, 0
		if len(queue1) != 0 {
			q1 = queue1[len(queue1)-1]
			queue1 = queue1[:len(queue1)-1]
		}

		if len(queue2) != 0 {
			q2 = queue2[len(queue2)-1]
			queue2 = queue2[:len(queue2)-1]
		}

		sum := q1 + q2 + isCarry
		if sum >= 10 {
			sum -= 10
			isCarry = 1
		} else {
			isCarry = 0
		}
		cur = &ListNode{
			Val:  sum,
			Next: nil,
		}
		cur.Next = pre
		pre = cur
	}

	if isCarry == 1 {
		return &ListNode{
			Val:  1,
			Next: cur,
		}
	}
	return cur
}

func reorderList(head *ListNode) {
	l := make([]*ListNode, 0)
	for head != nil {
		l = append(l, head)
		head = head.Next
	}

	for i := 0; i < len(l)/2; i++ {
		l[i].Next = l[len(l)-1-i]

		if l[len(l)-1-i] != l[i+1] {
			l[len(l)-1-i].Next = l[i+1]
		} else {
			l[len(l)-1-i].Next = nil
		}
	}
	if len(l)%2 == 1 {
		l[len(l)/2].Next = nil
	}
}

//回文联表
func isPalindrome1(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}
	m := make([]int, 0)
	for head != nil {
		m = append(m, head.Val)
		head = head.Next
	}

	if len(m) > 0 {
		for i := 0; i < len(m)/2; i++ {
			if m[i] != m[len(m)-1-i] {
				return false
			}
		}
	}
	return true
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	cur := root

	var travel func()
	travel = func() {
		for cur.Child != nil {
			if cur.Next != nil {
				queue = append(queue, cur.Next)
			}
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
			cur = cur.Next
		}

		for cur.Next != nil || cur.Child != nil {
			if cur.Child == nil {
				cur = cur.Next
			} else {
				if cur.Next != nil {
					queue = append(queue, cur.Next)
				}
				cur.Next = cur.Child
				cur.Child.Prev = cur
				cur.Child = nil
				cur = cur.Next
			}
		}
	}

	travel()

	for len(queue) != 0 {
		q := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		q.Prev = cur
		cur.Next = q
		cur = q
		travel()
	}
	return root
}

//自定义排序 + 循环的时候怎么取下标  通常使用i%len(m)来循环取下标
func insert(aNode *ListNode, x int) *ListNode {
	if aNode == nil {
		l := &ListNode{
			Val:  x,
			Next: nil,
		}
		l.Next = l
		return l
	} else if aNode.Next == aNode {
		aNode.Next = &ListNode{
			Val:  x,
			Next: aNode,
		}
		return aNode
	}
	var pre *ListNode
	cur := aNode
	m := make([]*ListNode, 0)

	m = append(m, cur)
	cur = cur.Next

	for cur != aNode {
		m = append(m, cur)
		cur = cur.Next
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].Val < m[j].Val
	})

	for i := 0; i < len(m); i++ {
		if m[i].Val >= x {
			pre = m[(i%len(m)+len(m)-1)%len(m)]
			cur = m[i]
			break
		}
	}

	if pre == nil {
		pre = m[len(m)-1]
		cur = m[0]
	}
	l := &ListNode{
		Val:  x,
		Next: cur,
	}

	pre.Next = l

	return aNode
}
