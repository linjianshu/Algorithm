package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	//permute([]int{1, 2, 3})

	//i := backtrack([]int{1, 2, 3})
	//fmt.Println(i)

	//lis := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	//fmt.Println(lis)

	//array := maxSubArray([]int{-3, 4, -1, 2, -6, 1, 4})
	//fmt.Println(array)

	//lcs := LCS("babcde", "acze")
	//fmt.Println(lcs)

	//node := &Node{
	//	pre:  nil,
	//	next: nil,
	//	key:  1,
	//	val:  1,
	//}
	//node2 := &Node{
	//	pre:  nil,
	//	next: nil,
	//	key:  2,
	//	val:  2,
	//}
	//node3 := &Node{
	//	pre:  nil,
	//	next: nil,
	//	key:  3,
	//	val:  3,
	//}
	//node4 := &Node{
	//	pre:  nil,
	//	next: nil,
	//	key:  1,
	//	val:  4,
	//}
	//lru := NewLRU(2)
	//lru.Set(node)
	//lru.Set(node2)
	//get := lru.Get(node.key)
	//fmt.Println(get)
	//lru.Set(node3)
	//lru.Get(node2.key)
	//lru.Set(node4)

	//element := nextGreaterElements([]int{2, 1, 2, 4, 3})
	//fmt.Println(element)

	//temperatures := dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70})
	//fmt.Println(temperatures)

	//window := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//fmt.Println(window)

	//node := ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val: 5,
	//					Next: &ListNode{
	//						Val:  6,
	//						Next: nil,
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	////palindrome := isPalindrome(&node)
	////fmt.Println(palindrome)
	//
	//n := reverseN(&node, 3)
	//fmt.Println(n)

	//l := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}
	//nodes := nextLargerNodes(l)
	//fmt.Println(nodes)

	//i := subsets([]int{1, 2, 3, 4, 5})
	//fmt.Println(i)

	//i := combine(1, 1)
	//fmt.Println(i)

	//depth := maxDepth("(1+(2*3)+((8)/4))+1")
	//fmt.Println(depth)

	//sum := subarraySum([]int{1}, 0)
	//fmt.Println(sum)

	//primes := countPrimes(10)
	//fmt.Println(primes)

	//pow := superPow(2, []int{2, 0})
	//fmt.Println(pow)

	//speed := minEatingSpeed([]int{1000000000}, 2)
	//fmt.Println(speed)

	//fmt.Println(len("376L3000297044521443622-01-051042400FE130BOSCH#"))
	//i := trap([]int{4, 2, 0, 3, 2, 5})
	//fmt.Println(i)

	//duplicates := removeDuplicates([]int{1, 1, 2})
	//fmt.Println(duplicates)

	//jump := canJump([]int{2, 0, 0})
	//fmt.Println(jump)

	i := beibao(3, 4, []int{2, 1, 3}, []int{4, 2, 3})
	fmt.Println(i)
}

//回溯问题
func permute(nums []int) {
	final := make([][]int, 0)

	//定义递归
	var travel func(nums []int, temp []int)
	travel = func(nums []int, temp []int) {
		//结束条件
		if len(nums) == len(temp) {
			res := make([]int, len(nums))
			copy(res, temp)
			final = append(final, res)
			return
		}

		//for循环是关键  做选择
		for _, v := range nums {
			//用来排除已经选择过的路径
			b := false
			for _, tem := range temp {
				if v == tem {
					b = true
				}
			}
			if b {
				continue
			} else {
				//添加到列表中
				temp = append(temp, v)
				travel(nums, temp)
				//撤销选择
				temp = temp[:len(temp)-1]
			}

		}
	}

	temp := make([]int, 0, len(nums))
	travel(nums, temp)
	fmt.Println(final)
}

