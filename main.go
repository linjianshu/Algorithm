package main

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//ints := reversePrint2(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  2,
	//			Next: nil,
	//		},
	//	},
	//})
	//fmt.Println(ints)

	//list := reverseList(&ListNode{
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
	//})
	//
	//fmt.Println(list)

	//lists := mergeTwoLists(&ListNode{Val: 1, Next: &ListNode{2, &ListNode{4, nil}}}, &ListNode{Val: 1, Next: &ListNode{3, &ListNode{4, nil}}})
	//fmt.Println(lists)

	//remaining := lastRemaining(6, 7)
	//fmt.Println(remaining)

	//numbers := printNumbers(1)
	//fmt.Println(numbers)

	//printNumbers1(1)
	//t := &TreeNode{
	//	Val: 3,
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
	//depth := maxDepth(t)
	//fmt.Println(depth)
	//depth1 := maxDepth(t)
	//fmt.Println(depth1)

	//words := reverseLeftWords("lrloseumgh", 6)
	//fmt.Println(words)
	//space := replaceSpace("We are happy.")
	//fmt.Println(space)

	//end := getKthFromEnd(&ListNode{
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
	//}, 2)
	//fmt.Println(end)

	//element := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	//fmt.Println(element)

	//largest := kthLargest(&TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: &TreeNode{2, nil, nil},
	//	},
	//	Right: &TreeNode{
	//		Val:   4,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}, 1)

	//fmt.Println(largest)

	//weight := hammingWeight(4294967293)
	//fmt.Println(weight)

	//words := findWords([]string{"asdddafadsfa"})
	//fmt.Println(words)

	//sequence := findContinuousSequence(15)
	//fmt.Println(sequence)

	//Next := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val:  9,
	//			Next: nil,
	//		},
	//	},
	//}

	//l := &ListNode{
	//	Val:  4,
	//	Next: Next,
	//}
	//fmt.Println(l)
	//deleteNode(Next)

	//obj := Constructor()
	//obj.AppendTail(3)
	//head := obj.DeleteHead()
	//fmt.Println(head)
	//deleteHead := obj.DeleteHead()
	//fmt.Println(deleteHead)
	//i := obj.DeleteHead()
	//fmt.Println(i)

	//node := TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 4,
	//			Left: &TreeNode{
	//				Val:   8,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   9,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val: 5,
	//			Left: &TreeNode{
	//				Val:   10,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   11,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//	Right: nil,
	//}
	//p := TreeNode{
	//	Val:   8,
	//	Left:  nil,
	//	Right: nil,
	//}
	//q := TreeNode{
	//	Val:   10,
	//	Left:  nil,
	//	Right: nil,
	//}
	//ancestor := lowestCommonAncestor2(&node, &p, &q)
	//fmt.Println(ancestor)

	//node := TreeNode{
	//	Val: 1,
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
	//		Val:  2,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	//symmetric := isSymmetric2(&node)
	//fmt.Println(symmetric)

	//node = TreeNode{
	//	Val:   1,
	//	Left:  &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{
	//			Val:   3,
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
	//		Val:   2,
	//		Left:  &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//symmetric := isSymmetric2(&node)
	//fmt.Println(symmetric)
	//
	//t := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//
	//fmt.Println(*t.Left==*t.Right)
	//i := new(*TreeNode)
	//i=&node.Left
	//fmt.Println(*i == node.Left)

	//node := TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}
	//
	//balanced := isBalanced(&node)
	//fmt.Println(balanced)

	//element := majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2})
	//fmt.Println(element)

	//duration := findPoisonedDuration1([]int{1, 2}, 2)
	//fmt.Println(duration)

	//same := &ListNode{
	//	Val: 8,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//node1 := ListNode{
	//	Val: 4,
	//	Next: &ListNode{
	//		Val:  1,
	//		Next: same,
	//	},
	//}
	//
	//node2 := ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 0,
	//		Next: &ListNode{
	//			Val:  1,
	//			Next: same,
	//		},
	//	},
	//}
	//
	//node := getIntersectionNode2(&node2, &node1)
	//fmt.Println(node)

	//number := findRepeatNumber([]int{2, 3, 1, 2, 5, 3})
	//fmt.Println(number)
	//
	//use := detectCapitalUse("USA")
	//fmt.Println(use)

	//sum := twoSum1([]int{2, 7, 11, 15}, 9)
	//fmt.Println(sum)

	//ints := exchange1([]int{1, 2, 3, 4})
	//fmt.Println(ints)
	//
	//char := firstUniqChar("dddccdbba")
	//fmt.Println(char)

	//array := maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4})
	//fmt.Println(array)

	//i := add1(1, 1)
	//fmt.Println(i)

	//numbers1 := getLeastNumbers1([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	//fmt.Println(numbers1)

	//A := &TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 1,
	//		},
	//		Right: &TreeNode{
	//			Val: 2,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 5,
	//	},
	//}
	//B := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 1,
	//	},
	//}
	//structure := isSubStructure(A, B)
	//fmt.Println(structure)

	//sequences := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})
	//fmt.Println(sequences)

	//nums := sumNums(9)
	//fmt.Println(nums)

	//order1 := levelOrder1(&TreeNode{
	//	Val: 3,
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
	//})
	//fmt.Println(order1)
	//t := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val: 11,
	//			Left: &TreeNode{
	//				Val:   7,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   2,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:  8,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val: 4,
	//			Left: &TreeNode{
	//				Val:   5,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   1,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//}
	//t = &TreeNode{
	//	Val:  -6,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val: -3,
	//		Left: &TreeNode{
	//			Val: -6,
	//			Left: &TreeNode{
	//				Val:   -6,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val: -5,
	//				Left: &TreeNode{
	//					Val:   1,
	//					Left:  nil,
	//					Right: nil,
	//				},
	//				Right: &TreeNode{
	//					Val:   7,
	//					Left:  nil,
	//					Right: nil,
	//				},
	//			},
	//		},
	//		Right: nil,
	//	},
	//}
	//sum := pathSum(t, -21)
	//fmt.Println(sum)
	i := permutation2("aab")
	fmt.Println(i)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//使用递归1 不够数量 还得拿刀int[]后反转
