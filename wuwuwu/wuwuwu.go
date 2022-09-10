package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//n := 0
	//fmt.Scanf("%d", &n)
	//oldsss := make([]int, 0, 0)
	//for i := 0; i <= n; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	oldsss = append(oldsss, num)
	//}
	//oldsss = oldsss[1:]
	//
	//target := make([]int, 0, 0)
	//for i := 0; i <= n; i++ {
	//	num := 0
	//	fmt.Scanf("%d", &num)
	//	target = append(target, num)
	//}
	//target = target[1:]

	//fmt.Println(oldsss)
	//fmt.Println(target)
	//m := make(map[int]int)
	//for i := 0; i < len(oldsss); i++ {
	//	m[oldsss[i]] = i
	//}

	//res := 0
	//var travel func(old []int, target []int)
	//travel = func(old []int, target []int) {
	//	if len(old) == 0 {
	//		return
	//	}
	//	if old[0] != target[0] {
	//		res++
	//		old[len(old)-1], oldsss[m[target[0]]] = oldsss[m[target[0]]], old[len(old)-1]
	//		travel(old[:len(old)-1], target[1:])
	//	} else {
	//		travel(old[1:], target[1:])
	//	}
	//}
	//travel(oldsss, target)
	//fmt.Println(res)
	//time := convertTime("11:00", "11:01")
	//fmt.Println(time)

	//winners := findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}})
	//fmt.Println(winners)

	//candies := maximumCandies([]int{5, 8, 6}, 3)
	//candies := maximumCandies([]int{4, 7, 5}, 16)
	//candies := maximumCandies([]int{1, 2, 6, 8, 6, 7, 3, 5, 2, 5}, 3)
	//fmt.Println(candies)

	//obj := Constructor([]byte{'a', 'b', 'c', 'd'}, []string{"ei", "zf", "ei", "am"}, []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"})
	//param_1 := obj.Encrypt("abcd")
	//fmt.Println(param_1)
	//param_2 := obj.Decrypt("eizfeiam")
	//fmt.Println(param_2)

	//t := &TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:  4,
	//			Left: nil,
	//			Right: &TreeNode{
	//				Val:   2,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//		Right: nil,
	//	},
	//}
	//recoverTree(t)
	//fmt.Println(t)

	//change := coinChange([]int{1, 2, 5}, 11)
	//change := coinChange([]int{2}, 3)
	//change := coinChange([]int{186, 419, 83, 408}, 6249)
	//fmt.Println(change)

	//i := permute([]int{1, 2, 3})
	//fmt.Println(i)

	//queens := solveNQueens(4)
	//queens := solveNQueens(6)
	//fmt.Println(queens)

	//solveSudoku([][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}})

	//lock := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	//lock := openLock([]string{"8888"}, "0009")
	//lock := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	//lock := openLock([]string{"0000"}, "8888")
	//fmt.Println(lock)

	//ints := guibingSort([]int{1, 3, 5, 4, 2})
	//fmt.Println(ints)

	//l := &ListNode{
	//	Val: 4,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 1,
	//			Next: &ListNode{
	//				Val:  3,
	//				Next: nil,
	//			},
	//		},
	//	},
	//}
	//list := sortList(l)
	//fmt.Println(list)

	//window := minWindow("ADOBECODEBANC", "ABC")
	//window := minWindow("a", "aa")
	//window := minWindow("bba", "ab")
	//fmt.Println(window)

	//inclusion := checkInclusion("a", "ab")
	//inclusion := checkInclusion("ab", "eidbaooo")
	//inclusion := checkInclusion("ab", "eidboaoo")
	//fmt.Println(inclusion)

	//anagrams := findAnagrams("cbaebabacd", "abc")
	//anagrams := findAnagrams("abab", "ab")
	//anagrams := findAnagrams("aa", "bb")
	//anagrams := findAnagrams("aaaaaaaaaa", "aaaaaaaaaaaaa")
	//fmt.Println(anagrams)

	//substring := lengthOfLongestSubstring("abcabcbb")
	//substring := lengthOfLongestSubstring("bbbbb")
	//substring := lengthOfLongestSubstring("pwwkew")
	//fmt.Println(substring)

	//lis := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	//lis := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	//lis := lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	//fmt.Println(lis)

	//array := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	//fmt.Println(array)

	//integer := largestInteger(1234)
	//integer := largestInteger(247)
	//integer := largestInteger(65875)
	//fmt.Println(integer)

	//result := minimizeResult("247+38")
	//result := minimizeResult("12+34")
	//result := minimizeResult("999+999")
	//fmt.Println(result)

	//product := maximumProduct([]int{0, 4}, 5)
	//product := maximumProduct([]int{6, 3, 3, 2}, 2)
	//product := maximumProduct([]int{92, 36, 15, 84, 57, 60, 72, 86, 70, 43, 16}, 62)
	//fmt.Println(product)

	//text := deleteText("Singing dancing in the rain", 10)
	//text := deleteText("Hello", 0)
	//text := deleteText("Hello World", 2)
	//text := deleteText("Hello World", 6)
	//fmt.Println(text)

	//sticks := lightSticks(1, 2, []int{3})
	//sticks := lightSticks(1, 1, []int{0, 3})
	//sticks := lightSticks(2, 2, []int{2, 5, 6, 7, 8, 10, 11})
	//fmt.Println(sticks)

	//bianliZheng()
	//bianliFan()
	//bianliXieYou()
	//bianliXieDouble()
	//bianliZuoxiajiao()
	//bianlidowotoUpLeftToRight()

	//subsequence := longestCommonSubsequence("abcde", "ace")
	//fmt.Println(subsequence)

	//distance := minDistance("horse", "ros")
	//distance := minDistance("intention", "execution")
	//fmt.Println(distance)

	//partition := canPartition([]int{1, 5, 11, 5})
	//partition := canPartition([]int{9, 5})
	//partition := canPartition([]int{1, 2, 3, 6})
	//partition := canPartition([]int{1, 2, 3, 5})
	//partition : = canPartition([]int{1, 1, 2, 2})
	//partition := canPartition([]int{1, 1})
	//partition := canPartition([]int{1, 2, 3, 4, 5, 6, 7})
	//fmt.Println(partition)

	//sack := knapSack(4, 3, []int{2, 1, 3}, []int{4, 2, 3})
	//fmt.Println(sack)

	//sum := nSum([]int{1, 5, 11, 5}, 11)
	//sum := nSum([]int{1, 2, 3, 6}, 11)
	//sum := nSum([]int{1, 1, 2, 2}, 11)
	//sum := nSum([]int{1, 2, 3, 4, 5, 6, 7}, 11)
	//fmt.Println(sum)

	//change := coinChange([]int{1, 2, 5}, 11)
	//change := coinChange([]int{2}, 3)
	//change := coinChange([]int{1}, 0)
	//fmt.Println(change)

	//subseq := longestPalindromeSubseq("bbbab")
	//fmt.Println(subseq)

	//insertions := minInsertions("mbadm")
	//insertions := minInsertions("leetcode")
	//insertions := minInsertions("zzazz")
	//fmt.Println(insertions)

	//match := isMatch("aa", "a*")
	//match := isMatch("aa", "a")
	//match := isMatch("aa", "aa")
	//match := isMatch("ab", ".*")
	//match := isMatch("zaaab", ".a*b")
	//match := isMatch("cb", ".a*b")
	//match := isMatch("amnb", "a..b")
	//match := isMatch("ab", ".b")
	//fmt.Println(match)

	//i := change(5, []int{1, 2, 5})
	//i := change(3, []int{2})
	//i := change(10, []int{10})
	//fmt.Println(i)

	//i := rob([]int{2, 1, 7, 9, 3, 1})
	//i := rob([]int{1, 2, 3, 1})
	//i := rob([]int{1, 2, 3})
	//i := rob([]int{2, 3, 2})
	//fmt.Println(i)

	//i := rob(&TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 4,
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
	//		Val:  5,
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
	//ways := findTargetSumWays([]int{1}, 2)
	//ways := findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1)
	//ways := findTargetSumWays([]int{100}, -200)
	//fmt.Println(ways)

	//form := findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
	//form := findMaxForm([]string{"10", "0", "1"}, 1, 1)
	//form := findMaxForm([]string{"00", "000"}, 1, 10)
	//form := findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 3, 4)
	//form := findMaxForm([]string{"100010101100110010101", "0110111011", "0011", "100101", "01000100011110110100111111", "00001011111011111101011110"}, 2, 24)
	//fmt.Println(form)

	//b := wordBreak("applepenapple", []string{"apple", "pen"})
	//b := wordBreak("leetcode", []string{"leet", "code"})
	//b := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	//fmt.Println(b)

	//sum4 := combinationSum4([]int{1, 2, 3}, 4)
	//sum4 := combinationSum4([]int{9}, 3)
	//sum4 := combinationSum4([]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}, 10)
	//sum4 := combinationSum4([]int{3, 1, 2, 4}, 4)
	//fmt.Println(sum4)

	//number := findClosestNumber([]int{-4, -2, 1, 4, 8})
	//number := findClosestNumber([]int{2, -1, 1})
	//fmt.Println(number)

	//pencils := waysToBuyPensPencils(20, 10, 5)
	//pencils := waysToBuyPensPencils(5, 10, 10)
	//fmt.Println(pencils)

	//constructor := Constructor()
	//constructor.Deposit([]int{0, 0, 1, 2, 1})
	//withdraw := constructor.Withdraw(600)
	//fmt.Println(withdraw)
	//constructor.Deposit([]int{0, 1, 0, 1, 1})
	//ints := constructor.Withdraw(600)
	//fmt.Println(ints)
	//i := constructor.Withdraw(550)
	//fmt.Println(i)
	//
	//constructor := Constructor()
	//constructor.Deposit([]int{100000, 100000, 100000, 100000, 100000})
	//withdraw := constructor.Withdraw(1000000000)
	//fmt.Println(withdraw)

	//sum := digitSum("11111222223", 3)
	//sum := digitSum("00000000", 3)
	//fmt.Println(sum)

	//rounds := minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})
	//rounds := minimumRounds([]int{5, 5, 5, 5})
	//rounds := minimumRounds([]int{2, 3, 3})
	//rounds := minimumRounds([]int{7, 7, 7, 7, 7, 7})
	//fmt.Println(rounds)

	//l := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//
	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: l,
	//		},
	//	},
	//}

	//list := reverseList(l2)
	//list := reverseN(l2, 2)
	//fmt.Println(list)

	//between := reverseBetween(l2, 2, 3)
	//fmt.Println(between)

	//s := ""
	//t := ""
	//fmt.Scanf("%s\n", &s)
	//fmt.Scanf("%s\n", &t)
	//i := distance(s, t)
	//fmt.Println(i)

	//digit := removeDigit("123", '3')
	//digit := removeDigit("2998589353917872714814599237991174513476623756395992135212546127959342974628712329595771672911914471", '3')
	//digit := removeDigit("1231", '1')
	//digit := removeDigit("551", '5')
	//fmt.Println(digit)

	//pickup := minimumCardPickup([]int{95, 11, 8, 65, 5, 86, 30, 27, 30, 73, 15, 91, 30, 7, 37, 26, 55, 76, 60, 43, 36, 85, 47, 96, 6})
	//pickup := minimumCardPickup([]int{3, 4, 2, 3, 4, 7})
	//pickup := minimumCardPickup([]int{1, 0, 5, 3})
	//fmt.Println(pickup)

	distinct := countDistinct([]int{2, 3, 3, 2, 2}, 2, 2)
	//distinct := countDistinct([]int{1, 2, 3, 4}, 4, 1)
	fmt.Println(distinct)
}

