package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//matches := numberOfMatches(14)
	//fmt.Println(matches)

	//days := shipWithinDays([]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, 8)
	//fmt.Println(days)

	//frequent := mostFrequent([]int{2, 2, 2, 2, 3}, 2)
	//fmt.Println(frequent)

	//jumbled := sortJumbled([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	//fmt.Println(jumbled)

	//inRange := cellsInRange("A1:F1")
	//fmt.Println(inRange)

	//tree := createBinaryTree([][]int{{39, 70, 1}, {13, 39, 1}, {85, 74, 1}, {74, 13, 1}, {38, 82, 1}, {82, 85, 1}})
	//fmt.Println(tree)

	//parenthesis := generateParenthesis(3)
	//fmt.Println(parenthesis)

	//t := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	//
	//t2 := &TreeNode{
	//	Val: 2,
	//	Left: &TreeNode{
	//		Val:  1,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//trees := mergeTrees(t, t2)
	//fmt.Println(trees)

	//node := connect(&Node{
	//	Val: 1,
	//	Left: &Node{
	//		Val: 2,
	//		Left: &Node{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Right: &Node{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Next: nil,
	//	},
	//	Right: &Node{
	//		Val: 3,
	//		Left: &Node{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Right: &Node{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//			Next:  nil,
	//		},
	//		Next: nil,
	//	},
	//	Next: nil,
	//})
	//fmt.Println(node)

	//matrix := updateMatrix([][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 1, 1, 1}})
	//fmt.Println(matrix)

	//lists := mergeTwoLists(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}, &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//})
	//fmt.Println(lists)

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

	//indices := findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2)
	//fmt.Println(indices)

	//artifacts := digArtifacts(2, [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 1}})
	//fmt.Println(artifacts)

	//top := maximumTop([]int{0, 1, 2}, 3)
	//fmt.Println(top)

	//rotting := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	//fmt.Println(rotting)

	//i := combine(5, 4)
	//fmt.Println(i)

	//i := permute([]int{1, 2, 3})
	//fmt.Println(i)

	//permutation := letterCasePermutation("a1b2")
	//fmt.Println(permutation)

	//i := rob([]int{2, 1, 1, 2})
	//fmt.Println(i)

	//total := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 2}, {-4, 1, 8, 4}})
	//fmt.Println(total)

	//bits := reverseBits(0b00000010100101000001111010011100)
	//bits := reverseBits(0b11000101)
	//fmt.Println(0b10 & 0b10)
	//fmt.Println(bits)
	//fmt.Printf("%b\n", bits)

	//i := tribonacci(25)
	//fmt.Println(i)

	//lis := lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	//fmt.Println(lis)

	//consecutive := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	//consecutive := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	//fmt.Println(consecutive)

	//s := multiply("567", "1234")
	//s := multiply("498828660196", "840477629533")
	//s := multiply("123456789", "987654321")
	//s := multiply("3", "38")
	//fmt.Println(s)

	//largest := findKthLargest([]int{5, 2, 4, 1, 3, 6, 0}, 4)
	//fmt.Println(largest)

	//array := divideArray([]int{1, 2, 3, 4})
	//fmt.Println(array)

	//count := maximumSubsequenceCount("fwymvreuftzgrcrxczjacqovduqaiig", "yy")
	//count := maximumSubsequenceCount("zigfj", "ju")
	//count := maximumSubsequenceCount("abdcdbc", "ac")
	//fmt.Println(count)

	//array := halveArray([]int{32, 98, 23, 14, 67, 40, 26, 9, 96, 96, 91, 76, 4, 40, 42, 2, 31, 13, 16, 37, 62, 2, 27, 25, 100, 94, 14, 3, 48, 56, 64, 59, 33, 10, 74, 47, 73, 72, 89, 69, 15, 79, 22, 18, 53, 62, 20, 9, 76, 64})
	//fmt.Println(array)
	//fmt.Println(tiles)

	//valley := countHillValley([]int{6, 6, 5, 5, 4, 1})
	//fmt.Println(valley)

	//collisions := countCollisions("LLRR")
	//collisions := countCollisions("RLRSLL")
	//collisions := countCollisions("SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR")
	//fmt.Println(collisions)

	//points := maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0})
	//points := maximumBobPoints(89, []int{3, 2, 28, 1, 7, 1, 16, 7, 3, 13, 3, 5})
	//fmt.Println(points)

	//envelopes := maxEnvelopes([][]int{{46, 89}, {50, 53}, {52, 68}, {72, 45}, {77, 81}})
	//fmt.Println(envelopes)

	//subsequence1 := longestCommonSubsequence("bsbininm", "jmjkbkjkv")
	//fmt.Println(subsequence1)

	//distance := minDistance("dinitrophenylhydrazine", "acetylphenylhydrazine")
	//distance := minDistance("intention", "execution")
	//distance := minDistance("horse", "ros")
	//distance := minDistance("", "a")
	//distance := minDistance("sea", "eat")
	//fmt.Println(distance)

	//ints := pancakeSort([]int{1, 3, 2})
	//fmt.Println(ints)

	//sum := subarraySum([]int{1, -1, 0}, 0)
	//sum := subarraySum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0)
	//sum := subarraySum([]int{1}, 0)
	//sum := subarraySum([]int{-1, -1, 1}, 0)
	//sum := subarraySum([]int{1, 2, 3, 5, 8, 2}, 10)
	//sum := subarraySum([]int{1, -1, 0}, 0)
	//fmt.Println(sum)

	//primes := countPrimes(10)
	//primes := countPrimes(499979)
	//primes := countPrimes(5000000)
	//fmt.Println(primes)

	//pow := superPow(2147483647, []int{2, 0, 0})
	//fmt.Println(pow)

	//speed := minEatingSpeed([]int{3, 6, 7, 11}, 8)
	//speed := minEatingSpeed([]int{30, 11, 23, 4, 20}, 5)
	//fmt.Println(speed)

	//days1 := shipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3)
	//fmt.Println(days1)

	//i := trap([]int{4, 2, 0, 3, 2, 5})
	//i := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	//fmt.Println(i)

	//duplicates := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	//fmt.Println(duplicates)

	//palindrome := longestPalindrome("babad")
	//fmt.Println(palindrome)

	//jump := canJump([]int{2, 3, 1, 1, 4})
	//jump := canJump([]int{3, 2, 1, 0, 4})
	//jump := canJump([]int{1, 1, 2, 2, 0, 1, 1})
	//jump := canJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0})
	//jump := canJump([]int{1, 1, 0, 1})
	//jump := canJump([]int{2, 0, 0})
	//fmt.Println(jump)

	//i := jump([]int{2, 3, 1, 1, 4})
	//i := jump([]int{2, 3, 0, 1, 4})
	//i := jump([]int{1, 2})
	//fmt.Println(i)

	//difference := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	//fmt.Println(difference)

	//deletion := minDeletion([]int{1, 1, 2, 2, 3, 3})
	//fmt.Println(deletion)

	//palindrome := kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3)
	//palindrome := kthPalindrome([]int{2, 4, 6}, 4)
	//palindrome := kthPalindrome([]int{2, 201429812, 8, 520498110, 492711727, 339882032, 462074369, 9, 7, 6}, 1)
	//fmt.Println(109883145 > 10e7-10e6)
	//fmt.Println(16017953 > 10e8-10e7)
	//palindrome := kthPalindrome([]int{109883145, 67184890, 615116035, 730676834, 13, 700172947}, 7)

	//palindrome := kthPalindrome([]int{94, 18, 388481008, 38, 16017953, 16, 223505660, 78, 123699557, 346244579, 2}, 8)
	//fmt.Println(palindrome)

	//coins := maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 2)
	//coins := maxValueOfCoins([][]int{{1, 100, 3}, {7, 8, 9}}, 4)
	//coins := maxValueOfCoins([][]int{{100}, {100}, {100}, {100}, {100}, {100}, {1, 1, 1, 1, 1, 1, 700}}, 7)
	//fmt.Println(coins)

	//intervals := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	//intervals := eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}})
	//intervals := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}})
	//intervals := eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}})
	//fmt.Println(intervals)

	//i := rob([]int{2, 1, 1, 8})
	//i := rob([]int{8, 1, 1, 2})
	//fmt.Println(i)

	//earn := deleteAndEarn([]int{2, 2, 3, 3, 3, 4})
	//earn := deleteAndEarn([]int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10})
	//fmt.Println(earn)

	//validString := checkValidString("()")
	//validString := checkValidString("**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))")
	//validString := checkValidString("(*)")
	//validString := checkValidString("(*))")
	//validString := checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()")
	//validString := checkValidString("(())()())(*))(((((())()*))**))**()(*(()()*)(**(())()**)((**(()(((()()**)())*(*))(())()()*")
	//validString := checkValidString("(")
	//fmt.Println(validString)

	//valid := isValid("()")
	//valid := isValid("()[]{}")
	//valid := isValid("{[]}")
	//fmt.Println(valid)

	//game := stoneGame([]int{3, 7, 2, 3})
	//fmt.Println(game)

	//i := bulbSwitch(6)
	//fmt.Println(i)

	//tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//tree := buildTree([]int{-1}, []int{-1})
	//tree := buildTree([]int{3, 9, 7, 8, 20}, []int{7, 9, 8, 3, 20})
	//fmt.Println(tree)

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

	//t := &TreeNode{
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

	//t := &TreeNode{
	//	Val: -3,
	//}

	//t := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   -2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   3,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}

	//t := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: -2,
	//		Left: &TreeNode{
	//			Val: 1,
	//			Left: &TreeNode{
	//				Val:   -1,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//			Right: nil,
	//		},
	//		Right: &TreeNode{Val: 3},
	//	},
	//	Right: &TreeNode{
	//		Val: -3,
	//		Left: &TreeNode{
	//			Val:   -2,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//}
	//sum := maxPathSum(t)
	//fmt.Println(sum)

	//flips := minBitFlips(3, 4)
	//fmt.Println(flips)

	//sum := triangularSum([]int{1, 2, 3, 4, 5})
	//fmt.Println(sum)

	//ways := numberOfWays("001101")
	//ways := numberOfWays("11100")
	//ways := numberOfWays("0001100100")
	//fmt.Println(ways)

	scores := sumScores("babab")
	//scores := sumScores("azbazbzaz")
	fmt.Println(scores)
}

func numberOfMatches(n int) int {
	res := 0
	var travel func(n int)
	travel = func(n int) {
		if n == 1 {
			return
		} else if n%2 == 0 {
			res += n / 2
			travel(n / 2)
		} else {
			res += (n - 1) / 2
			travel((n-1)/2 + 1)
		}
	}
	travel(n)
	return res
}

