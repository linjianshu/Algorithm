package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//stack := Constructor()
	//stack.Push(-2)
	//stack.Push(0)
	//stack.Push(-3)
	//min := stack.Min()
	//fmt.Println(min)
	//stack.Pop()
	//top := stack.Top()
	//fmt.Println(top)
	//i := stack.Min()
	//fmt.Println(i)

	//i:= search([]int{2,2}, 2)
	//fmt.Println(i)
	//array := minArray([]int{10, 10, 10, 1, 10})
	//fmt.Println(array)

	//straight := isStraight([]int{0,0,1,2,5})
	//fmt.Println(straight)
	//ways := numWays2(6)
	//fmt.Println(ways)

	//words := reverseWords2("  hello world!  ")
	//fmt.Println(len(words))
	//fmt.Println(words)

	//symmetric := isSymmetric(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
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
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
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
	//})
	//fmt.Println(symmetric)

	//depth := minDepth(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//})
	//fmt.Println(depth)

	//tree := buildTree([]int{1, 4, 2, 7, 5, 6}, []int{2, 7, 4, 1, 6, 5})
	//fmt.Println(tree)
	//matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//array := findNumberIn2DArray1(matrix, 5)
	//fmt.Println(array)

	//b := exist2([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	//b := exist2([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS")
	//b := exist2([][]byte{{'a', 'b'}}, "ba")
	//b := exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB")
	//b := exist([][]byte{{'a', 'b'}, {'c', 'd'}}, "abcd")
	//b := exist([][]byte{{'a'}}, "a")
	//fmt.Println(b)

	//count := movingCount3(16, 8, 4)
	//fmt.Println(count)

	//rope := cuttingRope(21)
	//fmt.Println(rope)
	//rope1 := cuttingRope1(21)
	//fmt.Println(rope1)

	//fmt.Printf("%b\n", 2.11)
	//fmt.Println(1 >> 1)
	//pow := myPow1(2.00000, 10)
	//fmt.Println(pow)

	number := isNumber("992700.5513600757592")
	fmt.Println(number)
}

// MinStack 要思考它的难点在哪里
type MinStack struct {
	a   []int
	min []int
}

// Constructor /** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		a:   []int{},
		min: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.a = append(this.a, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else if this.min[len(this.min)-1] >= x {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.a[len(this.a)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.a = this.a[:len(this.a)-1]
}

func (this *MinStack) Top() int {
	return this.a[len(this.a)-1]
}

func (this *MinStack) Min() int {
	return this.min[len(this.min)-1]
}

//看题目 有技巧的 什么是排序数组 你想什么鸡巴呢你
func search1(nums []int, target int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	return m[target]
}

func search2(nums []int, target int) int {
	count := 0

	var find func(middle int, b bool)
	find = func(middle int, b bool) {
		for {
			if middle < 0 || middle >= len(nums) {
				break
			}
			if nums[middle] == target {
				count++
				if b {
					middle++
				} else {
					middle--
				}
			} else if nums[middle] != target {
				break
			}
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			count++
			find(mid+1, true)
			find(mid-1, false)
			break
		}
	}

	return count

}

//所有的for都可以和递归替换 替换后直接速度快起来了 97.58 61.31
func search(nums []int, target int) int {
	count := 0

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			count++
			temp := mid + 1
			for {
				if temp <= len(nums)-1 {
					if nums[temp] != target {
						break
					} else {
						count++
						temp++
					}
				} else {
					break
				}
			}
			temp = mid - 1
			for {
				if temp >= 0 {
					if nums[temp] != target {
						break
					} else {
						count++
						temp--
					}
				} else {
					break
				}
			}
			break
		}
	}

	return count

}

func minArray1(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	right := len(numbers) - 1
	left := len(numbers) - 2
	for left >= 0 {
		if numbers[left] <= numbers[right] {
			left--
			right--
		} else {
			return numbers[right]
		}
	}
	return numbers[0]
}

//想到二分法 但是又放弃了 没想出来 字节的面试题哦  二分法又忘记了 二分法的精髓是 left right 不断的用mid去替换left或者right
//完全把labuladong讲的忘记了 需要这么麻烦吗???
func minArray2(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	left, right := 0, len(numbers)-1
	for left <= right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			if right == left+1 {
				return numbers[left]
			}
			right = mid
		} else if numbers[mid] > numbers[right] {
			if right == left+1 {
				return numbers[right]
			}
			left = mid
		} else if numbers[mid] == numbers[right] {
			right--
		}
	}
	return numbers[right+1]
}