func convertTime(current string, correct string) int {
	s1 := strings.Split(current, ":")
	s2 := strings.Split(correct, ":")
	hc, _ := strconv.Atoi(s1[0])
	mc, _ := strconv.Atoi(s1[1])
	hcc, _ := strconv.Atoi(s2[0])
	mcc, _ := strconv.Atoi(s2[1])
	low := hc*60 + mc
	high := hcc*60 + mcc
	res := 0
	for low != high {
		if low+60 <= high {
			low += 60
		} else if low+15 <= high {
			low += 15
		} else if low+5 <= high {
			low += 5
		} else if low+1 <= high {
			low += 1
		}
		res++

	}
	return res
}

func findWinners(matches [][]int) [][]int {
	res := make([][]int, 0, 0)
	winner := make(map[int]int)
	loser := make(map[int]int)
	for i := 0; i < len(matches); i++ {
		winner[matches[i][0]]++
		loser[matches[i][1]]++
	}

	w := make([]int, 0, 0)
	l := make([]int, 0, 0)
	for key, v := range loser {
		if v >= 2 {
			continue
		} else if v == 1 {
			l = append(l, key)
		}
	}
	sort.Ints(l)

	for key := range winner {
		if _, ok := loser[key]; !ok {
			w = append(w, key)
		}
	}
	sort.Ints(w)
	res = append(res, w, l)
	return res
}

func maximumCandies(candies []int, k int64) int {
	sum := int64(0)
	for i := 0; i < len(candies); i++ {
		sum += int64(candies[i])
	}
	right := sum / k
	if right == 1 {
		return 1
	}
	//left := sum / (k + 1)
	left := int64(1)
	//if right == 0 {
	//	return 0
	//}

	var travel func(mid int) bool
	travel = func(mid int) bool {
		count := int64(0)
		for i := 0; i < len(candies); i++ {
			count += int64(candies[i] / mid)
			if count >= k {
				return true
			}
		}
		return false
	}

	for left <= right {
		mid := left + (right-left)/2
		if travel(int(mid)) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return int(left - 1)

}

//
//type Encrypter struct {
//	KeysToValues map[byte]string
//	mDic         map[string]int
//}
//
//func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
//	m := make(map[byte]string)
//	for i := 0; i < len(keys); i++ {
//		m[keys[i]] = values[i]
//	}
//
//	mDic := make(map[string]int)
//
//	for i := 0; i < len(dictionary); i++ {
//		bytes := make([]byte, 0, 0)
//		for j := 0; j < len(dictionary[i]); j++ {
//			bytes = append(bytes, m[dictionary[i][j]]...)
//		}
//		mDic[string(bytes)]++
//	}
//
//	return Encrypter{
//		KeysToValues: m,
//		mDic:         mDic,
//	}
//}
//
//func (this *Encrypter) Encrypt(word1 string) string {
//	bytes := make([]byte, 0, 0)
//	for i := 0; i < len(word1); i++ {
//		bytes = append(bytes, this.KeysToValues[word1[i]]...)
//	}
//	return string(bytes)
//}
//
//func (this *Encrypter) Decrypt(word2 string) int {
//	return this.mDic[word2]
//}
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//二叉搜索树性质:通过中序遍历可以从大到小排列  构造两个数组 然后sort来对比 对比完去修改
func recoverTree1(root *TreeNode) {
	t := make([]*TreeNode, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		travel(root.Left)
		t = append(t, root)
		travel(root.Right)
	}

	travel(root)
	ints := make([]int, len(t), len(t))
	for i := 0; i < len(t); i++ {
		ints[i] = t[i].Val
	}

	sort.Ints(ints)
	for i := 0; i < len(t); i++ {
		if ints[i] != t[i].Val {
			t[i].Val = ints[i]
		}
	}

}

//二叉搜索树性质:通过中序遍历可以从大到小排列  构造一个数组 然后再对比 对比完数组中直接修改
func recoverTree2(root *TreeNode) {
	t := make([]*TreeNode, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		travel(root.Left)
		t = append(t, root)
		travel(root.Right)
	}

	travel(root)

	a := -1
	b := -1

	for i := 0; i < len(t)-1; i++ {
		if t[i].Val > t[i+1].Val {
			if a == -1 {
				a = i
			} else {
				b = i + 1
			}
		}
	}
	if b != -1 {
		t[a].Val, t[b].Val = t[b].Val, t[a].Val
	} else {
		t[a].Val, t[a+1].Val = t[a+1].Val, t[a].Val
	}

}

//太牛了 还暂时看不懂咱  迭代版本的中序遍历  还有一个mirrors中序 先放一放叭有点难
func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

//使用备忘录 来消除重复子问题 很慢很慢啊...
func coinChange1(coins []int, amount int) int {
	res := -1

	m := make(map[int]int)
	var travel func(target int) int
	travel = func(target int) int {
		if v, ok := m[target]; ok {
			return v
		}

		if target < 0 {
			return -1
		} else if target == 0 {
			return 0
		}
		ans := math.MaxInt
		for i := len(coins) - 1; i >= 0; i-- {
			var temp int
			temp = travel(target - coins[i])
			if temp == -1 {
				continue
			} else if temp == math.MaxInt {
				continue
			} else {
				if ans > temp+1 {
					ans = temp + 1
				}
			}
		}

		m[target] = ans
		return ans
	}

	res = travel(amount)
	if res == math.MaxInt {
		return -1
	}
	return res
}

//dp数组 这里定义很简单 问题是最好应该画一下递归树 来更好的了解重叠子问题来着
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] == i {
				dp[i] = 1
				break
			}
			if i-coins[j] >= 0 && dp[i-coins[j]] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

//这里感觉没有使用回溯 使用copy来复制一遍剩下的数组 使用ans.append 传递的是新的
//函数在返回的时候 不会对原来的ans有任何影响
func permute1(nums []int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(ans []int, left []int)
	travel = func(ans []int, left []int) {
		if len(left) == 0 {
			res = append(res, ans)
			return
		}

		for i := 0; i < len(left); i++ {
			temp := make([]int, len(left)-1, len(left)-1)
			//这里返回了实际上拷贝了多少个
			has := copy(temp, left[:i])
			//以便于后续进一步拷贝  其实有一个很好的办法是 把当前这个移动到最后面或者最前面  然后递归 然后回来的时候再调回来
			copy(temp[has:], left[i+1:])
			travel(append(ans, left[i]), temp)
		}
	}

	travel([]int{}, nums)
	return res
}

