package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//distance := maxDistance([]int{0,1})
	//fmt.Println(distance)

	//plants := wateringPlants([]int{7,7,7,7,7,7,7}, 8)
	//fmt.Println(plants)

	//rangeFreqQuery := Constructor([]int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56})
	//query := rangeFreqQuery.Query(1, 2, 4)
	//fmt.Println(query)

	//mirror := kMirror(2, 5)
	//fmt.Println(mirror)

	//indices := targetIndices([]int{1, 2, 5, 2, 3}, 4)
	//fmt.Println(indices)

	//averages := getAverages([]int{7,4,3,9,1,8,5,2,6}, 3)
	//fmt.Println(averages)

	//deletions := minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6})
	//fmt.Println(deletions)

	//numbers := findEvenNumbers([]int{2, 2, 8, 8, 2})
	//fmt.Println(numbers)
	//l := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 4,
	//			Next: &ListNode{
	//				Val: 7,
	//				Next: &ListNode{
	//					Val: 1,
	//					Next: &ListNode{
	//						Val: 2,
	//						Next: &ListNode{
	//							Val:  6,
	//							Next: nil,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	//fmt.Println(l)
	//middle := deleteMiddle(&ListNode{
	//	Val:  1,
	//	Next: nil,
	//})
	//fmt.Println(middle)

	//directions := getDirections(&TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 1,
	//		Left: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   2,
	//		Left:  &TreeNode{4, nil, nil},
	//		Right: &TreeNode{6, nil, nil},
	//	},
	//}, 3, 6)
	//fmt.Println(directions)

	//subsequence := maxSubsequence([]int{-1, -2, 3, 4}, 3)
	//fmt.Println(subsequence)

	//bank := goodDaysToRobBank([]int{77, 9, 111, 138, 139, 183, 112, 127, 38, 123, 198, 86, 163, 50, 140, 106, 99, 140, 152, 176, 124, 181, 14, 113, 53, 186, 76, 165, 32, 26, 137, 4, 186, 193, 130, 157, 173, 52, 27, 101, 154, 129, 34, 200, 51, 184, 127, 147, 197, 151, 190, 95, 62, 182, 77, 34, 174, 162, 7, 84, 105, 103, 128}, 2)
	//bank := goodDaysToRobBank([]int{5, 3, 3, 3, 5, 6, 2}, 2)
	//fmt.Println(bank)

	//points := countPoints("G4")
	//fmt.Println(points)

	//ranges := subArrayRanges([]int{4, -2, -3, 4, 1})
	//fmt.Println(ranges)

	//refill := minimumRefill([]int{2, 2, 5, 2, 2}, 5, 5)
	//fmt.Println(refill)

	//palindrome := firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"})
	//fmt.Println(palindrome)

	//spaces := addSpaces("icodeinpython", []int{1, 5, 7, 9})
	//fmt.Println(spaces)

	//periods := getDescentPeriods([]int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 4, 3, 10, 9, 8, 7})
	//fmt.Println(periods)
	//
	//println(jc(10))

	//found := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"})
	//fmt.Println(found)

	//recipes := findAllRecipes([]string{"bread", "sandwich"},
	//	[][]string{{"yeast", "flour"}, {"bread", "meat"}},
	//	[]string{"yeast", "flour", "meat"})
	//fmt.Println(recipes)

	//valid := canBeValid("))()))", "010100")
	//fmt.Println(valid)

	//reversals := isSameAfterReversals(526)
	//fmt.Println(reversals)

	//instructions := executeInstructions(3, []int{0, 1},
	//	"RRDDLU")
	//fmt.Println(instructions)

	//distances := getDistances([]int{2, 1, 3, 1, 2, 3, 3})
	//fmt.Println(distances)

	//b := checkString("a")
	//fmt.Println(b)

	//beams := numberOfBeams([]string{"000", "111", "000"})
	//fmt.Println(beams)

	//destroyed := asteroidsDestroyed(5, []int{4, 9, 23, 4})
	//fmt.Println(destroyed)

	//title := capitalizeTitle("i lOve leetcode")
	//fmt.Println(title)

	//pairSum()

	//palindrome := longestPalindrome([]string{"cl", "lc", "dd", "aa", "bb", "dd", "aa", "dd", "bb", "dd", "aa", "cc", "bb", "cc", "dd", "cc"})
	//palindrome := longestPalindrome([]string{"em", "pe", "mp", "ee", "pp", "me", "ep", "em", "em", "me"})
	//fmt.Println(palindrome)
	//
	//valid := checkValid([][]int{{15, 7, 18, 11, 19, 10, 14, 16, 8, 2, 3, 6, 5, 1, 17, 12, 9, 4, 13}, {17, 15, 9, 8, 11, 13, 7, 6, 5, 1, 3, 16, 12, 19, 10, 2, 4, 14, 18}, {19, 14, 12, 10, 8, 9, 17, 16, 4, 3, 13, 18, 1, 5, 7, 11, 2, 15, 6}, {4, 2, 10, 15, 19, 16, 8, 9, 5, 3, 1, 11, 13, 14, 6, 18, 12, 17, 7}, {13, 19, 9, 16, 5, 8, 6, 12, 14, 11, 18, 10, 7, 2, 3, 4, 15, 17, 1}, {4, 7, 18, 11, 17, 16, 5, 12, 10, 1, 15, 13, 14, 6, 19, 2, 3, 9, 8}, {14, 5, 15, 1, 18, 6, 12, 7, 8, 9, 3, 13, 2, 10, 19, 4, 11, 16, 17}, {10, 3, 1, 8, 14, 19, 11, 18, 15, 13, 9, 12, 16, 17, 7, 4, 5, 2, 6}, {14, 13, 19, 18, 7, 2, 4, 8, 10, 17, 12, 5, 15, 1, 6, 9, 11, 3, 16}, {19, 8, 10, 18, 16, 12, 11, 17, 4, 9, 7, 2, 5, 13, 15, 3, 6, 1, 14}, {1, 10, 6, 14, 7, 18, 3, 9, 4, 16, 5, 11, 13, 17, 15, 8, 19, 2, 12}, {13, 10, 5, 16, 1, 19, 17, 3, 9, 11, 7, 8, 12, 6, 4, 2, 14, 15, 18}, {17, 2, 1, 6, 9, 19, 18, 14, 4, 11, 12, 13, 16, 5, 8, 7, 3, 10, 15}, {1, 4, 10, 5, 13, 6, 18, 11, 3, 2, 15, 14, 16, 12, 17, 19, 8, 9, 7}, {2, 14, 3, 12, 16, 17, 11, 9, 1, 6, 5, 19, 10, 13, 4, 18, 7, 15, 8}, {15, 9, 8, 18, 14, 13, 4, 12, 5, 17, 6, 1, 11, 16, 19, 3, 7, 2, 10}, {15, 8, 12, 16, 13, 2, 6, 19, 18, 14, 10, 5, 11, 9, 7, 1, 3, 17, 4}, {15, 6, 17, 7, 5, 3, 1, 9, 19, 12, 10, 11, 16, 14, 18, 8, 2, 13, 4}, {6, 11, 10, 14, 2, 13, 16, 1, 9, 15, 8, 19, 17, 3, 5, 18, 7, 4, 12}})
	//fmt.Println(valid)

	//h := "hello"
	//bytes := []byte(h)
	//sort.Slice(bytes, func(i, j int) bool {
	//	if bytes[i] < bytes[j] {
	//		return false
	//	}
	//	return true
	//})
	//s := string(h)
	//fmt.Println(s)
	//count := wordCount([]string{"mox", "bj", "rsy", "jqsh"}, []string{"trk", "vjb", "jkr"})
	//fmt.Println(count)

	//i := divideString("abcdefghi", 3, 97)
	//fmt.Println(i)

	//moves := minMoves(1000000000, 1)
	//fmt.Println(moves)

	//points := mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}})
	points := mostPoints([][]int{{29, 1}, {90, 5}, {41, 5}, {46, 3}, {49, 5}, {49, 2}, {6, 5}, {36, 5}, {83, 1}, {60, 2}, {97, 3}, {54, 2}, {42, 5}, {42, 2}, {73, 4}, {38, 4}, {16, 4}, {44, 2}, {81, 2}, {76, 3}, {60, 4}, {83, 4}, {94, 2}, {13, 5}, {7, 2}, {77, 2}, {18, 2}, {91, 2}, {43, 4}, {84, 2}, {45, 1}, {42, 5}, {54, 4}, {18, 4}, {96, 5}, {44, 3}, {55, 4}, {49, 3}, {86, 2}, {41, 3}, {54, 3}, {66, 2}, {22, 3}, {35, 5}, {89, 3}, {61, 2}, {1, 3}, {72, 1}, {13, 3}, {70, 4}, {12, 4}, {35, 5}, {16, 3}, {67, 3}, {70, 3}, {5, 4}, {74, 4}, {36, 1}, {47, 2}, {72, 1}})
	//points := mostPoints([][]int{{62, 62}, {1, 84}, {44, 43}, {15, 95}, {18, 35}, {9, 45}, {7, 98}, {64, 78}, {77, 31}, {39, 93}, {19, 8}, {90, 82}, {69, 87}, {27, 30}, {44, 97}, {22, 95}, {97, 97}, {50, 52}, {93, 72}, {26, 37}, {52, 34}, {26, 21}, {66, 67}, {97, 41}, {27, 86}, {31, 46}, {10, 31}, {30, 45}, {98, 57}, {5, 67}, {33, 42}, {79, 87}, {27, 88}, {33, 22}, {94, 45}, {2, 29}, {87, 98}, {70, 31}, {54, 53}, {90, 82}, {85, 52}, {8, 17}, {7, 36}, {67, 81}, {67, 60}, {14, 54}, {93, 4}, {75, 50}, {31, 13}, {39, 3}, {69, 71}, {63, 36}, {65, 86}, {97, 38}, {1, 3}, {50, 2}, {26, 10}, {97, 16}, {34, 12}, {43, 59}, {19, 60}, {27, 82}, {56, 79}, {57, 30}, {77, 45}, {67, 78}, {44, 84}, {45, 72}, {88, 79}, {16, 23}, {95, 48}, {40, 63}, {68, 72}, {45, 27}, {81, 80}, {62, 26}, {23, 82}, {74, 89}, {67, 87}, {58, 25}, {65, 84}, {51, 89}, {96, 100}, {77, 95}, {93, 86}, {21, 76}, {40, 94}, {47, 60}})

	fmt.Println(points)
}