func shipWithinDays1(weights []int, days int) int {
	min := math.MaxInt64
	max := 0
	for i := 0; i < len(weights); i++ {
		if weights[i] < min {
			min = weights[i]
		}
		if weights[i] > max {
			max = weights[i]
		}
	}
	min = max
	max = (2 * max * len(weights) / days) + 1
	for min < max {
		mid := min + (max-min)/2
		res := 1

		temp := mid
		for i := 0; i < len(weights); i++ {
			if temp >= weights[i] {
				temp -= weights[i]
			} else {
				res++
				temp = mid - weights[i]
			}
		}
		if res > days {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

func mostFrequent(nums []int, key int) int {
	res := 0
	m := make(map[int]int)
	for i, num := range nums {
		if num == key && i+1 < len(nums) {
			m[nums[i+1]]++
		}
	}
	temp := 0
	for k, value := range m {
		if value > temp {
			res = k
			temp = value
		}
	}
	return res
}

func sortJumbled(mapping []int, nums []int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		mod := 10
		x := 0
		i := 0
		if num == 0 {
			m[nums[index]] = mapping[0]
			continue
		}
		for num != 0 {
			temp := num % mod
			x += mapping[temp] * int(math.Pow(10, float64(i)))
			i++
			num = num / mod
		}
		m[nums[index]] = x
	}

	sort.Slice(nums, func(i, j int) bool {
		if nums[i] == nums[j] {
			return false
		} else {
			return m[nums[i]] < m[nums[j]]
		}
	})

	return nums

}

func getAncestors1(n int, edges [][]int) [][]int {
	m := make(map[int][]int, 0)
	for i := 0; i < len(edges); i++ {
		if m[edges[i][1]] == nil {
			m[edges[i][1]] = make([]int, 0)
		}
		m[edges[i][1]] = append(m[edges[i][1]], edges[i][0])
	}
	res := make([][]int, 0, n)

	var travel func(k int, res []int) []int
	travel = func(k int, res []int) []int {
		if m[k] == nil || len(m[k]) == 0 {
			return res
		}
		for _, v := range m[k] {
			if res == nil {
				res = make([]int, 0)
			}
			res = append(res, v)
			res = travel(v, res)
		}
		return res
	}
	for i := 0; i < n; i++ {
		if m[i] == nil {
			res = append(res, []int{})
			continue
		} else {

		}
		temp := make([]int, 0)
		temp = append(temp, m[i]...)
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] > len(res) {
				for _, value := range m[m[i][j]] {
					temp = append(temp, value)
					temp = travel(value, temp)
				}
				continue
			}
			temp = append(temp, res[m[i][j]]...)
		}
		sort.Ints(temp)
		m1 := make(map[int]bool)
		temp1 := make([]int, 0)
		for k := 0; k < len(temp); k++ {
			if !m1[temp[k]] {
				temp1 = append(temp1, temp[k])
				m1[temp[k]] = true
			}
		}
		res = append(res, temp1)
	}
	return res
}

func getAncestors(n int, edges [][]int) [][]int {
	m := make(map[int][]int, 0)
	for i := 0; i < len(edges); i++ {
		if m[edges[i][1]] == nil {
			m[edges[i][1]] = make([]int, 0)
		}
		m[edges[i][1]] = append(m[edges[i][1]], edges[i][0])
	}

	mVisit := make(map[int]bool)

	res := make([][]int, 0, n)

	var travel func(ints []int, mTemp map[int]bool, k int)
	travel = func(ints []int, mTemp map[int]bool, k int) {
		if mVisit[k] {
			return
		}
		if len(ints) == 0 {
			return
		}
		if mTemp == nil {
			mTemp = make(map[int]bool)
		}

		for _, v := range ints {
			if !mTemp[v] {
				mTemp[v] = true
				travel(m[v], mTemp, v)
				ints = append(ints, m[v]...)
			} else {
				ints = append(ints, m[v]...)
			}
		}

		mm := make(map[int]bool)
		ans := make([]int, 0)
		for _, v := range ints {
			if !mm[v] {
				mm[v] = true
				ans = append(ans, v)
			} else {
				continue
			}
		}
		sort.Ints(ans)
		m[k] = ans
		mVisit[k] = true
	}

	for i := 0; i < n; i++ {
		travel(m[i], nil, i)
		ints := m[i]
		res = append(res, ints)
	}

	return res
}

func cellsInRange(s string) []string {
	res := make([]string, 0)
	split := strings.Split(s, ":")
	col1 := split[0][0]
	col2 := split[1][0]
	row1 := split[0][1:]
	row2 := split[1][1:]
	r1, _ := strconv.ParseInt(row1, 10, 32)
	r2, _ := strconv.ParseInt(row2, 10, 32)
	for j := col1 - 'a'; j <= col2-'a'; j++ {
		for i := r1; i <= r2; i++ {

			sprintf := fmt.Sprintf("%c%d", j+'a', i)
			res = append(res, sprintf)
		}
	}
	return res
}

//要思考 当超出数组的时候怎么办 这时候有个办法 就是改变数组啊啊啊啊啊 如果解决不了问题 就解决掉提出问题的人
func minimalKSum(a []int, k int) int64 {
	var res int64
	//增加哨兵节点
	a = append(a, 0, 2e9)
	sort.Ints(a)
	for i := 1; ; i++ {
		count := a[i] - a[i-1] - 1
		if count <= 0 || a[i] <= 1 {
			continue
		}
		if count <= k {
			res += int64((a[i] + a[i-1]) * count / 2)
			k -= count
		} else if count > k {
			res += int64((a[i-1] + 1 + a[i-1] + k) * k / 2)
			break
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode, 0)
	var root *TreeNode

	mParent := make(map[*TreeNode]*TreeNode)

	for _, v := range descriptions {
		if m[v[0]] == nil {
			parent := new(TreeNode)
			parent.Val = v[0]
			m[v[0]] = parent
		}

		t := new(TreeNode)
		if m[v[1]] == nil {
			m[v[1]] = t
		}
		t = m[v[1]]
		if v[2] == 1 {
			m[v[0]].Left = t
		} else if v[2] == 0 {
			m[v[0]].Right = t
		}

		if root == nil {
			root = m[v[0]]
		}
		mParent[m[v[1]]] = m[v[0]]
	}

	for mParent[root] != nil {
		root = mParent[root]
	}
	return root
}

func longestCommonSubsequence1(text1 string, text2 string) int {
	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)
	for i := 0; i < len(text1); i++ {
		m1[text1[i]] = i
	}

	for i := 0; i < len(text2); i++ {
		m2[text2[i]] = i
	}

	m3 := make(map[uint8]int)
	text3 := make([]uint8, 0)
	m4 := make(map[uint8]int)
	text4 := make([]uint8, 0)
	for i := 0; i < len(text1); i++ {
		if v, ok := m2[text1[i]]; ok {
			m3[text1[i]] = i
			text3 = append(text3, text1[i])
			m4[text1[i]] = v
			text4 = append(text4, text1[i])
		}
	}

	return 0
}

//找不到规律就完了 优化的点在于 拼接S的时候可以考虑使用[]byte
func generateParenthesis1(n int) []string {
	res := make([]string, 0)
	left, right := n, n
	var travel func(int, int, string)
	travel = func(left int, right int, s string) {
		if left == right && left == 0 {
			res = append(res, s)
			return
		} else if left > right {
			return
		}

		if left > 0 {
			travel(left-1, right, s+"(")
		}

		if right > 0 {
			travel(left, right-1, s+")")
		}
	}

	travel(left, right, "")

	return res
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var travel func(*TreeNode, *TreeNode) *TreeNode
	travel = func(node1 *TreeNode, node2 *TreeNode) *TreeNode {
		if node1 == nil && node2 == nil {
			return nil
		} else if node1 == nil && node2 != nil {
			return node2
		} else if node1 != nil && node2 == nil {
			return node1
		}
		newNode := &TreeNode{
			Val:   node1.Val + node2.Val,
			Left:  nil,
			Right: nil,
		}
		newNode.Left = travel(node1.Left, node2.Left)
		newNode.Right = travel(node1.Right, node2.Right)
		return newNode
	}
	node := travel(root1, root2)
	return node
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 == nil && root2 != nil {
		return root2
	} else if root1 != nil && root2 == nil {
		return root1
	}

	node := root1
	queue1 := make([]*TreeNode, 0)
	queue2 := make([]*TreeNode, 0)
	queue1 = append(queue1, root1)
	queue2 = append(queue2, root2)
	l1 := len(queue1)
	l2 := len(queue2)
	for l1 > 0 || l2 > 0 {
		for l1 != 0 || l2 != 0 {
			node1 := queue1[0]
			node2 := queue2[0]
			node1.Val += node2.Val

			if node1.Left != nil || node2.Left != nil {
				if node1.Left == nil {
					node1.Left = &TreeNode{}
				} else if node2.Left == nil {
					node2.Left = &TreeNode{}
				}
				queue1 = append(queue1, node1.Left)
				queue2 = append(queue2, node2.Left)
			}

			if node1.Right != nil || node2.Right != nil {
				if node1.Right == nil {
					node1.Right = &TreeNode{}
				} else if node2.Right == nil {
					node2.Right = &TreeNode{}
				}
				queue1 = append(queue1, node1.Right)
				queue2 = append(queue2, node2.Right)
			}
			queue1 = queue1[1:]
			queue2 = queue2[1:]
			l1--
			l2--
		}
		l1 = len(queue1)
		l2 = len(queue2)
	}

	return node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

//完美二叉树 就用性质就可以解决了 抓到切片里 根据索引干 比较废空间
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		root.Next = nil
		return root
	}

	m := make([]*Node, 2, 10)
	m[1] = root

	round := 1
	for i := 1; i < len(m); i++ {

		if m[i].Left != nil && m[i].Right != nil {
			m = append(m, m[i].Left, m[i].Right)
		}

		if (1<<round - 1) == i {
			m[i].Next = nil
			round++
		} else {
			m[i].Next = m[i+1]
		}
	}

	return root
}

//栈每次都会缩短 好像没有啥用
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		root.Next = nil
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)

	round := 1

	count := 1
	for len(queue) != 0 {
		node := queue[0]
		if node.Left != nil && node.Right != nil {
			queue = append(queue, node.Left, node.Right)
		}
		queue = queue[1:]
		if (1<<round - 1) == count {
			node.Next = nil
			round++
		} else {
			node.Next = queue[0]
		}
		count++
	}

	return root
}

