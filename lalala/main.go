package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//a := Node{
	//	Val:    7,
	//	Next:   nil,
	//	Random: nil,
	//}
	//b := Node{
	//	Val:    13,
	//	Next:   nil,
	//	Random: nil,
	//}
	//c := Node{
	//	Val:    11,
	//	Next:   nil,
	//	Random: nil,
	//}
	//d := Node{
	//	Val:    10,
	//	Next:   nil,
	//	Random: nil,
	//}
	//e := Node{
	//	Val:    1,
	//	Next:   nil,
	//	Random: nil,
	//}
	//a.Next = &b
	//b.Next = &c
	//c.Next = &d
	//d.Next = &e
	//b.Random = &a
	//c.Random = &e
	//d.Random = &c
	//e.Random = &a
	//list := copyRandomList(&a)
	//fmt.Println(list)

	//value := maxValue([][]int{{5, 0, 1, 1, 2, 1, 0, 1, 3, 6, 3, 0, 7, 3, 3, 3, 1}, {1, 4, 1, 8, 5, 5, 5, 6, 8, 7, 0, 4, 3, 9, 9, 6, 0}, {2, 8, 3, 3, 1, 6, 1, 4, 9, 0, 9, 2, 3, 3, 3, 8, 4}, {3, 5, 1, 9, 3, 0, 8, 3, 4, 3, 4, 6, 9, 6, 8, 9, 9}, {3, 0, 7, 4, 6, 6, 4, 6, 8, 8, 9, 3, 8, 3, 9, 3, 4}, {8, 8, 6, 8, 3, 3, 1, 7, 9, 3, 3, 9, 2, 4, 3, 5, 1}, {7, 1, 0, 4, 7, 8, 4, 6, 4, 2, 1, 3, 7, 8, 3, 5, 4}, {3, 0, 9, 6, 7, 8, 9, 2, 0, 4, 6, 3, 9, 7, 2, 0, 7}, {8, 0, 8, 2, 6, 4, 4, 0, 9, 3, 8, 4, 0, 4, 7, 0, 4}, {3, 7, 4, 5, 9, 4, 9, 7, 9, 8, 7, 4, 0, 4, 2, 0, 4}, {5, 9, 0, 1, 9, 1, 5, 9, 5, 5, 3, 4, 6, 9, 8, 5, 6}, {5, 7, 2, 4, 4, 4, 2, 1, 8, 4, 8, 0, 5, 4, 7, 4, 7}, {9, 5, 8, 6, 4, 4, 3, 9, 8, 1, 1, 8, 7, 7, 3, 6, 9}, {7, 2, 3, 1, 6, 3, 6, 6, 6, 3, 2, 3, 9, 9, 4, 4, 8}})
	//fmt.Println(value)

	//number := nthUglyNumber(259)
	//fmt.Println(number)

	//location, _ := time.LoadLocation("Asia/Shanghai")
	//date := time.Date(2019, 11, 9, 14, 13, 24, 0, location)
	//t2 := time.Date(2019, 11, 11, 9, 13, 50, 0, location)
	//sub := date.Sub(t2)
	//println(sub.Hours())
	//println(sub.Minutes())
	//fmt.Println(sub)
	//t := &TreeNode{

	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   7,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//bst := insertIntoBST(t, 5)
	//fmt.Println(bst)
	//t3 := &TreeNode{
	//	Val:   7,
	//	Left:  nil,
	//	Right: nil,
	//}
	//t2 := &TreeNode{
	//	Val:   6,
	//	Left:  nil,
	//	Right: t3,
	//}
	//
	//t := &TreeNode{
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
	//	Right: t2,
	//}
	//
	////node := deleteNode(t, 3)
	////fmt.Println(node)
	//
	//nodes := countNodes(t)
	//fmt.Println(nodes).

	//ser := Constructor()
	//serialize := ser.serialize(t)
	//fmt.Println(serialize)
	//deserialize := ser.deserialize(serialize)
	//fmt.Println(deserialize)
	//
	//t2 := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//
	//serialize1 := ser.serialize(t2)
	//fmt.Println(serialize1)
	//deserialize1 := ser.deserialize(serialize1)
	//fmt.Println(deserialize1)

	//t2 := &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val:   -7,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: -3,
	//		Left: &TreeNode{
	//			Val: -9,
	//			Left: &TreeNode{
	//				Val: 9,
	//				Left: &TreeNode{
	//					Val: 6,
	//					Left: &TreeNode{
	//						Val:  0,
	//						Left: nil,
	//						Right: &TreeNode{
	//							Val:   -1,
	//							Left:  nil,
	//							Right: nil,
	//						},
	//					},
	//					Right: &TreeNode{
	//						Val: 6,
	//						Left: &TreeNode{
	//							Val:   -4,
	//							Left:  nil,
	//							Right: nil,
	//						},
	//						Right: nil,
	//					},
	//				},
	//				Right: nil,
	//			},
	//			Right: &TreeNode{
	//				Val: -7,
	//				Left: &TreeNode{
	//					Val: -6,
	//					Left: &TreeNode{
	//						Val:   5,
	//						Left:  nil,
	//						Right: nil,
	//					},
	//					Right: nil,
	//				},
	//				Right: &TreeNode{
	//					Val: -6,
	//					Left: &TreeNode{
	//						Val: 9,
	//						Left: &TreeNode{
	//							Val:   -2,
	//							Left:  nil,
	//							Right: nil,
	//						},
	//						Right: nil,
	//					},
	//					Right: nil,
	//				},
	//			},
	//		},
	//		Right: &TreeNode{
	//			Val: -3,
	//			Left: &TreeNode{
	//				Val:   -4,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//	},
	//}
	//ser := Constructor()
	//deser := Constructor()
	//data := ser.serialize(t)
	//fmt.Println(data)
	//ans := deser.deserialize(data)
	//fmt.Println(ans)
	//ancestor := lowestCommonAncestor(t, t3, t2)
	//fmt.Println(ancestor)

	//i := fib(5)
	//fmt.Println(i)

	//number := singleNumber([]int{3, 4, 3, 3})
	//fmt.Println(number)

	//parenthesis := generateParenthesis(1)
	//fmt.Println(parenthesis)

	//number := minNumber([]int{3, 30, 34, 5, 9})
	//fmt.Println(number)

	//sum := getSum(1, 3)
	//fmt.Println(sum)
	//l := &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val: 9,
	//				Next: &ListNode{
	//					Val: 9,
	//					Next: &ListNode{
	//						Val: 9,
	//						Next: &ListNode{
	//							Val:  9,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//
	//l2 := &ListNode{
	//	Val: 9,
	//	Next: &ListNode{
	//		Val: 9,
	//		Next: &ListNode{
	//			Val: 9,
	//			Next: &ListNode{
	//				Val:  9,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//numbers := addTwoNumbers(l, l2)
	//
	//fmt.Println(numbers)

	//profit := maxProfit([]int{7, 6, 4, 3, 1})
	//fmt.Println(profit)

	//num := translateNum(506)
	//fmt.Println(num)
	//cost := minimumCost([]int{1, 2, 3})
	//fmt.Println(cost)

	//arrays := numberOfArrays([]int{3, -4, 5, 1, -2}, -4, 5)
	//fmt.Println(arrays)

	items := highestRankedKItems([][]int{{1, 1, 1}, {0, 0, 1}, {2, 3, 4}}, []int{2, 3}, []int{0, 0}, 3)
	fmt.Println(items)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//这里最容易错的地方在于 1.原链表的指针直接赋值给现在链表的random或者next 2.改变了原来链表的值或者指针
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	m1 := make(map[*Node]int)
	m2 := make(map[int]*Node)
	index := 0
	n := new(Node)
	temp := n
	temp1 := head
	for temp1 != nil {
		m1[temp1] = index
		newNode := new(Node)
		newNode.Val = temp1.Val
		temp.Next = newNode

		temp = temp.Next
		m2[index] = temp
		index++

		temp1 = temp1.Next
	}

	temp1 = head
	temp = n.Next
	for temp1 != nil {
		if temp1.Random == nil {
			temp.Random = nil
		} else {
			temp.Random = m2[m1[temp1.Random]]
		}
		temp1 = temp1.Next
		temp = temp.Next
	}

	return n.Next
}