func maxDistance(colors []int) int {
	final := 0
	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors); j++ {
			if colors[i] != colors[j] {
				temp := int(math.Abs(float64(i - j)))
				if temp > final {
					final = temp
				}
			}
		}
	}
	return final
}

func wateringPlants(plants []int, capacity int) int {
	final := 0
	temp := capacity
	for i := 0; i < len(plants); i++ {
		if temp >= plants[i] {
			temp -= plants[i]
			final += 1
		} else if temp < plants[i] {
			temp = capacity
			final += i
			final += i + 1
			temp -= plants[i]
		}
	}
	return final
}

type RangeFreqQuery struct {
	slice   *[]int
	m       *map[int][]int
	account map[int]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int, 10)
	account := make(map[int]int)
	for i, v := range arr {
		m[v] = append(m[v], i)
		account[v]++
	}
	return RangeFreqQuery{
		slice:   &arr,
		m:       &m,
		account: account,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	final := 0

	if left == right {
		if (*this.slice)[left] == value {
			return 1
		} else {
			return 0
		}
	}

	if this.account[value] > right-left {
		for _, v := range (*this.slice)[left : right+1] {
			if v == value {
				final++
			}
		}
	} else {
		for _, v := range (*this.m)[value] {
			if v >= left && v <= right {
				final++
			}
		}
	}
	return final
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */

func kMirror(k int, n int) int64 {
	account := 0
	var sum int64
	for account < n {
		for i := 1; i < math.MaxInt; i++ {
			itoa := strconv.Itoa(i)
			b := true

			for k := 0; k < len(itoa)/2; k++ {
				if itoa[k] != itoa[len(itoa)-1-k] {
					b = false
					break
				}
			}
			if b {
				var s string
				temp := i
				for temp/k != 0 {
					s += strconv.Itoa(temp % k)
					temp = temp / k
				}
				s += strconv.Itoa(temp % k)

				b1 := true
				for k := 0; k < len(s)/2; k++ {
					if s[k] != s[len(s)-1-k] {
						b1 = false
						break
					}
				}
				if b1 {
					account++
					sum += int64(i)
				}
			}
		}
	}
	return sum
}

func targetIndices(nums []int, target int) []int {
	final := make([]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			final = append(final, i)
		}
	}
	return final
}

