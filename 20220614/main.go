package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	//t := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//
	//t2 := &TreeNode{
	//	Val: -10,
	//	Left: &TreeNode{
	//		Val:   9,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val:   15,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	//t3 := &TreeNode{
	//	Val:  -3,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	//sum := maxPathSum(t2)

	//fmt.Println(sum)

	//tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//tree := buildTree([]int{1, 2, 3}, []int{3, 2, 1})
	//fmt.Println(tree)

	//change := coinChange([]int{1, 2, 5}, 11)
	//change := coinChange([]int{2}, 3)
	//change := coinChange([]int{1}, 0)
	//change := coinChange([]int{186, 419, 83, 408}, 62491)
	//change := coinChange([]int{2}, 1)
	//fmt.Println(change)

	//i := fib(5)
	//i := fib(45)
	//fmt.Println(i)

	//i := permute([]int{1, 2, 3})
	//fmt.Println(i)

	//bst := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	//bst := sortedArrayToBST([]int{9, 12, 14, 17, 19, 23, 50, 54, 67, 72, 76})
	//fmt.Println(bst)

	//depth := minDepth(&TreeNode{
	//	Val:  2,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:  4,
	//			Left: nil,
	//			Right: &TreeNode{
	//				Val:  5,
	//				Left: nil,
	//				Right: &TreeNode{
	//					Val:   6,
	//					Left:  nil,
	//					Right: nil,
	//				},
	//			},
	//		},
	//	},
	//})

	//fmt.Println(depth)

	//lock := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	//lock := openLock([]string{"8888"}, "0009")
	//lock := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	//fmt.Println(lock)

	//a := &ListNode{
	//	Val:  3,
	//	Next: nil,
	//}
	//b := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//c := &ListNode{
	//	Val:  0,
	//	Next: nil,
	//}
	//d := &ListNode{
	//	Val:  -4,
	//	Next: nil,
	//}
	//a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = b
	//cycle := detectCycle(a)
	//fmt.Println(cycle)

	//i := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	//i := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	//fmt.Println(i)

	//sum := twoSum([]int{3, 2, 4}, 6)
	//fmt.Println(sum)

	//inclusion := checkInclusion2("ab", "eidbaooo")
	//inclusion := checkInclusion2("ab", "eidboaoo")
	//inclusion := checkInclusion2("ab", "a")
	//fmt.Println(inclusion)

	//substring := lengthOfLongestSubstring("abcabcbb")
	//substring := lengthOfLongestSubstring("bbbbb")
	//substring := lengthOfLongestSubstring("pwwkew")
	//substring := lengthOfLongestSubstring(" ")
	//fmt.Println(substring)

	//window := minWindow("ADOBECODEBANC", "ABC")
	//window := minWindow("a", "a")
	//window := minWindow("a", "aa")
	//fmt.Println(window)

	//lis := 	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	//lis := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	//lis := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	//fmt.Println(lis)

	//envelopes := maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}})
	//fmt.Println(envelopes)

	//xiezhebianli([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	//daozhebianli([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})

	//subsequence := longestCommonSubsequence("abcde", "ace")
	//subsequence := longestCommonSubsequence("abc", "abc")
	//subsequence := longestCommonSubsequence("abc", "def")
	//fmt.Println(subsequence)

	//subseq := longestPalindromeSubseq("bbbab")
	//subseq := longestPalindromeSubseq("cbbd")
	//fmt.Println(subseq)

	//i := zeroonebeibao(3, 4, []int{2, 1, 3}, []int{4, 2, 3})
	//fmt.Println(i)

	//partition := canPartition([]int{1, 5, 11, 5})
	//fmt.Println(partition)

	//i := change(5, []int{1, 2, 5})
	//fmt.Println(i)

	//i := rob([]int{2, 7, 9, 3, 1})
	//fmt.Println(i)

	//i := rob2([]int{2, 3, 2})
	//i := rob2([]int{0})
	//fmt.Println(i)

	//i := rob3(&TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:  2,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//})
	//fmt.Println(i)

	//ways := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	//ways := findTargetSumWays2([]int{1, 1, 1, 1, 1}, 3)
	//ways := findTargetSumWays2([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1)
	//ways := findTargetSumWays1([]int{1}, 2)
	//ways := findTargetSumWays2([]int{1}, 2)
	//fmt.Println(ways)

	//bst := isValidBST(&TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//})
	//
	//fmt.Println(bst)

	//node := deleteNode(&TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   2,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  6,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}, 3)
	//fmt.Println(node)

	//nodes1 := countNodes1(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{Val: 5},
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//})

	//fmt.Println(nodes1)

	//constructor := Constructor()
	//serialize := constructor.serialize5(&TreeNode{
	//	Val:   2,
	//	Left:  &TreeNode{Val: 1},
	//	Right: &TreeNode{Val: 3},
	//})
	//deserialize := constructor.deserialize5(serialize)
	//fmt.Println(deserialize)

	//t := &TreeNode{
	//	Val:   4,
	//	Left:  nil,
	//	Right: nil,
	//}
	//
	//left := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val:   6,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: t,
	//	},
	//}
	//
	//right := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   8,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//ancestor := lowestCommonAncestor(&TreeNode{
	//	Val:   3,
	//	Left:  left,
	//	Right: right}, left, right)
	//fmt.Println(ancestor)

	//element := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	//element := nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4})
	//fmt.Println(element)

	//elements := nextGreaterElements([]int{1, 2, 1})
	//fmt.Println(elements)

	//temperatures := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	//fmt.Println(temperatures)

	//window := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//window := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	//window := maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4)
	//fmt.Println(/window)

	//substrings := countSubstrings("abc")
	//substrings := countSubstrings("aaa")
	//fmt.Println(substrings)

	//palindrome := longestPalindrome("abccccdd")
	//palindrome := longestPalindrome("ccc")
	//fmt.Println(palindrome)

	//palindrome := isPalindrome(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//})
	//
	//fmt.Println(palindrome)

	//between := reverseBetween(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}, 2, 4)
	//between := reverseBetween(&ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val:  5,
	//		Next: nil,
	//	},
	//}, 1, 2)
	//fmt.Println(between)

	//group := reverseKGroup(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 3,
	//			Next: &ListNode{
	//				Val: 4,
	//				Next: &ListNode{
	//					Val:  5,
	//					Next: nil,
	//				},
	//			},
	//		},
	//	},
	//}, 3)
	//fmt.Println(group)

	//i := subsets([]int{1, 2, 3})
	//i := subsets([]int{9, 0, 3, 5, 7})
	//fmt.Println(i)

	//i := combine(4, 2)
	//fmt.Println(i)

	//sum4 := combinationSum4([]int{1, 2, 3}, 4)
	//fmt.Println(sum4)

	//i := permute([]int{1, 2, 3})
	//fmt.Println(i)

	//parenthesis := generateParenthesis(3)
	//fmt.Println(parenthesis)

	//area := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	//area := largestRectangleArea([]int{2, 4})
	//area := largestRectangleArea([]int{0, 0})
	//area := largestRectangleArea([]int{2, 3})
	//fmt.Println(area)

	line := 0
	fmt.Scanf("%d\n", &line)
	m, n := 0, 0
	fmt.Scanf("%d %d\n", &m, &n)
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//count, _ := strconv.Atoi(lines[0])
		if len(lines) == line {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	matrix := make([][]int, line, line)
	for i := 0; i < len(lines); i++ {
		matrix[i] = make([]int, 0, 0)
		for _, s := range strings.Fields(lines[i]) {
			atoi, _ := strconv.Atoi(s)
			matrix[i] = append(matrix[i], atoi)
		}
	}

	fmt.Println(matrix)
	i, i2 := b(matrix, m, n)
	fmt.Println(i)
	for _, ints := range i2 {
		for _, v := range ints {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}

func b(matrix [][]int, m, n int) (int, [][]int) {
	res1, res2 := math.MaxInt, [][]int{}

	var travel func(i, temp int, road []int)
	travel = func(i, temp int, road []int) {
		if i == n {
			if temp < res1 {
				res1 = temp
				ans := make([]int, len(road))
				copy(ans, road)
				res2 = [][]int{ans}
			} else if temp == res1 {
				ans := make([]int, len(road))
				copy(ans, road)
				res2 = append(res2, ans)
			}
			return
		}
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > 0 {
				travel(j, temp+matrix[i][j], append(road, j))
			}
		}
	}

	travel(m, 0, []int{m})
	return res1, res2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//124  递归
func maxPathSum(root *TreeNode) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	ans := math.MinInt

	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(travel(root.Left), 0)
		right := max(travel(root.Right), 0)
		//判断全局变量和节点的左右中哪个大 更新一下res
		ans = max(ans, left+right+root.Val)

		//返回的是左中或者右中
		return max(left, right) + root.Val
	}

	travel(root)
	return ans
}