func reversePrint(head *ListNode) []int {
	result := make([]int, 0, 10000)
	var pushValue func(node *ListNode)
	pushValue = func(node *ListNode) {
		if node == nil {
		} else {
			result = append(result, node.Val)
			pushValue(node.Next)
		}
	}

	pushValue(head)
	lenth := len(result)
	for i := 0; i < lenth/2; i++ {
		temp := result[i]
		result[i] = result[lenth-1-i]
		result[lenth-1-i] = temp
	}
	return result
}

//使用递归 调整递归的顺序 不需要反转数组
func reversePrint1(head *ListNode) []int {
	var result = make([]int, 0, 10000)
	var pushValue func(head *ListNode)
	pushValue = func(head *ListNode) {
		if head == nil {
		} else {
			pushValue(head.Next)
			result = append(result, head.Val)
		}
	}
	pushValue(head)

	return result
}

//遍历链表
func reversePrint2(head *ListNode) []int {
	var result = make([]int, 0, 10000)
	if head == nil {
		return []int{}
	}
	i := 0
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
		i++
	}
	lenth := len(result)
	for i := 0; i < lenth/2; i++ {
		temp := result[i]
		result[i] = result[lenth-1-i]
		result[lenth-1-i] = temp
	}
	return result
}

func reverseList(head *ListNode) *ListNode {
	temp := head
	if head == nil {
		return nil
	} else {
		temp = head.Next
		head.Next = nil
	}

	for temp != nil {
		temp1 := temp.Next
		temp.Next = head
		head = temp
		temp = temp1
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var final = &ListNode{}
	var cur = &ListNode{}
	final.Next = cur
	var pre = &ListNode{}
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			pre = l1
			l1 = l1.Next
		} else {
			pre = l2
			l2 = l2.Next
		}
		cur.Next = pre
		cur = cur.Next
	}
	final = final.Next.Next
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return final
}

//约瑟夫环问题 暂时还不会啊啊啊啊啊啊啊啊
func lastRemaining(n int, m int) int {
	if n == 1 {
		return n
	}
	if m == 1 {
		return n - 1
	}
	node := &ListNode{}
	cur := &ListNode{}
	node.Next = cur
	for i := 1; i < n; i++ {
		pre := &ListNode{Val: i}
		cur.Next = pre
		cur = pre
	}

	node = node.Next
	cur.Next = node

	var jump int
	jump = m % n
	for i := 0; i < n-1; i++ {
		if m > n {
			jump = m % (n - i)
			if jump == 1 {
				jump = m
			}
		} else {
			jump = m % n
		}

		count := 1
		temp := node
		for count < jump-1 {
			temp = temp.Next
			count++
		}
		temp.Next = temp.Next.Next
		node = temp.Next
	}
	return node.Val
}

func printNumbers(n int) []int {
	s := "1"
	for i := 0; i < n; i++ {
		s += "0"
	}
	maxint, _ := strconv.Atoi(s)
	final := make([]int, 0, maxint-1)
	for i := 1; i < maxint; i++ {
		final = append(final, i)
	}
	return final
}