func backtrack(nums []int) [][]int {
	final := make([][]int, 0)
	res := make([]int, 0)
	//备忘录
	add := make(map[int]bool, len(nums))

	var travel func(nums []int, res []int)
	travel = func(nums []int, res []int) {
		//结束条件
		if len(res) == len(nums) {
			final = append(final, res)
			return
		}

		//做选择
		for _, v := range nums {
			if !add[v] {
				//已经选过了
				add[v] = true

				res = append(res, v)
				//递归
				travel(nums, res)

				//撤销选择 并且恢复原样
				res = res[:len(res)-1]
				add[v] = false
			}

		}
	}
	travel(nums, res)
	return final
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]int, len(nums))
	m[nums[0]] = 1
	for i := 1; i < len(nums); i++ {
		m[nums[i]] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				m[nums[i]] = int(math.Max(float64(m[nums[i]]), float64(m[nums[j]]+1)))
			}
		}
	}

	res := 1
	for _, v := range m {
		if v > res {
			res = v
		}
	}
	return res
}

//使用动态规划 dp数组的概念来做
//假设dp数组的dp[i]是第i个结尾的子数组最大值 来归纳第n个
func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	final := make([]int, 0, len(nums))
	final = append(final, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+final[i-1] {
			final = append(final, nums[i]+final[i-1])
		} else {
			final = append(final, nums[i])
		}
	}

	res := final[0]
	for i := 1; i < len(final); i++ {
		if final[i] > res {
			res = final[i]
		}
	}
	return res
}

//状态压缩
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp_0 := nums[0]
	res := dp_0
	var dp_1 int

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i]+dp_0 {
			dp_1 = nums[i] + dp_0
		} else {
			dp_1 = nums[i]
		}

		if dp_1 > res {
			res = dp_1
		}

		dp_0 = dp_1
	}

	return res
}

func LCS(str1 string, str2 string) int {
	l1 := len(str1)
	l2 := len(str2)
	var m [][]int
	b := l1 > l2
	if !b {
		temp := str1
		str1 = str2
		str2 = temp
		l1 = len(str1)
		l2 = len(str2)
	}
	m = make([][]int, l2+1, l2+1)
	m[0] = make([]int, l1+1, l1+1)

	for i := 1; i < len(m); i++ {
		m[i] = make([]int, l1+1, l1+1)
		for j := 0; j < len(m[i]); j++ {
			if j == 0 {
				m[i][j] = 0
			} else if str2[i-1] == str1[j-1] {
				res := m[i][j-1] + 1
				if res <= i {
					m[i][j] = res
				} else {
					m[i][j] = i
				}
			} else {
				res := m[i][j-1]
				res1 := m[i-1][j]
				max := math.Max(float64(res), float64(res1))
				m[i][j] = int(max)
			}
		}
	}
	return m[l2][l1]
}

// Node 基础数据结构 每一个节点 包含了前置和后置指针 以及自身的key和value
type Node struct {
	pre      *Node
	next     *Node
	key, val int
}

// DoubleLinkList 一整条链表 包含链表的头节点和尾巴节点 头结点用于删除头部的元素 尾巴节点用于增加新的元素 size定义为链表能容纳的长度
type DoubleLinkList struct {
	head *Node
	tail *Node
	size int
}

func NewDoubleLinkList() *DoubleLinkList {
	head := &Node{}
	tail := &Node{}
	head.pre = nil
	head.next = tail
	tail.pre = head
	tail.next = nil

	return &DoubleLinkList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (d *DoubleLinkList) AddLast(node *Node) {
	d.tail.pre.next = node
	node.pre = d.tail.pre
	node.next = d.tail
	d.tail.pre = node
	d.size++
}

func (d *DoubleLinkList) DeleteCurrent() *Node {
	if d.size < 1 {
		return nil
	}
	deletedNode := d.head.next
	d.head.next.next.pre = d.head
	d.head.next = d.head.next.next
	d.size--
	return deletedNode
}

// DeleteNode 这一步很关键 这个我漏了 在上层实现了 不好 这一步是因为有pre前置节点才能够做到快速删除 时间复杂度o(1)
func (d *DoubleLinkList) DeleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	d.size--
}