//这里最容易错的地方在于 1.原链表的指针直接赋值给现在链表的random或者next 2.改变了原来链表的值或者指针
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m1 := make(map[*Node]*Node)
	cur := head
	final := &Node{
		Val:    0,
		Next:   nil,
		Random: nil,
	}
	cur1 := final
	for cur != nil {
		temp := &Node{Val: cur.Val}
		cur1.Next = temp
		cur1 = cur1.Next
		m1[cur] = cur1
		cur = cur.Next
	}

	cur = head
	cur1 = final.Next

	for cur != nil {
		if cur.Random == nil {
			cur1.Random = nil
		} else {
			cur1.Random = m1[cur.Random]
		}
		cur1 = cur1.Next
		cur = cur.Next
	}

	return final.Next
}

//递归超时
func maxValue1(grid [][]int) int {
	var travel func(sum, i, j int) int
	travel = func(sum, i, j int) int {
		if i > len(grid)-1 || j > len(grid[0])-1 {
			return sum
		} else {
			sum += grid[i][j]
		}
		return int(math.Max(float64(travel(sum, i+1, j)), float64(travel(sum, i, j+1))))
	}
	return travel(0, 0, 0)
}

//还是超时
func maxValue2(grid [][]int) int {
	final := 0
	var travel func(sum, i, j int)
	travel = func(sum, i, j int) {
		if i > len(grid)-1 || j > len(grid[0])-1 {
			if i != len(grid)-1 && j != len(grid[0])-1 {
				return
			} else {
				if final < sum {
					final = sum
				}
				return
			}
		} else {
			sum += grid[i][j]
		}
		travel(sum, i+1, j)
		travel(sum, i, j+1)
	}
	travel(0, 0, 0)
	return final
}