//位运算
func printNumbers1(n int) []int {
	var maxint int
	maxint = 1
	for i := 0; i < n; i++ {
		maxint = maxint<<3 + maxint<<1
	}
	final := make([]int, 0, maxint-1)
	for i := 1; i < maxint; i++ {
		final = append(final, i)
	}
	return final
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft := maxDepth(root.Left) + 1
	depthRight := maxDepth(root.Right) + 1
	if depthRight > depthLeft {
		return depthRight
	} else {
		return depthLeft
	}
}

//func maxDepth1(root *TreeNode) int {
//	i:=-1
//	q := queue.New()
//	q.Add(root)
//	temp := queue.New()
//	temp.Add(root)
//	for temp.Length() != 0 {
//		for q.Length() != 0 {
//			q.Remove()
//			//q.Add(peek.Left)
//			//q.Add(peek.Right)
//
//		}
//
//		var peek *TreeNode
//		peek = temp.Remove().(*TreeNode)
//		if peek!=nil {
//			q.Add(peek.Left)
//			q.Add(peek.Right)
//		}
//		if temp.Length() == 0 {
//			if peek!=nil {
//				temp.Add(peek.Left)
//				temp.Add(peek.Right)
//			}
//			i++
//		}
//	}
//	return i
//}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root != nil {
		tempLeft := mirrorTree(root.Left)
		tempRight := mirrorTree(root.Right)
		root.Left = tempRight
		root.Right = tempLeft
	}
	return root
}

func reverseLeftWords(s string, n int) string {
	if n == 0 || n >= len(s) {
		return s
	} else {
		return s[n:] + s[:n]
	}
}

func replaceSpace1(s string) string {
	return strings.Replace(s, " ", "%20", len(s))
}
func replaceSpace(s string) string {
	var final string
	//index :=0
	for i := range s {
		if s[i] == ' ' {
			final += "%20"
			//index = i
		} else {
			final += string(s[i])
			//index = i
		}
	}
	return final

}

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	cur := head
	pre := &ListNode{}
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	var slice = make([]int, 0, k)
	for i := 0; i < k; i++ {
		slice = append(slice, pre.Val)
		pre = pre.Next
	}

	final := &ListNode{
		Val:  slice[len(slice)-1],
		Next: nil,
	}
	cur = final
	for i := len(slice) - 1; i >= 1; i-- {
		temp := &ListNode{
			Val:  slice[i-1],
			Next: nil,
		}
		cur.Next = temp
		cur = temp
	}
	return final
}
func getKthFromEnd2(head *ListNode, k int) *ListNode {
	pre := head
	for i := 1; i < k; i++ {
		pre = pre.Next
	}

	cur := head
	for pre.Next != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return cur
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
	cur := head
	for i := 1; i < k; i++ {
		head = head.Next
	}

	for head.Next != nil {
		head = head.Next
		cur = cur.Next
	}
	return cur
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, len(nums2))
	for i := 0; i < len(nums2)-1; i++ {
		b := false
		for j := i + 1; j < len(nums2); j++ {
			if nums2[i] < nums2[j] {
				m[nums2[i]] = nums2[j]
				b = true
				break
			}
		}
		if !b {
			m[nums2[i]] = -1
		}
	}
	m[nums2[len(nums2)-1]] = -1

	final := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		final = append(final, m[v])
	}
	return final
}

//很傻逼的做法 直接全部遍历 全部取出 全部排序 取索引
func kthLargest1(root *TreeNode, k int) int {
	array := make([]int, 0, 10)
	var getValue func(tree *TreeNode)
	getValue = func(tree *TreeNode) {
		if tree != nil {
			array = append(array, tree.Val)
			getValue(tree.Left)
			getValue(tree.Right)
		}
	}
	getValue(root)
	sort.Ints(array)
	return array[len(array)-k]
}

//看了题解 知道二叉搜索树中序遍历的结果已经是排好序的 但是还是全部遍历了 啥子
func kthLargest2(root *TreeNode, k int) int {
	array := make([]int, 0, 10)
	var getValue func(tree *TreeNode)
	getValue = func(tree *TreeNode) {
		if tree != nil {
			getValue(tree.Left)
			array = append(array, tree.Val)
			getValue(tree.Right)
		}
	}
	getValue(root)
	return array[len(array)-k]
}