//重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {

	var travel func(preorder []int, inorder []int) *TreeNode
	travel = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		} else if len(preorder) == 1 {
			return &TreeNode{
				Val:   preorder[0],
				Left:  nil,
				Right: nil,
			}
		}

		var index int
		for i := 0; i < len(inorder); i++ {
			if preorder[0] == inorder[i] {
				index = i
				break
			}
		}

		node := &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
		node.Left = travel(preorder[1:index+1], inorder[:index])
		node.Right = travel(preorder[index+1:], inorder[index+1:])
		return node
	}

	return travel(preorder, inorder)
}

func recoverTree1(root *TreeNode) {

	m := make(map[int]*TreeNode)
	s1 := make([]int, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		s1 = append(s1, root.Val)
		m[root.Val] = root
		travel(root.Right)
	}

	travel(root)

	s2 := make([]int, len(s1), len(s1))
	copy(s2, s1)
	sort.Ints(s2)

	var index int

	for i := 0; i < len(s2); i++ {
		if s2[i] != s1[i] {
			index = i
			break
		}
	}

	old := s1[index]
	target := s2[index]
	m[old].Val = target
	m[target].Val = old

}

func recoverTree(root *TreeNode) {
	var pre *TreeNode

	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		travel(root.Left)

		if pre == nil {
			pre = root
		}
		if root.Val < pre.Val {
			pre.Val, root.Val = root.Val, pre.Val
		}

		pre = root

		travel(root.Right)
	}

	travel(root)
}