//太猛了 用递归 记得使用别把题目让我链上的next忘记了呀  这个很想mysql里的数据结构了
func connect(root *Node) *Node {
	if root == nil {
		return nil
	} else if root.Left == nil && root.Right == nil {
		root.Next = nil
		return root
	}

	var travel func(parent *Node, child *Node)
	travel = func(parent *Node, child *Node) {
		if parent == nil {

		} else if child == nil {
			return
		} else if parent.Right != child {
			child.Next = parent.Right
		} else {
			if parent.Next != nil {
				child.Next = parent.Next.Left
			}
		}
		travel(child, child.Left)
		travel(child, child.Right)
	}

	travel(nil, root)
	return root
}

//虽然我写的丑 但是运行很快啊
//直接修改原数组 int[][] matrix 来记录距离和标志是否访问的
func updateMatrix1(mat [][]int) [][]int {
	res := make([][]int, 0, len(mat))
	queue := make([][2]int, 0)

	var IsoutOfRange func(x, y int, target int) bool
	IsoutOfRange = func(x, y int, target int) bool {
		if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[x]) && mat[x][y] > target {
			return true
		}
		return false
	}

	for i := 0; i < len(mat); i++ {
		line := make([]int, len(mat[i]), len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				//考虑用备忘录记录一下
				if IsoutOfRange(i, j+1, 0) {
					queue = append(queue, [2]int{i, j + 1})
				}
				if IsoutOfRange(i, j-1, 0) {
					queue = append(queue, [2]int{i, j - 1})
				}
				if IsoutOfRange(i-1, j, 0) {
					queue = append(queue, [2]int{i - 1, j})
				}
				if IsoutOfRange(i+1, j, 0) {
					queue = append(queue, [2]int{i + 1, j})
				}

			} else {
				mat[i][j] = math.MaxInt
				line[j] = math.MaxInt
			}
		}
		res = append(res, line)
	}

	dis := 1
	l := len(queue)
	last := make([][2]int, 0, l)
	for l != 0 {
		for l != 0 {
			ints := queue[0]
			x := ints[0]
			y := ints[1]
			res[x][y] = dis
			mat[x][y] = dis
			queue = queue[1:]
			last = append(last, ints)
			l--
		}
		for _, ints := range last {
			x := ints[0]
			y := ints[1]
			if IsoutOfRange(x, y+1, dis) {
				queue = append(queue, [2]int{x, y + 1})
			}
			if IsoutOfRange(x, y-1, dis) {
				queue = append(queue, [2]int{x, y - 1})
			}
			if IsoutOfRange(x-1, y, dis) {
				queue = append(queue, [2]int{x - 1, y})
			}
			if IsoutOfRange(x+1, y, dis) {
				queue = append(queue, [2]int{x + 1, y})
			}
		}
		last = last[:0]
		l = len(queue)
		dis++
	}
	return res
}

//Tree 只有 1 个 root，而图可以有多个源点，所以首先需要把多个源点都入队；
//Tree 是有向的因此不需要标识是否访问过，而对于无向图来说，必须得标志是否访问过哦！并且为了防止某个节点多次入队，需要在其入队之前就将其设置成已访问！【 看见很多人说自己的 BFS 超时了，坑就在这里哈哈哈
func updateMatrix2(mat [][]int) [][]int {
	queue := make([][2]int, 0)
	m := make(map[[2]int]bool)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				m[[2]int{i, j}] = true
			}
		}
	}

	var isInRange func(x, y int, num int)
	isInRange = func(x, y int, num int) {
		if !m[[2]int{x, y}] && x >= 0 && x < len(mat) && y >= 0 && y < len(mat[x]) && mat[x][y] != 0 {
			m[[2]int{x, y}] = true
			mat[x][y] = num
			queue = append(queue, [2]int{x, y})
			return
		}
		return
	}

	num := 1
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			ints := queue[0]
			x := ints[0]
			y := ints[1]
			isInRange(x-1, y, num)
			isInRange(x+1, y, num)
			isInRange(x, y-1, num)
			isInRange(x, y+1, num)
			queue = queue[1:]
		}
		num++
	}
	return mat
}

//用map很容易超时 那么我们就用新的res数组 如果res数组值为0 说明还没有被赋值过
//很重要 有时候 1.可以修改元数组的值 达到判断的目的
//很重要 有时候 2.可以弄一个新的数组 利用新数组和元素组的差异 来判断 哪些已经被赋值过了
func updateMatrix3(mat [][]int) [][]int {
	res := make([][]int, len(mat), len(mat))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(mat[i]), len(mat[i]))
	}

	queue := make([][2]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	var isInRange func(x, y int, num int)
	isInRange = func(x, y int, num int) {
		// res[x][y] == 0 和 mat[x][y] != 0 用来区分 那些没有被修改过的
		if x >= 0 && x < len(mat) && y >= 0 && y < len(mat[x]) && res[x][y] == 0 && mat[x][y] != 0 {
			res[x][y] = num
			queue = append(queue, [2]int{x, y})
			return
		}
		return
	}

	num := 1
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			ints := queue[0]
			x := ints[0]
			y := ints[1]
			isInRange(x-1, y, num)
			isInRange(x+1, y, num)
			isInRange(x, y-1, num)
			isInRange(x, y+1, num)
			queue = queue[1:]
		}
		num++
	}
	return res
}

//动态规划
func updateMatrix(mat [][]int) [][]int {
	res := make([][]int, len(mat), len(mat))

	for i := 0; i < len(mat); i++ {
		res[i] = make([]int, len(mat[i]), len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				res[i][j] = 0
				continue
			}
			res[i][j] = math.MaxInt / 2
		}
	}

	//从左上角开始 遍历左边和上边
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {

			if i >= 1 {
				//判断上边和本身
				res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i-1][j]+1)))
			}
			if j >= 1 {
				//判断左边和本身
				res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i][j-1]+1)))
			}
		}
	}

	//从右下角开始 遍历右边和下边
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[i]) - 1; j >= 0; j-- {
			if i < len(mat)-1 {
				//判断下边和本身
				res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i+1][j]+1)))
			}
			if j < len(mat[i])-1 {
				//判断右边和本身
				res[i][j] = int(math.Min(float64(res[i][j]), float64(res[i][j+1]+1)))
			}
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//额外创建了数据结构 能不能利用现有空间
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	root := new(ListNode)
	cur := root
	cur1, cur2 := list1, list2
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			cur.Next = cur2
			cur = cur.Next
			cur2 = cur2.Next
		} else if cur2 == nil {
			cur.Next = cur1
			cur = cur.Next
			cur1 = cur1.Next
		} else {
			if cur1.Val <= cur2.Val {
				cur.Next = cur1
				cur = cur.Next
				cur1 = cur1.Next
			} else {
				cur.Next = cur2
				cur = cur.Next
				cur2 = cur2.Next
			}
		}
	}
	return root.Next
}

//递归 递归又忘记了  x.next = travel()
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}

//迭代做法
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode
	cur := head
	nxt := head.Next
	cur.Next = nil
	for nxt != nil {
		temp := nxt.Next
		nxt.Next = cur
		if pre != nil {
			cur.Next = pre
		} else {
			cur.Next = nil
		}
		pre = cur
		cur = nxt
		nxt = temp
	}
	return cur
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {

		return head
	}

	var res *ListNode

	var travel func(head *ListNode)
	travel = func(head *ListNode) {
		if head.Next == nil {
			res = head
			return
		}
		if head.Next != nil {
			travel(head.Next)
			head.Next.Next = head
		}
	}
	travel(head)
	head.Next = nil

	return res
}

func findKDistantIndices(nums []int, key int, k int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			for j := 0; j <= k; j++ {
				if !m[i+j] && i+j >= 0 && i+j < len(nums) {
					res = append(res, i+j)
					m[i+j] = true
				}
				if !m[i-j] && i-j >= 0 && i-j < len(nums) {
					res = append(res, i-j)
					m[i-j] = true
				}
			}
		}
	}
	sort.Ints(res)
	return res
}

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	map1 := make([][]int, n, n)
	for i := 0; i < n; i++ {
		map1[i] = make([]int, n, n)
	}
	for _, ints := range dig {
		map1[ints[0]][ints[1]] = 1
	}
	res := 0

	for _, artifact := range artifacts {
		bre := true
		left := []int{artifact[0], artifact[1]}
		right := []int{artifact[2], artifact[3]}
		for i := left[0]; i <= right[0]; i++ {
			b := true
			for j := left[1]; j <= right[1]; j++ {
				if map1[i][j] != 1 {
					b = false
					break
				}
			}
			if !b {
				bre = false
				break
			}
		}
		if bre {
			res++
		}

	}
	return res
}

//分类讨论
func maximumTop(nums []int, k int) int {
	//如果k==0 那么就是首部
	if k == 0 {
		return nums[0]
	}

	//如果只有一个元素 那么k为奇数的时候就是空栈 -1 k是偶数的时候 就是这个元素
	if len(nums) == 1 {
		if k%2 == 0 {
			return nums[0]
		}
		return -1
	}

	temp := make([]int, 0)
	if k < len(nums) {
		//如果 k 比 数组长度短的话  例如有6个 k==4 实际上 通过枚举可以发现 可以取到的数是 1 2 3 和 5
		temp = append(temp, nums[:k-1]...)
		temp = append(temp, nums[k])
		sort.Ints(temp)
		return temp[len(temp)-1]
	} else if k == len(nums) {
		//如果正好相等  实际上 通过枚举可以发现 可以取到除了栈底之外的任何数
		temp = append(temp, nums[:k-1]...)
		sort.Ints(temp)
		return temp[len(temp)-1]
	} else {
		//如果k 大于数组的长度的话 任何数字的可以去得到
		sort.Ints(nums)
		return nums[len(nums)-1]
	}
}

func orangesRotting(grid [][]int) int {

	queue := make([][2]int, 0)
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				count++
			} else if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	if count == 0 {
		return 0
	}

	var isOutOfRange func(x, y int)
	isOutOfRange = func(x, y int) {
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == 1 {
			grid[x][y] = 2
			count--
			queue = append(queue, [2]int{x, y})
			return
		}
	}

	times := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			ints := queue[0]
			x := ints[0]
			y := ints[1]
			isOutOfRange(x-1, y)
			isOutOfRange(x+1, y)
			isOutOfRange(x, y-1)
			isOutOfRange(x, y+1)
			queue = queue[1:]
		}
		times++
	}
	if count != 0 {
		return -1
	}
	return times - 1
}

//优化之后的回溯算法 提高了很多很多
func combine(n int, k int) [][]int {
	res := [][]int{}
	var travel func(ans []int, input int)
	travel = func(ans []int, input int) {
		ans = append(ans, input)
		if len(ans) == k {
			temp := make([]int, k, k)
			copy(temp, ans)
			res = append(res, temp)
			return
		}

		//这里可以优化
		for i := input + 1; i <= n-k+len(ans)+1; i++ {
			travel(ans, i)
		}
	}

	//这里可以优化
	for i := 1; i <= n-k+1; i++ {
		travel([]int{}, i)
	}
	return res
}