func getAverages(nums []int, k int) []int {
	if k == 0 {
		return nums
	}
	if (len(nums)-1)/2 < k {
		for i := 0; i < len(nums); i++ {
			nums[i] = -1
		}
		return nums
	}
	final := make([]int, 0, len(nums))
	c := 2*k + 1
	temp := 0

	for i := 0; i < len(nums); i++ {
		if i < k || len(nums)-1-i < k {
			final = append(final, -1)
		} else {
			if temp > 0 {
				temp -= nums[i-k-1]
				temp += nums[i+k]
			} else {
				for j := i - k; j <= i+k; j++ {
					temp += nums[j]
				}
			}
			final = append(final, temp/c)

		}
	}
	return final
}

func minimumDeletions1(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	final := 0
	temp := make([]int, len(nums), cap(nums))
	copy(temp, nums)
	sort.Ints(temp)
	min := temp[0]
	max := temp[len(temp)-1]
	res1 := 0
	res2 := 0
	count := 0
	b1 := false //都是左边
	b2 := false
	for i, v := range nums {
		if v == min || v == max {
			if count == 0 {
				if i+1 <= len(nums)-i {
					res1 = i + 1
					b1 = true
				} else {
					res1 = len(nums) - i
				}
				count++
			} else if count == 1 {
				if i+1 <= len(nums)-i {
					res2 = i + 1
					b2 = true
				} else {
					res2 = len(nums) - i
				}
				break
			}
		}
	}
	if b1 == true && b2 == true {
		final = res1 + int(math.Abs(float64(res2-res1)))
	} else if b1 == false && b2 == false {
		final = res1 + int(math.Abs(float64(res2-res1)))
	} else {
		final = res1 + res2
	}
	return final
}

