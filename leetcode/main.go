package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//println(reverse(1563847412))
	//println(isPalindrome(1001))
	//fmt.Println(romanToInt("IV"))
	//println(romanToInt("MCMXCIV"))
	//println(climbStairs(5))
	//println(longestCommonPrefix([]string{"flower","flower","flower"}))
	//println(isValid("{[]}"))
	/*a:=&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val: 4,
				Next: nil,
			},
		},
	}
	b:=&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}*/
	//println(mergeTwoLists(a, b))
	//removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
	//println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
	//println(strStr("aaaaa","bba"))
	//println(searchInsert([]int{1,3,5,7}, 6))
	//println(maxSubArray([]int{-3,-2,0,-1}))
	//println(lengthOfLastWord("av "))
	//fmt.Println(plusOne([]int{9, 9}))

	//fmt.Println(addBinary("1010", "1011"))
	//duplicates := deleteDuplicates(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 1,
	//		Next: &ListNode{
	//			Val: 2,
	//			Next: &ListNode{
	//				Val: 3,
	//				Next: &ListNode{
	//					Val:  3,
	//					Next: nil,
	//				}},
	//		},
	//	},
	//})
	//for duplicates.Next != nil {
	//	fmt.Println(duplicates.Val)
	//	duplicates = duplicates.Next
	//}
	//fmt.Println(duplicates.Val)
	//merge([]int{4,5,6,0,0,0},3,[]int{1,2,3},3)

	//traversal := inorderTraversal1(&TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//})
	//fmt.Println(traversal)

	//result := isSameTree(&TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}, &TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//})
	//
	//fmt.Println(result)

	//i := maxDepth(&TreeNode{
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
	//fmt.Println(i)
	//depth := maxDepth(&TreeNode{
	//	Val:  1,
	//	Left: &TreeNode{
	//		Val:   5,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
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
	//})
	//fmt.Println(depth)

	//bst := sortedArrayToBST([]int{-10,-3,0,5,9})
	//fmt.Println(bst)

	//array := buildArray([]int{0, 2, 1, 5, 3, 4})
	//fmt.Println(array)

	i := permute([]int{1, 2, 3})
	fmt.Println(i)
}

//贪心算法或者动态规划算法好像在var全部小于0的时候实现不了 淦！！！
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var sum int
	var temp int
	//sum = nums[0]

	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if sum < temp {
			sum = temp
		} else if temp < 0 {
			temp = 0
		}
	}

	return sum
}

//从后向前 先找空 再找有的 减一下即可
func lengthOfLastWord(s string) int {
	if !strings.Contains(s, " ") {
		return len(s)
	}
	var first int
	var firstB bool
	var second int
	var secondB bool
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && !firstB {
			first = i
			firstB = true
		}
		if firstB && string(s[i]) == " " && !secondB {
			second = i
			secondB = true
		}
	}
	if !secondB {
		return first + 1
	}
	return first - second
}

//定义一个lastIndex如果进位就++ 并且使用lastIndex来寻找进位的位置在哪里
func plusOne(digits []int) []int {
	if len(digits) == 1 && digits[0] == 0 {
		return []int{1}
	}

	lastIndex := 0
	last := digits[len(digits)-1]
	if last+1 == 10 {
		digits[len(digits)-1] = 0
		lastIndex++
	} else {
		digits[len(digits)-1] += 1
	}
	for lastIndex > 0 {
		if lastIndex == len(digits) {
			break
		}
		digits[len(digits)-1-lastIndex]++
		if digits[len(digits)-1-lastIndex] == 10 {
			digits[len(digits)-1-lastIndex] = 0
			lastIndex++
		} else {
			lastIndex = 0
		}
	}
	if digits[0] == 0 {
		ret := []int{1}
		ret = append(ret, digits...)
		return ret
	} else {
		return digits
	}
}

//不能尝试转换为int 题目限制长度可以是10的4次方 一转就坏了
//更好的方法是不用判断哪个长度大 直接用j i 作为限制 i>=0 || j>=0 来控制就好了
func addBinary(a string, b string) string {
	aLength := len(a)
	bLength := len(b)
	if aLength > bLength {
		temp := a
		a = b
		b = temp
		tempLength := aLength
		aLength = bLength
		bLength = tempLength
	}

	up := uint8(0)
	final := ""
	for i := 0; i <= aLength-1; i++ {
		intb := b[bLength-1-i] - '0'
		inta := a[aLength-1-i] - '0'
		temp1 := intb + inta + up
		if temp1 >= 2 {
			final += strconv.Itoa(int(temp1) - 2)
			up = 1
		} else {
			final += strconv.Itoa(int(temp1))
			up = 0
		}
	}

	for i := bLength - aLength - 1; i >= 0; i-- {
		upu := b[i] - '0'
		temp2 := upu + up
		if temp2 == 2 {
			final += strconv.Itoa(int(temp2) - 2)
			up = 1
		} else {
			final += strconv.Itoa(int(temp2))
			up = 0
		}
	}

	if up == 1 {
		final += "1"
	}

	var finalfinal string
	for i := len(final) - 1; i >= 0; i-- {
		finalfinal += string(final[i])
	}
	return finalfinal

}