func (d *DoubleLinkList) Size() int {
	return d.size
}

// LRU 数据结构 包含哈希表 用于快速找到节点 包含双链表 便于增加和删除和修改 包含容量 用于初始化最大的容量界限
type LRU struct {
	m   map[int]*Node
	d   *DoubleLinkList
	cap int
}

func NewLRU(cap int) *LRU {
	list := NewDoubleLinkList()
	m := make(map[int]*Node)
	return &LRU{
		m:   m,
		d:   list,
		cap: cap,
	}
}

// Set 难点在于 保持一致性 如果链表中删除了 最好在map表中也删除
func (l *LRU) Set(node *Node) {
	if n, ok := l.m[node.key]; !ok {
		if l.d.size < l.cap {
			l.d.AddLast(node)
			l.m[node.key] = node
		} else {
			current := l.d.DeleteCurrent()
			delete(l.m, current.val)
			l.d.AddLast(node)
			l.m[node.key] = node
		}
	} else {
		l.d.DeleteNode(n)
		//n.pre.next = n.next
		//n.next.pre = n.pre
		//l.d.size--
		l.d.AddLast(node)
	}
}

func (l *LRU) Get(key int) int {
	if n, ok := l.m[key]; ok {
		//这方法包含了下面3个
		l.d.DeleteNode(n)
		//n.pre.next = n.next
		//n.next.pre = n.pre
		//l.d.size--
		l.d.AddLast(n)
		return n.val
	} else {
		return -1
	}
}

//使用单调栈来做
func nextGreaterElement(nums1 []int) (ans []int) {
	ans = make([]int, len(nums1))
	stack := make([]int, 0, len(nums1))
	for i := len(nums1) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums1[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums1[i])
	}
	return
}

//使用单调栈来做 要等几天 只是记录一下索引之间的距离罢了
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1] - i
		} else {
			ans[i] = 0
		}
		stack = append(stack, i)
	}
	return ans
}

func nextGreaterElements(nums []int) []int {
	c := make([]int, len(nums))
	copy(c, nums)
	nums = append(nums, c...)
	ans := make([]int, len(nums))
	stack := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		} else {
			ans[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return ans[:len(ans)/2]
}

//暴力解法
func maxSlidingWindow1(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	temp := make([]int, k)
	left := 0
	right := k - 1
	for right != len(nums) {
		index := 0
		for i := left; i <= right; i++ {
			temp[index] = nums[i]
			index++
		}
		sort.Ints(temp)
		ans[left] = temp[len(temp)-1]
		left++
		right++
	}
	return ans
}

//利用单调队列
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	list := &DoubleLink1{list: list.New()}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			//累积k-1个队列数目
			list.push(nums[i])
		} else {
			//push要求塞进去的时候要排序
			list.push(nums[i])
			//max要求空间复杂度o(1) 按照算法技巧 最前面的就是最大值
			m := list.max()
			ans[i-k+1] = m
			//把即将到达边界的数 从队列中剔除
			list.pop(nums[i-k+1])
		}
	}
	return ans
}

type DoubleLink1 struct {
	list *list.List
}

func (d *DoubleLink1) push(num int) {
	//for循环 把比num小的数字从队列中剔除 毕竟求最大值轮不到他们
	for d.list.Len() > 0 && d.list.Back().Value.(int) < num {
		d.list.Remove(d.list.Back())
	}

	//把自己加入到队列中去
	d.list.PushBack(num)
}

func (d *DoubleLink1) max() int {
	return d.list.Front().Value.(int)
}