//要求最大值 我懂了啊!!1hahahahahahah就是从右子树开始遍历 且遍历到长度=k为止就停下
func kthLargest3(root *TreeNode, k int) int {
	array := make([]int, 0, k)
	var getValue func(tree *TreeNode)
	getValue = func(tree *TreeNode) {
		if tree != nil {
			getValue(tree.Right)
			if len(array) >= k {
				return
			}
			array = append(array, tree.Val)

			getValue(tree.Left)
		}
	}
	getValue(root)
	return array[k-1]
}
func kthLargest(root *TreeNode, k int) int {
	var final int
	var getValue func(tree *TreeNode)
	getValue = func(tree *TreeNode) {
		if tree != nil {
			getValue(tree.Right)
			k = k - 1
			if k == 0 {
				final = tree.Val
				return
			} else if k < 0 {
				return
			}
			getValue(tree.Left)
		}
	}
	getValue(root)

	return final
}

func hammingWeight1(num uint32) int {
	sprintf := fmt.Sprintf("%b", num)
	final := make(map[int]int)

	//这里有个很关键的问题 就是字符串"1"--->int 不会转成1 而是会转成49 字符串"a" --->int 转成97  所以需要使用atoi
	for _, v := range sprintf {
		i := final[int(v)]
		i++
		final[int(v)] = i
	}
	i := final[int('1')]
	return i
}

func hammingWeight(num uint32) int {
	sprintf := fmt.Sprintf("%b", num)
	sum := 0
	for _, v := range sprintf {
		if v == '1' {
			sum++
		}
	}
	return sum
}

func findWords(words []string) []string {
	line1 := "qwertyuiop"
	line2 := "asdfghjkl"
	line3 := "zxcvbnm"
	final := make([]string, 0, len(words))

	for _, word := range words {
		wordLower := strings.ToLower(word)
		b := false
		if strings.Contains(line1, string(wordLower[0])) {
			for i := 1; i < len(wordLower); i++ {
				if !strings.Contains(line1, string(wordLower[i])) {
					b = true
					break
				}
			}
			if !b {
				final = append(final, word)
			}
		} else if strings.Contains(line2, string(wordLower[0])) {
			for i := 1; i < len(wordLower); i++ {
				if !strings.Contains(line2, string(wordLower[i])) {
					b = true
					break
				}
			}
			if !b {
				final = append(final, word)
			}
		} else if strings.Contains(line3, string(wordLower[0])) {
			for i := 1; i < len(wordLower); i++ {
				if !strings.Contains(line3, string(wordLower[i])) {
					b = true
					break
				}
			}
			if !b {
				final = append(final, word)
			}
		}
	}
	return final
}

//改成函数调用 内存空间增加了 ...
func findWords1(words []string) []string {
	line1 := "qwertyuiop"
	line2 := "asdfghjkl"
	line3 := "zxcvbnm"
	final := make([]string, 0, len(words))

	for _, word := range words {
		wordLower := strings.ToLower(word)
		b := false

		var a func(string, string)
		a = func(line string, wordLower string) {
			for i := 1; i < len(wordLower); i++ {
				if !strings.Contains(line, string(wordLower[i])) {
					b = true
					break
				}
			}
			if !b {
				final = append(final, word)
			}
		}

		if strings.Contains(line1, string(wordLower[0])) {
			a(line1, wordLower)
		} else if strings.Contains(line2, string(wordLower[0])) {
			a(line2, wordLower)
		} else if strings.Contains(line3, string(wordLower[0])) {
			a(line3, wordLower)
		}
	}
	return final
}

//增加的更多了...
func findWords2(words []string) []string {
	m := map[uint8]int{
		'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
		'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
		'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3,
	}
	final := make([]string, 0, len(words))

	for _, word := range words {
		b := false
		wordLower := strings.ToLower(word)
		result := m[wordLower[0]]
		for i := 1; i < len(wordLower); i++ {
			if m[wordLower[i]] != result {
				b = true
				break
			}
		}
		if !b {
			final = append(final, word)
		}
	}
	return final
}

func findContinuousSequence1(target int) [][]int {
	final := [][]int{}
	half := target / 2
	for i := 1; i <= half; i++ {
		sum := 0
		temp := []int{}
		for j := i; j <= half+1; j++ {
			temp = append(temp, j)
			sum += j
			if sum == target {
				final = append(final, temp)
			} else if sum > target {
				break
			}
		}
	}
	return final
}

//就变了一句话 内存直接砍掉2/3
func findContinuousSequence2(target int) [][]int {
	final := [][]int{}
	half := target / 2
	for i := 1; i <= half; i++ {
		sum := 0
		temp := []int{}
		for j := i; j <= half+1; j++ {
			sum += j
			if sum == target {
				for k := i; k <= j; k++ {
					temp = append(temp, k)
				}
				final = append(final, temp)
			} else if sum > target {
				break
			}
		}
	}
	return final
}