func permute(nums []int) [][]int {
	res := make([][]int, 0, 0)

	var travel func(ans []int, left []int)
	travel = func(ans []int, left []int) {
		if len(left) == 0 {
			res = append(res, ans)
			return
		}

		for i := 0; i < len(left); i++ {
			//通过调换位置的方式 移动到最前面 然后递归
			left[0], left[i] = left[i], left[0]
			//这里就传的是 1到最后的
			travel(append(ans, left[0]), left[1:])
			//回来的时候调换回来
			left[0], left[i] = left[i], left[0]
		}
	}

	travel([]int{}, nums)
	return res
}

//N皇后 回溯 前序遍历+后续遍历 不只是左上角左下角右上角右下角 而是左边对角线和右边对角线
//这里要注意!!! visited 哈希表最好没事儿就使用map[]int 用bool怕负负得正 两次调用就裂开了
func solveNQueens(n int) [][]string {
	col := make(map[int]bool)
	visited := make(map[[2]int]int)
	res := make([][]string, 0, 0)
	ans := make([][2]int, 0, 0)

	bytes := make([]byte, n, n)
	for i := 0; i < n; i++ {
		bytes[i] = '.'
	}

	var travel func(row int, ans [][2]int)
	travel = func(row int, ans [][2]int) {
		if len(ans) == n {
			temp := make([]string, len(ans), len(ans))
			for i := 0; i < len(ans); i++ {
				bytes[ans[i][1]] = 'Q'
				temp[i] = string(bytes)
				bytes[ans[i][1]] = '.'
			}
			res = append(res, temp)
			return
		}

		for j := 0; j < n; j++ {

			if !col[j] && visited[[2]int{row, j}] == 0 {
				col[j] = true
				visited[[2]int{row, j}]++

				for r, c := row+1, j-1; r < n && c >= 0; {
					visited[[2]int{r, c}]++
					r++
					c--
				}

				for r, c := row+1, j+1; r < n && c < n; {
					visited[[2]int{r, c}]++
					r++
					c++
				}

				travel(row+1, append(ans, [2]int{row, j}))

				col[j] = false
				visited[[2]int{row, j}]--

				for r, c := row+1, j-1; r < n && c >= 0; {
					visited[[2]int{r, c}]--
					r++
					c--
				}

				for r, c := row+1, j+1; r < n && c < n; {
					visited[[2]int{r, c}]--
					r++
					c++
				}

			}
		}

	}

	travel(0, ans)

	return res
}

//解数独 如果只需要返回一个满足条件的集合的话 那么就将返回值调整成bool类型 回溯到满足条件就返回
//或者笨比办法是 使用一个公共变量b 每次回溯完判断一下b是不是true 是的话就直接返回 不是就继续递归
func solveSudoku(board [][]byte) {
	visited := make(map[[2]int]bool)

	row := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = map[byte]bool{}
	}

	col := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		col[i] = map[byte]bool{}
	}

	//3*3 9个小块 内部只能1-9
	block := make([][3]map[byte]bool, 3)
	for i := 0; i < len(block); i++ {
		block[i] = [3]map[byte]bool{}
		for j := 0; j < 3; j++ {
			block[i][j] = map[byte]bool{}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				block[i/3][j/3][board[i][j]-'0'] = true

				row[i][board[i][j]-'0'] = true
				col[j][board[i][j]-'0'] = true

				visited[[2]int{i, j}] = true
			}
		}
	}

	var travel func(index int, jd int) bool
	travel = func(index int, jd int) bool {
		if index == 9 {
			return true
		}

		if jd == 9 {
			return travel(index+1, 0)
		}

		x := index / 3
		y := jd / 3
		if visited[[2]int{index, jd}] {
			return travel(index, jd+1)
		} else {
			for i := 1; i <= 9; i++ {
				if block[x][y][byte(i)] || row[index][byte(i)] || col[jd][byte(i)] {
					continue
				}

				board[index][jd] = byte('0' + i)
				visited[[2]int{index, jd}] = true
				block[x][y][byte(i)] = true
				row[index][byte(i)] = true
				col[jd][byte(i)] = true

				if travel(index, jd+1) {
					return true
				}

				board[index][jd] = byte('.')
				visited[[2]int{index, jd}] = false
				block[x][y][byte(i)] = false
				row[index][byte(i)] = false
				col[jd][byte(i)] = false

			}
			return false

		}
	}

	travel(0, 0)
}

//dfs
func minDepth(root *TreeNode) int {
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	} else {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	}
}

//bfs
func minDepth1(root *TreeNode) int {
	res := 0
	queue := make([]*TreeNode, 0, 0)
	if root != nil {
		queue = append(queue, root)
	}
	l := len(queue)
	for len(queue) != 0 {
		res++

		for l != 0 {
			q := queue[0]

			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}

			if q.Left == nil && q.Right == nil {
				return res
			}
			queue = queue[1:]
			l--
		}
		l = len(queue)
	}

	return res
}

//解密码锁 非常有启发性啊!!!  bfs 直接暴力穷举 然后备忘录剪枝!!!!!!!!!
func openLock1(deadends []string, target string) int {
	visited := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		visited[deadends[i]] = true
	}

	res := 0
	bytes := []byte{'0', '0', '0', '0'}

	if target == string(bytes) {
		return res
	}

	if visited[string(bytes)] {
		return -1
	}

	queue := make([][]byte, 0, 0)
	queue = append(queue, bytes)
	for len(queue) != 0 {
		l := len(queue)
		for l > 0 {
			q := queue[0]

			if string(q) == target {
				return res
			}

			for i := 0; i < 4; i++ {
				temp1 := make([]byte, 4, 4)
				temp2 := make([]byte, 4, 4)

				copy(temp1, q)
				temp1[i] = (temp1[i]-'0'+10+1)%10 + '0'
				if !visited[string(temp1)] {
					queue = append(queue, temp1)
					visited[string(temp1)] = true
				}

				copy(temp2, q)
				temp2[i] = (temp2[i]-'0'+10-1)%10 + '0'
				if !visited[string(temp2)] {
					queue = append(queue, temp2)
					visited[string(temp2)] = true
				}
			}
			queue = queue[1:]
			l--
		}
		res++
	}
	return -1
}

//双向bfs 特点:需要明确知道起点和终点 否则用不了 此外 用哈希表来替换队列 使用交替哈希表的形式来轮流探索
func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for i := 0; i < len(deadends); i++ {
		visited[deadends[i]] = true
	}

	res := 0
	bytes := []byte{'0', '0', '0', '0'}

	if target == string(bytes) {
		return res
	}

	if visited[string(bytes)] {
		return -1
	}

	q1 := make(map[string]bool)
	q2 := make(map[string]bool)

	q1[string(bytes)] = true

	q2[target] = true

	for len(q1) != 0 && len(q2) != 0 {
		temp := make(map[string]bool)
		for key := range q1 {
			if q2[key] {
				return res
			}
			visited[key] = true

			bs := make([]byte, 4, 4)

			copy(bs, key)

			for i := 0; i < 4; i++ {
				bs[i] = (key[i]-'0'+10+1)%10 + '0'

				if !visited[string(bs)] {
					temp[string(bs)] = true
				}

				bs[i] = (key[i]-'0'+10-1)%10 + '0'

				if !visited[string(bs)] {
					temp[string(bs)] = true
				}

				bs[i] = key[i]
			}
		}

		res++

		q1 = q2
		q2 = temp
	}

	return -1
}

func guibingSort(ints []int) []int {

	var merge func(pre, lat []int) []int
	merge = func(pre, lat []int) []int {
		ans := make([]int, 0, len(pre)+len(lat))
		p, l := 0, 0
		for p != len(pre) || l != len(lat) {
			if p == len(pre) {
				ans = append(ans, lat[l])
				l++
			} else if l == len(lat) {
				ans = append(ans, pre[p])
				p++
			} else {
				if pre[p] < lat[l] {
					ans = append(ans, pre[p])
					p++
				} else {
					ans = append(ans, lat[l])
					l++
				}
			}
		}
		return ans
	}

	var mergeSort func(ints []int) []int
	mergeSort = func(ints []int) []int {
		if len(ints) == 1 || len(ints) == 0 {
			return ints
		}
		mid := len(ints) / 2

		pre := mergeSort(ints[:mid])
		lat := mergeSort(ints[mid:])
		return merge(pre, lat)
	}

	return mergeSort(ints)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	var merge func(pre *ListNode, latter *ListNode) *ListNode
	merge = func(pre *ListNode, latter *ListNode) *ListNode {
		res := &ListNode{
			Val:  0,
			Next: nil,
		}
		cur := res

		p, l := pre, latter
		for p != nil || l != nil {
			if p == nil {
				cur.Next = l
				cur = cur.Next
				l = l.Next
			} else if l == nil {
				cur.Next = p
				cur = cur.Next
				p = p.Next
			} else {
				if p.Val < l.Val {
					cur.Next = p
					cur = cur.Next
					p = p.Next
				} else {
					cur.Next = l
					cur = cur.Next
					l = l.Next
				}
			}
		}

		return res.Next
	}

	var mergeSort func(head *ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		last := head
		slow, fast := head, head
		for !(fast == nil || fast.Next == nil) {
			last = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		last.Next = nil
		pre := mergeSort(head)
		latter := mergeSort(slow)
		return merge(pre, latter)
	}

	return mergeSort(head)
}

//s = "ADOBECODEBANC", t = "ABC"
//BANC
//最小覆盖子串  滑动窗口
func minWindow1(s string, t string) string {
	length := math.MaxInt
	res := ""
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}

	mvisited := make(map[byte]int)

	var checked func() bool
	checked = func() bool {
		b := true
		for key, v := range m {
			if mvisited[key] >= v {

			} else {
				b = false
				break
			}
		}
		return b
	}

	left, right := 0, 0
	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			mvisited[s[right]]++
		}
		right++

		for checked() {
			temp := s[left:right]

			if len(temp) < length {
				length = len(temp)
				res = temp
			}

			u := s[left]
			if _, ok := m[u]; !ok {
				left++
			} else {
				mvisited[s[left]]--
				left++
			}
		}
	}

	return res
}