//原来打算用map来判断 是否已经加入了
//在这里传递的时候 是通过构造了新的切片来传递的 可能会耗时
func permute1(nums []int) [][]int {
	res := make([][]int, 0, len(nums)*(len(nums)-1))
	var travel func(ans []int, nums []int)
	travel = func(ans []int, numsNow []int) {
		if len(ans) == len(nums) {
			temp := make([]int, len(nums), len(nums))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(numsNow); i++ {
			ans = append(ans, numsNow[i])
			temp := make([]int, len(nums)-len(ans), len(nums)-len(ans))
			//如果不会改变原数组结构 可以用这种方式 他并不是拼接起来
			copy(temp, numsNow[:i])
			copy(temp[i:], numsNow[i+1:])
			travel(ans, temp)
			//记得要还原回来
			ans = ans[:len(ans)-1]
		}
	}

	travel([]int{}, nums)

	return res
}

//这里传递的时候 永远认为前一个已经被干完了 通过swap交换的形式
func permute(nums []int) [][]int {
	res := make([][]int, 0, len(nums)*(len(nums)-1))
	var travel func(ans []int, nums []int)
	travel = func(ans []int, numsNow []int) {
		if len(ans) == len(nums) {
			temp := make([]int, len(nums), len(nums))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		for i := 0; i < len(numsNow); i++ {
			//利用交换 来把当前的放到index为0 的位置上
			numsNow[0], numsNow[i] = numsNow[i], numsNow[0]
			ans = append(ans, numsNow[0])

			travel(ans, numsNow[1:])

			//回溯完成后 记得还原回来 否则会影响后面的运算
			numsNow[i], numsNow[0] = numsNow[0], numsNow[i]
			//由于ans是切片引用类型 所以递归的时候已经被加入过了 回溯完成后 记得把ans去掉 否则会影响后面的运算
			ans = ans[:len(ans)-1]
		}
	}

	travel([]int{}, nums)

	return res
}

func letterCasePermutation1(s string) []string {
	res := []string{}
	var travel func(ans []byte, index int)
	travel = func(ans []byte, index int) {
		if index == len(s) {
			res = append(res, string(ans))
			return
		}

		if unicode.IsNumber(rune(s[index])) {
			ans = append(ans, s[index])
			travel(ans, index+1)
			return
		}

		ans = append(ans, s[index])

		travel(ans, index+1)
		ans = ans[:len(ans)-1]

		if s[index] >= 65 && s[index] <= 91 {
			ans = append(ans, byte(unicode.ToLower(rune(s[index]))))
		} else {
			ans = append(ans, byte(unicode.ToUpper(rune(s[index]))))
		}

		travel(ans, index+1)
		ans = ans[:len(ans)-1]
	}
	travel([]byte{}, 0)

	return res
}

//根据bits 位集合来做 既然每个字母都有2种情况 那么3个字母 就会有2的三次方种排列情况
func letterCasePermutation(s string) []string {
	res := []string{}
	count := 0
	for _, v := range s {
		if unicode.IsLetter(v) {
			count++
		}
	}

	//一共有 1<<count种情况
	for i := 0; i < 1<<count; i++ {
		ans := []byte{}
		carry := 0

		for j := 0; j < len(s); j++ {
			//如果是数字 就跳过
			if unicode.IsNumber(rune(s[j])) {
				ans = append(ans, s[j])
				continue
			}

			//如果&后为0 就用小写的
			if 1<<carry&i == 0 {
				ans = append(ans, byte(unicode.ToLower(rune(s[j]))))
			} else {
				//如果 &后为1 就用大写的
				ans = append(ans, byte(unicode.ToUpper(rune(s[j]))))
			}

			//改变移动的量
			carry++
		}
		res = append(res, string(ans))
	}

	return res
}

func climbStairs1(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	for i := 3; i <= n; i++ {
		m[i] = m[i-2] + m[i-1]
	}
	return m[n]
}

//动态规划 只记录有用的
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	pre1 := 1
	pre2 := 2

	var res int
	for i := 3; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}

	return res
}

//回溯行不通 时间复杂度太高了
func rob1(nums []int) int {
	res := 0
	var travel func([]int, int)
	travel = func(ints []int, cash int) {
		if len(ints) == 0 {
			if cash > res {
				res = cash
			}
			return
		}
		for i := 0; i < len(ints); i++ {
			if i+2 < len(ints) {
				travel(ints[i+2:], cash+ints[i])
			} else {
				if cash+ints[i] > res {
					res = cash + ints[i]
				}
			}
		}
	}

	travel(nums, 0)
	return res
}

//dp数组代表当前打劫店家的最大金额  这个金额是 2种情况种的最大值 1.这家打劫的钱+dp[i-2] 2.这家不打劫 也就是dp[i-1]
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < len(nums); i++ {
		temp := dp[i-2] + nums[i]
		if temp > dp[i-1] {
			dp[i] = temp
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-1]
}

//动态规划 就是一层一层求 写出状态转移方程 可以优化的点在于空间上只需要这一层和上一层
func minimumTotal(triangle [][]int) int {
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	if len(triangle) == 1 {
		return triangle[0][0]
	} else if len(triangle) == 2 {
		return min(triangle[1][0], triangle[1][1]) + triangle[0][0]
	}

	res := make([][]int, len(triangle), len(triangle))
	for i := 0; i < len(triangle); i++ {
		if i == 0 {
			res[0] = append(res[0], triangle[0]...)
			continue
		}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				res[i] = append(res[i], res[i-1][0]+triangle[i][j])
				continue
			} else if j == len(triangle[i])-1 {
				res[i] = append(res[i], res[i-1][len(res[i-1])-1]+triangle[i][j])
				continue
			}
			res[i] = append(res[i], min(res[i-1][j-1]+triangle[i][j], res[i-1][j]+triangle[i][j]))
		}
	}

	sort.Ints(res[len(triangle)-1])
	return res[len(triangle)-1][0]
}

//n&(n-1) 为了去掉最后一位1 如果只有一位为1 n&(n-1) 就变成0了
func isPowerOfTwo1(n int) bool {
	return n&(n-1) == 0 && n > 0
}

//-n 是反码+1 所以 n和-n与运算的结果等于n的话就说明只有一个1
func isPowerOfTwo(n int) bool {
	return n&(-n) == n && n > 0
}

func hammingWeight1(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num != 0 {
			res++
		}
	}
	return res
}

//位运算  原来的数字不动 每次通过1左移的方式来判断 原数字那一位是不是1 如果是1 再通过左移的方式更新res
func reverseBits1(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		if (num & (1 << i)) != 0 {
			res |= 1 << (31 - i)
		}
	}
	return res
}

//位运算  n>> 可能会提前结束
func reverseBits2(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		if num&1 != 0 {
			res |= 1 << (31 - i)
		}
		num >>= 1
		if num == 0 {
			break
		}
	}
	return res
}

//拆分 分治 最后连接
func reverseBits(num uint32) uint32 {
	var concact func(num1, num2 uint32, chai int) uint32
	concact = func(num1, num2 uint32, chai int) uint32 {
		//默认传递进来num1是原数组左边的 num2是原数组右边的 现在连接的时候要颠倒一下位置 那么就通过左移<<来达到
		return num2<<chai | num1
	}
	var travel func(num uint32, chai int) uint32
	travel = func(num uint32, chai int) uint32 {
		if chai == 0 {
			return num
		}

		//通过左移获取左半部分的
		num1 := travel(num>>chai, chai/2)
		//通过笨比方法获取右半部分的
		var temp uint32
		for i := 0; i < chai; i++ {
			temp = temp | ((num & 1) << i)
			num = num >> 1
		}
		num2 := travel(temp, chai/2)
		return concact(num1, num2, chai)
	}
	return travel(num, 16)
}

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	}
	pre1 := 0
	pre2 := 1
	pre3 := 1
	for i := 3; i <= n; i++ {
		temp := pre1 + pre2 + pre3
		pre1 = pre2
		pre2 = pre3
		pre3 = temp
	}
	return pre3
}

//动态规划 dp数组的定义很重要 dp[x]=上到x台阶的花费=dp[x-1]+cost(x-1)或者dp[x-2]+cost(x-2)
//dp的特征在于 不走回头路 之前计算过的dp不会再被修改
func minCostClimbingStairs1(cost []int) int {
	dp := make([]int, len(cost)+1, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(dp)-1]
}

//空间优化
func minCostClimbingStairs(cost []int) int {
	pre := 0
	cur := 0
	var min func(x, y int) int
	min = func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	for i := 2; i < len(cost)+1; i++ {
		temp := cur
		cur = min(cur+cost[i-1], pre+cost[i-2])
		pre = temp
	}
	return cur
}

//dp数组的定义很重要
func lengthOfLIS(nums []int) int {

	res := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = 1
			continue
		}

		res[i] = 1

		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if res[i] < res[j]+1 {
					res[i] = res[j] + 1
				}
			}
		}
	}
	sort.Ints(res)
	return res[len(res)-1]
}

//动态规划 dp数组的定义很重要 dp[i] = max{ dp[i-1]+nums[i] , nums[i] }
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var ans int

	res := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			res[i] = nums[i]
			ans = res[i]
			continue
		}
		if res[i-1]+nums[i] >= nums[i] {
			res[i] = res[i-1] + nums[i]
		} else {
			res[i] = nums[i]
		}
		if ans < res[i] {
			ans = res[i]
		}

	}
	return ans
}

//空间优化 不使用额外空间了
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var ans int

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nums[i] = nums[i]
			ans = nums[i]
			continue
		}
		if nums[i-1]+nums[i] >= nums[i] {
			nums[i] = nums[i-1] + nums[i]
		} else {
			nums[i] = nums[i]
		}
		if ans < nums[i] {
			ans = nums[i]
		}

	}
	return ans
}

//空间优化 只使用前一个和这一个  dp[i]定义为 前i个包含i的最大连续子序列和的值
func maxSubArray(nums []int) int {
	pre := nums[0]
	ans := pre

	for i := 1; i < len(nums); i++ {
		if nums[i]+pre > nums[i] {
			pre = nums[i] + pre
		} else {
			pre = nums[i]
		}

		if pre > ans {
			ans = pre
		}
	}

	return ans
}

//用visited 复杂了  利用hash表的去重特性
func longestConsecutive1(nums []int) int {
	res := 0
	m := make(map[int]bool)
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	for key := range m {
		if visited[key] {
			continue
		}
		visited[key] = true
		k := key + 1
		temp := 1
		for m[k] {
			visited[k] = true
			k++
			temp++
		}
		if temp > res {
			res = temp
		}
	}

	return res
}