func minimumDeletions(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	temp := make([]int, len(nums), cap(nums))
	copy(temp, nums)
	sort.Ints(temp)
	min := temp[0]
	max := temp[len(temp)-1]
	res1 := 0
	res2 := 0
	count := 0
	for i, v := range nums {
		if v == min || v == max {
			if count == 0 {
				res1 = i
				count++
			} else {
				res2 = i
				break
			}
		}
	}
	maxi := -1
	mini := -1
	if res2 >= res1 {
		maxi = res2
		mini = res1
	} else {
		maxi = res1
		mini = res2
	}

	a := maxi + 1
	b := len(nums) - mini
	c := mini + 1 + len(nums) - maxi
	var temp1 int
	if a > b {
		temp1 = b
	} else {
		temp1 = a
	}
	if temp1 > c {
		temp1 = c
	} else {

	}

	return temp1
}

func findEvenNumbers(digits []int) []int {
	final := make([]int, 0)
	sort.Ints(digits)
	m := make(map[int]bool)
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(digits); k++ {
				if k == i || k == j {
					continue
				}
				sprintf := fmt.Sprintf("%v%v%v", digits[i], digits[j], digits[k])
				atoi, _ := strconv.Atoi(sprintf)
				if atoi < 100 || atoi%2 == 1 || m[atoi] {
					continue
				}
				m[atoi] = true
				final = append(final, atoi)
			}
		}
	}
	return final
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	pre := head
	slow := head
	fast := head
	var temp *ListNode
	for fast != nil && fast.Next != nil {
		temp = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	temp.Next = slow.Next
	return pre
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {

	var b bool
	var travel func(root *TreeNode, startValue int, destValue int, direction string) string
	travel = func(root *TreeNode, startValue int, destValue int, direction string) string {
		res := ""
		if root == nil {
			return ""
		}
		left := travel(root.Left, startValue, destValue, "L")

		if b {
			res += direction
		}

		if root.Val == startValue {
			b = true
		}
		if root.Val == destValue {
			return res
		}
		if left != "" {
			left += "U"
		}
		right := travel(root.Right, startValue, destValue, "R")
		if right != "" {
			right += "U"
		}
		s := left + right

		return s
	}
	return travel(root, startValue, destValue, "")
}

func maxSubsequence(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}
	copy1 := make([]int, len(nums))
	copy(copy1, nums)
	sort.Ints(copy1)
	deleted := copy1[0 : len(nums)-k]
	for _, d := range deleted {
		for i, v := range nums {
			if v == d {
				temp := nums[:i]
				temp1 := nums[i+1:]
				nums = append(temp, temp1...)
				break
			}
		}
	}
	return nums
}

//之前从来没有遇到过 前缀和思想 从前往后记录下降的个数 从后往前记录下降的个数 然后两个个数都符合>time才行
func goodDaysToRobBank(security []int, time int) []int {
	final := make([]int, 0)
	if time == 0 {
		for i := 0; i < len(security); i++ {
			final = append(final, i)
		}
		return final
	}

	if (len(security)-1)/2 < time {
		return []int{}
	}

	count1 := 0
	base1 := security[0]
	ok1 := make([]int, 0)
	for i := 1; i < len(security); i++ {
		if security[i] < base1 {
			base1 = security[i]
			count1++
		} else if security[i] == base1 {
			base1 = security[i]
			count1++
		} else if security[i] > base1 {
			count1 = 0
			base1 = security[i]
		}
		if count1 >= time {
			ok1 = append(ok1, i)
		}
	}
	if len(ok1) == 0 {
		return []int{}
	}
	for i := 0; i < len(ok1); i++ {
		b := true
		q := security[ok1[i]]
		for j := 1; j <= time; j++ {
			if ok1[i]+time > len(security)-1 {
				b = false
				break
			}
			if q > security[ok1[i]+j] {
				b = false
				break
			}
			q = security[ok1[i]+j]
		}
		if b {
			final = append(final, ok1[i])
		}
	}
	return final
}