//动态规划 实在给搞蒙了
func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] = grid[i][j-1] + grid[i][j]
				continue
			} else if j == 0 {
				grid[i][j] = grid[i-1][j] + grid[i][j]
				continue
			}
			grid[i][j] = int(math.Max(float64(grid[i-1][j]), float64(grid[i][j-1]))) + grid[i][j]
		}
	}
	return grid[m-1][n-1]
}

func nthUglyNumber(n int) int {
	m := make([]int, 0)
	f := make([]int, 0)
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 3
	} else if n == 4 {
		return 4
	} else if n == 5 {
		return 5
	}
	m = append(m, 2)
	m = append(m, 3)
	m = append(m, 4)
	m = append(m, 5)
	i := 6
	for len(m)+1 != n {
		b := false
		b1 := false
		if len(f) > 0 {
			for _, v := range f {
				if i%v == 0 {
					i++
					b1 = true
					break
				}
			}
		}
		if b1 {
			continue
		}
		for _, v := range m {
			if i%v == 0 {
				b = true
				m = append(m, i)
				i++
				break
			}
		}
		if !b {
			f = append(f, i)
			i++
		}
	}
	return i - 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	var travel func(root *TreeNode, min *TreeNode, max *TreeNode) bool
	travel = func(root *TreeNode, min *TreeNode, max *TreeNode) bool {
		if root == nil {
			return true
		}
		if min != nil && root.Val <= min.Val {
			return false
		}
		if max != nil && root.Val >= max.Val {
			return false
		}
		return travel(root.Left, min, root) && travel(root.Right, root, max)
	}
	return travel(root, nil, nil)
}

func searchBST1(root *TreeNode, val int) *TreeNode {

	var travel func(root *TreeNode, val int) *TreeNode
	travel = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return nil
		} else if root.Val == val {
			return root
		}
		node := travel(root.Left, val)
		if node != nil {
			return node
		} else {
			return travel(root.Right, val)
		}
	}
	return travel(root, val)
}

