package main

import (
	"fmt"
	"math"
)

func main() {
	//depth1 := maxDepth1(&Node{
	//	Val:      1,
	//	Children: []*Node{{3, []*Node{{5, nil}, {6, nil}}}, {2, nil}, {4, nil}},
	//})
	//
	//fmt.Println(depth1)

	//balanced := isBalanced(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//			Left: &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val:   4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//})
	//
	//fmt.Println(balanced)

	//traversal := postorderTraversal(&TreeNode{
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
	//fmt.Println(traversal)

	//preorder := bstFromPreorder([]int{4, 2})
	//fmt.Println(preorder)

	//postorder := verifyPostorder([]int{1, 6, 3, 2, 5})
	//postorder := verifyPostorder([]int{4, 6, 7, 5})
	//fmt.Println(postorder)

	//sum := pathSum(&TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 4,
	//		Left: &TreeNode{
	//			Val:  11,
	//			Left: &TreeNode{Val: 7},
	//			Right: &TreeNode{
	//				Val:   2,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 8,
	//		Left: &TreeNode{
	//			Val:   13,
	//			Left:  nil,
	//			Right: nil,
	//		},
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
	//}, 22)
	//fmt.Println(sum)

	//sum := maxPathSum(&TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val:   -1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: nil,
	//})
	//
	//fmt.Println(sum)
	count := huiwenCount("a")
	fmt.Println(count)
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
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	if left > right {
		return left
	}

	return right
}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth1(root *Node) int {
	if root == nil {
		return 0
	}

	max := 1
	for _, child := range root.Children {
		temp := maxDepth1(child) + 1
		if max < temp {
			max = temp
		}
	}

	return max
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := travel(root.Left) + 1
		right := travel(root.Right) + 1
		if left > right {
			return left
		}
		return right
	}
	left := travel(root.Left)
	right := travel(root.Right)
	return (math.Abs(float64(right)-float64(left)) <= 1.0) && isBalanced(root.Left) && isBalanced(root.Right)
}

func invertTree(root *TreeNode) *TreeNode {
	var travel func(root *TreeNode) *TreeNode
	travel = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		left := travel(root.Left)
		right := travel(root.Right)

		root.Left = right
		root.Right = left
		return root
	}

	return travel(root)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var travel func(root, p, q *TreeNode) *TreeNode
	travel = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root != p && root != q {
			left := travel(root.Left, p, q)
			right := travel(root.Right, p, q)
			if left == nil && right == nil {
				return nil
			} else if left != nil && right == nil {
				return left
			} else if right != nil && left == nil {
				return right
			} else {
				return root
			}
		} else if root == p || root == q {
			return root
		} else {
			return nil
		}
	}
	node := travel(root, p, q)
	return node
}

func preorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		travel(root.Left)
		travel(root.Right)
	}
	travel(root)

	return res
}

//前序遍历 用栈来做
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {

		if root == nil {
			return
		}

		travel(root.Left)
		res = append(res, root.Val)
		travel(root.Right)
	}

	travel(root)
	return res
}

//用栈来做 而且还要分片来做  先传左边的 左边的结束了再传右边的
func inorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	cur := root
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		res = append(res, node.Val)
		stack = stack[:len(stack)-1]
		cur = node.Right
	}
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	cur := root
	var pre *TreeNode
	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		node := stack[len(stack)-1]
		if node.Right != nil && node.Right != pre {
			stack = append(stack, node.Right)
			cur = node.Right
		} else {
			res = append(res, node.Val)
			pre = node
			node = nil
		}
		stack = stack[:len(stack)-1]

	}
	return res
}

func postorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			res = append([]int{root.Val}, res...)
			root = root.Right
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		root = node.Left
	}

	return res
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			res = append(res, root.Val)
			root = root.Right
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Left
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0, 0)
	stack := make([]*TreeNode, 0, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	var travel func(preorder []int, postorder []int) *TreeNode

	travel = func(preorder []int, postorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		} else if len(preorder) == 1 {
			return &TreeNode{Val: preorder[0]}
		}

		m := make(map[int]int)
		for i := 0; i < len(postorder); i++ {
			m[postorder[i]] = i
		}

		pivot := m[preorder[1]]

		root := &TreeNode{Val: preorder[0]}
		root.Left = travel(preorder[1:1+pivot+1], postorder[:pivot+1])
		root.Right = travel(preorder[1+pivot+1:], postorder[pivot+1:len(postorder)-1])

		return root
	}
	return travel(preorder, postorder)
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	var travel func([]int, []int) *TreeNode
	travel = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		} else if len(preorder) == 1 {
			return &TreeNode{Val: preorder[0]}
		}

		root := &TreeNode{Val: preorder[0]}
		m := make(map[int]int)
		for i := 0; i < len(inorder); i++ {
			m[inorder[i]] = i
		}
		pivot := m[preorder[0]]

		root.Left = travel(preorder[1:pivot+1], inorder[:pivot])
		root.Right = travel(preorder[pivot+1:], inorder[pivot+1:])
		return root
	}

	return travel(preorder, inorder)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var travel func([]int, []int) *TreeNode

	travel = func(inorder []int, postorder []int) *TreeNode {
		if len(inorder) == 0 {
			return nil
		} else if len(inorder) == 1 {
			return &TreeNode{Val: inorder[0]}
		}

		root := &TreeNode{Val: postorder[len(postorder)-1]}
		m := make(map[int]int)
		for i := 0; i < len(inorder); i++ {
			m[inorder[i]] = i
		}
		pivot := m[postorder[len(postorder)-1]]
		root.Left = travel(inorder[:pivot], postorder[:pivot])
		root.Right = travel(inorder[pivot+1:], postorder[pivot:len(postorder)-1])
		return root
	}

	return travel(inorder, postorder)
}

func bstFromPreorder(preorder []int) *TreeNode {
	var travel func(preorder []int) *TreeNode
	travel = func(preorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		} else if len(preorder) == 1 {
			return &TreeNode{Val: preorder[0]}
		}

		root := &TreeNode{Val: preorder[0]}
		index := len(preorder)
		for i := 1; i < len(preorder); i++ {
			if preorder[i] > preorder[0] {
				index = i
				break
			}
		}

		root.Left = travel(preorder[1:index])
		root.Right = travel(preorder[index:])
		return root
	}

	return travel(preorder)
}

func verifyPostorder(postorder []int) bool {
	var travel func(postorder []int, min int, max int) bool
	travel = func(postorder []int, min int, max int) bool {
		if len(postorder) == 0 {
			return true
		} else if len(postorder) == 1 {
			if postorder[len(postorder)-1] > min && postorder[len(postorder)-1] < max {
				return true
			}
			return false
		}

		index := len(postorder) - 1 //待定
		for i := 0; i < len(postorder)-1; i++ {
			if postorder[i] > postorder[len(postorder)-1] {
				index = i
				break
			}
		}

		return (postorder[len(postorder)-1] > min && postorder[len(postorder)-1] < max) && travel(postorder[:index], min, postorder[len(postorder)-1]) && travel(postorder[index:len(postorder)-1], postorder[len(postorder)-1], max)
	}

	return travel(postorder, math.MinInt, math.MaxInt)
}

func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root
	last := root
	for node != nil {
		if node.Val > val {
			last = node
			node = node.Left
		} else {
			last = node
			node = node.Right
		}
	}
	if last.Val > val {
		last.Left = &TreeNode{Val: val}
	} else {
		last.Right = &TreeNode{Val: val}
	}

	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var travel func(root *TreeNode, val int) *TreeNode
	travel = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		}

		if root.Val > val {
			root.Left = travel(root.Left, val)
		} else {
			root.Right = travel(root.Right, val)
		}
		return root
	}
	return travel(root, val)
}

func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0, 0)
	var travel func(root *TreeNode, target int, ans []int)
	travel = func(root *TreeNode, target int, ans []int) {
		if root == nil {
			return
		}
		if target-root.Val == 0 && root.Left == nil && root.Right == nil {
			temp := make([]int, len(ans), cap(ans))
			copy(temp, ans)
			temp = append(temp, root.Val)
			res = append(res, temp)
			return
		}

		travel(root.Left, target-root.Val, append(ans, root.Val))
		travel(root.Right, target-root.Val, append(ans, root.Val))

	}
	travel(root, target, []int{})
	return res
}