func countPoints(rings string) int {
	count := 0
	m := make(map[byte]map[byte]bool)
	for i := 0; i < len(rings); i += 2 {
		if _, ok := m[rings[i+1]]; !ok {
			m2 := make(map[byte]bool)
			m2[rings[i]] = true
			m[rings[i+1]] = m2
		} else {
			m[rings[i+1]][rings[i]] = true
		}
	}
	for _, m2 := range m {
		if len(m2) == 3 {
			count++
		}
	}
	return count
}

func subArrayRanges1(nums []int) int64 {
	var final int64 = 0
	for i := 1; i < len(nums); i++ {
		slow, fast := 0, i
		for fast < len(nums) {
			min, max := math.MaxInt64, math.MinInt64
			for j := slow; j <= fast; j++ {
				if nums[j] > max {
					max = nums[j]
				}
				if nums[j] < min {
					min = nums[j]
				}
			}
			final += int64(max - min)
			slow++
			fast++
		}

	}
	return final
}

func subArrayRanges(nums []int) int64 {
	if len(nums) == 1 || len(nums) == 0 {
		return 0
	}
	var final int64 = 0
	for i := 0; i < len(nums); i++ {
		slow, fast := i, i+1

		min, max := nums[slow], nums[slow]

		for fast < len(nums) {
			if nums[fast] > max {
				max = nums[fast]
			}
			if nums[fast] < min {
				min = nums[fast]
			}
			final += int64(max - min)
			fast++
		}
	}
	return final
}

func minimumRefill1(plants []int, capacityA int, capacityB int) int {
	count := 0
	temp1 := capacityA
	temp2 := capacityB
	before, after := 0, len(plants)-1
	for before <= after {
		if before == after {
			if capacityA >= capacityB {
				capacityA -= plants[before]
			} else {
				capacityB -= plants[after]
			}
		} else {
			capacityA -= plants[before]
			capacityB -= plants[after]
		}

		if capacityA < 0 {
			count++
			capacityA += temp1
		}
		if capacityB < 0 {
			count++
			capacityB += temp2
		}
		before++
		after--
	}
	return count
}

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	count := 0
	temp1 := capacityA
	temp2 := capacityB
	before, after := 0, len(plants)-1
	for before <= after {
		if before == after {
			if capacityA >= capacityB {
				capacityA -= plants[before]
			} else {
				capacityB -= plants[after]
			}
		} else {
			capacityA -= plants[before]
			capacityB -= plants[after]
		}

		if capacityA < 0 {
			count++
			capacityA = temp1 - plants[before]
		}
		if capacityB < 0 {
			count++
			capacityB = temp2 - plants[after]
		}
		before++
		after--
	}
	return count
}

func firstPalindrome(words []string) string {
	for _, v := range words {
		for i := 0; i <= len(v)/2; i++ {
			if v[i] == v[len(v)-i-1] {
				if i == len(v)/2 {
					return v
				}
			} else {
				break
			}
		}
	}
	return ""
}

func addSpaces1(s string, spaces []int) string {
	final := ""
	index := 0
	temp := make([]string, 0, len(spaces)+1)
	for _, v := range spaces {
		temp = append(temp, s[index:v])
		index = v
	}
	temp = append(temp, s[index:])
	for i := 0; i < len(temp); i++ {
		final += temp[i] + " "
	}
	final = final[:len(final)-1]
	return final
}

func addSpaces2(s string, spaces []int) string {
	final := s
	index := 0
	for _, v := range spaces {
		final = final[:v+index] + " " + final[v+index:]
		index++
	}
	return final
}

func addSpaces3(s string, spaces []int) string {
	final := ""
	temp := make([]string, 0, len(spaces)+1)
	for i := len(spaces) - 1; i >= 0; i-- {
		temp = append(temp, s[spaces[i]:])
		s = s[:spaces[i]]
	}

	for i := len(temp) - 1; i >= 0; i-- {
		final += temp[i]
		final += " "
	}
	final = s[:spaces[0]] + " " + final
	final = final[:len(final)-1]

	return final
}

func addSpaces(s string, spaces []int) string {
	temp := make([]string, 0, len(spaces)+1)
	count := 0
	index := 0
	for i := 0; i <= len(spaces)-1; i++ {
		temp = append(temp, s[:spaces[i]-index])
		s = s[spaces[i]-index:]
		index = spaces[i]
		count++
	}
	temp = append(temp, s)
	final := strings.Join(temp, " ")
	return final
}

func getDescentPeriods1(prices []int) int64 {
	var sum int64 = 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if i == j {
				sum++
				continue
			}
			if prices[j] == prices[j-1]-1 {
				sum++
				continue
			} else {
				break
			}
		}
	}
	return sum
}