func minArray(numbers []int) int {
	left, right := 0, len(numbers)-1
	//闭区间
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			//发现mid大于right的 肯定就不是mid了 再因为是闭区间 那就mid++
			left = mid + 1
		} else if numbers[mid] == numbers[right] {
			right--
		}
	}
	return numbers[left]
}

func isStraight(nums []int) bool {
	m := map[int]int{}
	min := 13
	max := 1
	for _, v := range nums {
		//重复的直接pass 不包含0
		if m[v] == 1 && v != 0 {
			return false
		}
		//求最大值
		if v >= max {
			max = v
		}
		//求不为0的最小值
		if v <= min && v != 0 {
			min = v
		}
		//记在map中
		m[v]++
	}

	//不存在大小王的情况下 肯定是23456  56789 首尾相差4
	if m[0] == 0 {
		if max-min != 4 {
			return false
		}
		return true
	} else {
		//有大小王情况下 超过4肯定pass 不超过4随便弥补
		//可以这么想 用反证法 A=>B  即 超过4=>pass
		//         也即 非B=>非A  即 不pass的都是不超过4的
		//		   随便来个顺子 肯定不pass 肯定都是首尾不超过4
		if max-min > 4 {
			return false
		}
		return true
	}
}

func missingNumber1(nums []int) int {
	var final int
	//如果对应索引和值不对应 就说明肯定是它
	for i := range nums {
		if nums[i] != i {
			return i
		}
		//如果找到最后都发现没有 就说明没有的值是n
		if i == len(nums)-1 {
			final = i + 1
		}
	}
	return final
}

//排序数组的查找 首选二分法 复杂度位logN  这里的关键在于 选到的mid如果索引不在和值不相等就说明 它本身或者它的左边有问题
//如果索引和值相等 就向右移动 说明左边的都没有问题 !!!
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] != mid {
			right = mid
			if left == right {
				return left
			}
		} else if nums[mid] == mid {
			left = mid + 1
		}
	}
	return left
}

//在if中多套了几个if 有问题有风险  边界还是不知道怎么处理  可以投机地试一下 -1 或者 +1
func missingNumber2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] != mid {
			right = mid - 1
		} else if nums[mid] == mid {
			left = mid + 1
		}
	}
	return left
}

func numWays(n int) int {
	//备忘录
	m := make(map[int]int)
	var travel func(n int) int
	travel = func(n int) int {
		if _, ok := m[n]; ok {
			return m[n]
		}

		if n == 0 {
			m[0] = 1
			return 1
		} else if n == 1 {
			m[1] = 1
			return 1
		} else if n == 2 {
			m[2] = 2
			return 2
		}
		m[n] = (travel(n-1) + travel(n-2)) % 1000000007
		//由于只要上一个和上上个数据 去掉多余的
		delete(m, n-2)
		return m[n]
	}
	return travel(n)
}

func numWays1(n int) int {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 1
	i := 2
	for i <= n {
		m[i] = (m[i-1] + m[i-2]) % 1000000007
		i++
	}
	return m[n]
}

func numWays2(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	}
	pre := 1
	cur := 1
	i := 2
	for i <= n {
		temp := cur
		cur = (cur + pre) % 1000000007
		pre = temp
		i++
	}
	return cur
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	var cur, count int
	var final string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 {
				continue
			} else {
				cur = i
				final += s[cur+1 : cur+1+count]
				final += " "
				count = 0
			}
		} else {
			cur = i
			count++
		}
	}
	if count != 0 {
		final += s[:count]
	} else if len(final) > 0 {
		final = final[:len(final)-1]
	}
	return final
}

func reverseWords1(s string) string {
	if len(s) == 0 {
		return ""
	}
	var count int
	var final string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 {
				continue
			} else {
				if len(final) > 0 {
					final += " "
				}
				final += s[i+1 : i+1+count]
				count = 0
			}
		} else {
			count++
		}
	}
	if count != 0 {
		if len(final) > 0 {
			final += " "
		}
		final += s[:count]
	}
	return final
}

func reverseWords2(s string) string {
	if len(s) == 0 {
		return ""
	}
	var final string
	split := strings.Split(s, " ")
	for _, v := range split {
		if v == "" {
			continue
		}
		if len(final) == 0 {
			final = v
		} else {
			final = v + " " + final
		}
	}
	return final
}