//降低空间复杂度  o(2*n)=o(n) 只要是常数次遍历n 都可以视为o(n)
func longestConsecutive2(nums []int) int {
	res := 0
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	for key := range m {
		if m[key-1] {
			continue
		}
		k := key + 1
		temp := 1
		for m[k] {
			k++
			temp++
		}
		if temp > res {
			res = temp
		}
	}

	return res
}

//123 456  首当其冲的 解决的就是溢出的问题 其次 解决如何转换的问题
//转换的巧妙就在于 '7' 的 7 = '7'-'0'
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := ""
	//两字符串相加问题
	var travel func(s string)
	travel = func(s string) {
		if res == "" {
			res = s
			return
		}
		cur := ""
		add := 0

		count := 0
		resl := len(res)
		sl := len(s)

		//从最后一位开始加起
		for ; resl-count > 0; count++ {
			temp := int(res[resl-1-count]-'0'+s[sl-1-count]-'0') + add
			cur = strconv.Itoa(temp%10) + cur
			add = int(temp) / 10
		}

		for i := len(s) - len(res) - 1; i >= 0; i-- {
			temp := int(s[i]-'0') + add
			//百分号别忘记了!!!
			cur = strconv.Itoa(temp%10) + cur
			add = int(temp) / 10
		}

		//防止头部丢了
		if add > 0 {
			cur = strconv.Itoa(add) + cur
		}
		res = cur
	}

	l1 := len(num1)
	l2 := len(num2)
	//模拟乘法  记一下进位add 叠加一下cur
	for i := l2 - 1; i >= 0; i-- {
		add := 0
		cur := ""

		//补0
		for j := l2 - 1; j > i; j-- {
			cur += "0"
		}
		nums2 := num2[i] - '0'
		for j := l1 - 1; j >= 0; j-- {
			nums1 := num1[j] - '0'
			product := int(nums1*nums2) + add

			cur = strconv.Itoa(product%10) + cur
			add = product / 10
		}

		//防止头部丢了
		if add > 0 {
			cur = strconv.Itoa(add) + cur
		}

		//和res相加
		travel(cur)
	}
	return res
}

//你很恶心诶 ??? 了不起?
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//只排序k个 其他的不管了
func findKthLargest2(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		temp := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[temp] {
				temp = j
			}
		}
		nums[i], nums[temp] = nums[temp], nums[i]
		count++
		if count == k {
			break
		}
	}
	return nums[count-1]
}

//用堆排序
func findKthLargest3(nums []int, k int) int {
	m := new(myheap)
	for i := 0; i < len(nums); i++ {
		heap.Push(m, nums[i])
	}

	//每一次pop都会排序一下 ?
	for i := 0; i < k-1; i++ {
		heap.Pop(m)
	}
	return (*m)[0]
}

//归并排序
func findKthLargest4(nums []int, k int) int {
	var merge func(ints1 []int, ints2 []int) []int
	merge = func(ints1, ints2 []int) []int {
		res := make([]int, 0, len(ints1)+len(ints2))
		i, j := 0, 0
		for i != len(ints1) || j != (len(ints2)) {
			if i == len(ints1) {
				res = append(res, ints2[j])
				j++
				continue
			} else if j == len(ints2) {
				res = append(res, ints1[i])
				i++
				continue
			} else if ints1[i] < ints2[j] {
				res = append(res, ints1[i])
				i++
			} else {
				res = append(res, ints2[j])
				j++
			}
		}
		return res
	}

	var mergeSort func(nums []int) []int
	mergeSort = func(nums []int) []int {
		if len(nums) == 1 {
			return nums
		}
		left := 0
		right := len(nums)
		mid := left + (right-left)/2
		ints1 := mergeSort(nums[:mid])
		ints2 := mergeSort(nums[mid:])
		ints := merge(ints1, ints2)
		return ints
	}

	nums = mergeSort(nums)

	return nums[len(nums)-k]
}

//快速排序
func findKthLargest(nums []int, k int) int {
	var travel func(nums []int, i int) int
	travel = func(nums []int, i int) int {
		pivot := nums[0]
		left, right := 0, len(nums)-1
		for left < right {
			for right > left && nums[right] >= pivot {
				right--
			}
			if left != right && right >= 0 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
			}
			for left < right && nums[left] <= pivot {
				left++
			}
			if left != right && left < len(nums)-1 {
				nums[left], nums[right] = nums[right], nums[left]
				right--
			}
		}
		return left
	}

	var quickSort func(nums []int)
	quickSort = func(nums []int) {
		if len(nums) == 0 || len(nums) == 1 {
			return
		}
		mid := travel(nums, 0)
		quickSort(nums[:mid+1])
		quickSort(nums[mid+1:])
	}

	quickSort(nums)
	return nums[len(nums)-k]
}

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	if len(m) > len(nums)/2 {
		return false
	}

	count := 0
	for _, v := range m {
		count += v / 2
	}
	if count == len(nums)/2 {
		return true
	}
	return false
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	pre := pattern[0]
	lat := pattern[1]

	if pre == lat {
		solu := fmt.Sprintf("%c", pre) + text
		m := make(map[byte]int)
		for i := 0; i < len(solu); i++ {
			if solu[i] == pre {
				m[pre]++
			}
		}

		return int64(m[pre] * (m[pre] - 1) / 2)
	}
	solu1 := fmt.Sprintf("%c", pre) + text
	count1 := int64(0)
	m1 := make(map[byte]int)
	for i := 0; i < len(solu1); i++ {
		if solu1[i] == pre {
			m1[pre]++
		} else if solu1[i] == lat {
			count1 += int64(m1[pre])
		}
	}

	solu2 := text + fmt.Sprintf("%c", lat)
	count2 := int64(0)
	m2 := make(map[byte]int)
	for i := 0; i < len(solu2); i++ {
		if solu2[i] == pre {
			m2[pre]++
		} else if solu2[i] == lat {
			count2 += int64(m2[pre])
		}
	}

	if count1 > count2 {
		return count1
	}
	return count2
}

func halveArray(nums []int) int {
	m := new(myheap1)
	sum := float32(0)
	for i := 0; i < len(nums); i++ {
		heap.Push(m, float32(nums[i]))
		sum += float32(nums[i])
	}
	count := 0
	total := float32(0)
	for 2*total < sum {
		pop := heap.Pop(m)
		i := pop.(float32)
		i = i / 2
		total += i
		heap.Push(m, i)
		count++
	}
	return count
}

type myheap []int

func (m *myheap) Len() int {
	return len(*m)
}

func (m *myheap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *myheap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *myheap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	temp := make([]int, len(floor), len(floor))
	count := 0
	for i, c := range floor {
		if c == '0' {
			count = 0
			temp[i] = 0
		} else if c == '1' {
			count++
			temp[i] = count
		}
	}

	times := 0
	for times < numCarpets {
		max := math.MinInt
		loc := -1
		for i := 0; i < len(temp); i++ {
			if temp[i] > max {
				max = temp[i]
				loc = i
			}
		}

		for i := 0; i < carpetLen; i++ {
			if loc-i >= 0 {
				temp[loc-i] = 0
			}
		}
		times++
	}

	sum := 0
	for i := 0; i < len(temp); i++ {
		if temp[i] >= 1 {
			if i+1 == len(temp) || temp[i+1] == 0 {
				sum += temp[i]
			}
		}
	}
	return sum
}

func countHillValley(nums []int) int {
	temp := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			temp = append(temp, nums[i])
			continue
		}
		if nums[i] != nums[i-1] {
			temp = append(temp, nums[i])
		}
	}

	res := 0
	for i := 0; i < len(temp); i++ {
		if i == 0 || i == len(temp)-1 {
			continue
		}

		left := temp[i-1]
		right := temp[i+1]
		if left < temp[i] && right < temp[i] {
			res++
		} else if left > temp[i] && right > temp[i] {
			res++
		}
	}
	return res
}

func countCollisions1(directions string) int {
	res := 0
	bytes := []byte(directions)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 'L' {
			if i-1 >= 0 {
				if bytes[i-1] == 'S' {
					res = res + 1
					bytes[i] = 'S'
				} else if bytes[i-1] == 'R' {
					res = res + 2
					bytes[i] = 'S'
				} else {

				}
			} else {

			}
		} else if bytes[i] == 'R' {
			if i+1 <= len(bytes)-1 {
				if bytes[i+1] == 'S' {
					res = res + 1
					bytes[i] = 'S'
				} else if bytes[i+1] == 'L' {
					res = res + 2
					bytes[i] = 'S'
					bytes[i+1] = 'S'
				} else {

				}
			} else {

			}
		}
	}
	return res
}

func countCollisions(directions string) int {
	res := 0
	bytes := []byte(directions)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == 'L' {
			if i-1 >= 0 {
				if bytes[i-1] == 'S' {
					res = res + 1
					bytes[i] = 'S'
				} else if bytes[i-1] == 'R' {
					res = res + 2
					bytes[i] = 'S'
					bytes[i-1] = 'S'
				} else {

				}
			} else {

			}
		} else if bytes[i] == 'R' {
			for j := i + 1; j < len(bytes); j++ {
				if bytes[j] == 'R' {
					continue
				} else if bytes[j] == 'S' {
					res = res + 1
					bytes[j] = 'S'
					for k := j - 1; k > i; k-- {
						bytes[k] = 'S'
						res++
					}
					bytes[i] = 'S'
					break
				} else if bytes[j] == 'L' {
					res = res + 2
					bytes[j] = 'S'
					for k := j - 1; k > i; k-- {
						bytes[k] = 'S'
						res++
					}
					bytes[i] = 'S'
					break
				}
			}
		}
	}
	return res
}

type myheap1 []float32

func (m *myheap1) Len() int {
	return len(*m)
}

func (m *myheap1) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

func (m *myheap1) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap1) Push(x interface{}) {
	*m = append(*m, x.(float32))
}

func (m *myheap1) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}

//暴力回溯  复杂度并不是很高 可以使用
func maximumBobPoints1(numArrows int, aliceArrows []int) []int {
	res := make([]int, len(aliceArrows), len(aliceArrows))
	resCount := 0
	temp := [12]int{}

	var travel func(temp [12]int, i int, left int)
	travel = func(temp [12]int, i int, left int) {
		if left < 0 {
			return
		} else if i == 0 {
			sum := 0
			for key, v := range temp {
				if v > 0 {
					sum += key
				}
			}
			if sum > resCount {
				resCount = sum
				res = temp[:]
				res[0] = left
			}
			return
		} else if left == 0 {
			sum := 0
			for key, v := range temp {
				if v > 0 {
					sum += key
				}
			}
			if sum > resCount {
				resCount = sum
				res = temp[:]
			}
			return
		}

		AShot := aliceArrows[i]
		temp[i] = AShot + 1
		travel(temp, i-1, left-AShot-1)

		temp[i] = 0
		travel(temp, i-1, left)
	}
	travel(temp, 11, numArrows)

	return res
}