func getDescentPeriods(prices []int) int64 {
	var sum int64 = 0
	var jiecheng int64 = 1
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			if j == 0 {
				continue
			}
			if prices[j] == prices[j-1]-1 {
				jiecheng++
				continue
			} else {
				break
			}
		}
		sum += jc(int64(jiecheng)) - jiecheng
		i = i + int(jiecheng) - 1
		jiecheng = 1
	}
	sum += int64(len(prices))
	return sum
}

func jc(n int64) int64 {
	if n > 0 {
		return n + jc(n-1)
	}
	return n
}

func mostWordsFound(sentences []string) int {
	final := 0
	for _, sentence := range sentences {
		split := strings.Split(sentence, " ")
		i := len(split)
		if final < i {
			final = i
		}
	}
	return final
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	final := make([]string, 0)
	iToR := make(map[string][]string, 0)
	Rsum := make(map[string]int)

	for i, recipe := range recipes {
		for _, v := range ingredients[i] {
			iToR[v] = append(iToR[v], recipe)
			Rsum[recipe]++
		}
	}

	for len(supplies) > 0 {
		supply := supplies[0]
		supplies = supplies[1:]
		for _, r := range iToR[supply] {
			Rsum[r]--
			if Rsum[r] == 0 {
				supplies = append(supplies, r)
				final = append(final, r)
			}
		}
	}
	return final
}

//func canBeValid(s string, locked string) bool {
//	if len(locked)%2 == 1 {
//		return false
//	}
//	final := make([]string, len(locked))
//	for i, v := range locked {
//		if v == '1' {
//			final[i] = fmt.Sprintf("%c", s[i])
//		}
//	}
//	if final[0] == "" {
//		final[0] = "("
//	}
//	index := 0
//	for i, v := range final {
//		if i == 0 {
//			continue
//		}
//		if v == ")" {
//			200
//		}
//	}
//	return true
//}

func isSameAfterReversals(num int) bool {
	itoa := strconv.Itoa(num)
	temp := ""
	for i := 0; i < len(itoa); i++ {
		temp += fmt.Sprintf("%c", itoa[len(itoa)-1-i])
	}
	final := ""
	i, _ := strconv.Atoi(temp)
	temp = strconv.Itoa(i)
	for i := 0; i < len(temp); i++ {
		final += fmt.Sprintf("%c", temp[len(temp)-1-i])
	}
	atoi, _ := strconv.Atoi(final)
	if atoi == num {
		return true
	}
	return false

}

func executeInstructions(n int, startPos []int, s string) []int {
	final := make([]int, len(s))
	if n <= 1 {
		return final
	}
	ii := startPos[0]
	j := startPos[1]

	var travel func(start []int, s string) int
	travel = func(start []int, s string) int {
		if len(s) == 0 {
			return 0
		}
		if s[0] == 'U' {
			start[0]--
		} else if s[0] == 'D' {
			start[0]++
		} else if s[0] == 'L' {
			start[1]--
		} else if s[0] == 'R' {
			start[1]++
		}
		if start[0] >= n || start[0] < 0 || start[1] >= n || start[1] < 0 {
			return 0
		} else {
			return travel(start, s[1:]) + 1
		}
	}

	for i := 0; i < len(s); i++ {
		final[i] = travel(startPos, s[i:])
		startPos[0] = ii
		startPos[1] = j
	}

	return final
}

func getDistances1(arr []int) []int64 {
	final := make([]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		target := arr[i]
		var tempSum int64
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			if arr[j] == target {
				tempSum += int64(math.Abs(float64(j - i)))
			}
		}
		final[i] = tempSum
	}
	return final
}

func getDistances2(arr []int) []int64 {
	final := make([]int64, len(arr))
	m := make(map[int]float64)
	for i := 0; i < len(arr); i++ {
		target := arr[i]
		var tempSum float64
		tempSum += m[i]
		for j := i; j < len(arr); j++ {
			if j == i {
				continue
			}
			if arr[j] == target {
				abs := math.Abs(float64(j - i))
				m[j] += abs
				tempSum += abs
			}
		}
		final[i] = int64(tempSum)
	}
	return final
}

func getDistances(arr []int) []int64 {
	final := make([]int64, len(arr))
	slice := make(map[int][]int)
	m := make(map[int]float64)

	for i := 0; i < len(arr); i++ {
		slice[arr[i]] = append(slice[arr[i]], i)
	}

	for _, ints := range slice {

		for j := 0; j < len(ints); j++ {
			var tempSum float64
			for k := j; k < len(ints); k++ {
				if ints[j] == ints[k] {
					tempSum += m[ints[k]]
					continue
				}

				abs := math.Abs(float64(ints[k] - ints[j]))

				m[ints[k]] += abs
				tempSum += abs
			}
			final[ints[j]] += int64(tempSum)
		}
	}

	return final
}