//func spiralOrder1(matrix [][]int) []int {
//	final := []int{}
//	if len(matrix) != 0 && len(matrix[0]) != 0 {
//		for i := 0; i < len(matrix); i++ {
//			if i == 0 {
//				final = append(final, matrix[0]...)
//				break
//			} else if i != len(matrix)-1 {
//				final = append(final, matrix[i][len(matrix[i])-1])
//			} else if i == len(matrix)-1 {
//				for j := len(matrix[i]) - 1; j >= 0; j-- {
//					final = append(final, matrix[i][j])
//				}
//			}
//		}
//		matrix = matrix[1 : len(matrix)-1]
//		for _, v := range matrix {
//			v = v[:len(v)-1]
//		}
//	}
//}

//func spiralOrder(matrix [][]int) []int {
//	if len(matrix) == 0 {
//		return []int{}
//	} else if len(matrix) == 1 {
//		return matrix[0]
//	}
//
//	final := []int{}
//	m := make(map[int]bool)
//	step := len(matrix[0])
//	init := []int{}
//	for _, v := range matrix {
//		init = append(init, v...)
//	}
//
//	i := 0
//	b1 := false
//	for i < len(init) {
//		if i < step {
//			final = append(final, init[i])
//			m[init[i]] = true
//			i++
//		} else if i%step == 0 {
//			i = i - 1 + step
//			b1 = true
//		} else if (i+1)%step == 0 {
//			final = append(final, init[i])
//			i = i + step
//			if i == len(init)-1 {
//				b1 = true
//			}
//		}
//	}
//	for i >= step {
//		for j := 0; j < step-1; j++ {
//			i--
//			final = append(final, init[i])
//		}
//		if (i+1)%step == 0 {
//			i = i + 1 - step
//		} else if i%step == 0 {
//			final = append(final, init[i])
//			i = i - step
//		}
//	}
//	i++
//
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Right != nil && root.Left == nil {
		return false
	}

	left := make([]int, 0)
	right := make([]int, 0)
	var travel func(root *TreeNode, b bool)
	travel = func(root *TreeNode, b bool) {
		if root != nil {
			if b {
				travel(root.Left, b)
				travel(root.Right, b)
				left = append(left, root.Val)

			} else {
				travel(root.Right, b)
				travel(root.Left, b)
				right = append(right, root.Val)
			}
		}
		if root == nil {
			if b {
				left = append(left, 0)
			} else {
				right = append(right, 0)
			}
		}
	}
	travel(root.Left, true)
	travel(root.Right, false)
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

//递归 优化内存 因为原本我以为定义递归很难有返回值 但是现在可以将返回值放在bool中
//只需要将递归方法的返回值加上&&连接起来就可以了
func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var travel func(a, b *TreeNode) bool
	travel = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		} else if a == nil && b != nil {
			return false
		} else if a != nil && b == nil {
			return false
		}
		return a.Val == b.Val && travel(a.Left, b.Right) && travel(a.Right, b.Left)
	}
	return travel(root.Left, root.Right)
}

//传递的时候 不能修改值然后传入 不然返回的时候恢复值很难受的啊 !
//targetSum-root.Val例如这样传值 返回的时候就能直接拿到targetSum的值 不需要+root.val才能拿到原始值!!!!
func hasPathSum(root *TreeNode, targetSum int) bool {
	var travel func(root *TreeNode, targetSum int) bool
	travel = func(root *TreeNode, targetSum int) bool {
		if root == nil {
			return false
		}
		if root.Val == targetSum && root.Right == nil && root.Left == nil {
			return true
		}
		return travel(root.Left, targetSum-root.Val) || travel(root.Right, targetSum-root.Val)
	}
	return travel(root, targetSum)
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	var travel func(root *TreeNode, final int) bool
	travel = func(root *TreeNode, final int) bool {
		if root == nil {
			return false
		}
		if final+root.Val == targetSum && root.Left == nil && root.Right == nil {
			return true
		}
		return travel(root.Left, final+root.Val) || travel(root.Right, final+root.Val)
	}
	return travel(root, 0)
}