func (d *DoubleLink1) pop(num int) {
	//判断要pop掉的是不是已经不在队列中了 如果在的话肯定是在第一位 如果不在可能是由于push的时候被剔除掉了
	if d.list.Front().Value.(int) == num {
		d.list.Remove(d.list.Front())
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//超出时间限制了
func isPalindrome1(head *ListNode) bool {
	final := ""
	for head != nil {
		final += strconv.Itoa(head.Val)
		head = head.Next
	}
	left, right := 0, len(final)-1
	for left < right {
		if final[left] != final[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//尽量别用strconv.Itoa 很耗时间
func isPalindrome2(head *ListNode) bool {
	final := make([]int, 0)
	for head != nil {
		final = append(final, head.Val)
		head = head.Next
	}
	left, right := 0, len(final)-1
	for left < right {
		if final[left] != final[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Reserve 利用后序遍历  这个只能打到反转的效果 不行
func Reserve(head *ListNode) bool {
	l := new(ListNode)
	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return l
		}
		node := travel(head.Next)
		node.Next = head
		head.Next = nil
		return head
	}

	travel(head)
	return l == head
}

//利用后序遍历和双指针 依次比较链表
func isPalindrome3(head *ListNode) bool {
	left := head

	var travel func(head *ListNode) bool
	travel = func(head *ListNode) bool {
		if head == nil {
			return true
		}
		res := travel(head.Next)
		if res && head.Val == left.Val {
			left = left.Next
			return true
		} else {
			return false
		}
	}
	return travel(head)
}

func reverseN(head *ListNode, n int) *ListNode {
	var mid *ListNode
	var begin *ListNode
	if n == 0 {
		return head
	} else if n == 1 {
		begin = head.Next
		mid = head.Next.Next
	} else {
		temp := head
		for n > 1 {
			temp = temp.Next
			n--
		}
		begin = temp
		mid = begin.Next
	}

	temp := head
	next := temp.Next
	temp.Next = nil
	for next != mid {
		node := next.Next
		next.Next = temp
		temp = next
		next = node
	}

	head.Next = mid
	return temp
}

//单调栈解法
func nextLargerNodes(head *ListNode) []int {
	init := make([]int, 0)

	var travel func(head *ListNode)
	travel = func(head *ListNode) {
		if head == nil {
			return
		}
		travel(head.Next)
		init = append(init, head.Val)
	}

	travel(head)

	ans := make([]int, len(init))
	queue := make([]int, 0)

	for i := 0; i < len(init); i++ {
		for len(queue) > 0 && queue[len(queue)-1] <= init[i] {
			queue = queue[:len(queue)-1]
		}
		if len(queue) > 0 {
			ans[len(init)-i-1] = queue[len(queue)-1]
		} else {
			ans[len(init)-i-1] = 0
		}
		queue = append(queue, init[i])
	}
	return ans
}

//利用数学归纳法 递归来做
func subsets1(nums []int) [][]int {

	var travel func(nums []int) [][]int
	travel = func(nums []int) [][]int {
		if len(nums) == 0 {
			return [][]int{{}}
		}
		i := travel(nums[:len(nums)-1])
		for _, ints := range i {
			c := make([]int, len(ints), cap(ints))
			copy(c, ints)
			res := append(c, nums[len(nums)-1])
			i = append(i, res)
		}
		return i
	}
	return travel(nums)
}

//使用回溯来做 做选择 撤销选择 这里有点不一样 就是为了防止1,2,3 和 1,3,2这种 直接使用fori i+1这样就不用担心排序的问题了  不知道
func subsets(nums []int) [][]int {
	pow := math.Pow(2.0, float64(len(nums)))
	final := make([][]int, 0, int(pow))
	final = append(final, []int{})

	var travel func(nums []int, temp []int)
	travel = func(nums []int, temp []int) {
		if len(nums) == 0 {
			return
		}

		for i := 0; i < len(nums); i++ {
			c := make([]int, len(temp), cap(temp))
			copy(c, temp)
			res := append(c, nums[i])
			final = append(final, res)
			travel(nums[i+1:], res)
		}

	}
	travel(nums, []int{})
	return final
}

func combine(n int, k int) [][]int {
	final := make([][]int, 0)

	var travel func(j, k int, res []int)
	travel = func(j, k int, res []int) {
		if j > n {
			return
		}
		for i := j; i <= n; i++ {
			c := make([]int, len(res), cap(res))
			copy(c, res)
			ints := append(c, i)
			if len(ints) == k {
				final = append(final, ints)
				continue
			}
			travel(i+1, k, ints)
		}
	}
	travel(1, k, []int{})
	return final
}

func maxDepth(s string) int {
	res := make([]rune, 0)
	count := 0
	max := 0
	for _, i := range s {
		if i == '(' {
			res = append(res, i)
			count++
			if count > max {
				max = count
			}
		} else if i == ')' {
			if res[len(res)-1] == '(' {
				res = res[:len(res)-1]
				count--
			}
		}
	}

	return max
}

//前缀和思想
func subarraySum1(nums []int, k int) int {
	count := 0
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0

	for i := 0; i < len(nums); i++ {
		preSum[i+1] = nums[i] + preSum[i]
	}

	for i := 0; i < len(preSum); i++ {
		for j := i + 1; j < len(preSum); j++ {
			if preSum[j]-preSum[i] == k {
				count++
			}
		}
	}
	return count
}

//前缀和思想
func subarraySum(nums []int, k int) int {
	count := 0
	preSum := make(map[int]int, len(nums)+1)
	preSum[0] = 1
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = temp + nums[i]
		count += preSum[temp-k]
		preSum[temp] += 1
	}

	return count
}

//普通算法
func countPrimes1(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		b := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				b = false
				break
			}
		}
		if b {
			res++
		}
	}
	return res
}

//截止到sqrt
func countPrimes2(n int) int {
	res := 0
	for i := 2; i < n; i++ {
		b := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				b = false
				break
			}
		}
		if b {
			res++
		}
	}
	return res
}