func checkString(s string) bool {
	a := strings.LastIndex(s, "a")
	b := strings.Index(s, "b")
	if a < b {
		return true
	} else if b == -1 {
		return true
	}
	return false
}

func numberOfBeams(bank []string) int {
	queue := make([]int, 0)
	for i := 0; i < len(bank); i++ {
		count := 0
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				count++
			}
		}
		if count == 0 {
			continue
		}
		queue = append(queue, count)
	}

	sum := 0
	if len(queue) == 0 {
		return 0
	}
	q := queue[0]
	queue = queue[1:]
	for len(queue) != 0 {
		q1 := queue[0]
		sum += q * q1
		q = q1
		queue = queue[1:]
	}
	return sum
}

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	for _, asteroid := range asteroids {
		if mass >= asteroid {
			mass += asteroid
		} else {
			return false
		}
	}
	return true
}

func capitalizeTitle(title string) string {
	split := strings.Split(title, " ")
	final := ""
	for _, s := range split {
		if len(s) <= 2 {
			final += strings.ToLower(s) + " "
		} else {
			final += strings.ToUpper(s[0:1]) + strings.ToLower(s[1:]) + " "
		}
	}
	return final[:len(final)-1]
}

func pairSum(head *ListNode) int {
	final := make([]int, 0)
	for head != nil {
		final = append(final, head.Val)
		head = head.Next
	}
	max := math.MinInt64
	for i := 0; i < len(final)/2; i++ {
		temp := final[i] + final[len(final)-i-1]
		if temp > max {
			max = temp
		}
	}

	return max
}

func longestPalindrome(words []string) int {
	m := make(map[string]int)
	res := 0
	for _, word := range words {
		if word[0] == word[1] {
			m[word]++
			//if res == 0 {
			//	res += 1
			//}
		} else {
			m[word]++
		}
	}

	b := false
	b1 := false

	res1 := 0
	for k, v := range m {
		if v >= 1 && k[1:] != k[:1] && m[k[1:]+k[:1]] >= 1 {
			min := math.Min(float64(v), float64(m[k[1:]+k[:1]]))
			res += int(min)
		} else if v >= 1 && k[1:] == k[:1] {
			b = true
			if v%2 == 1 {
				b1 = true
			}
			res1 += v / 2
		}
	}
	if !b {
		return res * 2
	}
	if !b1 {
		return (res1*2)*2 + res*2
	}
	return (res1*2+1)*2 + res*2
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m := make(map[[2]int]bool)
	final := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if m[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == 1 {
				m[[2]int{i, j}] = true
				final++
				continue
			}
			if i+stampHeight-1 <= len(grid) && j+stampWidth-1 <= len(grid[i]) {
				for k := i; k < i+stampHeight-1; k++ {
					for n := j; n < j+stampWidth-1; n++ {
						if !m[[2]int{k, n}] {
							m[[2]int{k, n}] = true
							final++
						}
					}
				}
			}
		}
	}
	if final == len(grid)*len(grid[0]) {
		return true
	}
	return false

}

func checkValid(matrix [][]int) bool {
	//line := make([]map[int]int, 0)
	height := len(matrix)
	width := len(matrix[0])
	//column := make([]map[int]int, width)
	for i := 0; i < height; i++ {
		line1 := make(map[int]int)
		for j := 0; j < width; j++ {
			line1[matrix[i][j]]++
		}
		if len(line1) != width {
			return false
		}
	}

	for i := 0; i < width; i++ {
		column1 := make(map[int]int)
		for j := 0; j < height; j++ {
			column1[matrix[j][i]]++
		}
		if len(column1) != height {
			return false
		}
	}
	return true
}

func wordCount(startWords []string, targetWords []string) int {
	res := 0
	m := make(map[int][]string)
	for _, word := range startWords {
		l := len(word)
		m[l] = append(m[l], word)
	}

	for _, word := range targetWords {
		l := len(word)
		for _, str := range m[l-1] {
			temp := str
			//length := len(temp)
			count := 0
			for _, byt := range word {

				if strings.ContainsRune(temp, byt) {
					continue
				} else {
					count += 1
					if count == 2 {
						break
					}
					//length++
					//if length > l {
					//	break
					//}
				}
			}
			if count == 1 {
				res++
				break
			}
			//if length == len(word) {
			//	res++
			//	break
			//}
		}
	}
	return res
}