func minDepth(root *TreeNode) int {
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil || (root.Left == nil && root.Right == nil) {
			return 0
		}
		if root.Left == nil {
			return travel(root.Right) + 1
		} else if root.Right == nil {
			return travel(root.Left) + 1
		}
		return min(travel(root.Left)+1, travel(root.Right)+1)
	}

	if root == nil {
		return 0
	}

	return travel(root) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		m[inorder[i]] = i
	}

	var recur func(root, start, end int) *TreeNode
	recur = func(root, start, end int) *TreeNode {
		if start == end {
			return &TreeNode{Val: inorder[start]}
		}
		if start > end {
			return nil
		}
		t := new(TreeNode)
		t.Val = preorder[root]
		t.Left = recur(root+1, start, m[preorder[root]]-1)
		t.Right = recur(root+1, end, m[preorder[root-1]]-1)
		return t
	}
	var t TreeNode
	t.Val = preorder[0]
	t.Left = recur(1, 0, m[preorder[0]]-1)
	t.Right = recur(1, m[preorder[0]]+1, len(preorder)-1)

	return &t
}

//想法还不够巧妙 这里已经用了限制的框框框起来了 然后再框框里去遍历
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	rows := matrix[0]
	columns := make([]int, 0, len(matrix))
	for i := 0; i < len(matrix); i++ {
		columns = append(columns, matrix[i][0])
	}
	var right, bottom = len(rows) - 1, len(columns) - 1
	for i := 0; i < len(rows); i++ {
		if rows[i] == target {
			return true
		}
		if rows[i] > target {
			right = i - 1
			break
		}
	}
	for i := 1; i < len(columns); i++ {
		if columns[i] == target {
			return true
		}
		if columns[i] > target {
			bottom = i - 1
			break
		}
	}
	for i := 1; i <= bottom; i++ {
		for j := 0; j <= right; j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

//利用左下角的性质来干
func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	i := len(matrix) - 1
	j := 0
	for j <= len(matrix[0])-1 && i >= 0 {
		temp := matrix[i][j]
		if temp > target {
			i--
		} else if temp < target {
			j++
		} else if temp == target {
			return true
		}
	}
	return false
}

//我感觉关键在于备忘录的操作和如果不行的话就恢复到初始状态的操作
func exist(board [][]byte, word string) bool {
	m := make(map[[2]int]bool)
	result := [][2]int{}
	Sub := 1

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				result = append(result, [2]int{i, j})
			}
		}
	}

	var search = func(i, j, index int) bool {
		if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
			if board[i][j] == word[index] && m[[2]int{i, j}] == false {
				return true
			}
		}
		return false
	}

	var judge func(point [2]int) bool
	judge = func(point [2]int) bool {
		if Sub == len(word) {
			return true
		}
		m[[2]int{point[0], point[1]}] = true
		tempresult := make([][2]int, 0)
		if search(point[0]-1, point[1], Sub) {
			tempresult = append(tempresult, [2]int{point[0] - 1, point[1]})
		}
		if search(point[0], point[1]-1, Sub) {
			tempresult = append(tempresult, [2]int{point[0], point[1] - 1})
		}
		if search(point[0], point[1]+1, Sub) {
			tempresult = append(tempresult, [2]int{point[0], point[1] + 1})
		}
		if search(point[0]+1, point[1], Sub) {
			tempresult = append(tempresult, [2]int{point[0] + 1, point[1]})
		}

		Sub++

		for i := 0; i < len(tempresult); i++ {
			if judge(tempresult[i]) {
				return true
			}
		}
		//如果不行的话 要恢复为原来的状态
		m[[2]int{point[0], point[1]}] = false
		Sub--
		return false
	}

	for i := 0; i < len(result); i++ {
		m = make(map[[2]int]bool)
		m[result[i]] = true
		if judge(result[i]) {
			return true
		}
	}
	return false
}