//动态规划
func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	ans := [12]int{}
	dp := make([][]int, 12, 12)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 12, 12)
	}
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < 12; i++ {
		aa := aliceArrows[i]
		for j := 1; j <= numArrows; j++ {
			if j < aa+1 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j-aa-1]+i, dp[i-1][j]) //状态转移
			}
		}
	}

	for i := 11; i > 0; i-- { //路径恢复
		if dp[i][numArrows] > dp[i-1][numArrows] {
			aa := aliceArrows[i]
			ans[i] = aa + 1
			numArrows -= aa + 1
		}
	}

	ans[0] = numArrows
	return ans[:]
}

func maximumBobPoints3(numArrows int, aliceArrows []int) []int {
	res := make([]int, 0, 12)
	var max func(x, y int) int
	max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, 12, 12)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 12, 12)
	}

	for i := 1; i < 12; i++ {
		for j := 1; j <= numArrows; j++ {
			aa := aliceArrows[i] + 1
			if j < aa {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-aa]+i)
			}
		}
	}

	return res
}

//时间复杂度爆炸 告辞
func maxEnvelopes1(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	answer := make([]int, len(envelopes), len(envelopes))
	for i := 0; i < len(answer); i++ {
		answer[i] = envelopes[i][1]
	}

	res := 0
	dp := make([]int, len(envelopes), len(envelopes))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(envelopes); i++ {
		for j := 0; j < i; j++ {
			if answer[i] > answer[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})

	answer := make([]int, 0)
	for _, envelope := range envelopes {
		//巧妙二分 如果在内部 就更新 防止扩大
		ints := sort.SearchInts(answer, envelope[1])
		if ints >= len(answer) {
			answer = append(answer, envelope[1])
		} else {
			//如果在外部 就说明应该扩大
			answer[ints] = envelope[1]
		}
	}
	return len(answer)
}

//动态规划
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
				//这里很重要 发现了规律了吗 和背包问题类似 如果要选择他 那么就要扣掉当前的行和列 防止重复
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//不选择
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}

//递归 超时
func minDistance1(word1 string, word2 string) int {
	res := math.MaxInt

	var travel func(cost int, index int, word string)
	travel = func(cost int, index int, word string) {
		if index == len(word2) {
			cost += len(word)
			if res > cost {
				res = cost
			}
			return
		}

		if len(word) == 0 {
			//插入
			travel(cost+1, index+1, word)
		} else if word[0] == word2[index] {
			//相等
			travel(cost, index+1, word[1:])
		} else {
			//删除
			travel(cost+1, index, word[1:])

			//插入
			travel(cost+1, index+1, word)

			//替换
			travel(cost+1, index+1, word[1:])
		}
	}

	travel(0, 0, word1)

	return res
}

//这种题目有一点困难啊  完全不懂   动态规划
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1, len(word2)+1)
	}

	var min func(a, b, c int) int
	min = func(a, b, c int) int {
		return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = i
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				//相等的时候就只需要知道前面是什么就好
				dp[i][j] = dp[i-1][j-1]
			} else {
				//不等的时候 得需要找找
				//[i-1][j] 是删除  [i-1][j-1]是替换 [i][j-1]是新增
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
			}
		}
	}

	return dp[len(word1)][len(word2)]
}

//煎饼排序
func pancakeSort(arr []int) []int {
	res := make([]int, 0, 0)

	var travel func(ints []int)
	travel = func(ints []int) {
		//递归的停止条件
		if len(ints) == 1 {
			return
		}

		//找出最大值
		max := math.MinInt
		matIndex := -1
		for i := 0; i < len(ints); i++ {
			if ints[i] > max {
				max = ints[i]
				matIndex = i
			}
		}

		//如果最大值本身就在数组的末尾 就不管他
		if len(ints)-1 == matIndex {
			travel(ints[:len(ints)-1])
			return
		}

		//否则 更新res 翻转数组 把最大值甩到最上面
		res = append(res, matIndex+1)
		for i := 0; i < (matIndex+1)/2; i++ {
			ints[i], ints[matIndex-i] = ints[matIndex-i], ints[i]
		}

		//更新res  翻转数组 把最大值甩到当前数组的最下面
		res = append(res, len(ints))
		for i := 0; i < len(ints)/2; i++ {
			ints[i], ints[len(ints)-1-i] = ints[len(ints)-1-i], ints[i]
		}

		//开始考虑最后上一层的烧饼
		travel(ints[:len(ints)-1])
	}

	travel(arr)
	return res
}

//前缀和
func subarraySum1(nums []int, k int) int {
	res := 0
	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 1; i < len(preSum); i++ {
		if preSum[i] == k {
			res++
		}

		for j := 1; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				res++
			}
		}

	}
	return res
}

//前缀和 优化 实际上 preSum[i] = preSum[i]-preSum[0]
func subarraySum2(nums []int, k int) int {
	res := 0
	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 1; i < len(preSum); i++ {
		//从0开始 presum[i] - presum[0]  == presum[i]
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				res++
			}
		}
	}
	return res
}

//前缀和 哈希优化 preSum[i]-preSum[j] == k  只要是等式就可以尝试使用哈希表来代替
func subarraySum3(nums []int, k int) int {
	res := 0
	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 1; i < len(preSum); i++ {
		if i != 1 {
			m[preSum[i-1]]++
		}
		if preSum[i] == k {
			res++
		}
		if m[preSum[i]-k] != 0 {
			res += m[preSum[i]-k]
		}
	}
	return res
}

//前缀和 哈希优化 preSum[i]-preSum[j] == k  只要是等式就可以尝试使用哈希表来代替
//去掉preSum数组 用一个变量直接代替
func subarraySum(nums []int, k int) int {
	res := 0
	m := make(map[int]int)
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]

		//如果当前数组 也就是前缀和==k 那么加加
		if preSum == k {
			res++
		}

		//如果是前缀和中的某一部分 那么加这一部分存在的次数
		j := preSum - k
		res += m[j]
		m[preSum]++
	}
	return res
}

//筛数法还是超时
func countPrimes1(n int) int {
	res := 0
	if n == 0 || n == 1 {
		return res
	}
	visited := make(map[int]bool)
	for i := 2; i < n; i++ {
		if visited[i] {
			continue
		}

		res++
		//筛数法这里有个优化细节 就是可以从i平方开始筛 因为之前的应该都被筛过了  例如当2位质数的时候  6被筛过了 而3位质数的时候 6已经筛过了 所以还是从9开始筛把
		for k := i * i; k < n; k += i {
			visited[k] = true
		}
	}

	return res
}

//没什么事儿就用[]bool 数组来代替hashmap
//我觉得应该是因为数组在内存中线性连续的.所以用接近常量v-a的定位算法速度最快.这里的算法效率是o(1).
//而map在内存中存储并不是线性连续的,而且要经过hash计算index所以就慢了一些.虽然hash也是o(1)的效率但是在计算步骤上要比v-a复杂, 所以速度上不如.
func countPrimes(n int) int {
	res := 0
	visited := make([]bool, n+1, n+1)
	for i := 2; i < n; i++ {
		if visited[i] {
			continue
		}

		res++
		for k := i; k < n; k += i {
			//额外判断一下也可以 但是就是有点耗时了
			if !visited[k] {
				visited[k] = true
			}
		}
	}

	return res
}

//快速幂运算 如何防止溢出  part1和part2可以看作是原数字的拆分 所以原数字的mod 肯定是等于part的mod 和part2的mod的mod 然后在 mod1337
// a*b %1337 = a%1337 * b%1337 %1337  所以我认为实际上 part1 已经是a %1337的结果了
func superPow1(a int, b []int) int {
	var travel func(x int, y int) int

	travel = func(x int, y int) int {
		if y == 0 {
			return 1
		}
		if y == 1 {
			return x
		}

		return ((x % 1337) * travel(x, y-1) % 1337) % 1337
	}

	last := b[len(b)-1]
	b = b[:len(b)-1]
	var part1 int
	var part2 int
	if len(b) == 0 {
		part1 = travel(a, last)
		part2 = travel(a, 0)
	} else {
		part1 = travel(a, last)
		part2 = travel(superPow(a, b), 10)
	}
	return (part1 * part2) % 1337
}

//快速幂  a的b次方 如果b是偶数 就对半拆 如果b是奇数  就拆成 1和b-1
func superPow(a int, b []int) int {
	var travel func(x int, y int) int

	travel = func(x int, y int) int {
		if y == 0 {
			return 1
		}

		if y == 1 {
			return x % 1337
		}

		if y%2 == 0 {
			return (travel(x, y/2) * travel(x, y/2)) % 1337
		} else {
			return (travel(x, 1) * travel(x, y-1)) % 1337
		}
	}

	last := b[len(b)-1]
	b = b[:len(b)-1]
	var part1 int
	var part2 int
	if len(b) == 0 {
		part1 = travel(a, last)
		part2 = travel(a, 0)
	} else {
		part1 = travel(a, last)
		part2 = travel(superPow(a, b), 10)
	}
	return (part1 * part2) % 1337
}