//利用二叉搜索树的性质 去掉一边
func searchBST(root *TreeNode, val int) *TreeNode {

	var travel func(root *TreeNode, val int) *TreeNode
	travel = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return nil
		} else if root.Val == val {
			return root
		}

		if val > root.Val {
			return travel(root.Right, val)
		} else {
			return travel(root.Left, val)
		}
	}
	return travel(root, val)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {

	var travel func(root *TreeNode, val int) *TreeNode
	travel = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		} else if root.Val == val {
			return root
		} else if root.Val < val {
			root.Right = travel(root.Right, val)
		} else {
			root.Left = travel(root.Left, val)
		}
		return root
	}
	return travel(root, val)
}

//go中如果已经分配了内存地址 不能直接赋值a=nil 这样是不行的 因此要知道他的上级 通过上级.right = nil 来间接赋值
func deleteNode(root *TreeNode, key int) *TreeNode {
	var getRightMin func(root *TreeNode) *TreeNode
	getRightMin = func(root *TreeNode) *TreeNode {
		if root.Left != nil {
			return getRightMin(root.Left)
		} else {
			return root
		}
	}

	if root == nil {
		return nil
	}
	if root.Val == key {
		if root == nil {
			return nil
		} else if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Right == nil {
			return root.Left
		} else if root.Left == nil {
			return root.Right
		} else {
			MinNode := getRightMin(root.Right)
			root.Val = MinNode.Val
			//这里究极无敌他妈的重要
			root.Right = deleteNode(root.Right, MinNode.Val)
			return root
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

//普通的二叉树的节点计算 复杂度可能就是遍历所有吧 就是oN呗
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

//利用完全二叉树的性质 完全二叉树的一个分支 肯定有一个是满二叉树 也就是PerfectBinaryTree 满二叉树的节点数量其实就是高度的指数关系
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root.Left, root.Right
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}

	if hl == hr {
		return int(math.Pow(2.0, float64(1+hl)) - 1)
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}

type Codec struct {
	slice []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string. 前序遍历
func (this *Codec) serialize1(root *TreeNode) string {
	var final string
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			final += "# "
			return
		} else {
			final += fmt.Sprintf("%v ", root.Val)
		}

		travel(root.Left)
		travel(root.Right)
	}

	travel(root)
	return final
}

// Deserializes your encoded data to tree. 前序遍历
func (this *Codec) deserialize1(data string) *TreeNode {
	data = data[:len(data)-1]
	this.slice = append(this.slice, strings.Split(data, " ")...)
	var travel func() *TreeNode
	travel = func() *TreeNode {
		if len(this.slice) > 0 {
			s := this.slice[0]
			this.slice = this.slice[1:]
			if s != "#" {
				i, _ := strconv.Atoi(s)
				t := &TreeNode{
					Val: i,
				}
				t.Left = travel()
				t.Right = travel()
				return t
			}
			return nil
		}
		return nil
	}
	return travel()
}

// Serializes a tree to a single string. 后序遍历
func (this *Codec) serialize2(root *TreeNode) string {
	var final string
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			final += "# "
			return
		}
		travel(root.Left)
		travel(root.Right)
		final += fmt.Sprintf("%v ", root.Val)
	}
	travel(root)
	return final
}

// Deserializes your encoded data to tree. 后序遍历  后续遍历生成树的时候 要先找到根节点 这很重要 最后一位是根节点 从后往前 先生成右子树 再生成左子树
func (this *Codec) deserialize2(data string) *TreeNode {
	data = data[:len(data)-1]
	this.slice = append(this.slice, strings.Split(data, " ")...)

	var travel func() *TreeNode
	travel = func() *TreeNode {
		if len(this.slice) == 0 {
			return nil
		}
		s := this.slice[len(this.slice)-1]
		this.slice = this.slice[:len(this.slice)-1]
		if s == "#" {
			return nil
		}
		atoi, _ := strconv.Atoi(s)
		t := &TreeNode{
			Val:   atoi,
			Left:  nil,
			Right: nil,
		}
		t.Right = travel()
		t.Left = travel()
		return t
	}

	return travel()
}

// Serializes a tree to a single string. 层序遍历
func (this *Codec) serialize(root *TreeNode) string {
	var final string
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		t := queue[0]
		queue = queue[1:]
		if t == nil {
			final += "# "
			continue
		}
		queue = append(queue, t.Left)
		queue = append(queue, t.Right)
		final += strconv.Itoa(t.Val) + " "
	}
	return final
}