//滑动窗口 优化了一下 主要是框架  记住框架结构
func minWindow(s string, t string) string {
	length := math.MaxInt
	res := ""

	valid := 0
	visited := make(map[byte]int)

	need := 0
	m := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		m[t[i]]++
		need++
	}

	var checked func() bool
	checked = func() bool {
		return valid == need
	}

	left, right := 0, 0
	for right < len(s) {
		if _, ok := m[s[right]]; ok {

			visited[s[right]]++
			if visited[s[right]] <= m[s[right]] {
				valid++
			}
		}
		right++

		for checked() {
			temp := s[left:right]

			if len(temp) < length {
				length = len(temp)
				res = temp
			}

			u := s[left]
			if _, ok := m[u]; !ok {
				left++
			} else {

				visited[u]--
				if visited[u] < m[u] {
					valid--
				}
				left++
			}
		}
	}

	return res
}

//字符串排列 滑动窗口 还是写的太丑了点 没有掌握精髓  尤其是左边窗口收缩的时机!!!
func checkInclusion1(s1 string, s2 string) bool {
	m := make(map[byte]int)
	need := 0
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		need++
	}

	window := make(map[byte]int)
	isvalid := 0

	if len(s1) > len(s2) {
		return false
	}

	left, right := 0, len(s1)-1
	for i := left; i <= right; i++ {
		if _, ok := m[s2[i]]; !ok {
			continue
		} else {
			window[s2[i]]++
			if window[s2[i]] <= m[s2[i]] {
				isvalid++
			}
		}
	}

	if isvalid == need {
		return true
	}

	for right < len(s2)-1 {
		right++
		if _, ok := m[s2[right]]; !ok {

		} else {
			window[s2[right]]++
			if window[s2[right]] <= m[s2[right]] {
				isvalid++
			}
		}

		if _, ok := m[s2[left]]; !ok {

		} else {
			window[s2[left]]--
			if window[s2[left]] < m[s2[left]] {
				isvalid--
			}
		}
		left++

		if isvalid == need {
			return true
		}
	}
	return false

}

//滑动窗口 优化 变成模版类型 关键在于左边何时开始收缩 开始收缩的时机  还可以优化  在于 能否用数组来做呢 26个字母如果规定好
func checkInclusion2(s1 string, s2 string) bool {
	m := make(map[byte]int)
	need := 0
	for i := 0; i < len(s1); i++ {
		m[s1[i]]++
		need++
	}

	window := make(map[byte]int)
	isvalid := 0

	if len(s1) > len(s2) {
		return false
	}

	left, right := 0, 0
	for right < len(s2) {
		if _, ok := m[s2[right]]; !ok {

		} else {
			window[s2[right]]++
			//这里要很注意 如果++之后发现 窗口里的值小于等于的 那就说明是个有效的字母 isvalid就要++
			if window[s2[right]] <= m[s2[right]] {
				isvalid++
			}
		}
		//别忘了自增
		right++

		//这里很关键  不同的题目这里就可能不同 因为窗口的大小是固定的 所以当右-左=窗口大小的时候 就要进入判断是否满足要求
		for right-left >= len(s1) {
			if isvalid == need {
				return true
			}

			if _, ok := m[s2[left]]; !ok {

			} else {
				//这里要很注意  如果--之后 发现窗口的值小于 那就说明原来是个有效的字母 isvalid就要--
				window[s2[left]]--
				if window[s2[left]] < m[s2[left]] {
					isvalid--
				}
			}
			//别忘了自增
			left++
		}
	}

	return false
}

//不用map来做 而是用数组来做 数组的效率很高 好像还是一般般诶
func checkInclusion3(s1 string, s2 string) bool {
	ints1 := [26]int{}
	for i := 0; i < len(s1); i++ {
		ints1[s1[i]-'a']++
	}

	ints2 := [26]int{}
	left, right := 0, 0
	for right < len(s2) {
		ints2[s2[right]-'a']++
		right++

		for right-left >= len(s1) {
			if ints1 == ints2 {
				return true
			}

			ints2[s2[left]-'a']--
			left++
		}
	}

	return false
}

//数组 关键还是如何处理  利用go数组可以做相等比较的特性 来整活啊
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	ints1 := [26]int{}
	ints2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		ints1[s1[i]-'a']++
		ints2[s2[i]-'a']++
	}

	if ints1 == ints2 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		ints2[s2[i]-'a']++
		ints2[s2[i-len(s1)]-'a']--

		if ints1 == ints2 {
			return true
		}
	}

	return false
}

//异位词  滑动窗口 使用map 来做的 主要记住框架  优化可以使用数组 我试试看
func findAnagrams1(s string, p string) []int {
	res := make([]int, 0, 0)
	m := make(map[byte]int)
	need := len(p)
	for i := 0; i < len(p); i++ {
		m[p[i]]++
	}

	Visited := make(map[byte]int)
	valid := 0

	left, right := 0, 0
	for right < len(s) {
		if _, ok := m[s[right]]; !ok {

		} else {
			Visited[s[right]]++
			if Visited[s[right]] <= m[s[right]] {
				valid++
			}
		}
		right++

		for right-left >= len(p) {
			if need == valid {
				res = append(res, left)
			}

			if _, ok := m[s[left]]; !ok {

			} else {
				Visited[s[left]]--
				if Visited[s[left]] < m[s[left]] {
					valid--
				}
			}

			left++
		}
	}

	return res
}

//使用数组 很快很快 滑动窗口的框架记住了!!!!!!!  左边窗口收缩的时机
func findAnagrams2(s string, p string) []int {
	res := make([]int, 0, 0)
	ints1 := [26]int{}
	ints2 := [26]int{}
	for i := 0; i < len(p); i++ {
		ints1[p[i]-'a']++
	}

	left, right := 0, 0
	for right < len(s) {
		ints2[s[right]-'a']++
		right++

		for right-left >= len(p) {
			if ints1 == ints2 {
				res = append(res, left)
			}

			ints2[s[left]-'a']--
			left++
		}
	}

	return res
}

//思路也很好
func findAnagrams(s string, p string) []int {
	res := make([]int, 0, 0)
	ints1 := [26]int{}
	ints2 := [26]int{}
	if len(p) > len(s) {
		return []int{}
	}

	left := 0
	for i := 0; i < len(p); i++ {
		ints1[p[i]-'a']++
		ints2[s[i]-'a']++
	}

	if ints1 == ints2 {
		res = append(res, left)
	}

	for i := len(p); i < len(s); i++ {
		ints2[s[i]-'a']++
		ints2[s[i-len(p)]-'a']--
		left++
		if ints1 == ints2 {
			res = append(res, left)
		}
	}

	return res
}

//abcabcbb
//3 滑动窗口  这样做很笨  得改一下 不能右边一划就判断 如果没问题就该继续划
func lengthOfLongestSubstring1(s string) int {
	res := 0
	m := make(map[byte]int)
	var check func() bool
	check = func() bool {
		b := true
		for _, value := range m {
			if value > 1 {
				return false
			}
		}
		return b
	}

	left, right := 0, 0
	for right < len(s) {

		if m[s[right]] < 1 {
			if right-left+1 > res {
				res = right - left + 1
			}
		}
		m[s[right]]++
		right++

		for !check() {
			m[s[left]]--

			left++
		}
	}

	return res
}

//修改了一下 如果没问题 就继续滑
func lengthOfLongestSubstring2(s string) int {
	res := 0
	m := make(map[byte]int)
	var check func() bool
	check = func() bool {
		b := true
		for _, value := range m {
			if value > 1 {
				return false
			}
		}
		return b
	}

	left, right := 0, 0
	for right < len(s) {

		m[s[right]]++

		//没问题就跳过检验环节
		if m[s[right]] <= 1 {
			if right-left+1 > res {
				res = right - left + 1
			}
		} else {
			//有问题就进入检验环节
			for !check() {
				m[s[left]]--

				left++
			}
		}
		right++

	}

	return res
}

//修改了一下 如果没问题 就继续滑
func lengthOfLongestSubstring(s string) int {
	res := 0
	m := make(map[byte]int)

	left, right := 0, 0
	for right < len(s) {

		m[s[right]]++

		//没问题就跳过检验环节
		if m[s[right]] <= 1 {
			if right-left+1 > res {
				res = right - left + 1
			}
		} else {
			//有问题就进入检验环节
			for m[s[right]] > 1 {
				m[s[left]]--
				left++
			}
		}
		right++

	}

	return res
}