//动态规划 dp[j] = min { dp[j-coins[i] } +1
func coinChange(coins []int, amount int) int {
	//特殊情况处理
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1, amount+1)
	for i := 0; i < len(dp); i++ {
		//初始化处理
		dp[i] = -1
	}

	//base case
	dp[0] = 0
	for i := 0; i < len(coins); i++ {
		if coins[i] < len(dp) {
			dp[coins[i]] = 1
		}
	}

	for i := 0; i < len(dp); i++ {
		//已经处理过的就跳过
		if dp[i] != -1 {
			continue
		}

		min := math.MaxInt
		for j := 0; j < len(coins); j++ {
			//防止出现[2],1的情况 导致索引越界
			if i-coins[j] < 0 {
				continue
			}
			temp := dp[i-coins[j]]
			//有 并且不等-1
			if temp < min && temp != -1 {
				min = temp
			}
		}

		if min != math.MaxInt {
			dp[i] = min + 1
		}
	}
	return dp[amount]
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	dp := make([]int, n+1, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % (1e9 + 7)
	}

	return dp[n]
}

//func permute(nums []int) [][]int {
//	res := make([][]int, 0, len(nums)*(len(nums)-1)/2)
//
//	var travel func(nums []int, temp []int)
//	travel = func(nums []int, temp []int) {
//		if len(nums) == 0 {
//			res = append(res, temp)
//			return
//		}
//
//		for i := 0; i < len(nums); i++ {
//			nums[0], nums[i] = nums[i], nums[0]
//			//temp = append(temp, nums[0])
//			travel(nums[1:], append(temp, nums[0]))
//			nums[0], nums[i] = nums[i], nums[0]
//		}
//	}
//
//	travel(nums, []int{})
//	return res
//}

func sortedArrayToBST(nums []int) *TreeNode {

	var travel func(nums []int) *TreeNode
	travel = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		if len(nums) == 1 {
			return &TreeNode{Val: nums[0]}
		}
		i := len(nums) / 2
		node := &TreeNode{Val: nums[i]}
		node.Left = travel(nums[:i])
		node.Right = travel(nums[i+1:])
		return node
	}
	return travel(nums)
}

//BFS广度优先
func minDepth(root *TreeNode) int {
	res := 1
	queue := make([]*TreeNode, 0)
	if root == nil {
		return 0
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return res
			} else {
				queue = queue[1:]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
		}

		res += 1
	}
	return res
}

func openLock(deadends []string, target string) int {
	m := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		m[deadends[i]] = true
	}

	queue := make([][]byte, 0)
	queue = append(queue, []byte{'0', '0', '0', '0'})

	res := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			start := queue[0]
			queue = queue[1:]

			if m[string(start)] {
				continue
			}

			m[string(start)] = true
			if string(start) == target {
				return res
			}

			for j := 0; j < 4; j++ {
				add := make([]byte, 4)
				copy(add, start)
				add[j] = (start[j]-'0'+1+10)%10 + '0'

				sub := make([]byte, 4)
				copy(sub, start)
				sub[j] = (start[j]-'0'-1+10)%10 + '0'

				if !m[string(add)] {
					queue = append(queue, add)
				}

				if !m[string(sub)] {
					queue = append(queue, sub)
				}
			}
		}
		res++
	}
	return -1
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre, cur := head, head
	for i := 0; i < k; i++ {
		cur = cur.Next
	}

	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}

	return pre
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{i, v}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	m := make(map[uint8]int)
	l := len(s1)
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
	}

	left, right := 0, 0
	for right < len(s2) {

		for right-left < l {
			m[s2[right]]--
			if m[s2[right]] == 0 {
				delete(m, s2[right])
			}
			right++
		}

		if len(m) == 0 {
			return true
		}

		m[s2[left]]++
		if m[s2[left]] == 0 {
			delete(m, s2[left])
		}
		left++
	}

	return false
}

func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need := make(map[uint8]int)
	window := make(map[uint8]int)
	count := 0
	l := len(s1)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	left, right := 0, 0
	for right < len(s2) {
		window[s2[right]]++

		if window[s2[right]] == need[s2[right]] {
			count++
		}
		right++

		for right-left >= l {
			if len(need) == count {
				return true
			}

			if window[s2[left]] == need[s2[left]] {
				count--
			}
			window[s2[left]]--
			left++
		}
	}

	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	a, b := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		a[s1[i]-'a']++
		b[s2[i]-'a']++
	}

	if a == b {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		b[s2[i-len(s1)]-'a']--
		b[s2[i]-'a']++
		if a == b {
			return true
		}
	}

	return false
}

//s = "cbaebabacd", p = "abc"
func findAnagrams(s string, p string) []int {
	res := make([]int, 0, 0)
	need, window := make(map[uint8]int), make(map[uint8]int)
	isValid := 0

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	for right < len(s) {
		window[s[right]]++
		if window[s[right]] == need[s[right]] {
			isValid++
		}
		right++

		for right-left >= len(p) {
			if len(need) == isValid {
				res = append(res, left)
			}
			if window[s[left]] == need[s[left]] {
				isValid--
			}
			window[s[left]]--
			left++
		}
	}

	return res
}