func wordCount1(startWords []string, targetWords []string) int {
	res := 0
	m := make(map[int][]string)
	for _, word := range startWords {
		l := len(word)
		m[l] = append(m[l], word)
	}

	for _, word := range targetWords {
		l := len(word)
		for _, str := range m[l-1] {
			temp := str
			if strings.ContainsAny(word, temp) {
				res++
				break
			}
			//for _, byt := range word {
			//
			//	if strings.Contains(temp, string(byt)) {
			//		continue
			//	} else {
			//		temp += string(byt)
			//	}
			//}
			//if len(temp) == len(word) {
			//	res++
			//	break
			//}
		}
	}
	return res
}

func divideString(s string, k int, fill byte) []string {
	final := make([]string, 0)
	i := len(s) / k
	for j := 0; j < i; j++ {
		s2 := s[j*k : (j+1)*k]
		final = append(final, s2)
	}

	if i*k < len(s) {
		temp := s[i*k:]
		index := k - len(temp)
		for index > 0 {
			temp += string(fill)
			index--
		}
		final = append(final, temp)
	}

	return final
}

func minMoves1(target int, maxDoubles int) int {
	sum := 0
	var travel func(target int, maxDoubles int)
	travel = func(target int, maxDoubles int) {
		if target == 1 {
			return
		}
		if target%2 == 1 {
			sum += 1
			target -= 1
			travel(target, maxDoubles)
		} else if target%2 == 0 && maxDoubles > 0 {
			sum += 1
			target /= 2
			maxDoubles -= 1
			travel(target, maxDoubles)
		} else {
			sum += 1
			target -= 1
			travel(target, maxDoubles)
		}
	}
	travel(target, maxDoubles)
	return sum
}

func minMoves(target int, maxDoubles int) int {
	sum := 0

	for target != 1 {
		if target == 1 {
			break
		}
		if maxDoubles == 0 {
			return sum + target - 1
		}
		if target%2 == 1 {
			sum += 1
			target -= 1
		} else if target%2 == 0 && maxDoubles > 0 {
			sum += 1
			target /= 2
			maxDoubles -= 1
		} else {
			sum += 1
			target -= 1
		}
	}

	return sum
}

func mostPoints1(questions [][]int) int64 {
	var final int64
	var travel func(questions [][]int, index int, b bool, sum int64)
	travel = func(questions [][]int, index int, b bool, sum int64) {
		if index >= len(questions) {
			if sum > final {
				final = sum
			}
			return
		}
		if b {
			ints := questions[index]
			travel(questions, index+1+ints[1], true, sum+int64(ints[0]))
			travel(questions, index+1+ints[1], false, sum+int64(ints[0]))
		} else {
			travel(questions, index+1, true, sum)
			travel(questions, index+1, false, sum)
		}
	}

	travel(questions, 0, true, 0)
	travel(questions, 0, false, 0)
	return final
}

func mostPoints2(questions [][]int) int64 {
	var final int64
	var travel func(questions [][]int, index int, b bool, sum int64)
	travel = func(questions [][]int, index int, b bool, sum int64) {
		if index >= len(questions) {
			if sum > final {
				final = sum
			}
			return
		}
		ints := questions[index]

		if b {
			if ints[0] < ints[1] && index+ints[1] < len(questions) {
				travel(questions, index+1, true, sum)
				travel(questions, index+1, false, sum)
			} else {
				travel(questions, index+1+ints[1], true, sum+int64(ints[0]))
				travel(questions, index+1+ints[1], false, sum+int64(ints[0]))
			}
		} else {
			travel(questions, index+1, true, sum)
			travel(questions, index+1, false, sum)
		}
	}

	travel(questions, 0, true, 0)
	travel(questions, 0, false, 0)
	return final
}

func mostPoints3(questions [][]int) int64 {
	var final int64
	var travel func(questions [][]int, index int, sum int64)
	travel = func(questions [][]int, index int, sum int64) {
		if index >= len(questions) {
			if sum > final {
				final = sum
			}
			return
		}
		ints := questions[index]
		if ints[0] < ints[1] && index+ints[1] < len(questions) {
			travel(questions, index+1, sum)
		} else {
			travel(questions, index+1+ints[1], sum+int64(ints[0]))
			travel(questions, index+1, sum)
		}

	}

	travel(questions, 0, 0)
	return final
}

func mostPoints(questions [][]int) int64 {
	var final int64
	var travel func(index int, sum int64)
	travel = func(index int, sum int64) {
		if index >= len(questions) {
			if sum > final {
				final = sum
			}
			return
		}
		ints := questions[index]
		if ints[0] < ints[1] && index+ints[1] < len(questions) {
			travel(index+1, sum)
		} else {
			travel(index+1+ints[1], sum+int64(ints[0]))
			travel(index+1, sum)
		}

	}

	travel(0, 0)
	return final
}