//动态规划
func lengthOfLIS1(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	dp := make([]int, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	res := math.MinInt
	for i := 0; i < len(dp); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

//二分法来做  时间复杂度nlog(n)
func lengthOfLIS(nums []int) int {
	res := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		index := sort.SearchInts(res, nums[i])
		if index > len(res)-1 {
			//如果很大很大 那么就说明是新的 可以增加长度
			res = append(res, nums[i])
		} else {
			//如果可以被替换 那就替换掉没关系
			res[index] = nums[i]
		}
	}

	return len(res)
}

//动态规划
func maxSubArray1(nums []int) int {
	dp := make([]int, len(nums), len(nums))
	if len(nums) == 0 {
		return 0
	}

	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	sort.Ints(dp)
	return dp[len(dp)-1]
}

//状态压缩+一次遍历 动态规划
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	cur := 1
	temp := nums[0]
	res := temp
	for cur != len(nums) {
		if temp < 0 {
			temp = nums[cur]
		} else {
			temp = temp + nums[cur]
		}

		if res < temp {
			res = temp
		}
		cur++
	}

	return res
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	root := &TreeNode{
		Val:   nums[maxIndex],
		Left:  nil,
		Right: nil,
	}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])

	return root
}

func largestInteger1(num int) int {
	itoa := strconv.Itoa(num)
	ji := make([]int, 0, 0)
	ou := make([]int, 0, 0)
	for i := 0; i < len(itoa); i++ {
		if i%2 == 0 {
			ou = append(ou, int(itoa[i]-'0'))
		} else {
			ji = append(ji, int(itoa[i]-'0'))
		}
	}
	sort.Ints(ji)
	sort.Ints(ou)
	res := make([]byte, 0, 0)
	for i := 0; i < len(itoa); i++ {
		if i%2 == 1 {
			res = append(res, byte(ji[len(ji)-1]+'0'))
			ji = ji[:len(ji)-1]
		} else {
			res = append(res, byte(ou[len(ou)-1]+'0'))
			ou = ou[:len(ou)-1]
		}
	}
	s := string(res)
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func largestInteger(num int) int {
	itoa := strconv.Itoa(num)
	ji := make([]int, 0, len(itoa))
	jilocation := make([]int, 0, len(itoa))
	ou := make([]int, 0, len(itoa))
	oulocation := make([]int, 0, len(itoa))
	for i := 0; i < len(itoa); i++ {
		if (itoa[i]-'0')%2 == 0 {
			oulocation = append(oulocation, i)
			ou = append(ou, int(itoa[i]-'0'))
		} else {
			jilocation = append(jilocation, i)
			ji = append(ji, int(itoa[i]-'0'))
		}
	}

	sort.Ints(oulocation)
	sort.Ints(jilocation)
	sort.Ints(ou)
	sort.Ints(ji)
	res := make([]byte, 0, 0)
	for i := 0; i < len(itoa); i++ {
		if len(jilocation) == 0 {
			res = append(res, byte(ou[len(ou)-1]+'0'))
			ou = ou[:len(ou)-1]
			oulocation = oulocation[1:]
		} else if len(oulocation) == 0 {
			res = append(res, byte(ji[len(ji)-1]+'0'))
			ji = ji[:len(ji)-1]
			jilocation = jilocation[1:]
		} else {
			if oulocation[0] < jilocation[0] {
				res = append(res, byte(ou[len(ou)-1]+'0'))
				ou = ou[:len(ou)-1]
				oulocation = oulocation[1:]
			} else {
				res = append(res, byte(ji[len(ji)-1]+'0'))
				ji = ji[:len(ji)-1]
				jilocation = jilocation[1:]
			}
		}
	}

	atoi, _ := strconv.Atoi(string(res))
	return atoi
}

func minimizeResult(expression string) string {
	res := ""
	min := math.MaxInt
	var calcu func(ex string)
	calcu = func(ex string) {
		left := strings.Index(ex, "(")
		right := strings.Index(ex, ")")
		plus := strings.Index(ex, "+")

		one := ex[left+1 : plus]
		two := ex[plus+1 : right]

		an1 := ex[:left]
		an2 := ex[right+1:]

		oneone, _ := strconv.Atoi(one)
		twotwo, _ := strconv.Atoi(two)
		an1an1, _ := strconv.Atoi(an1)
		an2an2, _ := strconv.Atoi(an2)

		ans := 0
		if an1an1 == 0 && an2an2 == 0 {
			ans = oneone + twotwo
		} else if an1an1 == 0 {
			ans = (oneone + twotwo) * an2an2
		} else if an2an2 == 0 {
			ans = (oneone + twotwo) * an1an1
		} else {
			ans = (oneone + twotwo) * an1an1 * an2an2
		}
		if ans < min {
			min = ans
			res = ex
		}
	}

	index := strings.Index(expression, "+")
	for i := 0; i < index; i++ {
		for j := index + 1; j < len(expression); j++ {

			temp := expression[:i] + "(" + expression[i:j+1] + ")" + expression[j+1:]
			calcu(temp)

		}
	}

	return res
}

type myheap []int

func (m *myheap) Len() int {
	return len(*m)
}

func (m *myheap) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *myheap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *myheap) Pop() interface{} {
	i := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return i
}

func maximumProduct(nums []int, k int) int {
	m := new(myheap)
	for i := 0; i < len(nums); i++ {
		heap.Push(m, nums[i])
	}

	res := 1
	for i := 0; i < k; i++ {
		x := heap.Pop(m).(int)
		x++
		heap.Push(m, x)
	}

	for i := 0; i < len(nums); i++ {
		res = (res % int(math.Pow(10, 9)+7)) * ((*m)[i] % int(math.Pow(10, 9)+7))
	}
	res = res % int(math.Pow(10, 9)+7)
	return res
}

func deleteText(article string, index int) string {
	u := article[index]
	if u == ' ' {
		return article
	} else {
		pre := strings.LastIndex(article[:index], " ")
		cur := strings.Index(article[index:], " ")
		var a string

		if pre == -1 {
			a = ""
		} else {
			a = article[:pre]
		}

		var b string
		if cur == -1 {
			b = ""
		} else {
			b = article[index+cur:]
		}

		if pre == -1 && cur == -1 {
			return ""
		} else if pre == -1 {
			return b[1:]
		}
		return a + b
	}
}