//使用滑动窗口 实在是太秀了鸭!!!我去! 牛皮
func findContinuousSequence(target int) [][]int {
	final := [][]int{}
	half := target/2 + 1
	i, j := 1, 2
	sum := i + j
	for j <= half {
		if sum < target {
			j++
			sum += j
		} else if sum == target {
			temp := make([]int, 0, j-i+1)
			for k := i; k <= j; k++ {
				temp = append(temp, k)
			}
			final = append(final, temp)
			sum -= i
			i++
		} else {
			sum -= i
			i++
		}
	}
	return final
}

func deleteNode(node *ListNode) {
	temp := node
	for temp != nil && temp.Next != nil {
		temp.Val = temp.Next.Val
		if temp.Next.Next == nil {
			temp.Next = nil
		}
		temp = temp.Next
	}
}

type CQueue struct {
	left  []int
	right []int
}

func Constructor() CQueue {
	cQueue := CQueue{
		left:  make([]int, 0, 10),
		right: make([]int, 0, 10),
	}
	return cQueue
}

func (this *CQueue) AppendTail(value int) {
	this.left = append(this.left, value)
}

func (this *CQueue) DeleteHead() int {
	var final int
	if len(this.right) != 0 {
		len := len(this.right)
		final = this.right[len-1]
		this.right = this.right[:len-1]
	} else if len(this.left) != 0 {
		this.right = this.left
		this.left = []int{}

		len := len(this.right)
		for i := 0; i < len/2; i++ {
			temp := this.right[i]
			this.right[i] = this.right[len-1-i]
			this.right[len-1-i] = temp
		}

		final = this.right[len-1]
		this.right = this.right[:len-1]
	} else {
		return -1
	}

	return final
}

//递归实在是不会 搞了两天了
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)

	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return nil
	}
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

//实在是不会 方法清晰 先遍历一遍 然后记录下来祖先 然后遍历某个节点的祖先 然后遍历另一个节点的时候去判断祖先是否已经遍历过了
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
		}
		travel(root.Left)
		travel(root.Right)
	}
	travel(root)

	node := p
	visited[node.Val] = true
	for node.Val != root.Val {
		node = parent[node.Val]
		visited[node.Val] = true
	}
	visited[root.Val] = true

	node1 := q
	for !visited[node1.Val] {
		node1 = parent[node1.Val]
	}
	return node1
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parent[root.Left.Val] = root
		}
		if root.Right != nil {
			parent[root.Right.Val] = root
		}
		travel(root.Left)
		travel(root.Right)
	}
	travel(root)

	visited[p.Val] = true
	for p.Val != root.Val {
		p = parent[p.Val]
		visited[p.Val] = true
	}
	visited[root.Val] = true

	for !visited[q.Val] {
		q = parent[q.Val]
	}
	return q
}

//这题思路到了 但是还是不会
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	final := [][]int{}
	var travel func(root []*TreeNode)
	travel = func(root []*TreeNode) {
		if len(root) == 0 {
			return
		}
		//临时存储数组的
		temp := []int{}
		//临时存储某个节点下挂着的节点
		tempQueue := []*TreeNode{}
		for _, item := range root {
			temp = append(temp, item.Val)
			if item.Left != nil {
				tempQueue = append(tempQueue, item.Left)
			}
			if item.Right != nil {
				tempQueue = append(tempQueue, item.Right)
			}
		}
		final = append(final, temp)
		queue = tempQueue
	}

	//循环的跑
	for len(queue) != 0 {
		travel(queue)
	}

	return final
}

//这里我觉得最关键的一点就是在于 通过变量b来控制是先向左打印还是先向右打印
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	} else if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	var final = make([]int, 0, 10)
	var final1 = make([]int, 0, 10)
	var travel func(root *TreeNode, final *[]int)
	b := true
	travel = func(root *TreeNode, final *[]int) {
		if root == nil {
			*final = append(*final, 0)
			return
		}
		*final = append(*final, root.Val)
		if b {
			travel(root.Left, final)
			travel(root.Right, final)
		} else {
			travel(root.Right, final)
			travel(root.Left, final)
		}
	}
	travel(root.Left, &final)
	b = false
	travel(root.Right, &final1)
	if len(final) != len(final1) {
		return false
	}
	for k := range final {
		if (final1)[k] != (final)[k] {
			return false
		}
	}
	return true
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	} else if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}
	var temp *TreeNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		temp = root.Left
		root.Left = root.Right
		root.Right = temp
		travel(root.Left)
		travel(root.Right)
	}

	travel(root.Left)

	return reflect.DeepEqual(*root.Left, *root.Right)
}

//这题有点难 主要是最后的思路没到位 最后都会回归到nil点上判断 判断完返回什么呢 用什么来表示错误呢
func isBalanced(root *TreeNode) bool {
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := travel(root.Left)
		right := travel(root.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if left == right {
			return left + 1
		} else if math.Abs(float64(left-right)) == 1 {
			return max(left, right) + 1
		} else {
			return -1
		}
	}
	return travel(root) != -1
}