type ListNode1 struct {
	Val  int
	Next *ListNode
}

//83  这题不会啊 还是要多练 头指针保存下来 用别的变量来改变指针里的内容
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//这里发现自己冒泡排序和插入排序一点都不会 另外可以利用排序好的特点 使用双指针来往新的数组里插入东西 先判断再插入
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		//nums1 = nums2
		copy(nums1, nums2)
		return
	} else if n == 0 {
		return
	} else {
		ints := make([]int, m+n, m+n)
		ints = nums1[:m]
		ints = append(ints, nums2[:n]...)
		//插入排序
		for i := 1; i < len(ints); i++ {
			for j := 0; j < i; j++ {
				if ints[i] < ints[j] {
					temp := ints[j]
					ints[j] = ints[i]
					ints[i] = temp
				}
			}
		}

		//冒泡排序
		//for i := 0; i < len(ints); i++ {
		//	for j := i+1; j < len(ints); j++ {
		//		if ints[i]>ints[j] {
		//			temp:=ints[j]
		//			ints[j] = ints[i]
		//			ints[i] = temp
		//		}
		//	}
		//}
		copy(nums1, ints)
		//fmt.Println(ints)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var finalint = make([]int, 0, 5)
	if root.Left != nil {
		intLeft := bianli(root.Left)
		if intLeft != nil {
			finalint = append(finalint, intLeft...)
			temp = make([]int, 0, 5)
		}
	}
	finalint = append(finalint, root.Val)
	if root.Right != nil {
		intRight := bianli(root.Right)
		if intRight != nil {
			finalint = append(finalint, intRight...)
			temp = make([]int, 0, 5)
		}
	}
	return finalint
}

var temp = make([]int, 0, 5)

func bianli(son *TreeNode) []int {
	if son.Left != nil {
		bianli(son.Left)
	}
	temp = append(temp, son.Val)
	if son.Right != nil {
		bianli(son.Right)
	}
	return temp
}

//能够用迭代来做呢 递归大家都会
func inorderTraversal1(root *TreeNode) []int {
	var final = make([]int, 0, 5)
	var inorder func(son *TreeNode)
	inorder = func(son *TreeNode) {
		if son == nil {
			return
		}
		if son.Left != nil {
			inorder(son.Left)
		}
		final = append(final, son.Val)

		if son.Right != nil {
			inorder(son.Right)
		}
	}

	inorder(root)

	return final
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//用递归 还用的不够好啊
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	var bianli func(a *TreeNode, b *TreeNode) bool

	var inout bool

	bianli = func(a *TreeNode, b *TreeNode) bool {
		inout = false
		if a.Val != b.Val {
			return false
		}

		if (a.Left == nil && b.Left != nil) || (a.Left != nil && b.Left == nil) || (a.Right == nil && b.Right != nil) || (a.Right != nil && b.Right == nil) {
			return false
		}

		if a.Left == nil && b.Right == nil && a.Right == nil && b.Right == nil {
			inout = true
			return true
		}

		var result1 bool
		if a.Left != nil && b.Left != nil {
			inout = true
			result1 = bianli(a.Left, b.Left)
		} else if a.Left == nil && b.Left == nil {
			result1 = true
		}
		var result2 bool
		if a.Right != nil && b.Right != nil {
			inout = true
			result2 = bianli(a.Right, b.Right)
		} else if a.Right == nil && b.Right == nil {
			result2 = true
		}
		if inout {
			return result1 && result2
		} else {
			return a.Val == b.Val && a.Left == nil && b.Right == nil && a.Left == nil && b.Right == nil
		}
	}

	var inorNot bool
	var bLeft bool

	if p.Left != nil && q.Left != nil {
		inorNot = true
		bLeft = bianli(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		bLeft = true
	}
	var bRight bool
	if p.Right != nil && q.Right != nil {
		inorNot = true
		bRight = bianli(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		bRight = true
	}

	if inorNot {
		return bLeft && bRight && q.Val == p.Val
	} else {
		return q.Val == p.Val && p.Left == nil && p.Right == nil && q.Left == nil && q.Right == nil
	}
}

//递归用的还不够好唉
func maxDepth(root *TreeNode) int {
	i := 1
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return maxDepth(root.Right) + i
	} else if root.Left != nil && root.Right == nil {
		return maxDepth(root.Left) + i
	} else {
		x := maxDepth(root.Left) + i
		y := maxDepth(root.Right) + i
		if x >= y {
			return x
		} else {
			return y
		}
	}
}

//要很清楚的知道如何构建二叉树 高度相差1
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	binaryIndex := len(nums) / 2
	var treenode TreeNode
	treenode.Val = nums[binaryIndex]
	if len(nums) == 1 {
		return &treenode
	} else {
		left := sortedArrayToBST(nums[0:binaryIndex])
		var right *TreeNode
		if binaryIndex+1 >= len(nums) {
			right = nil
		} else {
			right = sortedArrayToBST(nums[binaryIndex+1:])
		}
		treenode.Left = left
		treenode.Right = right
		return &treenode
	}
}

func buildArray(nums []int) []int {
	if nums == nil {
		return nil
	}
	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[nums[i]])
	}
	return result
}