func findAnagrams1(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	res := make([]int, 0, 0)
	a, b := [26]int{}, [26]int{}
	for i := 0; i < len(p); i++ {
		a[p[i]-'a']++
		b[s[i]-'a']++
	}

	if a == b {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		b[s[i-len(p)]-'a']--
		b[s[i]-'a']++

		if a == b {
			res = append(res, i-len(p)+1)
		}
	}

	return res
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]int)

	left, right := 0, 0
	for right < len(s) {
		m[s[right]]++

		for m[s[right]] > 1 {

			m[s[left]]--
			left++
		}

		right++
		temp := right - left
		if res < temp {
			res = temp
		}
	}

	return res
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need, window := make(map[byte]int), make(map[byte]int)
	isValid := 0
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	res := ""
	left, right := 0, 0
	for right < len(s) {
		window[s[right]]++
		if need[s[right]] > 0 && window[s[right]] == need[s[right]] {
			isValid++
		}

		for isValid == len(need) {
			if res == "" || len(res) > right-left {
				res = s[left : right+1]
			}

			window[s[left]]--
			if window[s[left]] < need[s[left]] {
				isValid--
			}
			left++
		}

		right++
	}

	return res
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	sort.Ints(dp)
	return dp[len(dp)-1]
}

func lengthOfLIS1(nums []int) int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		index := sort.SearchInts(res, nums[i])
		if index >= len(res) {
			res = append(res, nums[i])
		} else {
			res[index] = nums[i]
		}
	}

	return len(res)
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return false
		}
	})

	res := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		index := sort.SearchInts(res, envelopes[i][1])
		if index >= len(res) {
			res = append(res, envelopes[i][1])
		} else {
			res[index] = envelopes[i][1]
		}
	}
	return len(res)
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	sort.Ints(dp)
	return dp[len(dp)-1]
}

func xiezhebianli(nums [][]int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			fmt.Println(nums[j][j+i])
		}
	}
}

func daozhebianli(nums [][]int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 1; j < len(nums)-i; j++ {
			fmt.Println(nums[i][i+j])
		}
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1, len(text2)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text1)+1, len(text1)+1)
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if text2[i-1] == text1[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}

	}

	return dp[len(text2)][len(text1)]
}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s), len(s))
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s), len(s))
	}

	for i := 0; i < len(dp); i++ {
		dp[i][i] = 1
	}

	for i := len(dp) - 1; i >= 0; i-- {
		for j := 1; j < len(dp)-i; j++ {
			if s[i] == s[j+i] {
				dp[i][j+i] = dp[i+1][j+i-1] + 2
			} else {
				dp[i][j+i] = max(dp[i+1][j+i], dp[i][j+i-1])
			}
		}
	}

	return dp[0][len(s)-1]
}

//定义很重要 前i个物品限重为j的时候 而不是可以装i个物品
func zeroonebeibao(n, w int, wt, val []int) int {
	dp := make([][]int, n+1, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1, w+1)
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if j >= wt[i-1] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt[i-1]]+val[i-1])
			} else {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}

		}
	}

	return dp[n][w]
}

func canPartition(nums []int) bool {
	mid, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	mid = sum / 2

	dp := make([][]bool, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, mid+1, mid+1)
	}

	dp[0][0] = true

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)][mid]
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(coins)][amount]
}

func rob(nums []int) int {
	dp := make([]int, len(nums)+1, len(nums)+1)
	dp[0] = 0
	if len(dp) > 1 {
		dp[1] = nums[0]
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(nums)]
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	a := rob(nums[1:])
	b := rob(nums[:len(nums)-1])
	if a > b {
		return a
	}
	return b
}

func rob3(root *TreeNode) int {

	m := make(map[*TreeNode]int)

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if v, ok := m[root]; ok {
			return v
		}

		if root == nil {
			return 0
		}
		a := root.Val
		if root.Left != nil {
			a += travel(root.Left.Left)
			a += travel(root.Left.Right)
		}
		if root.Right != nil {
			a += travel(root.Right.Left)
			a += travel(root.Right.Right)
		}

		b := travel(root.Left) + travel(root.Right)

		res := max(a, b)
		m[root] = res
		return res
	}

	return travel(root)
}

//输入：nums = [1,1,1,1,1], target = 3
//输出：5
func findTargetSumWays(nums []int, target int) int {
	res := 0

	var travel func(nums []int, sum int)
	travel = func(nums []int, sum int) {
		if len(nums) == 0 {
			if sum == target {
				res++
			}

			return
		}

		travel(nums[1:], sum+nums[0])
		travel(nums[1:], sum-nums[0])
	}

	travel(nums, 0)

	return res
}

func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= 0 {
			sum += nums[i]
		} else {
			sum += -nums[i]
		}
	}

	if target > sum || target < -sum {
		return 0
	}

	dp := make([][]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum*2+1)
	}
	dp[0][sum] = 1

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j-nums[i-1] >= 0 && j-nums[i-1] < len(dp[i]) {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}

			if j+nums[i-1] >= 0 && j+nums[i-1] < len(dp[i]) {
				dp[i][j] += dp[i-1][j+nums[i-1]]
			}
		}
	}

	return dp[len(nums)][sum+target]
}