func majorityElement(nums []int) int {
	m := make(map[int]int)
	l := len(nums) / 2
	var final int
	for _, v := range nums {
		m[v]++
		if m[v] > l {
			final = v
			break
		}
	}
	return final
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	low := -1
	var final int
	for _, v := range timeSeries {
		if v <= low {
			//如果v比low小 说明区间重合了 那就扣除掉之间的不需要+1
			final += (v + duration - 1) - low
		} else {
			//否则 直接加上持续时间
			final += duration
		}
		//更新low的值
		low = v + duration - 1
	}
	return final
}

func findPoisonedDuration1(timeSeries []int, duration int) int {
	var final int
	for i := range timeSeries {
		if i == 0 {
			final += duration
			timeSeries[i] = timeSeries[i] + duration - 1
			continue
		}
		if timeSeries[i] <= timeSeries[i-1] {
			final += (timeSeries[i] + duration - 1) - timeSeries[i-1]
		} else {
			final += duration
		}
		timeSeries[i] = timeSeries[i] + duration - 1
	}
	return final
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	//备忘录 两边同时移动记录有米有已经访问过的
	visited := make(map[*ListNode]bool)
	var final *ListNode
	var travel func(root *ListNode)
	travel = func(root *ListNode) {
		for root != nil {
			if _, ok := visited[root]; ok {
				final = root
				break
			}
			visited[root] = true
			root = root.Next
		}
	}

	travel(headA)
	travel(headB)

	return final
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	visited := make(map[*ListNode]bool)
	var final *ListNode
	tempA := headA
	tempB := headB
	for final == nil && (tempA != nil || tempB != nil) {
		if tempA != nil {
			if _, ok := visited[tempA]; ok {
				final = tempA
				break
			}
			visited[tempA] = true
			tempA = tempA.Next
		}
		if tempB != nil {
			if _, ok := visited[tempB]; ok {
				final = tempB
				break
			}
			visited[tempB] = true
			tempB = tempB.Next
		}
	}
	return final
}

//双指针这才是名正言顺 链表可以想象为 A走了A步  B走了B步 如何相等 那就是A走过A+B的路 B走过B+A的路 这样A+B=B+A我们就相遇了
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var final *ListNode
	tempA := headA
	tempB := headB
	for {
		//如果相等 nil也是相等 就跳出
		if tempB == tempA {
			final = tempB
			break
		}
		//如果不为空就赋值 否则就赋B
		if tempA != nil {
			tempA = tempA.Next
		} else {
			tempA = headB
		}
		//如果不为空就赋值 否则就赋A
		if tempB != nil {
			tempB = tempB.Next
		} else {
			tempB = headA
		}
	}
	return final
}

//别指望用下一步去判断 如果相遇就在这一步呢? 另外 既然tempa可以作为返回值,就不要声明final了
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	tempA := headA
	tempB := headB
	for {
		//如果相等 nil也是相等 就跳出
		if tempB == tempA {
			break
		}
		//如果不为空就赋值 否则就赋B
		if tempA != nil {
			tempA = tempA.Next
		} else {
			tempA = headB
		}
		//如果不为空就赋值 否则就赋A
		if tempB != nil {
			tempB = tempB.Next
		} else {
			tempB = headA
		}
	}
	return tempA
}

func deleteNode1(head *ListNode, val int) *ListNode {
	//特例 要删除的就在头部
	if head.Val == val {
		return head.Next
	}
	pre, next := head, head
	for next != nil {
		//如果相等 就将上个pre的指针指向next的下一个指针 打到删除的效果
		if next.Val == val {
			pre.Next = next.Next
			break
		} else {
			//如果next的值不等,就移动指针
			pre = next
			next = next.Next
		}
	}
	return head
}

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] == 1 {
			return v
		} else {
			m[v]++
		}
	}
	return 0
}

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	//如果最后一个字母是大写的 那么前面的也必须都是大写的 否则就false
	if strings.ToUpper(string(word[len(word)-1])) == string(word[len(word)-1]) {
		for i := len(word) - 2; i >= 0; i-- {
			if !(strings.ToUpper(string(word[i])) == string(word[i])) {
				return false
			}
		}
		return true
	} else if strings.ToLower(string(word[0])) == string(word[0]) {
		//如果第一个字母是小写的 那么后面的也必须都是小写的 否则就false
		for i := 1; i < len(word); i++ {
			if !(strings.ToLower(string(word[i])) == string(word[i])) {
				return false
			}
		}
		return true
	} else {
		//否则 不满足前两种的话 就是最后一个字母是小写的 第一个字母是大写的 那就判断中间的是不是都是小写的
		for i := 1; i < len(word)-1; i++ {
			if !(strings.ToLower(string(word[i])) == string(word[i])) {
				return false
			}
		}
		return true
	}
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v] = v
	}
	for _, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{v, target - v}
		}
	}
	return []int{}
}