//筛数法
func countPrimes(n int) int {
	isPrimes := make([]bool, n+1)

	for i := range isPrimes {
		isPrimes[i] = true
	}

	res := 0
	for i := 2; i < n; i++ {
		b := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				b = false
				break
			}
		}
		if b {
			for k := i * 2; k < len(isPrimes); k += i {
				isPrimes[k] = false
			}
		}
	}

	for i := 2; i < n; i++ {
		if isPrimes[i] {
			res++
		}
	}
	return res
}

//超级次方 递归 有点晦涩难懂
func superPow1(a int, b []int) int {
	var modPow func(m int, n int) int
	modPow = func(m int, n int) int {
		base := m % 1337
		res := 1
		for i := 0; i < n; i++ {
			res *= base
			res = res % 1337
		}
		return res
	}
	var travel func(a int, b []int) int
	travel = func(a int, b []int) int {
		if len(b) >= 1 {
			last := b[len(b)-1]
			part1 := modPow(a, last)
			b = b[:len(b)-1]
			part2 := modPow(travel(a, b), 10)
			return (part1 * part2) % 1337
		} else {
			return 1
		}
	}
	return travel(a, b)
}

func superPow(a int, b []int) int {
	var modPow func(m int, n int) int
	modPow = func(m int, n int) int {
		if n == 0 {
			return 1
		}
		base := m % 1337
		if n%2 == 1 {
			return (base * modPow(m, n-1)) % 1337
		} else {
			sub := modPow(m, n/2)
			return (sub * sub) % 1337
		}
	}
	var travel func(a int, b []int) int
	travel = func(a int, b []int) int {
		if len(b) >= 1 {
			last := b[len(b)-1]
			part1 := modPow(a, last)
			b = b[:len(b)-1]
			part2 := modPow(travel(a, b), 10)
			return (part1 * part2) % 1337
		} else {
			return 1
		}
	}
	return travel(a, b)
}