//分为两个子集 A 和 B
//A分配+ B分配-
// sum(A)-sum(B)=target
// sum(A) = target +sum(B)
// sum(A)+sum(A) = target + sum(B) + sum(A)
// 2sum(A) = target + sum(A+B)
func findTargetSumWays2(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if (target+sum)%2 != 0 {
		return 0
	}

	a := (target + sum) / 2
	dp := make([][]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, a+1, a+1)
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j < nums[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][a]
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isValidBST(root *TreeNode) bool {
	pre, cur := math.MinInt, math.MinInt

	res := true
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		travel(root.Left)

		if !res {
			return
		}

		pre = cur
		cur = root.Val
		if cur <= pre {
			res = false
			return
		}

		travel(root.Right)
	}

	travel(root)

	return res
}

func isValidBST1(root *TreeNode) bool {
	var travel func(root *TreeNode, max int, min int) bool
	travel = func(root *TreeNode, max int, min int) bool {
		if root == nil {
			return true
		}

		if root.Val >= max || root.Val <= min {
			return false
		}

		return travel(root.Left, root.Val, min) && travel(root.Right, max, root.Val)
	}

	return travel(root, math.MaxInt, math.MinInt)

}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	} else if root.Val == val {
		return root
	} else {
		return searchBST(root.Left, val)
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	res := root

	node := &TreeNode{Val: val}
	last := root
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		last = root

		if root.Val < val {
			return travel(root.Right)
		} else {
			return travel(root.Left)
		}
	}

	if travel(root) == nil {
		if last.Val < val {
			last.Right = node
		} else {
			last.Left = node
		}
	}

	return res
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		} else if root.Val > val {
			root.Left = travel(root.Left)
			return root
		} else {
			root.Right = travel(root.Right)
			return root
		}
	}

	return travel(root)

}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var findLarger func(root *TreeNode) *TreeNode
	findLarger = func(root *TreeNode) *TreeNode {
		if root.Left == nil {
			return root
		}
		return findLarger(root.Left)
	}

	if root == nil {
		return nil
	} else if root.Val == key {
		if root.Right == nil && root.Left == nil {
			return nil
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Right != nil && root.Left == nil {
			return root.Right
		} else {
			larger := findLarger(root.Right)
			root.Val, larger.Val = larger.Val, root.Val
			root.Right = deleteNode(root.Right, larger.Val)
			return root
		}

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes1(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}

	for r != nil {
		r = r.Right
		hr++
	}

	if hr == hl {
		return int(math.Pow(2.0, float64(hl)) - 1)
	} else {
		return 1 + countNodes1(root.Left) + countNodes1(root.Right)
	}
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	strs := make([]string, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			strs = append(strs, " ")
			return
		}

		strs = append(strs, fmt.Sprintf("%d", root.Val))

		travel(root.Left)

		travel(root.Right)
	}
	travel(root)
	return strings.Join(strs, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")

	var travel func() *TreeNode
	travel = func() *TreeNode {
		if len(split) == 0 {
			return nil
		}

		if split[0] == " " {
			split = split[1:]
			return nil
		}

		atoi, _ := strconv.Atoi(split[0])
		split = split[1:]

		node := &TreeNode{Val: atoi}
		node.Left = travel()
		node.Right = travel()
		return node
	}

	return travel()
}

// Serializes a tree to a single string.
func (this *Codec) serialize1(root *TreeNode) string {
	strs := make([]string, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		strs = append(strs, fmt.Sprintf("%d", root.Val))
		travel(root.Left)
		travel(root.Right)
	}
	travel(root)
	return strings.Join(strs, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize1(data string) *TreeNode {
	if data == "" {
		return nil
	}
	split := strings.Split(data, " ")
	ints := make([]int, 0, 0)
	for i := 0; i < len(split); i++ {
		atoi, _ := strconv.Atoi(split[i])
		ints = append(ints, atoi)
	}

	a := make([]int, len(ints), cap(ints))
	copy(a, ints)

	sort.Ints(a)

	var travel func(ints, a []int) *TreeNode
	travel = func(ints, a []int) *TreeNode {
		if len(ints) == 0 {
			return nil
		}
		m := make(map[int]int)
		for i := 0; i < len(a); i++ {
			m[a[i]] = i
		}

		node := &TreeNode{Val: ints[0]}

		node.Left = travel(ints[1:m[node.Val]+1], a[:m[node.Val]])
		node.Right = travel(ints[m[node.Val]+1:], a[m[node.Val]+1:])
		return node
	}

	return travel(ints, a)
}

func (Codec) serialize3(root *TreeNode) string {
	arr := []string{}
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		arr = append(arr, strconv.Itoa(node.Val))
	}
	postOrder(root)
	return strings.Join(arr, " ")
}

func (Codec) deserialize3(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

// Serializes a tree to a single string.
func (this *Codec) serialize4(root *TreeNode) string {
	res := make([]string, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			res = append(res, " ")
			return
		}

		travel(root.Left)
		travel(root.Right)
		res = append(res, fmt.Sprintf("%d", root.Val))
	}

	travel(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize4(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")

	var travel func() *TreeNode
	travel = func() *TreeNode {
		if len(strs) == 0 {
			return nil
		}

		if strs[len(strs)-1] == " " {
			strs = strs[:len(strs)-1]
			return nil
		}

		atoi, _ := strconv.Atoi(strs[len(strs)-1])
		strs = strs[:len(strs)-1]
		node := &TreeNode{Val: atoi}
		node.Right = travel()
		node.Left = travel()
		return node
	}

	return travel()
}

// Serializes a tree to a single string.
func (this *Codec) serialize5(root *TreeNode) string {
	res := make([]string, 0, 0)
	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, " ")
		} else {
			res = append(res, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize5(data string) *TreeNode {
	if data == "" {
		return nil
	}

	strs := strings.Split(data, ",")
	if strs[0] == " " {
		return nil
	}

	atoi, _ := strconv.Atoi(strs[0])
	node := &TreeNode{Val: atoi}
	strs = strs[1:]
	queue := make([]*TreeNode, 0, 0)
	queue = append(queue, node)
	for len(strs) != 0 {
		node := queue[0]
		queue = queue[1:]

		if strs[0] == " " {
			node.Left = nil
		} else {
			i, _ := strconv.Atoi(strs[0])
			newnode := &TreeNode{Val: i}
			node.Left = newnode
			queue = append(queue, newnode)
		}
		if strs[1] == " " {
			node.Right = nil
		} else {
			i, _ := strconv.Atoi(strs[1])
			newnode := &TreeNode{Val: i}
			node.Right = newnode
			queue = append(queue, newnode)
		}
		strs = strs[2:]
	}

	return node
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	m := make(map[*TreeNode]*TreeNode)
	var travel func(parent, child *TreeNode)
	travel = func(parent, child *TreeNode) {
		if child == nil {
			return
		}
		m[child] = parent
		travel(child, child.Left)
		travel(child, child.Right)
	}

	travel(root, root.Left)
	travel(root, root.Right)

	parentP, parentQ := p, q
	visited := make(map[*TreeNode]int)
	visited[parentP]++
	visited[parentQ]++

	var res *TreeNode
	for visited[parentP] != 2 && visited[parentQ] != 2 {
		if v, ok := m[parentP]; ok {
			parentP = v
			visited[parentP]++
		}

		if v, ok := m[parentQ]; ok {
			parentQ = v
			visited[parentQ]++
		}
	}

	if visited[parentP] == 2 {
		res = parentP
	} else {
		res = parentQ
	}
	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var travel func(root, p, q *TreeNode) *TreeNode
	travel = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		left := travel(root.Left, p, q)
		right := travel(root.Right, p, q)

		if left != nil && right != nil {
			return root
		} else if left == nil && right != nil {
			return right
		} else if left != nil && right == nil {
			return left
		} else {
			return nil
		}
	}

	return travel(root, p, q)
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, 0)
	m := make(map[int]int)
	res := make([]int, 0, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for i := 0; i < len(nums1); i++ {
		res = append(res, m[nums1[i]])
	}
	return res
}

func nextGreaterElements1(nums []int) []int {
	res := make([]int, 2*len(nums), 2*len(nums))
	ints := make([]int, len(nums), len(nums)*2)
	copy(ints, nums)
	ints = append(ints, nums...)
	stack := make([]int, 0, 0)

	for i := len(ints) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= ints[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}

		stack = append(stack, ints[i])
	}
	return res[:len(nums)]
}

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	stack := make([]int, 0, 0)
	for i := len(nums)*2 - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%len(nums)] {
			stack = stack[:len(stack)-1]
		}

		if i < len(res) {
			if len(stack) == 0 {
				res[i] = -1
			} else {
				res[i] = stack[len(stack)-1]
			}
		}

		stack = append(stack, nums[i%len(nums)])
	}
	return res[:len(nums)]
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))
	stack := make([]int, 0, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return res
}