//没想到双指针的做法 但是做的时候发现 定义好刚刚开始的索引位置很重要！！！for的循环条件也很重要！
func twoSum1(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		}
	}
	return []int{}
}

func exchange(nums []int) []int {
	m := make(map[int]bool)
	final := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 1 {
			m[v] = true
			final = append(final, v)
		}
	}
	for _, v := range nums {
		if !m[v] {
			final = append(final, v)
		}
	}
	return final
}

//nums[i],nums[j] = nums[j],nums[i]  还有这么骚的swap !!!
func exchange1(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 == 0 || nums[j]%2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 1 {
			i++
		}
		if nums[j]%2 == 0 {
			j--
		}
	}
	return nums
}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	final := make([]int, 0, 10)
	var travel func(root *Node)
	travel = func(root *Node) {
		if root == nil {
			return
		}
		for _, child := range root.Children {
			travel(child)
		}
		final = append(final, root.Val)

	}
	travel(root)

	return final
}

//不能从数组中扣除的原因是 如果有3个相同的 删除了俩 还有留下1个
func firstUniqChar(s string) byte {
	m := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		if m[s[i]] == false {
			if i == len(s)-1 {
				return s[i]
			}
			for j := i + 1; j < len(s); j++ {
				if s[i] != s[j] && j == len(s)-1 {
					return s[i]
				}
				if s[i] == s[j] {
					m[s[i]] = true
					break
				}
			}
		}
	}
	return ' '
}

//如果前一个小于零 并且比后面的一个小 那么就给后一个 如果前一个小于零后面不如它小就给这一个 如果前一个为零后面大于0 就舍弃这个
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	m := make([]int, 0)
	a, b := 0, 1
	final := nums[0]
	m = append(m, final)
	for b <= len(nums)-1 {
		if final < 0 {
			if final < nums[b] {
				final = nums[b]
			}
			a++
			b++
			m = append(m, final)
		} else {
			final += nums[b]
			a++
			b++
			m = append(m, final)
		}
	}
	a = m[0]
	for _, v := range m {
		if v > a {
			a = v
		}
	}
	return a
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	final, b := nums[0], 1
	temp := nums[0]
	//定义一个指针
	for b <= len(nums)-1 {
		//如果temp小于零
		if temp < 0 {
			//如果还小于第b个数 那么就重新赋值给第b个数  否则就不动
			if temp < nums[b] {
				temp = nums[b]
			}
		} else {
			//如果temp>=0 就先加起来
			temp += nums[b]
		}
		//移动指针
		b++
		//判断加起来之后会不会比原来的大 会的话就更新 不会的话就不动原来的最大值
		if temp > final {
			final = temp
		}
	}
	return final
}

func add(a int, b int) int {
	return int(big.NewInt(int64(a)).Add(big.NewInt(int64(a)), big.NewInt(int64(b))).Int64())
}

//实在是不会 使用递归就不要用for  使用for就不要用递归 怎么两个混合这用???
//另外 位运算要学习一下 这题就相当于找规律了
func add1(a int, b int) int {
	var final int
	var travel func(a, b int)
	travel = func(a, b int) {
		for a&b != 0 {
			temp1 := (a & b) << 1
			temp2 := a ^ b
			a = temp1
			b = temp2
		}
		final = a ^ b
	}
	travel(a, b)
	return final
}

//会被知乎大佬笑话
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