func lightSticks(height int, width int, indices []int) []int {
	min := math.MaxInt
	res := make([]int, 0, 0)
	matrix := make([][]int, height, height)
	mVisited := make(map[[2]int]int)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, width, width)
	}

	m := make(map[[3]int]int)

	mod := width + width + 1
	for i := 0; i < len(indices); i++ {
		x := indices[i] % mod
		if x >= width {
			xx := indices[i] / (height + width + 1)
			yy := x - width
			//向下
			m[[3]int{xx, yy, 2}]++
			m[[3]int{xx + 1, yy, 1}]++
		} else {
			//向右
			xx := indices[i] / (height + width + 1)
			yy := x
			m[[3]int{xx, yy, 4}]++
			m[[3]int{xx, yy + 1, 3}]++
		}
	}

	ints := make([][2]int, 0, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			ints = append(ints, [2]int{i, j})
		}
	}

	temp := make([][2]int, 0, 0)

	var travel func(x, y int, ans int) int
	travel = func(x, y int, ans int) int {
		if len(temp) == 0 {
			return ans
		}

		for len(temp) != 0 {
			l := len(temp)
			for i := 0; i < l; i++ {
				x1 := temp[0][0]
				y1 := temp[0][1]
				temp = temp[1:]

				if x1-1 >= 0 && m[[3]int{x1, y1, 1}] < 1 && mVisited[[2]int{x1 - 1, y1}] < 1 {
					if mVisited[[2]int{x1, y1}] >= 1 {
						continue
					}

					if x1-1 >= 0 {
						if mVisited[[2]int{x1 - 1, y1}] < 1 && m[[3]int{x1, y1, 1}] < 1 {
							temp = append(temp, [2]int{x1 - 1, y1})
						}
					}
					if x1+1 <= height {
						if mVisited[[2]int{x1 + 1, y1}] < 1 && m[[3]int{x1, y1, 2}] < 1 {
							temp = append(temp, [2]int{x1 + 1, y1})
						}
					}
					if y1-1 >= 0 {
						if mVisited[[2]int{x1, y1 - 1}] < 1 && m[[3]int{x1, y1, 3}] < 1 {
							temp = append(temp, [2]int{x1, y1 - 1})
						}
					}
					if y1+1 <= width {
						if mVisited[[2]int{x1, y1 + 1}] < 1 && m[[3]int{x1, y1, 4}] < 1 {
							temp = append(temp, [2]int{x1, y1 + 1})
						}
					}
					mVisited[[2]int{x1, y1}]++
				} else if x1+1 <= height && m[[3]int{x1, y1, 2}] < 1 && mVisited[[2]int{x1 + 1, y1}] < 1 {
					if mVisited[[2]int{x1, y1}] >= 1 {
						continue
					}

					if x1-1 >= 0 {
						if mVisited[[2]int{x1 - 1, y1}] < 1 && m[[3]int{x1, y1, 1}] < 1 {
							temp = append(temp, [2]int{x1 - 1, y1})
						}
					}
					if x1+1 <= height {
						if mVisited[[2]int{x1 + 1, y1}] < 1 && m[[3]int{x1, y1, 2}] < 1 {
							temp = append(temp, [2]int{x1 + 1, y1})
						}
					}
					if y1-1 >= 0 {
						if mVisited[[2]int{x1, y1 - 1}] < 1 && m[[3]int{x1, y1, 3}] < 1 {
							temp = append(temp, [2]int{x1, y1 - 1})
						}
					}
					if y1+1 <= width {
						if mVisited[[2]int{x1, y1 + 1}] < 1 && m[[3]int{x1, y1, 4}] < 1 {
							temp = append(temp, [2]int{x1, y1 + 1})
						}
					}
					mVisited[[2]int{x1, y1}]++
				} else if y1-1 >= 0 && m[[3]int{x1, y1, 3}] < 1 && mVisited[[2]int{x1, y1 - 1}] < 1 {
					if mVisited[[2]int{x1, y1}] >= 1 {
						continue
					}

					if x1-1 >= 0 {
						if mVisited[[2]int{x1 - 1, y1}] < 1 && m[[3]int{x1, y1, 1}] < 1 {
							temp = append(temp, [2]int{x1 - 1, y1})
						}
					}
					if x1+1 <= height {
						if mVisited[[2]int{x1 + 1, y1}] < 1 && m[[3]int{x1, y1, 2}] < 1 {
							temp = append(temp, [2]int{x1 + 1, y1})
						}
					}
					if y1-1 >= 0 {
						if mVisited[[2]int{x1, y1 - 1}] < 1 && m[[3]int{x1, y1, 3}] < 1 {
							temp = append(temp, [2]int{x1, y1 - 1})
						}
					}
					if y1+1 <= width {
						if mVisited[[2]int{x1, y1 + 1}] < 1 && m[[3]int{x1, y1, 4}] < 1 {
							temp = append(temp, [2]int{x1, y1 + 1})
						}
					}
					mVisited[[2]int{x1, y1}]++
				} else if y1+1 <= width && m[[3]int{x1, y1, 4}] < 1 && mVisited[[2]int{x1, y1 + 1}] < 1 {
					if mVisited[[2]int{x1, y1}] >= 1 {
						continue
					}

					if x1-1 >= 0 {
						if mVisited[[2]int{x1 - 1, y1}] < 1 && m[[3]int{x1, y1, 1}] < 1 {
							temp = append(temp, [2]int{x1 - 1, y1})
						}
					}
					if x1+1 <= height {
						if mVisited[[2]int{x1 + 1, y1}] < 1 && m[[3]int{x1, y1, 2}] < 1 {
							temp = append(temp, [2]int{x1 + 1, y1})
						}
					}
					if y1-1 >= 0 {
						if mVisited[[2]int{x1, y1 - 1}] < 1 && m[[3]int{x1, y1, 3}] < 1 {
							temp = append(temp, [2]int{x1, y1 - 1})
						}
					}
					if y1+1 <= width {
						if mVisited[[2]int{x1, y1 + 1}] < 1 && m[[3]int{x1, y1, 4}] < 1 {
							temp = append(temp, [2]int{x1, y1 + 1})
						}
					}
					mVisited[[2]int{x1, y1}]++
				} else {
					return math.MaxInt

				}

			}
			ans++
		}
		return ans
	}

	for i := 0; i <= height; i++ {
		for j := 0; j <= width; j++ {
			if i-1 >= 0 && m[[3]int{i, j, 1}] < 1 {
				temp = append(temp, [2]int{i, j})
				ans := travel(i, j, 0)
				if ans < min {
					res = []int{}
					xx := (i * (width + 1)) + j
					min = ans
					res = append(res, xx)
				} else if ans == min {
					xx := (i * (width + 1)) + j
					res = append(res, xx)
				}
				mVisited = map[[2]int]int{}
			} else if i+1 <= height && m[[3]int{i, j, 2}] < 1 {
				temp = append(temp, [2]int{i, j})
				ans := travel(i, j, 0)
				if ans < min {
					res = []int{}
					xx := (i * (width + 1)) + j
					min = ans
					res = append(res, xx)
				} else if ans == min {
					xx := (i * (width + 1)) + j
					res = append(res, xx)
				}
				mVisited = map[[2]int]int{}
			} else if j-1 >= 0 && m[[3]int{i, j, 3}] < 1 {
				temp = append(temp, [2]int{i, j})
				ans := travel(i, j, 0)
				if ans < min {
					res = []int{}
					xx := (i * (width + 1)) + j
					min = ans
					res = append(res, xx)
				} else if ans == min {
					xx := (i * (width + 1)) + j
					res = append(res, xx)
				}
				mVisited = map[[2]int]int{}
			} else if j+1 <= width && m[[3]int{i, j, 4}] < 1 {
				temp = append(temp, [2]int{i, j})
				ans := travel(i, j, 0)
				if ans < min {
					res = []int{}
					xx := (i * (width + 1)) + j
					min = ans
					res = append(res, xx)
				} else if ans == min {
					xx := (i * (width + 1)) + j
					res = append(res, xx)
				}
				mVisited = map[[2]int]int{}
			}

		}
	}

	if len(res) == 0 || min == math.MaxInt {
		return []int{}
	}
	return res
}

var matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

func bianliZheng() {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func bianliFan() {
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

//斜着遍历  这里l充当的是一个偏移量 试想一下 如果l为1 那么i就和j相等 就是对角线 然后l变大 j每次都给的偏移量比i大一点
//就逐步向右边滑动
func bianliXieYou() {
	for l := 1; l <= len(matrix); l++ {
		for i := 0; i <= len(matrix)-l; i++ {
			j := l + i - 1
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

}

//斜着遍历  这里l充当的是一个偏移量 试想一下 如果l为1 那么i就和j相等 就是对角线 然后l变大 j每次都给的偏移量比i大一点
func bianliXieDouble() {
	for l := 1; l <= len(matrix); l++ {
		for i := 0; i <= len(matrix)-l; i++ {
			j := l + i - 1
			fmt.Printf("%d\t", matrix[i][j])
			if i != j {
				fmt.Printf("%d\t", matrix[j][i])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

//调整顺序就行了其实
func bianliZuoxiajiao() {
	for l := len(matrix); l >= 1; l-- {
		for i := 0; i <= len(matrix)-l; i++ {
			j := l + i - 1
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

//对角线右边 从下到上 从左到右
func bianlidowotoUpLeftToRight() {
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := i + 1; j < len(matrix[i]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
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
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text2[i-1] == text1[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp[len(text2)][len(text1)]
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1, len(word2)+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b

	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1]+1, dp[i][j-1]+1), dp[i-1][j]+1)
			}
		}

	}

	return dp[len(word1)][len(word2)]
}

//其实就是动态规划  给定数组 求是否有数组中和为x 的情况
//完全背包问题 主要在于dp数组的构造 非常吊 取max 计算sum  计算一半target 如果sum不是偶数或者max>target 或者个数<2 都说明不行
//dp数组的定义为  dp[i][j] 在索引 nums[0-i]中 能够刚好计算等于j的是true还是false
//所以 j最大就是target  而i最大就是nums的最大索引
func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	sum := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	if max > target {
		return false
	}

	dp := make([][]bool, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1, target+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if nums[i] <= j {
				//不选择
				a := dp[i-1][j]

				//选择的话 就要确保前面没有选过 什么叫做前面没有选过呢 那当然就是[i-1][j-nums[i]]  例如nums[2]=2 dp[2][5]表示索引0-2中有无和恰好为5的
				//那么有两种情况
				//1.就是dp[1][5]  即不选择 索引0-1中是否有和恰好为5的
				//2.dp[2-1][5-nums[2]]=dp[1][3] 也就是选择 选择就会占用 占用就要去寻找dp[1][3] 即索引0-1中有没有和恰好为3的 只要有 那么就代表凑到了!
				b := dp[i-1][j-nums[i]]
				dp[i][j] = a || b
			} else {
				//选择不了 太大了 一选就超
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][target]
}

//N=3 W=4  呜呜呜我怎么可以这么 菜呜呜
//wt=[2 1 3] val =[4 2 3]   背包问题 啊啊啊啊啊 dp数组的定义很关键啊啊啊啊 dp[i][j]是前i个物品中 限重j时 的最大价值
func knapSack(w, n int, wt, val []int) int {
	dp := make([][]int, n+1, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, w+1, w+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}

	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 0
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
			if wt[i-1] > j {
				//选不了 太大了
				dp[i][j] = dp[i-1][j]
			} else {
				//不选择了
				temp1 := dp[i-1][j]

				temp2 := dp[i-1][j-wt[i-1]] + val[i-1]
				dp[i][j] = max(temp1, temp2)
			}
		}
	}

	return dp[n][w]
}

//其实就是动态规划  给定数组 求是否有数组中和为x 的情况
func nSum(nums []int, target int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	target = sum / 2
	dp := make([][]bool, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, target+1, target+1)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	dp[0][nums[0]] = true

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if j >= nums[i-1] {
				//不选
				a := dp[i-1][j]

				//选择
				b := dp[i-1][j-nums[i-1]]
				dp[i][j] = a || b
			} else {
				//选不了啊 根本 因为当前索引的值太大了 大到都超过j了
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)][target]
}

//完全背包问题  在于硬币的数量可以有很无限个 dp[i][j]定义为 索引为0-i中的数组,能够构造为j的结果是多少  先构造base case
func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1, amount+1)
	}

	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = -1
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}
	var min func(a, b int) int
	min = func(a, b int) int {
		if a == -1 && b == -1 {
			return -1
		} else if a == -1 {
			return b
		} else if b == -1 {
			return a
		} else if a < b {
			return a
		}
		return b
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if coins[i-1] > j {
				//选不了 太大了  只能用上一层的了
				dp[i][j] = dp[i-1][j]
			} else {
				//选的了的情况下 判断前面能不能构造的出来
				if dp[i][j-coins[i-1]] != -1 {
					//能构造的出来就 去比谁小
					dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
				} else {
					//构造不出来就用曾经的
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return dp[len(coins)][amount]
}

//从下到上 从左到右 遍历  最长回文子序列  dp[i][j]的定义: 索引i-j处最长回文子序列的结果是多少
func longestPalindromeSubseq1(s string) int {
	dp := make([][]int, len(s), len(s))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s), len(s))
		dp[i][i] = 1
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := len(dp) - 1; i >= 0; i-- {
		for j := i + 1; j < len(dp[i]); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}

//对角线遍历  给一个l的偏移量  状态压缩还不会 没关系~
func longestPalindromeSubseq2(s string) int {
	dp := make([][]int, len(s), len(s))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s), len(s))
		dp[i][i] = 1
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for l := 2; l <= len(dp); l++ {
		for i := 0; i <= len(dp)-l; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}

	return dp[0][len(s)-1]
}