type Queue struct {
	linklist *list.List
}

func (q *Queue) Push(x int) {
	for q.linklist.Len() != 0 && q.linklist.Back().Value.(int) < x {
		q.linklist.Remove(q.linklist.Back())
	}
	q.linklist.PushBack(x)
}

func (q *Queue) Pop(x int) {
	if q.linklist.Front().Value.(int) == x {
		q.linklist.Remove(q.linklist.Front())
	}
}

func (q *Queue) Max() int {
	return q.linklist.Front().Value.(int)
}

func maxSlidingWindow1(nums []int, k int) []int {
	res := make([]int, 0, 0)
	window := Queue{linklist: &list.List{}}
	l := 0

	for i := 0; i < len(nums); i++ {
		if l < k {
			window.Push(nums[i])
			l++
			if l == k {
				res = append(res, window.Max())
			}
		} else {
			window.Pop(nums[i-k])
			window.Push(nums[i])
			res = append(res, window.Max())
		}
	}

	return res
}

type myheap struct {
	ints []int
}

func (m *myheap) Len() int {
	return len(m.ints)
}

func (m *myheap) Less(i, j int) bool {
	return (*m).ints[i] > (*m).ints[j]
}

func (m *myheap) Swap(i, j int) {
	(*m).ints[i], (*m).ints[j] = (*m).ints[j], (*m).ints[i]
}

func (m *myheap) Push(x interface{}) {
	(*m).ints = append((*m).ints, x.(int))
}