//for循环的时候 如果是线性的 那么可以考虑优化的点就在于 可以使用二分法从中间开始截断
func minEatingSpeed(piles []int, h int) int {
	c := make([]int, 0)
	c = append(c, piles...)
	sort.Ints(c)
	index := h - len(piles)

	min := 1
	max := c[len(c)-1]

	if len(c)-1-index >= 0 {
		min = len(c) - 1 - index
	}

	for min < max {
		mid := min + (max-min)/2
		res := 0
		for j := 0; j < len(piles); j++ {
			if piles[j]%mid == 0 {
				res += piles[j] / mid
			} else {
				res += piles[j]/mid + 1
			}
		}

		if res <= h {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return min
}

//接雨水问题 实质上看单个局部的雨水量 就是左边的最大值 和右边的最大值相减的结果
func trap1(height []int) int {
	res := 0
	var getLeftMax func(i int) int
	getLeftMax = func(i int) int {
		final := height[i]
		for j := i; j >= 0; j-- {
			if height[j] > final {
				final = height[j]
			}
		}
		return final
	}

	var getRightMax func(i int) int
	getRightMax = func(i int) int {
		final := height[i]
		for j := i; j < len(height); j++ {
			if height[j] > final {
				final = height[j]
			}
		}
		return final
	}

	var getMin func(left, right int) int
	getMin = func(left, right int) int {
		if left > right {
			return right
		} else {
			return left
		}
	}

	for i := 1; i < len(height); i++ {
		left := getLeftMax(i)
		right := getRightMax(i)
		if height[i] >= left || height[i] >= right {
			continue
		}
		min := getMin(left, right)
		res += min - height[i]
	}
	return res
}

func trap(height []int) int {
	res := 0

	m1 := make(map[int]int)
	m1[len(height)-1] = height[len(height)-1]

	m := make(map[int]int)
	m[0] = height[0]

	for i := 1; i < len(height); i++ {
		if m1[len(height)-i] > height[len(height)-1-i] {
			m1[len(height)-1-i] = m1[len(height)-i]
		} else {
			m1[len(height)-1-i] = height[len(height)-1-i]
		}

		if m[i-1] > height[i] {
			m[i] = m[i-1]
		} else {
			m[i] = height[i]
		}
	}

	var getMin func(left, right int) int
	getMin = func(left, right int) int {
		if left > right {
			return right
		} else {
			return left
		}
	}

	for i := 1; i < len(height); i++ {
		left := m[i]
		right := m1[i]
		if height[i] >= left || height[i] >= right {
			continue
		}
		min := getMin(left, right)
		res += min - height[i]
	}
	return res
}

//双指针 快指针在前 满指针在后 前后不等的时候 满指针++ 快指针++ 快慢指针值对调 最后返回slow的索引+1
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			temp := nums[fast]
			//fast压根不需要替换
			//nums[fast] = nums[slow+1]
			nums[slow+1] = temp
			slow++
		}
		fast++
	}
	return slow + 1
}

//贪心思想 没看解析搞不懂
func canJump(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}
		temp := i + nums[i]
		if farthest < temp {
			farthest = temp
		}
	}
	if farthest >= len(nums)-1 {
		return true
	} else {
		return false
	}
}

//0-1背包问题 二维数组动态规划  dp[i][j] 定义为 当前i个包含i个物品时,背包重量为j时,最大能够装下的价值是多少
func beibao(n int, w int, wt []int, val []int) int {
	dp := make([][]int, n+1, n+1)
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1, w+1)
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			//j 代表重量  如果当前的重量超过背包的重量了 那就只能不装了
			if j-wt[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//否则 判断装还是不装  之所以不装反而会更大 可能是因为 这里有两个变量叭 不装的话我就去找前一个相同重量的 装的话我就得去找前一个中重量-当前这个重量的那个值+当前的物品的价值
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			}
		}
	}
	return dp[n][w]
}