//构造回文串  dp定义  扩张顺序  状态转移
func minInsertions(s string) int {
	dp := make([][]int, len(s), len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s), len(s))
	}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for l := 2; l <= len(dp); l++ {
		for i := 0; i <= len(dp)-l; i++ {
			j := l + i - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[0][len(s)-1]
}

//正则表达式匹配 做不了做不了
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1, len(p)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s)+1, len(s)+1)
	}
	if strings.Contains(p, ".*") {
		return true
	}

	dp[0][0] = true
	s = " " + s
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			if p[i-1] == '.' {
				if j == 0 {
					dp[i][j] = false
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else if p[i-1] == '*' {
				zero := dp[i-2][j]
				one := dp[i-1][j]
				more := false
				k := j

				for ; k-1 >= 0 && s[k-1] == s[k]; k-- {
					more = more || dp[i-1][k]
				}
				more = more || dp[i-1][k]
				dp[i][j] = zero || one || more

			} else if p[i-1] == s[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}

	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}

//打家劫舍1
func rob1(nums []int) int {
	dp := make([]int, len(nums)+1, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
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

//打家劫舍2
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return max(rob1(nums[:len(nums)-1]), rob1(nums[1:]))
}

//打家劫舍3
func rob(root *TreeNode) int {
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
		if root == nil {
			return 0
		}

		a := root.Val
		al := 0
		ar := 0
		if root.Left != nil {
			if v, ok := m[root.Left.Left]; ok {
				al = v
			} else {
				al = travel(root.Left.Left)
			}

			if v, ok := m[root.Left.Right]; ok {
				al += v
			} else {
				al += travel(root.Left.Right)
			}
		}
		if root.Right != nil {
			if v, ok := m[root.Right.Left]; ok {
				ar = v
			} else {
				ar = travel(root.Right.Left)
			}

			if v, ok := m[root.Right.Right]; ok {
				ar += v
			} else {
				ar += travel(root.Right.Right)
			}
		}

		b := 0
		if v, ok := m[root.Left]; ok {
			b = v
		} else {
			b = travel(root.Left)
		}

		c := 0
		if v, ok := m[root.Right]; ok {
			c = v
		} else {
			c = travel(root.Right)
		}

		m[root] = max(a+al+ar, b+c)
		return m[root]
	}

	return travel(root)
}

//动态规划 完全背包问题~ 树树酱牛皮!!! dp[i][j]定义为 索引0-i中 能够凑到j的结果有几种  这里的j 很难表达 因为有负数 所以得emm对
func findTargetSumWays1(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	//如果总和不够大 或者 总和不够小 说明就无了
	if sum < target || sum < -target {
		return 0
	}

	dp := make([][]int, len(nums), len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*sum+1, 2*sum+1)
	}

	//base case 这里要处理好 是++ 因为j需要代表负数 但是呢j又是索引从0开始的
	//比如 1 1 1 1 1 那么最小的j就是-5  最大的j就是5  如何用j来表达  j就是 0 1 2 3 4 5 6 7 8 9 10 11
	//其中0代表-5  11代表5  5代表0
	dp[0][2*sum-(sum-nums[0])]++
	dp[0][2*sum-(sum+nums[0])]++

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			//状态转移
			if j-nums[i] >= 0 && j+nums[i] < len(dp[i]) {
				//如果没有越界 那么就是前面的两边转移过来的 从j+nums[i] 或者j-nums[i]
				dp[i][j] = dp[i-1][j+nums[i]] + dp[i-1][j-nums[i]]
			} else if j-nums[i] >= 0 && j+nums[i] >= len(dp[i]) {
				//如果越界了 就是从前面的一边转移过来的
				dp[i][j] = dp[i-1][j-nums[i]]
			} else if j-nums[i] < 0 && j+nums[i] < len(dp[i]) {
				dp[i][j] = dp[i-1][j+nums[i]]
			}
		}

	}

	return dp[len(nums)-1][2*sum-(sum-target)]
}

//回溯 很慢~
func findTargetSumWays(nums []int, target int) int {
	res := 0
	var travel func(nums []int, target int)
	travel = func(nums []int, target int) {
		if len(nums) == 0 {
			if target == 0 {
				res++
			}
			return
		}

		x := nums[0]
		travel(nums[1:], target-x)
		travel(nums[1:], target+x)

	}
	travel(nums, target)
	return res
}

//题目有点理解错了 dp[i][j][k]定义为  前i个中 最多有j个0 k个1的最长子集长度是
func findMaxForm(strs []string, m int, n int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][][]int, len(strs)+1, len(strs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n+1, n+1)
		}
	}

	zeroOne := make([][2]int, len(strs), len(strs))
	for i := 0; i < len(strs); i++ {
		zero := strings.Count(strs[i], "0")
		one := len(strs[i]) - zero
		zeroOne[i] = [2]int{zero, one}
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			for k := 0; k < len(dp[i][j]); k++ {
				mm := zeroOne[i-1][0]
				nn := zeroOne[i-1][1]

				//如果超过了 简直选不了
				if mm > j || nn > k {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = max(dp[i-1][j][k], dp[i-1][j-mm][k-nn]+1)
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

//动态规划 单词拆分  dp[i][j]定义: i代表s中的前i个字符 j代表字典中的前j个
//dp[i][j] 就是 s中前i个字符 能否有 字典中的前j个东东构成
func wordBreak(s string, wordDict []string) bool {
	dp := make([][]bool, len(s)+1, len(s)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(wordDict)+1, len(wordDict)+1)
	}

	//base case 前0个当然可以由字典中的任何东东构成辣
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = true
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			str := s[:i]
			//选与不选

			if len(str) < len(wordDict[j-1]) {
				//选不了  太大了
				dp[i][j] = dp[i][j-1]
			} else {
				//选的话 就要逐个去判断 首先要保证字典里的串长度要小
				for k := 0; k < j; k++ {
					if i >= len(wordDict[k]) {
						sss := s[i-len(wordDict[k]) : i]
						//其次判断是否相等 并且前面的会不会被成功构成
						dp[i][j] = dp[i][j] || (sss == wordDict[k] && dp[i-len(wordDict[k])][j])
					}
				}
			}
		}
	}

	return dp[len(s)][len(wordDict)]
}

//还是不太会呜呜呜呜呜
func combinationSum4(nums []int, target int) int {
	dp := make([][]int, len(nums)+1, len(nums)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, target+1, target+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			for k := 0; k < i; k++ {
				if j < nums[k] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] += dp[i][j-nums[k]]
				}
			}
		}
	}

	return dp[len(nums)][target]
}

func findClosestNumber(nums []int) int {
	res := math.MaxInt
	var abs func(a int) int
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	for i := 0; i < len(nums); i++ {
		if abs(nums[i])-0 < abs(res) {
			res = nums[i]
		} else if abs(nums[i])-0 == abs(res) {
			if nums[i] > res {
				res = nums[i]
			} else {

			}
		}
	}
	return res
}

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	res := int64(0)
	count := total / cost1
	for i := 0; i <= count; i++ {
		res += int64((total-cost1*i)/cost2 + 1)
	}
	return res
}

type ATM struct {
	cash     []int
	maxInt   int
	totalSum int
}

func Constructor() ATM {
	return ATM{cash: make([]int, 5, 5), maxInt: 0}
}