func (m *myheap) Pop() interface{} {
	i := (*m).ints[len((*m).ints)-1]
	(*m).ints = (*m).ints[:len((*m).ints)-1]
	return i
}

func maxSlidingWindow(nums []int, k int) []int {
	m := make(map[int]int)
	h := &myheap{ints: []int{}}
	res := make([]int, 0, 0)
	count := 0
	for i := 0; i < len(nums); i++ {
		if count < k {
			heap.Push(h, nums[i])
			m[nums[i]] = i
			count++

			if count == k {
				res = append(res, (*h).ints[0])
			}
		} else {
			heap.Push(h, nums[i])
			m[nums[i]] = i
			for m[h.ints[0]] <= i-k {
				heap.Pop(h)
			}
			res = append(res, (*h).ints[0])
		}
	}
	return res
}

//aaa  6
func countSubstrings1(s string) int {
	res := len(s)
	dp := make([][]bool, len(s), len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s), len(s))
		dp[i][i] = true
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			if s[j] == s[j+i] && ((j+i)-(j) < 2 || dp[j+1][j+i-1]) {
				dp[j][j+i] = true
				res++
			} else {
				dp[j][j+i] = false
			}
		}
	}

	return res
}
func countSubstrings(s string) int {
	res := 0
	var travel func(l, r int)
	travel = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			res++
			l--
			r++
		}
	}

	for i := 0; i < len(s); i++ {
		travel(i, i)
		travel(i, i+1)
	}
	return res
}

func longestPalindrome(s string) int {
	res := 0
	a := [26]int{}
	A := [26]int{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' {
			a[s[i]-'a']++
		} else {
			A[s[i]-'A']++
		}
	}

	b := false
	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			res += a[i]
		} else {
			res += a[i] / 2 * 2
			b = true
		}
	}

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res += A[i]
		} else {
			res += A[i] / 2 * 2
			b = true
		}
	}
	if b {
		res++
	}
	return res
}

func isPalindrome1(head *ListNode) bool {
	ints := make([]int, 0, 0)
	var travel func(head *ListNode)
	travel = func(head *ListNode) {
		if head == nil {
			return
		}
		ints = append(ints, head.Val)
		travel(head.Next)
	}

	travel(head)
	c := make([]int, len(ints), cap(ints))
	copy(c, ints)
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-1-i] = c[len(c)-1-i], c[i]
	}

	for i := 0; i < len(c); i++ {
		if c[i] != ints[i] {
			return false
		}
	}

	return true
}

func isPalindrome2(head *ListNode) bool {
	ints := make([]int, 0, 0)
	rev := make([]int, 0, 0)
	var travel func(head *ListNode)
	travel = func(head *ListNode) {
		if head == nil {
			return
		}
		ints = append(ints, head.Val)
		travel(head.Next)
		rev = append(rev, head.Val)

	}

	travel(head)

	for i := 0; i < len(rev); i++ {
		if rev[i] != ints[i] {
			return false
		}
	}

	return true
}

func isPalindrome3(head *ListNode) bool {
	ints := make([]int, 0, 0)

	for head != nil {
		ints = append(ints, head.Val)
		head = head.Next
	}

	c := make([]int, len(ints), cap(ints))
	copy(c, ints)
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-1-i] = c[len(c)-1-i], c[i]
	}

	for i := 0; i < len(c); i++ {
		if c[i] != ints[i] {
			return false
		}
	}

	return true
}

func isPalindrome4(head *ListNode) bool {
	listnode := &ListNode{}
	cur := listnode

	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		cur.Next = &ListNode{Val: head.Val}
		cur = cur.Next

		root := travel(head.Next)

		if head.Next != nil {
			head.Next.Next = head
			head.Next = nil
		}
		if root == nil {
			return head
		}
		return root
	}
	head = travel(head)

	listnode = listnode.Next
	for head != nil && listnode != nil {
		if head.Val != listnode.Val {
			return false
		}
		head = head.Next
		listnode = listnode.Next
	}

	return true
}

func isPalindrome5(head *ListNode) bool {
	left := head

	var travel func(head *ListNode) bool
	travel = func(head *ListNode) bool {
		if head == nil {
			return true
		}

		res := travel(head.Next)

		if head.Val == left.Val && res {
			left = left.Next
			return true
		}
		return false
	}
	return travel(head)
}

func isPalindrome6(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		ans := travel(head.Next)
		if ans == nil {
			return head
		}

		head.Next.Next = head
		head.Next = nil
		return ans
	}

	node := travel(slow)

	left, right := head, node
	for left.Val == right.Val && right.Next != nil && left.Next != nil {
		left = left.Next
		right = right.Next
	}
	return left.Val == right.Val
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		if head.Next == nil {
			return head
		}

		pre, cur := head, head.Next
		pre.Next = nil
		for cur != nil {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}

		return pre
	}

	node := travel(slow)

	left, right := head, node
	for left.Val == right.Val && right.Next != nil && left.Next != nil {
		left = left.Next
		right = right.Next
	}
	return left.Val == right.Val
}