func SearchRange1(root *TreeNode, k1 int, k2 int) []int {
	temp := make([]int, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		temp = append(temp, root.Val)
		travel(root.Right)
	}
	travel(root)
	res := make([]int, 0, 0)
	for i := 0; i < len(temp); i++ {
		if temp[i] >= k1 && temp[i] <= k2 {
			res = append(res, temp[i])
		}
	}

	return res
}

func SearchRange(root *TreeNode, k1 int, k2 int) []int {
	res := make([]int, 0, 0)

	var travel func(root *TreeNode, k1 int, k2 int)
	travel = func(root *TreeNode, k1 int, k2 int) {
		if root == nil {
			return
		}

		if root.Val > k1 {
			travel(root.Left, k1, k2)
		}

		if root.Val >= k1 && root.Val <= k2 {
			res = append(res, root.Val)
		}

		if root.Val < k2 {
			travel(root.Right, k1, k2)
		}
	}
	travel(root, k1, k2)

	return res
}

func levelOrder(root *Node) [][]int {
	res := make([][]int, 0, 0)
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		ans := make([]int, 0, 0)
		for i := 0; i < l; i++ {
			q := queue[0]
			ans = append(ans, q.Val)
			for _, child := range q.Children {
				queue = append(queue, child)
			}
			queue = queue[1:]
		}
		res = append(res, ans)
	}

	return res
}

func maxPathSum1(root *TreeNode) int {

	res := math.MinInt
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		ans := 0
		if root == nil {
			return ans
		}

		left := travel(root.Left)
		right := travel(root.Right)

		if root.Val+left+right > res {
			res = root.Val + left + right
		}

		if root.Val > res {
			res = root.Val
		}

		if left < 0 && right < 0 {
			ans = root.Val
		} else if left < right {
			ans = root.Val + right
		} else {
			ans = root.Val + left
		}

		if ans > res {
			res = ans
		}
		return ans
	}

	sum := travel(root)
	if sum > res {
		res = sum
	}

	return res
}

func maxPathSum(root *TreeNode) int {

	res := math.MinInt

	var max func(int, int) int
	max = func(i int, i2 int) int {
		if i > i2 {
			return i
		}
		return i2
	}
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := max(travel(root.Left), 0)
		right := max(travel(root.Right), 0)

		res = max(res, root.Val)
		res = max(res, left+right+root.Val)

		return max(left, right) + root.Val
	}

	travel(root)

	return res
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var travel func(node *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := travel(root.Left)
		right := travel(root.Right)
		if left == 0 && right == 0 {
			res = max(res, 0)
		} else if left == 0 {
			res = max(res, right+1)
		} else if right == 0 {
			res = max(res, left+1)
		} else {
			res = max(res, right+left+2)
		}

		return max(left, right) + 1
	}
	return travel(root)
}

func huiwenCount(s string) int {
	res := 0
	var travel func(i int)
	travel = func(i int) {
		res++
		l, r := i-1, i+1
		for l >= 0 && r < len(s) {
			if s[l] != s[r] {
				break
			}
			l--
			r++
			res++
		}
		//aba

		//abba
		if i+1 < len(s) && s[i] == s[i+1] {
			res++
			l, r = i-1, i+2
			for l >= 0 && r < len(s) {
				if s[l] != s[r] {
					break
				}
				l--
				r++
				res++
			}
		}
	}
	for i := 0; i < len(s); i++ {
		travel(i)
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func jiaodian(a *ListNode, b *ListNode) *ListNode {
	cur := a
	m := make(map[*ListNode]bool)
	for cur != nil {
		m[cur] = true
		cur = cur.Next
	}

	cur = b
	for cur != nil {
		if !m[b] {
			cur = cur.Next
		} else {
			return cur
		}
	}

	return nil
}