func (this *ATM) Deposit(banknotesCount []int) {
	money := []int{20, 50, 100, 200, 500}

	for i := 0; i < len(banknotesCount); i++ {
		this.cash[i] += banknotesCount[i]
		this.totalSum += money[i] * banknotesCount[i]
		if this.cash[i] > 0 {
			this.maxInt = i
		}
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5, 5)

	if this.totalSum < amount {
		return []int{-1}
	}

	var travel func(b bool, left int) bool
	travel = func(b bool, left int) bool {
		if left == 0 {
			return true
		} else if left < 0 {
			return false
		}

		money := []int{20, 50, 100, 200, 500}

		if !b {

			//l, r := 1, this.cash[this.maxInt]
			//for l <= r {
			//	temp := this.maxInt
			//
			//	mid := l + (r-l)/2
			//	this.cash[this.maxInt] -= mid
			//	res[this.maxInt] += mid
			//
			//	if travel(true, left-money[this.maxInt]*mid) {
			//		return true
			//	}
			//
			//	if this.cash[this.maxInt] == 0 {
			//		for i := this.maxInt - 1; i >= 0; i-- {
			//			if this.cash[i] > 0 {
			//				this.maxInt = i
			//				break
			//			}
			//		}
			//	}
			//
			//	res[this.maxInt] -= temp
			//	this.cash[temp] += mid
			//	this.maxInt = temp
			//}

			i := left / money[this.maxInt]
			if i == 0 {
				return false
			} else if i > this.cash[this.maxInt] {
				temp := this.maxInt
				temp1 := this.cash[this.maxInt]

				res[this.maxInt] += this.cash[this.maxInt]
				this.cash[this.maxInt] = 0

				if this.cash[this.maxInt] == 0 {
					for i := this.maxInt - 1; i >= 0; i-- {
						if this.cash[i] > 0 {
							this.maxInt = i
							break
						}
					}
				}

				this.totalSum -= money[temp] * temp1

				if travel(true, left-money[temp]*temp1) {
					return true
				}

				this.totalSum += money[temp] * temp1

				res[temp] -= temp1
				this.cash[temp] += temp1
				this.maxInt = temp
				return false
			} else {
				temp := this.maxInt

				res[this.maxInt]++
				this.cash[this.maxInt]--

				if this.cash[this.maxInt] == 0 {
					for i := this.maxInt - 1; i >= 0; i-- {
						if this.cash[i] > 0 {
							this.maxInt = i
							break
						}
					}
				}

				this.totalSum -= money[temp]

				if travel(true, left-money[temp]) {
					return true
				}

				this.totalSum += money[temp]

				res[temp]--
				this.cash[temp]++
				this.maxInt = temp
			}

			return false
		}

		for i := len(this.cash) - 1; i >= 0; i-- {
			if this.cash[i] > 0 {
				this.cash[i]--
				res[i]++
				this.totalSum -= money[i]

				if travel(true, left-money[i]) {
					return true
				}

				this.totalSum += money[i]
				this.cash[i]++
				res[i]--

			}
		}
		return false
	}

	if !travel(false, amount) {
		return []int{-1}
	}
	return res
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */

func digitSum(s string, k int) string {
	res := make([]string, 0, 0)
	for len(s) > k {
		for len(s) >= k {
			res = append(res, s[:k])
			s = s[k:]
		}

		if len(s) > 0 {
			res = append(res, s)
		}

		s = ""
		for i := 0; i < len(res); i++ {
			temp := 0
			for j := 0; j < len(res[i]); j++ {
				temp += int(res[i][j] - '0')
			}
			itoa := strconv.Itoa(temp)
			s += itoa
		}
		res = []string{}
	}
	return s
}

func minimumRounds1(tasks []int) int {
	sort.Ints(tasks)
	dp := make([]int, len(tasks)+1)
	dp[0] = 0
	dp[1] = -1
	for i := 2; i < len(dp); i++ {
		if tasks[i-1] == tasks[i-2] {
			if dp[i-1] == -1 {
				if dp[i-2] == -1 {
					dp[i] = -1
				} else {
					dp[i] = dp[i-2] + 1
				}
			} else if tasks[i-1] == tasks[i-3] {
				if dp[i-1] == dp[i-2] {
					//dp[i] = -1
					//dp[i] = dp[i-1] + 1
					if i-3 >= 0 && dp[i-3] == -1 {
						dp[i] = dp[i-2] + 1
					} else {
						dp[i] = dp[i-1]
					}
				} else {
					dp[i] = dp[i-1]
				}
			}
		} else {
			dp[i] = -1
		}
	}

	return dp[len(tasks)]
}

func minimumRounds(tasks []int) int {
	sort.Ints(tasks)
	dp := make([]int, len(tasks)+1)
	dp[0] = 0
	dp[1] = -1
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 2; i < len(dp); i++ {
		if tasks[i-1] == tasks[i-2] {
			if i-3 < 0 {
				if dp[i-1] == -1 {
					dp[i] = 1
				} else {
					dp[i] = dp[i-1]
				}
			} else {
				if dp[i-1] == -1 {
					dp[i] = dp[i-2] + 1
				} else if dp[i-2] != -1 {
					if dp[i-3] == -1 {
						dp[i] = dp[i-2] + 1
					} else if dp[i-2] == -1 {
						dp[i] = dp[i-3] + 1
					} else {
						dp[i] = min(dp[i-3]+1, dp[i-2]+1)
					}
				} else {
					dp[i] = dp[i-1]
				}
			}
		} else {
			dp[i] = -1
			if dp[i] == dp[i-1] {
				return -1
			}

		}
	}

	return dp[len(tasks)]
}

func isPalindrome(head *ListNode) bool {
	left := head

	var travel func(root *ListNode) bool
	travel = func(root *ListNode) bool {
		if root == nil {
			return true
		}
		b2 := travel(root.Next)
		b := left.Val == root.Val && b2
		left = left.Next
		return b
	}

	return travel(head)
}

//返回原始头节点
func reverseList1(head *ListNode) *ListNode {

	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		node := travel(head.Next)
		if node != nil {
			node.Next = head
		}
		head.Next = nil
		return head
	}

	node := travel(head)
	return node
}

//返回现在头节点
func reverseList(head *ListNode) *ListNode {

	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		node := travel(head.Next)
		if node != nil {
			head.Next.Next = head
		}
		head.Next = nil
		if node == nil {
			return head
		}
		return node
	}

	node := travel(head)
	return node
}

//反转前N个链表
func reverseN(head *ListNode, n int) *ListNode {
	var nxt *ListNode
	var travel func(head *ListNode, i int) *ListNode
	travel = func(head *ListNode, i int) *ListNode {
		if i == n {
			nxt = head.Next
			return head
		}

		node := travel(head.Next, i+1)
		if node != nil {
			head.Next.Next = head
		}
		head.Next = nil
		if node == nil {
			return head
		}
		return node
	}

	node := travel(head, 1)
	head.Next = nxt
	return node
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var pre *ListNode
	cur := head

	for left != 1 {
		pre = cur
		cur = cur.Next
		left--
		right--
	}

	sum := 0
	var nxt *ListNode
	temp := cur
	for 1 != right {
		sum++
		right--
		temp = temp.Next
	}
	nxt = temp.Next
	n := reverseN(cur, sum+1)
	cur.Next = nxt

	pre.Next = n
	return head
}

func distance(s, t string) int {
	res := 0
	l1 := len(s)
	l2 := len(t)
	left := 0
	right := 0

	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a < b {
			return b - a
		}
		return a - b
	}
	for right < l2 {

		for right-left == l1-1 {
			i := 0
			for l := left; l <= right; l++ {
				res += abs(int(t[l]), int(s[i]))
				i++
			}
			left++
		}
		right++
	}

	return res
}

func removeDigit(number string, digit byte) string {

	count := strings.Count(number, string(digit))
	if count == 1 {
		index := strings.Index(number, string(digit))
		return number[:index] + number[index+1:]
	}

	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			if i == len(number)-1 {
				return number[:len(number)-1]
			} else if number[i+1] <= number[i] {
				continue
			} else {
				return number[:i] + number[i+1:]
			}
		}
	}

	index := strings.LastIndex(number, string(digit))
	return number[:index] + number[index+1:]
}

func minimumCardPickup(cards []int) int {
	left, right := 0, 0
	res := math.MaxInt
	m := make(map[int]int)

	for right < len(cards) {

		if v, ok := m[cards[right]]; ok {
			temp := right - v + 1
			if temp < res {
				res = temp
			}
			m[cards[right]] = right
			left++
		} else {
			m[cards[right]] = right
		}

		right++
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

type LiliNode struct {
	Next  *LiliNode
	Val   int
	isEnd bool
}

//哈希表模拟 啊啊啊啊  切片怎么去重  先转成数组 用map来实现 map一个特别大的数组 能够容纳下切片的容量就可以辣 然后用==来实现就行了呀！！！
func countDistinct(nums []int, k int, p int) int {
	res := 0

	m := make(map[[200]int]bool)
	for i := 0; i < len(nums); i++ {
		cnt, idx, ints := 0, 0, [200]int{}
		for j := i; j < len(nums); j++ {
			if nums[j]%p == 0 {
				cnt++
				if cnt > k {
					break
				}
			} else {

			}
			ints[idx] = nums[j]
			idx++
			if !m[ints] {
				m[ints] = true
				res++
			}
		}
	}

	return res
}