//能不能选一种正确的数据结构来存储 易于比较的内中 比如 堆 或者算法 快速排序什么的
func getLeastNumbers1(arr []int, k int) []int {
	final := make([]int, 0, len(arr))
	final = append(final, arr[0])
	for i := 1; i < len(arr); i++ {
		for k := range final {
			if arr[i] <= final[k] {
				//如果直接赋值就完了呀 陷阱！！！ slice切片是引用类型 进行append会影响在后面的元素的！！！
				//需要使用copy 并且 copy的时候不能指定容量啊
				temp := make([]int, len(final[k:]), cap(final[k:]))
				copy(temp, final[k:])
				final = append(final[:k], arr[i])
				final = append(final, temp...)
				break
			} else if arr[i] > final[k] && k == len(final)-1 {
				final = append(final, arr[i])
			}
		}
	}
	return final[:k]
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	//如果A是空肯定不行 如果B是空 按照题意也不行
	if A == nil || B == nil {
		return false
	}
	//记录相同值的A树中的树节点指针
	final := make([]*TreeNode, 0)
	var find func(A *TreeNode)
	find = func(A *TreeNode) {
		if A == nil {
			return
		}
		if A.Val == B.Val {
			final = append(final, A)
		}
		find(A.Left)
		find(A.Right)
	}
	find(A)

	//判断是否相同
	var isEqual func(A *TreeNode, B *TreeNode) bool
	isEqual = func(A *TreeNode, B *TreeNode) bool {
		//这里相当关键 也就是当B是nil的时候返回true 原因在于B的树结构可能会缺失一部分 但不影响不缺失的部分的相等性质
		if B == nil {
			return true
		}
		if A == nil && B != nil {
			return false
		} else if A != nil && B == nil {
			return false
		} else if A == nil && B == nil {
			return true
		}
		if A.Val != B.Val {
			return false
		} else {
			//递归的比较左子树和右子树
			return isEqual(A.Left, B.Left) && isEqual(A.Right, B.Right)
		}
	}

	res := false
	//从满足条件的集合中判断是否相等 如果相等就return true
	for _, node := range final {
		res = isEqual(node, B)
		if res {
			return true
		}
	}
	return false
}

//辅助栈 简单的呜呜
func validateStackSequences(pushed []int, popped []int) bool {
	//如果长度为1就是true
	if len(pushed) == 1 {
		return true
	}
	//构造一个stack实际上还是切片
	stack := make([]int, 0, len(pushed))
	for len(pushed) != 0 {
		//往切片里塞值
		stack = append(stack, pushed[0])
		//去掉塞进切片的部分
		pushed = pushed[1:]
		//如果切片的末位值等于popped的首位 就说明可以弹出来了
		for len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			//弹出
			stack = stack[:len(stack)-1]
			//去掉弹出的部分
			popped = popped[1:]
		}
	}
	if len(stack) != 0 || len(popped) != 0 {
		return false
	}
	return true
}

func sumNums1(n int) int {
	return (n + int(math.Pow(float64(n), float64(2)))) >> 1
}

//利用短路思想来做 如果短路了 就不用执行&&之后的内容了 从而避免了无限递归
func sumNums(n int) int {
	sum := 0
	var travel func(n int) bool
	travel = func(n int) bool {
		sum += n
		return n > 0 && travel(n-1)
	}
	travel(n)
	return sum
}

func levelOrder1(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	b := true
	final := [][]int{}
	for len(queue) != 0 {

		m := make([]int, 0)
		l := len(queue)

		for i := 0; i < l; i++ {
			if queue[i] == nil {
				continue
			}
			m = append(m, queue[i].Val)
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}

		queue = queue[l:]
		if b {

			if len(m) != 0 {
				final = append(final, m)
			}
		} else {
			m1 := make([]int, 0, len(m))
			for i := len(m) - 1; i >= 0; i-- {
				m1 = append(m1, m[i])
			}
			if len(m1) != 0 {
				final = append(final, m1)
			}
		}
		b = !b
	}

	return final
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return [][]int{}
	}
	final := make([][]int, 0)

	var travel func(root *TreeNode, cur []int, value int)
	travel = func(root *TreeNode, cur []int, value int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if value+root.Val == target {
				cur = append(cur, root.Val)
				temp := make([]int, len(cur))
				copy(temp, cur)
				final = append(final, temp)
				return
			} else {
				return
			}
		}
		cur = append(cur, root.Val)
		travel(root.Left, cur, value+root.Val)
		travel(root.Right, cur, value+root.Val)
	}

	travel(root, []int{}, 0)
	return final
}

func permutation1(s string) []string {
	final := make([]string, 0, len(s))
	m := make(map[int]bool)
	check := make(map[string]bool)
	var temp string
	var travel func(s string)
	travel = func(s string) {
		for i, v := range s {
			if len(temp) == len(s) && !check[temp] {
				check[temp] = true
				final = append(final, temp)
				return
			}
			if m[i] == false {
				m[i] = true
				temp += fmt.Sprintf("%c", v)
				travel(s)
				m[i] = false
				temp = temp[:len(temp)-1]
			}
		}
	}
	travel(s)
	return final
}

func permutation2(s string) []string {
	final := make([]string, 0, len(s))
	check := make(map[string]bool)
	var temp string
	var travel func(s1 string)
	travel = func(s1 string) {
		for i, v := range s1 {
			temp += fmt.Sprintf("%c", v)
			if len(temp) == len(s) && !check[temp] {
				check[temp] = true
				final = append(final, temp)
			}

			travel(s1[:i] + s1[i+1:])
			temp = temp[:len(temp)-1]
		}
	}
	travel(s)
	return final
}