func exist1(board [][]byte, word string) bool {
	m := make(map[[2]int]bool)
	result := [][2]int{}
	char := word[0]
	Sub := 1
	point := [2]int{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == char {
				result = append(result, [2]int{i, j})
			}
		}
	}

	var search = func(i, j, index int) bool {
		if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
			if board[i][j] == word[index] && m[[2]int{i, j}] == false {
				return true
			}
		}
		return false
	}

	var judge func(point [2]int) bool
	judge = func(point [2]int) bool {
		if Sub == len(word) {
			return true
		}
		row := point[0]
		column := point[1]
		m[[2]int{row, column}] = true
		tempresult := make([][2]int, 0)
		if search(row-1, column, Sub) {
			tempresult = append(tempresult, [2]int{row - 1, column})
		}
		if search(row, column-1, Sub) {
			tempresult = append(tempresult, [2]int{row, column - 1})
		}
		if search(row, column+1, Sub) {
			tempresult = append(tempresult, [2]int{row, column + 1})
		}
		if search(row+1, column, Sub) {
			tempresult = append(tempresult, [2]int{row + 1, column})
		}

		Sub++

		for i := 0; i < len(tempresult); i++ {
			if judge(tempresult[i]) {
				return true
			}
		}
		//如果不行的话 要恢复为原来的状态
		m[[2]int{row, column}] = false
		Sub--
		return false
	}

	for i := 0; i < len(result); i++ {
		m = make(map[[2]int]bool)
		m[result[i]] = true
		point = result[i]
		if judge(point) {
			return true
		}
	}
	return false
}

//关键还是在于不对的时候如何恢复 以及 可以将参数+1的方式穿进去 遇到string不需要单独拆分 , 传入index就可以了  备忘录功能
func exist2(board [][]byte, word string) bool {
	rows, columns := len(board), len(board[0])
	m := make(map[[2]int]bool)
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i > rows-1 || j < 0 || j > columns-1 || board[i][j] != word[k] || m[[2]int{i, j}] == true {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		m[[2]int{i, j}] = true
		b := dfs(i-1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i+1, j, k+1)
		if !b {
			m[[2]int{i, j}] = false
		}
		return b
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

//不使用备忘录 直接提高了好多 关键在于 不用备忘录防止往回走
func exist3(board [][]byte, word string) bool {
	rows, columns := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i > rows-1 || j < 0 || j > columns-1 || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		//不用备忘录防止往回走
		board[i][j] = ' '
		b := dfs(i-1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1) || dfs(i+1, j, k+1)
		board[i][j] = word[k]
		return b
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			//这里也很关键
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func movingCount(m int, n int, k int) int {
	sum := 0
	mapp := make(map[[2]int]bool)
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || mapp[[2]int{i, j}] {
			return
		}
		a := i % 10
		b := (i - a) / 10
		c := j % 10
		d := (j - c) / 10
		if a+b+c+d > k {
			return
		}
		mapp[[2]int{i, j}] = true
		sum++
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i+1, j)
	}

	dfs(0, 0)

	return sum
}

//更好看而已
func movingCount1(m int, n int, k int) int {
	mapp := make(map[[2]int]bool)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || mapp[[2]int{i, j}] {
			return 0
		}
		a := i % 10
		b := (i - a) / 10
		c := j % 10
		d := (j - c) / 10
		if a+b+c+d > k {
			return 0
		}
		mapp[[2]int{i, j}] = true
		return dfs(i-1, j) + dfs(i, j-1) + dfs(i, j+1) + dfs(i+1, j) + 1
	}

	return dfs(0, 0)
}

//广度优先搜索
func movingCount2(m int, n int, k int) int {
	sum := 0
	added := make(map[[2]int]bool)
	queue1 := make([][2]int, 0)
	queue2 := make([][2]int, 0)
	queue1 = append(queue1, [2]int{0, 0})
	added[[2]int{0, 0}] = true
	var a func(i, j int) bool
	a = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		a := i % 10
		b := (i - a) / 10
		c := j % 10
		d := (j - c) / 10
		if a+b+c+d > k {
			return false
		}

		sum++
		return true
	}
	for len(queue1) != 0 || len(queue2) != 0 {
		if len(queue1) != 0 {
			for _, v := range queue1 {
				i := v[0]
				j := v[1]
				if a(i, j) {
					top := [2]int{i - 1, j}
					left := [2]int{i, j - 1}
					right := [2]int{i, j + 1}
					bottom := [2]int{i + 1, j}
					if !added[top] {
						added[top] = true
						queue2 = append(queue2, top)
					}
					if !added[left] {
						added[left] = true

						queue2 = append(queue2, left)
					}
					if !added[right] {
						added[right] = true

						queue2 = append(queue2, right)
					}
					if !added[bottom] {
						added[bottom] = true

						queue2 = append(queue2, bottom)
					}
				}
				queue1 = queue1[1:]
			}
		} else if len(queue2) != 0 {
			for _, v := range queue2 {
				i := v[0]
				j := v[1]
				if a(i, j) {
					top := [2]int{i - 1, j}
					left := [2]int{i, j - 1}
					right := [2]int{i, j + 1}
					bottom := [2]int{i + 1, j}
					if !added[top] {
						added[top] = true

						queue1 = append(queue1, top)
					}
					if !added[left] {
						added[left] = true

						queue1 = append(queue1, left)
					}
					if !added[right] {
						added[right] = true

						queue1 = append(queue1, right)
					}
					if !added[bottom] {
						added[bottom] = true

						queue1 = append(queue1, bottom)
					}
				}
				queue2 = queue2[1:]
			}
		}
	}

	return sum
}