//线性的时候 使用二分快速定位
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for i := 0; i < len(piles); i++ {
		if piles[i] > max {
			max = piles[i]
		}
	}

	var travel func(eat int) bool
	travel = func(eat int) bool {
		temp := 0
		for i := 0; i < len(piles); i++ {
			if piles[i]%eat == 0 {
				temp += piles[i] / eat
			} else {
				temp += (piles[i] / eat) + 1
			}
		}
		if temp <= h {
			return true
		}
		return false

	}

	left := 1
	right := max
	for left <= right {
		mid := left + (right-left)/2
		if travel(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func shipWithinDays(weights []int, days int) int {
	var travel func(capability int) bool
	travel = func(capability int) bool {
		temp := 1
		c := capability
		for i := 0; i < len(weights); i++ {
			if c >= weights[i] {
				c -= weights[i]
			} else {
				temp++
				c = capability
				c -= weights[i]
			}
		}
		if temp <= days {
			return true
		}
		return false
	}

	max := weights[0]
	for i := 0; i < len(weights); i++ {
		if weights[i] > max {
			max = weights[i]
		}
	}
	left := max
	right := max * len(weights)
	for left <= right {
		mid := left + (right-left)/2
		if travel(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

//接雨水问题 找规律嘛 待优化
func trap1(height []int) int {
	res := 0

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < len(height)-1; i++ {
		left, right := height[i], height[i]
		for j := i - 1; j >= 0; j-- {
			if height[j] > left {
				left = height[j]
			}
		}
		if left == height[i] {
			continue
		}
		for j := i + 1; j < len(height); j++ {
			if height[j] > right {
				right = height[j]
			}
		}
		if right == height[i] {
			continue
		}
		res += min(left, right) - height[i]
	}
	return res
}

//优化有点不到位呀 就动态更新了一下 从左往右移动的时候 如果当前的比左边的大 不仅要跳过去 还要把左边的最大值更新一下
//如果当前的索引超过了右边最大值的索引  那右边就重新开始找当前右边的最大值
func trap2(height []int) int {
	res := 0

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	leftIndex, rightIndex := 0, 0
	for i := 1; i < len(height); i++ {
		if height[i] > height[rightIndex] {
			rightIndex = i
		}
	}
	for i := 1; i < len(height)-1; i++ {

		if height[leftIndex] <= height[i] {
			leftIndex = i
			continue
		}

		if rightIndex <= i {
			rightIndex = i
			for j := i + 1; j < len(height); j++ {
				if height[j] > height[rightIndex] {
					rightIndex = j
				}
			}
		}

		if height[rightIndex] <= height[i] {
			continue
		}

		res += min(height[leftIndex], height[rightIndex]) - height[i]
	}
	return res
}

//备忘录优化
//要求左边的最大值 那么就是索引-1的左边最大值和左边的数字比较 谁大就是目前索引的左边最大值!!!!
func trap(height []int) int {
	res := 0

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	left := make([]int, len(height), len(height))
	right := make([]int, len(height), len(height))

	left[0] = height[0]
	right[len(height)-1] = height[len(height)-1]

	//好补! 这里绝了 类似于动态规划了 要求左边的最大值 那么就是索引-1的左边最大值和左边的数字比较 谁大就是目前索引的左边最大值
	for i := 1; i < len(left); i++ {
		left[i] = max(height[i], left[i-1])
	}

	for i := len(right) - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}

	for i := 1; i < len(height)-1; i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}

//删除有序数组的重复元素 使用o1 复杂度
//双指针 注意替换规则 如果fast和slow不等 那么就slow+1 跃迁到未知的位置 然后把fast的值拷贝到slow中 然后不论如何fast都要++
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

//最长回文子串  双指针 中心拓扑
func longestPalindrome(s string) string {
	res := ""

	var travel func(i int)
	travel = func(i int) {
		temp := 0
		for i+temp < len(s) && i-temp >= 0 && s[i+temp] == s[i-temp] {
			temp++
		}
		if len(res) < (temp-1)*2+1 {
			res = s[i-temp+1 : i+temp]
		}

		temp1 := 0
		for i+1+temp1 < len(s) && i-temp1 >= 0 && s[i+1+temp1] == s[i-temp1] {
			temp1++
		}
		if len(res) < temp1*2 {
			res = s[i-temp1+1 : i+1+temp1]
		}
	}

	for i := 0; i < len(s); i++ {
		travel(i)
	}

	return res
}

//不知道做出了动态规划 ...
func canJump1(nums []int) bool {
	//初始化 最开始能跳跃的最远距离是0  后面更新
	farest := 0

	//记录上次的
	last := 0
	for farest < len(nums)-1 {
		//记录当前索引最多能跳多远
		num := nums[farest]
		//定义max 为当前的距离 为了和后面计算的最远距离作比较
		max := farest
		//从上次的开始 计算能够跳跃的这段距离里 这段距离中能跳出下一次最远距离的那个索引
		for i := last; i <= farest+num; i++ {

			if i >= len(nums)-1 {
				return true
			}

			jump := nums[i]
			if i+jump > max {
				max = i + jump
			}
		}

		//更新一下当前的最远距离 为下次计算做准备
		last = farest
		//更新最远距离
		farest = max
		//如果和上次的最远距离比较 没变 那么就表示跳不到最前面了
		if farest <= last {
			return false
		}
	}
	return farest >= len(nums)-1
}

//贪心思路
func canJump(nums []int) bool {
	//初始化 最开始能跳跃的最远距离是0  后面更新
	farest := 0

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums)-1; i++ {
		farest = max(i+nums[i], farest)

		if farest <= i {
			return false
		}
	}
	return farest >= len(nums)-1
}

//跳跃游戏Ⅱ 动态规划 +备忘录
func jump1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}

	m := make(map[int]int)

	var travel func(i int) int
	travel = func(i int) int {
		if i >= len(nums)-1 {
			return 0
		}
		res := math.MaxInt >> 1

		for j := i + 1; j <= i+nums[i]; j++ {
			if v, ok := m[j]; !ok {
				m[j] = travel(j) + 1
				res = min(res, m[j])
			} else {
				res = min(res, v)
			}
		}
		return res
	}

	return travel(0)
}

//跳跃游戏Ⅱ 贪心 贪心的性质是 每一步都能达到局部的最优  到最后就能打到全局的最优
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var travel func(i int) int
	travel = func(i int) int {
		if i+nums[i] >= len(nums)-1 {
			return 1
		}

		index := i
		for j := i; j <= i+nums[i]; j++ {
			//判断谁更有潜力
			if index+nums[index] < j+nums[j] {
				index = j
			}
		}

		//赋值的时候不能赋跳过去的最大值 而应该赋当前的索引值
		return travel(index) + 1
	}
	return travel(0)
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 0, 2)
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]] = true
	}
	for i := 0; i < len(nums2); i++ {
		m2[nums2[i]] = true
	}

	res1 := make([]int, 0, 0)
	for key := range m1 {
		if !m2[key] {
			res1 = append(res1, key)
		}
	}

	res2 := make([]int, 0, 0)
	for key := range m2 {
		if !m1[key] {
			res2 = append(res2, key)
		}
	}

	res = append(res, res1)
	res = append(res, res2)
	return res
}

func minDeletion(nums []int) int {
	res := 0
	var travel func([]int)
	travel = func(ints []int) {
		if len(ints) == 0 {
			return
		}
		if len(ints) == 1 {
			res++
			return
		}

		if ints[0] == ints[1] {
			res++
			travel(ints[1:])
		} else if ints[0] != ints[1] {
			travel(ints[2:])
		}
	}
	travel(nums)
	return res
}

//思路跟不上 构造回文  使用的是先拿到一半 然后再去添加构造
func kthPalindrome(queries []int, intLength int) []int64 {
	res := make([]int64, 0)

	var start int
	if intLength%2 == 0 {
		start = int(math.Pow(10, float64(intLength)/2-1))
	} else {
		start = int(math.Pow(10, float64(intLength/2)))
	}

	for i := 0; i < len(queries); i++ {
		num := start + queries[i] - 1
		bytes := []byte(fmt.Sprintf("%v", num))
		//复制新数组 不会影响原来的
		newbytes := append([]byte{}, bytes...)
		for j := 0; j < len(newbytes)/2; j++ {
			newbytes[j], newbytes[len(newbytes)-1-j] = newbytes[len(newbytes)-1-j], newbytes[j]
		}
		if intLength%2 == 0 {
			bytes = append(bytes, newbytes...)
		} else {
			bytes = append(bytes, newbytes[1:]...)
		}
		if len(bytes) > intLength {
			res = append(res, -1)
		} else {
			atoi, _ := strconv.Atoi(string(bytes))
			res = append(res, int64(atoi))
		}
	}
	return res
}

func kthPalindrome1(queries []int, intLength int) []int64 {
	res := make([]int64, 0, 0)
	for _, q := range queries {
		q--
		if intLength&1 == 0 {
			st := int(math.Pow10(intLength/2 - 1))
			v := st + q
			s := []byte(strconv.Itoa(v))
			t := append([]byte{}, s...)
			for i, n := 0, len(t); i < n/2; i++ {
				t[i], t[n-1-i] = t[n-1-i], t[i]
			}
			s = append(s, t...)
			if len(s) > intLength {
				res = append(res, -1)
				continue
			}
			w, _ := strconv.ParseInt(string(s), 10, 64)
			res = append(res, w)
		} else {
			st := int(math.Pow10(intLength / 2))
			v := st + q
			s := []byte(strconv.Itoa(v))
			t := append([]byte{}, s...)
			for i, n := 0, len(t); i < n/2; i++ {
				t[i], t[n-1-i] = t[n-1-i], t[i]
			}
			s = append(s, t[1:]...)
			if len(s) > intLength {
				res = append(res, -1)
				continue
			}
			w, _ := strconv.ParseInt(string(s), 10, 64)
			res = append(res, w)
		}
	}
	return res
}

//动态规划 01背包问题
func maxValueOfCoins(piles [][]int, k int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	m := len(piles)
	n := k + 1
	preSum := make([][]int, len(piles), len(piles))
	for i := 0; i < len(preSum); i++ {
		preSum[i] = make([]int, len(piles[i])+1, len(piles[i])+1)
		for j := 0; j < len(piles[i]); j++ {
			preSum[i][j+1] = piles[i][j] + preSum[i][j]
		}
	}

	dp := make([][]int, m+1, m+1)

	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n, n)
	}

	for i := 0; i < len(dp); i++ {
		dp[i][0] = 0
	}

	for i := 1; i < len(dp); i++ {
		ints := piles[i-1]
		for j := 1; j < k+1; j++ {
			if len(ints) == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else {
				//初始化为上一行的j个
				dp[i][j] = dp[i][j-1]
				for q := 0; q <= j && len(preSum[i-1])-1 >= q; q++ {
					//循环的和之前的做比较 比谁大  这里很烦的地方在于 索引老是搞错呀
					dp[i][j] = max(dp[i][j], dp[i-1][j-q]+preSum[i-1][q])
				}
			}
		}
	}
	return dp[len(piles)][k]
}

//贪心思想 区间重叠问题
func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	for i := 0; i < len(intervals)-1; {
		j := i + 1
		//这里非常关键哈 如果判定重复了 就继续下一个判定
		for j < len(intervals) && intervals[i][1] > intervals[j][0] {
			res++
			j++
		}
		//最后把j赋值给i  这样实际上复杂度还是o(n)
		i = j
	}
	return res
}

//贪心 最少的箭射气球 依然是计算重复区间的问题鸭!!!
func findMinArrowShots(points [][]int) int {
	res := 0
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	for i := 0; i < len(points); {
		j := i + 1
		for j < len(points) && points[i][1] >= points[j][0] {
			j++
		}
		res++
		i = j
	}
	return res
}