func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	i := 1
	root := head
	var start *ListNode
	var end *ListNode
	var pre *ListNode

	var preStart *ListNode
	var afterEnd *ListNode
	for start == nil || end == nil {
		if i == left {
			preStart = pre
			start = root
		}
		if i == right {
			end = root
			afterEnd = end.Next
		}

		pre = root
		root = root.Next
		i++
	}

	var reverse func(start, end *ListNode)
	reverse = func(start, end *ListNode) {
		var pre *ListNode
		cur := start

		for pre != end {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}
	}

	reverse(start, end)

	if preStart != nil {
		preStart.Next = end
	} else {
		head = end
	}
	start.Next = afterEnd
	return head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	var successor *ListNode
	var travel func(head *ListNode, n int) *ListNode
	travel = func(head *ListNode, n int) *ListNode {
		if n == 1 {
			successor = head.Next
			return head
		}

		last := travel(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return last
	}

	if left == 1 {
		return travel(head, right)
	} else {
		head.Next = reverseBetween(head.Next, left-1, right-1)
	}
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var travel func(head *ListNode, k int) *ListNode
	travel = func(head *ListNode, k int) *ListNode {
		if head == nil {
			return nil
		}

		var pre *ListNode
		cur := head

		for i := 0; i < k; i++ {
			if cur == nil {
				return head
			}
			cur = cur.Next
		}

		cur = head
		for i := 0; i < k && cur != nil; i++ {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}
		head.Next = travel(cur, k)
		return pre
	}
	return travel(head, k)
}

//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
func subsets1(nums []int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(ans, nums []int, count int)
	travel = func(ans, nums []int, count int) {
		if len(ans) == count {
			temp := make([]int, count, count)
			copy(temp, ans)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			nums[0], nums[i] = nums[i], nums[0]
			ans = append(ans, nums[0])
			travel(ans, nums[i+1:], count)
			ans = ans[:len(ans)-1]
			nums[0], nums[i] = nums[i], nums[0]
		}
	}

	for i := 0; i <= len(nums); i++ {
		travel([]int{}, nums, i)
	}

	return res
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(nums []int, start int, ans []int)
	travel = func(nums []int, start int, ans []int) {
		temp := make([]int, len(ans), len(ans))
		copy(temp, ans)
		res = append(res, temp)

		for i := start; i < len(nums); i++ {
			ans = append(ans, nums[i])
			travel(nums, i+1, ans)
			ans = ans[:len(ans)-1]
		}
	}

	travel(nums, 0, []int{})

	return res
}

//输入: n = 4, k = 2
//输出:
//[
//[2,4],
//[3,4],
//[2,3],
//[1,2],
//[1,3],
//[1,4],
//]
func combine(n int, k int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(ans []int, start, count int)
	travel = func(ans []int, start, count int) {
		if count == k {
			temp := make([]int, count, count)
			copy(temp, ans)
			res = append(res, temp)
			return
		}

		for i := start; i <= n; i++ {
			ans = append(ans, i)
			travel(ans, i+1, count+1)
			ans = ans[:len(ans)-1]
		}
	}

	travel([]int{}, 1, 0)

	return res
}

func combinationSum41(nums []int, target int) int {
	dp := make(map[int]int)

	var travel func(target int) int
	travel = func(target int) int {
		if v, ok := dp[target]; ok {
			return v
		}

		if target == 0 {
			return 1
		}
		ans := 0
		for i := 0; i < len(nums); i++ {
			if target >= nums[i] {
				ans += travel(target - nums[i])
			}
		}

		dp[target] = ans
		return ans
	}

	return travel(target)

}

func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)

	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]

}

func permute(nums []int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(ans []int, nums []int)
	travel = func(ans []int, nums []int) {
		if len(nums) == 0 {
			temp := make([]int, len(ans), len(ans))
			copy(temp, ans)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			nums[i], nums[0] = nums[0], nums[i]
			ans = append(ans, nums[0])
			travel(ans, nums[1:])
			nums[i], nums[0] = nums[0], nums[i]
			ans = ans[:len(ans)-1]
		}
	}

	travel([]int{}, nums)

	return res
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, 0)

	var travel func(ans string, left, right int)
	travel = func(ans string, left, right int) {
		if left == n && right == n {
			res = append(res, ans)
			return
		}

		if left > n || right > n || right > left {
			return
		}

		travel(ans+"(", left+1, right)

		travel(ans+")", left, right+1)
	}

	travel("", 0, 0)

	return res
}

//单调栈  哨兵节点加入之后一视同仁
func largestRectangleArea(heights []int) int {
	if len(heights) == 1 {
		return heights[0]
	}
	res := 0
	stack := make([]int, 0, 0)
	newHeight := make([]int, len(heights)+2, len(heights)+2)
	copy(newHeight[1:], heights)

	left := make([]int, len(newHeight), len(newHeight))
	right := make([]int, len(newHeight), len(newHeight))
	stack = append(stack, 0)
	for i := 0; i < len(newHeight); i++ {
		for len(stack) != 0 && newHeight[stack[len(stack)-1]] >= newHeight[i] {
			right[stack[len(stack)-1]] = i - stack[len(stack)-1] - 1
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = i - 1
		} else {
			left[i] = i - stack[len(stack)-1] - 1
		}
		stack = append(stack, i)
	}

	for i := 1; i < len(newHeight)-1; i++ {
		temp := (left[i] + right[i] + 1) * newHeight[i]
		if temp > res {
			res = temp
		}
	}

	return res
}

func a(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