//隐藏的优化策略 就是只需要向右和向下就可以了干
func movingCount3(m int, n int, k int) int {
	mapp := make(map[[2]int]bool)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || mapp[[2]int{i, j}] {
			return 0
		}
		a := i % 10
		b := (i - a) / 10
		c := j % 10
		d := (j - c) / 10
		if a+b+c+d > k {
			return 0
		}
		mapp[[2]int{i, j}] = true
		return dfs(i, j+1) + dfs(i+1, j) + 1
	}

	return dfs(0, 0)
}

//广度优先搜索 隐藏的优化策略 只需要向右和向下就可以了呜呜呜  实际上只要一个队列就可以了 ljs废物
func movingCount4(m int, n int, k int) int {
	sum := 0
	added := make(map[[2]int]bool)
	queue1 := make([][2]int, 0)
	queue2 := make([][2]int, 0)
	queue1 = append(queue1, [2]int{0, 0})
	added[[2]int{0, 0}] = true
	var a func(i, j int) bool
	a = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		a := i % 10
		b := (i - a) / 10
		c := j % 10
		d := (j - c) / 10
		if a+b+c+d > k {
			return false
		}

		sum++
		return true
	}
	for len(queue1) != 0 || len(queue2) != 0 {
		if len(queue1) != 0 {
			for _, v := range queue1 {
				i := v[0]
				j := v[1]
				if a(i, j) {
					right := [2]int{i, j + 1}
					bottom := [2]int{i + 1, j}
					if !added[right] {
						added[right] = true

						queue2 = append(queue2, right)
					}
					if !added[bottom] {
						added[bottom] = true

						queue2 = append(queue2, bottom)
					}
				}
				queue1 = queue1[1:]
			}
		} else if len(queue2) != 0 {
			for _, v := range queue2 {
				i := v[0]
				j := v[1]
				if a(i, j) {
					right := [2]int{i, j + 1}
					bottom := [2]int{i + 1, j}
					if !added[right] {
						added[right] = true

						queue1 = append(queue1, right)
					}
					if !added[bottom] {
						added[bottom] = true

						queue1 = append(queue1, bottom)
					}
				}
				queue2 = queue2[1:]
			}
		}
	}

	return sum
}

//太关键了 我去 这个问题的关键就在于 如果f(3)<3那么就用3 如果f(5)>5 那么就用f(5) 把多个拆分 还原成两个拆分
func cuttingRope(n int) int {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 1
	m[3] = 2
	m[4] = 4
	var travel func(n int) int
	travel = func(n int) int {
		//备忘录
		if v, ok := m[n]; ok {
			return v
		}
		var final int
		//例如 i=5 那么就枚举 14 23 即可   要知道f(2)=1 f(3)=2 f(4)=4 f(5)=max(1+4的情况的值,2+3情况的值)=6 f(6)=max(1+5,2+4,3+3)=9
		//关键的来了看过来各位 f(6)计算的时候 1+5可以分为啥呢 f(1)肯定是等于1的没话讲 f(5)就要和5本身做比较 因为 5可以拆呀,但是f(5)=6大于5本身 所以要取6
		//所以1+5 的最大值应该是6
		//验证一下 f(8)=max(1+7,2+6,3+5,4+4)
		//1+7 是取1+f(7) 因为f(7)等于12大于7本身 所以=12
		//2+6 是取2+f(6) 因为f(6)等于9 大于6本身 所以=2*9=18   这里为什么取2 是因为f(2)=1小于2本身
		//3+5 是取3+f(5) 因为f(5)等于6 大于5本身 所以=3*6=18   这里为什么取3 是因为f(3)=2小于3本身
		for i := 1; i <= n/2; i++ {
			j := n - i
			var temp = travel(j)
			m[j] = temp
			if j > travel(j) {
				temp = j
			}
			var temp1 = travel(i)
			if i > travel(i) {
				temp1 = i
			}
			m[i] = temp1
			k := temp * temp1
			if final < k {
				final = k
			}
		}
		return final
	}
	return travel(n)
}