// Deserializes your encoded data to tree. 层序遍历
func (this *Codec) deserialize(data string) *TreeNode {
	data = data[:len(data)-1]
	split := strings.Split(data, " ")
	queue := make([]*TreeNode, 0)
	if split[0] == "#" {
		return nil
	}
	atoi, _ := strconv.Atoi(split[0])
	root := &TreeNode{Val: atoi}
	queue = append(queue, root)
	split = split[1:]
	for len(queue) != 0 {
		node := queue[0]
		if split[0] != "#" {
			i, _ := strconv.Atoi(split[0])
			t := &TreeNode{Val: i}
			queue = append(queue, t)
			node.Left = t
		} else {
			node.Left = nil
		}
		split = split[1:]

		if split[0] != "#" {
			i, _ := strconv.Atoi(split[0])
			t := &TreeNode{Val: i}
			queue = append(queue, t)
			node.Right = t
		} else {
			node.Right = nil
		}
		split = split[1:]

		queue = queue[1:]
	}

	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var travel func(root, p *TreeNode, q *TreeNode) *TreeNode
	travel = func(root, p *TreeNode, q *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		if root == q || root == p {
			return root
		}
		left := travel(root.Left, p, q)
		right := travel(root.Right, p, q)
		if left != nil && right != nil {
			return root
		} else if left == nil && right == nil {
			return nil
		} else if left != nil {
			return left
		} else {
			return right
		}
	}
	return travel(root, p, q)
}

func fib(n int) int {
	i := 3
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	m := make(map[int]int)
	m[0] = 0
	m[1] = 1
	m[2] = 1
	for i <= n {
		m[i] = (m[i-1] + m[i-2]) % 1000000007
		i++
	}
	return m[n]
}

func singleNumber(nums []int) int {
	final := 0
	m := make(map[int]int)
	sort.Ints(nums)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v < 3 {
			return k
		}
	}
	return final
}

//用left 和 right 来记录剩余的个税或者使用过的左右括号个数 递归体中先写上递归的结束条件
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	var travel func(s string, left, right int)
	travel = func(s string, left, right int) {
		if left > right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if len(s) == 2*n {
			res = append(res, s)
		}

		travel(s+"(", left-1, right)
		travel(s+")", left, right-1)
	}
	travel("", n, n)
	return res
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		a := fmt.Sprintf("%d%d", nums[i], nums[j])
		b := fmt.Sprintf("%d%d", nums[j], nums[i])
		return a < b
	})
	res := ""
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return res
}

func getSum1(a int, b int) int {
	count := 0
	for b > 0 {
		count++
		b--
	}
	for b < 0 {
		count--
		b++
	}
	for a > 0 {
		count++
		a--
	}
	for a < 0 {
		count--
		a++
	}
	return count
}

//利用位运算
func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left != right {
		if numbers[left]+numbers[right] == target {
			return []int{left, right}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}
	return []int{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var travel func(l *ListNode) *ListNode
	travel = func(l *ListNode) *ListNode {
		if l.Next == nil || l == nil {
			return l
		}
		node := travel(l.Next)
		l.Next.Next = l
		l.Next = nil
		return node
	}

	node := travel(l1)
	listNode := travel(l2)
	fmt.Println(node, listNode)
	return nil

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp1 := l1
	temp2 := l2
	final := new(ListNode)
	res := final
	b := false
	var travel func(l1, l2 *ListNode) *ListNode
	travel = func(l1, l2 *ListNode) *ListNode {
		if l1 == nil && l2 == nil {
			return nil
		} else if l1 != nil && l2 == nil {
			sum := 0
			if !b {
				sum = l1.Val
			} else {
				sum = l1.Val + 1
			}
			if sum >= 10 {
				b = true
				sum -= 10
			} else {
				b = false
			}
			return &ListNode{Val: sum}

		} else if l2 != nil && l1 == nil {
			sum := 0
			if !b {
				sum = l2.Val
			} else {
				sum = l2.Val + 1
			}
			if sum >= 10 {
				b = true
				sum -= 10
			} else {
				b = false
			}
			return &ListNode{Val: sum}
		} else {
			sum := 0
			if !b {
				sum = temp1.Val + temp2.Val
			} else {
				sum = temp1.Val + temp2.Val + 1
			}
			if sum >= 10 {
				b = true
				sum -= 10
			} else {
				b = false
			}
			return &ListNode{Val: sum}
		}
	}
	for temp1 != nil || temp2 != nil {
		res.Next = travel(temp1, temp2)
		if temp1 != nil {
			temp1 = temp1.Next
		}
		if temp2 != nil {
			temp2 = temp2.Next
		}
		res = res.Next
	}
	if b {
		res.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return final.Next
}

func maxProfit(prices []int) int {
	res := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				continue
			} else {
				temp := prices[j] - prices[i]
				if res < temp {
					res = temp
				}
			}
		}
	}
	return res
}