//看起来多少有些花里胡哨了 实际上就是解决转圈圈问题 就是取 1到len  和0到len-1 找到最大值 进行比较
func rob3(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	temp1 := make([]int, len(nums)-1, len(nums)-1)
	temp2 := make([]int, len(nums)-1, len(nums)-1)
	temp1[0] = nums[0]
	temp2[0] = nums[1]
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums)-1; i++ {
		if i-2 >= 0 {
			temp1[i] = max(temp1[i-1], nums[i]+temp1[i-2])
		} else {
			temp1[i] = max(temp1[i-1], nums[i])
		}
	}

	for i := 2; i < len(nums); i++ {
		if i-2-1 >= 0 {
			temp2[i-1] = max(temp2[i-1-1], nums[i]+temp2[i-2-1])
		} else {
			temp2[i-1] = max(temp2[i-1-1], nums[i])
		}
	}

	sort.Ints(temp1)
	sort.Ints(temp2)

	return max(temp1[len(temp1)-1], temp2[len(temp2)-1])
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	temp1 := make([]int, len(nums)-1, len(nums)-1)
	copy(temp1, nums)
	temp2 := make([]int, len(nums)-1, len(nums)-1)
	copy(temp2, nums[1:])
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums)-1; i++ {
		if i-2 >= 0 {
			temp1[i] = max(temp1[i-1], temp1[i]+temp1[i-2])
		} else {
			temp1[i] = max(temp1[i-1], temp1[i])
		}
	}

	for i := 1; i < len(nums)-1; i++ {
		if i-2 >= 0 {
			temp2[i] = max(temp2[i-1], temp2[i]+temp2[i-2])
		} else {
			temp2[i] = max(temp2[i-1], temp2[i])
		}
	}

	sort.Ints(temp1)
	sort.Ints(temp2)

	return max(temp1[len(temp1)-1], temp2[len(temp2)-1])
}

//太暴力了
func deleteAndEarn1(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(nums)
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j]+1 == nums[i] {
				continue
			}
			if j < 0 {
				dp[i] = nums[i]
			} else {
				dp[i] = max(dp[j]+nums[i], dp[i])
			}
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}

//删除并获得点数  和打家劫舍差不多的其实  没事就对数组排个序 不吃亏
func deleteAndEarn(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(nums)
	temp := make([]int, nums[len(nums)-1]+1, nums[len(nums)-1]+1)
	//预处理一下  去重
	for i := 0; i < len(nums); i++ {
		temp[nums[i]] += nums[i]
	}

	dp := make([]int, len(temp), len(temp))
	dp[0] = temp[0]
	for i := 1; i < len(temp); i++ {
		dp[i] = temp[i]
		j := i - 2
		if j < 0 {
			continue
		} else {
			//状态转移方程
			dp[i] = max(dp[j]+temp[i], dp[i-1])
		}
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, 0)
	var travel func(left, right int, s string)
	travel = func(left, right int, s string) {
		if left == 0 && right == 0 {
			res = append(res, s)
			return
		}
		if right < left {
			return
		}

		if left > 0 {
			travel(left-1, right, s+"(")
		}

		if right > 0 {
			travel(left, right-1, s+")")
		}

	}
	travel(n, n, "")
	return res
}

//暴力递归做不出来呀 超时
func checkValidString1(s string) bool {
	if len(s) == 0 {
		return true
	}

	sss := make([]byte, 0, 0)
	for i := 0; i < len(s); i++ {
		if len(sss) == 0 {
			sss = append(sss, s[i])
		} else if sss[len(sss)-1] == '(' && s[i] == ')' {
			sss = sss[:len(sss)-1]
		} else {
			sss = append(sss, s[i])
		}
	}

	if len(sss) == 0 {
		return true
	}

	var travel func(i int, left int, right int) bool
	travel = func(i int, left int, right int) bool {
		if right > left {
			return false
		}

		if i == len(sss)-1 {
			return left == right
		}

		if sss[i] == '(' {
			return travel(i+1, left+1, right)
		} else if sss[i] == ')' {
			return travel(i+1, left, right+1)
		} else {
			blank := travel(i+1, left, right)
			if blank {
				return true
			}

			l := travel(i+1, left+1, right)
			if l {
				return true
			}

			r := travel(i+1, left, right+1)
			if r {
				return true
			}
			return false
		}
	}
	return travel(0, 0, 0)
}

//贪心 太秒了 我擦
//从前往后遍历 时刻判断 右括号的数量会不会远超过左括号+星号的数量
//从后往前遍历 时刻判断 左括号的数量会不会远超过右括号+星号的数量
func checkValidString(s string) bool {
	left, right, star := 0, 0, 0
	for i := 0; i < len(s); i++ {

		if s[i] == '(' {
			left++
		} else if s[i] == '*' {
			star++
		} else {
			right++
		}
		if right > left+star {
			return false
		}
	}

	left, right, star = 0, 0, 0
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '(' {
			left++
		} else if s[i] == '*' {
			star++
		} else {
			right++
		}
		if left > right+star {
			return false
		}
	}

	return true
}

func isValid(s string) bool {
	stack := make([]byte, 0, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		if s[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

//找规律  只要是4的倍数 基本没戏了 不是4的倍数 就可以进行调整
func canWinNim(n int) bool {
	return n%4 != 0
}

//石子游戏 找规律呀 有点智力题 智商跟不上了
//就是可以控制自己拿的是索引奇数堆还是索引偶数堆 所以只需要实现记录一下哪一堆多 就可以达到控制的效果
func stoneGame(piles []int) bool {
	return true
}

func bulbSwitch1(n int) int {
	res := 0
	bools := make([]bool, n, n)
	for i := 1; i <= n; i++ {
		for j := i - 1; j < n; j += i {
			bools[j] = !bools[j]
		}
	}
	for i := 0; i < len(bools); i++ {
		if bools[i] {
			res++
		}
	}
	return res
}

//4 的时候是1  为啥呢 1的时候翻了一次 2的时候翻了一次 4的时候翻了一次 翻了奇数次 所以是1
// 	n的开方就是指  n==4的时候 有哪些是1 的呢 显然  1 和4 是 因为1的时候1*1重复了  就会变成奇数次 4的时候2*2重复了 就会变成奇数次
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pre := 0
	var in int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[pre] {
			in = i
			break
		}
	}
	root := &TreeNode{
		Val: preorder[pre],
	}

	var travel func(a []int, b []int) *TreeNode
	travel = func(a []int, b []int) *TreeNode {
		if len(a) == 1 && len(b) == 1 {
			return &TreeNode{Val: a[0]}
		}
		if len(a) == 0 && len(b) == 0 {
			return nil
		}
		pre1 := 0
		var in1 int
		for i := 0; i < len(b); i++ {
			if b[i] == a[pre1] {
				in1 = i
				break
			}
		}
		root1 := &TreeNode{Val: a[pre1]}
		root1.Left = travel(a[pre1+1:in1+1], b[pre1:in1])
		root1.Right = travel(a[in1+1:], b[in1+1:])
		return root1
	}
	root.Left = travel(preorder[pre+1:in+1], inorder[pre:in])
	root.Right = travel(preorder[in+1:], inorder[in+1:])
	return root
}

//最大路径和 维护一个res 动态的更新
func maxPathSum1(root *TreeNode) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := math.MinInt >> 1
	var travel func(root *TreeNode, sum int) int
	travel = func(root *TreeNode, sum int) int {
		if root == nil {
			return sum
		}

		left := math.MinInt >> 1
		right := math.MinInt >> 1
		if root.Left != nil {
			left = travel(root.Left, 0)
		}
		if root.Right != nil {
			right = travel(root.Right, 0)
		}

		if left >= 0 && right >= 0 {
			larger := max(left, right)
			//如果连起来比左右任何一个都大 那就用连起来的和去和res去更新比较
			if left+right+root.Val >= larger {
				res = max(res, left+right+root.Val)
			} else {
				//否则就只拿大的一边去和res更新
				res = max(res, larger)
			}
			//返回的时候 只能连一边
			return larger + root.Val
		} else if left >= 0 && right < 0 {
			if root.Val > 0 {
				//如果节点大于0 那就加起来和res更新比较
				res = max(res, left+root.Val)
			} else {
				//否则 就只拿一遍去比较
				res = max(res, left)
			}
			//返回的时候 只能连一边
			return left + root.Val
		} else if right >= 0 && left < 0 {
			if root.Val > 0 {
				res = max(res, right+root.Val)
			} else {
				res = max(res, right)
			}
			return right + root.Val
		} else {
			//如果都小于0
			i := max(left, right)
			//如果节点值是最大的 那就和res更新比较
			if root.Val > i {
				res = max(res, root.Val)
			} else {
				//否则 就拿一个最大的更新比较
				res = max(res, i)
			}

			//由于两边都小于0 所以只能返回自己
			return root.Val

		}
	}

	travel(root, res)
	return res
}

//最大路径和 维护一个res 动态的更新
func maxPathSum(root *TreeNode) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := math.MinInt >> 1
	var travel func(root *TreeNode) int
	travel = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := travel(root.Left)
		right := travel(root.Right)
		res = max(res, left+right+root.Val)
		res = max(res, root.Val)
		if root.Val+max(left, right) <= 0 {
			return 0
		}
		return root.Val + max(left, right)
	}

	travel(root)
	return res
}

func minBitFlips(start int, goal int) int {
	res := 0
	ans := start ^ goal
	for i := 0; i < 32; i++ {
		if ans&1 != 0 {
			res++
		}
		ans >>= 1
	}
	return res
}

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp := nums
	for len(temp) != 1 {
		newTemp := make([]int, len(temp)-1, len(temp)-1)
		for i := 0; i < len(temp)-1; i++ {
			newTemp[i] = (temp[i] + temp[i+1]) % 10
		}
		temp = newTemp
	}

	return temp[0]
}

func sumScores1(s string) int64 {
	res := int64(0)
	res += int64(len(s))
	for i := 1; i < len(s); i++ {
		temp := s[i:]
		index := 0
		for j := 0; j < len(temp); j++ {
			if temp[j] == s[index] {
				index++
			} else {
				break
			}
		}
		res += int64(index)
		index = 0
	}

	return res
}

func numberOfWays(s string) int64 {
	res := int64(0)
	zeroCount := strings.Count(s, "0")
	oneCount := strings.Count(s, "1")

	a := 0
	for _, v := range s {
		if v == '0' {
			a++
		} else if v == '1' {
			res += int64(a * (zeroCount - a))
		}
	}

	b := 0
	for _, v := range s {
		if v == '1' {
			b++
		} else if v == '0' {
			res += int64(b * (oneCount - b))
		}
	}

	return res
}

func sumScores(s string) int64 {
	ans := 0
	calcZ := func(s []byte) []int {
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
		}
		z[0] = n
		return z
	}

	z := calcZ([]byte(s))

	for _, v := range z {
		ans += v
	}

	return int64(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