func cuttingRope1(n int) int {
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	}
	m := make(map[int]int)
	m[1] = 1
	m[2] = 1
	m[3] = 2
	m[4] = 4
	var final int
	q := 5
	for q != n+1 {
		for i := 1; i <= q/2; i++ {
			temp := m[i]
			if i > temp {
				temp = i
			}
			temp1 := m[q-i]
			if q-i > temp1 {
				temp1 = q - i
			}
			k := temp * temp1
			if k > m[q] {
				m[q] = k
				final = k
			}
		}
		q++
	}

	return final
}

func levelOrder(root *TreeNode) []int {
	final := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		if queue[0] != nil {
			final = append(final, queue[0].Val)
			queue = append(queue, queue[0].Left)
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
	}
	return final
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	b := false
	if n < 0 {
		n = -n
		b = true
	}
	base := x
	for i := 1; i < n; i++ {
		x = x * base
	}
	if b {
		x = 1 / x
	}
	return x
}

//快速幂算法 利用&运算来获得末位 利用右移运算来除2 利用快速幂减少计算次数
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 1 {
		return 1
	}
	var final float64 = 1
	b := false
	if n < 0 {
		n = -n
		b = true
	}
	for n != 0 {
		a := n & 1
		if a == 1 {
			final *= x
		}
		x *= x
		n = n >> 1
	}

	if b {
		final = 1 / final
	}
	return final
}

//什么鬼 直接全自动状态机 什么网易儿 我直接穷举了
func isNumber(s string) bool {
	//判断是不是整数
	var isInt func(sub string) bool
	isInt = func(sub string) bool {
		if len(sub) == 0 {
			return false
		}
		//符号个数多了就错了
		if strings.Count(sub, "+")+strings.Count(sub, "-") > 1 {
			return false
		}
		//去掉符号
		if sub[0] == '+' || sub[0] == '-' {
			sub = sub[1:]
		}
		//尝试转int 转错了就说明有问题
		_, err := strconv.Atoi(sub)
		if err != nil {
			return false
		}
		return true
	}

	//判断是不是小数
	var isdot func(sub string) bool
	isdot = func(sub string) bool {
		//符号个数多了就错了
		if strings.Count(sub, "+")+strings.Count(sub, "-") > 1 {
			return false
		}
		//去掉符号
		if sub[0] == '+' || sub[0] == '-' {
			sub = sub[1:]
		}
		//小数点多了就错了
		if strings.Count(sub, ".") > 1 {
			return false
		}
		i := strings.Index(sub, ".")
		if i >= 0 {
			//如果存在就去掉小数点
			sub = sub[:i] + sub[i+1:]
		}
		//如果去掉符号之后再去掉小数点之后长度为零 就错了
		if len(sub) == 0 {
			return false
		}
		//如果去掉符号之后再去掉小数点之后 还有+-符号就错了
		if sub[0] == '+' || sub[0] == '-' {
			return false
		}
		//尝试转int 转错了就说明有问题
		_, err := strconv.Atoi(sub)
		if err != nil {
			//如果问题是范围太大那没关系  233245234523452345.2345234523452345 合起来转确实会遇到这种问题
			if strings.Contains(err.Error(), "value out of range") {
				return true
			}
			return false
		}
		return true
	}

	//去空格
	s = strings.Trim(s, " ")
	//去完空格为0 就错了
	if len(s) == 0 {
		return false
	}

	temp1 := true
	temp2 := true

	//e的个数多了就错了
	if strings.Count(s, "e")+strings.Count(s, "E") > 1 {
		return false
	}

	index := 0

	if strings.Contains(s, "e") || strings.Contains(s, "E") {
		if strings.Index(s, "e") > strings.Index(s, "E") {
			index = strings.Index(s, "e")
		} else {
			index = strings.Index(s, "E")
		}
		//判断e后面是不是整数
		temp2 = isInt(s[index+1:])
	}

	//没有含e 或者 E的部分
	if index == 0 {
		temp1 = isdot(s) || isInt(s)
	} else {
		temp1 = isdot(s[:index]) || isInt(s[:index])
	}

	return temp1 && temp2
}