func translateNum(num int) int {
	itoa := strconv.Itoa(num)
	res := 0
	var travel func(s string)
	travel = func(s string) {
		if len(s) == 0 {
			res++
			return
		}
		travel(s[1:])
		if len(s) >= 2 && s[0] != '0' {
			atoi, _ := strconv.Atoi(s[:2])
			if atoi < 26 {
				travel(s[2:])
			}
		}

	}
	travel(itoa)
	return res
}

func minimumCost(cost []int) int {
	res := 0
	sort.Ints(cost)
	for i := len(cost) - 1; i >= 0; i -= 3 {
		if i >= 0 {
			res += cost[i]
		}
		if i-1 >= 0 {
			res += cost[i-1]
		}
	}
	return res
}

func numberOfArrays(differences []int, lower int, upper int) int {
	res := 0
	preSum := 0
	min := 0
	max := 0
	for i := 0; i < len(differences); i++ {
		preSum += differences[i]
		if preSum < min {
			min = preSum
		}
		if preSum > max {
			max = preSum
		}
	}

	if lower-min < lower || lower-min > upper || upper-max > upper || upper-max < lower || lower-min > upper-max {
		return 0
	}
	min = lower - min
	max = upper - max
	res = max - min + 1
	res = int(math.Abs(float64(res)))
	return res
}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	temp := make([][]int, 0)
	queue := make([][]int, 0)
	queue = append(queue, start)
	visited := make(map[[2]int]bool)
	for len(queue) != 0 {
		startup := queue[0]
		queue = queue[1:]
		if startup[0] < 0 || startup[0] >= len(grid) || startup[1] < 0 || startup[1] >= len(grid[0]) || grid[startup[0]][startup[1]] == 0 || visited[[2]int{startup[0], startup[1]}] {
			continue
		}

		visited[[2]int{startup[0], startup[1]}] = true

		if grid[startup[0]][startup[1]] >= pricing[0] && grid[startup[0]][startup[1]] <= pricing[1] {
			temp = append(temp, startup)
		}

		queue = append(queue, []int{startup[0] - 1, startup[1]})
		queue = append(queue, []int{startup[0] + 1, startup[1]})
		queue = append(queue, []int{startup[0], startup[1] - 1})
		queue = append(queue, []int{startup[0], startup[1] + 1})
	}
	sort.Slice(temp, func(i, j int) bool {
		if math.Abs(float64(start[0]-temp[i][0]+start[1]-temp[i][1])) != math.Abs(float64(start[0]-temp[j][0]+start[1]-temp[j][1])) {
			return math.Abs(float64(start[0]-temp[i][0]+start[1]-temp[i][1])) < math.Abs(float64(start[0]-temp[j][0]+start[1]-temp[j][1]))
		}
		if grid[temp[i][0]][temp[i][1]] != grid[temp[j][0]][temp[j][1]] {
			return grid[temp[i][0]][temp[i][1]] < grid[temp[j][0]][temp[j][1]]
		}
		if temp[i][0] != temp[j][0] {
			return temp[i][0] < temp[j][0]
		}
		return temp[i][1] < temp[j][1]
	})
	if len(temp) > k {
		temp = temp[:k]
	}
	return temp
}
