package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	//obj := Constructor()
	//fmt.Println(obj.Insert(0))
	//fmt.Println(obj.Insert(1))
	//fmt.Println(obj.Remove(0))
	//fmt.Println(obj.Insert(2))
	//fmt.Println(obj.GetRandom())
	//fmt.Println(obj.Remove(1))
	//fmt.Println(obj.Insert(2))
	//fmt.Println(obj.GetRandom())

	//lruCache := Constructor(2)
	//lruCache.Put(1, 1)
	//lruCache.Put(2, 2)
	//fmt.Println(lruCache.Get(1))
	//lruCache.Put(3, 3)
	//fmt.Println(lruCache.Get(2))
	//lruCache.Put(4, 4)
	//fmt.Println(lruCache.Get(1))
	//fmt.Println(lruCache.Get(3))
	//fmt.Println(lruCache.Get(4))

	//lruCache := Constructor(3)
	//lruCache.Put(1, 1)
	//lruCache.Put(2, 2)
	//lruCache.Put(3, 3)
	//lruCache.Put(4, 4)
	//fmt.Println(lruCache.Get(4))
	//fmt.Println(lruCache.Get(3))
	//fmt.Println(lruCache.Get(2))
	//fmt.Println(lruCache.Get(1))
	//lruCache.Put(5, 5)
	//fmt.Println(lruCache.Get(1))
	//fmt.Println(lruCache.Get(2))
	//fmt.Println(lruCache.Get(3))
	//fmt.Println(lruCache.Get(4))
	//fmt.Println(lruCache.Get(5))

	//lruCache := Constructor(1)
	//lruCache.Put(2, 1)
	//lruCache.Get(2)
	//lruCache.Put(3, 2)
	//lruCache.Get(2)
	//lruCache.Get(3)

	//anagram := isAnagram("anagram", "nagaram")
	//fmt.Println(anagram)

	//anagrams := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	//fmt.Println(anagrams)

	//anagrams := isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")
	//fmt.Println(anagrams)

	//difference := findMinDifference([]string{"00:00", "23:59", "00:00"})
	//fmt.Println(difference)

	//rpn := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	//fmt.Println(rpn)

	//collision := asteroidCollision([]int{8, -8})
	//fmt.Println(collision)

	//root := Constructor(&TreeNode{
	//	Val:   1,
	//	Left:  nil,
	//	Right: nil,
	//})
	//
	//insert := root.Insert(2)
	//fmt.Println(insert)
	//
	//get_root := root.Get_root()
	//fmt.Println(get_root)

	//values := rightSideView(&TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   5,
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
	//		Val:  2,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   9,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//})
	//fmt.Println(values)

	//tree := pruneTree(&TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val: 0,
	//		Left: &TreeNode{
	//			Val:   0,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//})
	//fmt.Println(tree)

	//numbers := sumNumbers(&TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 9,
	//		Left: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//})
	//fmt.Println(numbers)

	//sum := pathSum(&TreeNode{
	//	Val: 0,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}, 1)
	//
	//fmt.Println(sum)

	//bst := increasingBST(&TreeNode{
	//	Val: 5,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val: 2,
	//			Left: &TreeNode{
	//				Val:   1,
	//				Left:  nil,
	//				Right: nil,
	//			},
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
	//			Val:  8,
	//			Left: &TreeNode{Val: 7},
	//			Right: &TreeNode{
	//				Val:   9,
	//				Left:  nil,
	//				Right: nil,
	//			},
	//		},
	//	},
	//})
	//fmt.Println(bst)

	//stairs := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	//fmt.Println(stairs)

	//duplicate := containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3)
	//fmt.Println(duplicate)

	//constructor := Constructor(3, []int{4, 5, 8, 2})
	//add := constructor.Add(3)
	//fmt.Println(add)
	//i := constructor.Add(5)
	//fmt.Println(i)
	//i2 := constructor.Add(10)
	//fmt.Println(i2)
	//i3 := constructor.Add(9)
	//fmt.Println(i3)
	//i4 := constructor.Add(4)
	//fmt.Println(i4)

	//frequent := topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2)
	//fmt.Println(frequent)

	//pairs := kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 10)
	//fmt.Println(pairs)

	//constructor := Constructor()
	//constructor.Insert("apple")
	//fmt.Println(constructor.Search("apple"))
	//fmt.Println(constructor.Search("app"))
	//fmt.Println(constructor.StartsWith("app"))
	//constructor.Insert("app")
	//fmt.Println(constructor.Search("app"))
	//fmt.Println(constructor.Search("apple"))

	//words := replaceWords([]string{"cat", "bat", "rat"}, "the ca was rattled by the battery")
	//fmt.Println(words)

	//constructor := Constructor()
	//constructor.BuildDict([]string{"hello", "leetcode"})
	//search := constructor.Search("hello")
	//fmt.Println(search)
	//b := constructor.Search("hhllo")
	//fmt.Println(b)
	//b2 := constructor.Search("hell")
	//fmt.Println(b2)
	//b3 := constructor.Search("leetcoded")
	//fmt.Println(b3)

	//encoding := minimumLengthEncoding([]string{"qaa", "a", "aa"})
	//encoding := minimumLengthEncoding([]string{"me", "time", "bell"})
	//fmt.Println(encoding)

	//constructor := Constructor()
	//constructor.Insert("apple", 3)
	//sum := constructor.Sum("ap")
	//fmt.Println(sum)
	//constructor.Insert("app", 2)
	//constructor.Insert("apple", 2)
	//sum = constructor.Sum("ap")
	//fmt.Println(sum)

	//xor := findMaximumXOR([]int{2, 8, 10})
	//fmt.Println(xor)

	//insert := searchInsert([]int{1}, 0)
	//fmt.Println(insert)

	//array := peakIndexInMountainArray([]int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19})
	//fmt.Println(array)

	//duplicate := singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	//fmt.Println(duplicate)

	//constructor := Constructor([]int{1})
	//fmt.Println(constructor.PickIndex())
	//fmt.Println(constructor.PickIndex())
	//fmt.Println(constructor.PickIndex())
	//fmt.Println(constructor.PickIndex())

	//i := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	//fmt.Println(i)

	//array := relativeSortArray([]int{2, 21, 43, 38, 0, 42, 33, 7, 24, 13, 12, 27, 12, 24, 5, 23, 29, 48, 30, 31}, []int{2, 42, 38, 0, 43, 21})
	//fmt.Println(array)

	//largest := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	//fmt.Println(largest)

	//s := convert("AB", 1)
	//fmt.Println(s)

	//area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	//fmt.Println(area)

	//duplicate := containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	//fmt.Println(duplicate)

	//order := levelOrder(nil)
	//fmt.Println(order)

	//i := search([]int{-1, 0, 3, 5, 9, 12}, 2)
	//fmt.Println(i)

	//unique := isUnique("abc")
	//fmt.Println(unique)

	//sum := threeSum([]int{1, -1, -1, 0})
	//fmt.Println(sum)

	//profit := maxProfit([]int{3, 2, 6, 5, 0, 3})
	//fmt.Println(profit)

	//closest := threeSumClosest([]int{1, 1, 1, 1}, -100)
	//fmt.Println(closest)

	//squares := sortedSquares([]int{-7, -3, 2, 3, 11})
	//fmt.Println(squares)

	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(nums, 3)

	//subarrays := numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2)
	//fmt.Println(subarrays)

	//i := merge([][]int{{1, 4}, {4, 5}})
	//fmt.Println(i)

	//moveZeroes([]int{1, 3, 0, 12})
	//moveZeroes([]int{0, 0, 1})

	//sum3 := twoSum3([]int{3, 24, 50, 79, 88, 150, 345}, 200)
	//fmt.Println(sum3)

	//words := reverseWords("Let's take LeetCode contest")
	//fmt.Println(words)

	//array := sortArray([]int{5, 1, 1, 2, 0, 0})
	//fmt.Println(array)

	//removeNthFromEnd(&ListNode{
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

	//radius := findRadius([]int{282475249, 622650073, 984943658, 144108930, 470211272, 101027544, 457850878, 458777923}, []int{823564440, 115438165, 784484492, 74243042, 114807987, 137522503, 441282327, 16531729, 823378840, 143542612})
	//fmt.Println(radius)

	//substring := lengthOfLongestSubstring("pwwkew")
	//fmt.Println(substring)

	//inclusion := checkInclusion("ab", "eidbaoo")
	//fmt.Println(inclusion)

	//order := findDiagonalOrder([][]int{{1}, {3}, {4}, {5}, {6}})
	//fmt.Println(order)

	prefix := longestCommonPrefix([]string{"ab", "a"})
	fmt.Println(prefix)
}

//type RandomizedSet struct {
//	//hash链表
//	m map[int]int
//	l []int
//}
//
///** Initialize your data structure here. */
//func Constructor() RandomizedSet {
//	return RandomizedSet{
//		m: map[int]int{},
//		l: make([]int, 0),
//	}
//}
//
///** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
//func (this *RandomizedSet) Insert(val int) bool {
//	if _, ok := this.m[val]; !ok {
//		this.l = append(this.l, val)
//		this.m[val] = len(this.l) - 1
//		return true
//	}
//	return false
//}
//
////可以通过把下标里的值 换到最后 然后再删除最后一位 就可以了!!!!
///** Removes a value from the set. Returns true if the set contained the specified element. */
//func (this *RandomizedSet) Remove(val int) bool {
//	if i, ok := this.m[val]; ok {
//		//不仅要把切片调换 还要把map里的下标调换
//		temp := this.l[len(this.l)-1]
//		this.l[len(this.l)-1] = val
//		this.l[i] = temp
//		this.l = this.l[:len(this.l)-1]
//
//		//调换map
//		this.m[temp] = i
//		delete(this.m, val)
//		return true
//	}
//	return false
//}
//
///** Get a random element from the set. */
//func (this *RandomizedSet) GetRandom() int {
//	return this.l[rand.Intn(len(this.l))]
//}

// LRUCache 重要的地方在于 删除在头部删除 新增在尾部新增
type LRUCache struct {
	capacity int
	len      int
	head     *Node
	tail     *Node
	m        map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

//func Constructor(capacity int) LRUCache {
//	return LRUCache{
//		capacity: capacity,
//		len:      0,
//		m:        make(map[int]*Node),
//	}
//}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {

		//如果在尾节点 直接返回
		if node == this.tail {
			return node.value
		} else if node == this.head {
			//如果在头结点 头节点向下移动
			this.head.next.pre = nil
			this.head = this.head.next
		} else {
			//删除节点
			node.pre.next = node.next
			node.next.pre = node.pre
		}

		//接到尾节点
		this.tail.next = node
		node.pre = this.tail

		//换成尾节点
		this.tail = node

		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value = value

		//如果在尾节点 直接返回
		if node == this.tail {
			return
		} else if node == this.head {
			//如果在头结点 头节点向下移动
			this.head.next.pre = nil
			this.head = this.head.next
		} else {
			//删除节点
			node.pre.next = node.next
			node.next.pre = node.pre
		}

		//接到尾节点
		this.tail.next = node
		node.pre = this.tail

		//换成尾节点
		this.tail = node
	} else if this.len < this.capacity {
		this.len++

		newnode := &Node{
			key:   key,
			value: value,
			pre:   nil,
			next:  nil,
		}

		//接到尾节点
		if this.tail == nil {
			this.tail = newnode
			this.head = newnode
		} else {
			this.tail.next = newnode
			newnode.pre = this.tail
		}

		//换成尾节点
		this.tail = newnode

		this.m[key] = this.tail
	} else {
		newnode := &Node{
			key:   key,
			value: value,
			pre:   nil,
			next:  nil,
		}
		//接到尾节点
		this.tail.next = newnode
		newnode.pre = this.tail

		//换成尾节点
		this.tail = newnode

		delete(this.m, this.head.key)
		this.head = this.head.next

		this.m[key] = this.tail
	}
}

//直接弄空间出来比较
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})
	bytes := []rune(t)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})

	return string(bytes) == string(runes)
}

//重点在于利用26个单位的数组进行相等比较
func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	al := [26]int{}
	for i := 0; i < len(s); i++ {
		al[s[i]-'a']++
		al[t[i]-'a']--
	}
	return al == [26]int{}
}

//给每一个str都弄一个[26]int的数组 然后比较数组一不一样
func groupAnagrams1(strs []string) [][]string {
	res := make([][]string, 0)
	alstrs := make([][26]int, len(strs))
	for i := 0; i < len(strs); i++ {
		temp := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			temp[strs[i][j]-'a']++
		}
		alstrs[i] = temp
	}

	m := make(map[int]bool, len(strs))

	for i := 0; i < len(alstrs); i++ {
		if m[i] {
			continue
		}
		m[i] = true
		result := make([]string, 0)
		result = append(result, strs[i])
		for j := i + 1; j < len(alstrs); j++ {
			if m[j] {
				continue
			}
			if alstrs[i] == alstrs[j] {
				m[j] = true
				result = append(result, strs[j])
			}
		}
		res = append(res, result)
	}
	return res
}

//利用排序 选择一个通用的参考系 例如 都转换程a-z排序的 然后存到map中
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		s := []byte(strs[i])
		sort.Slice(s, func(i, j int) bool {
			return s[i] > s[j]
		})
		m[string(s)] = append(m[string(s)], strs[i])
	}

	for _, strings := range m {
		res = append(res, strings)
	}
	return res
}

func isAlienSorted1(words []string, order string) bool {
	m := make(map[uint8]int)
	index := 0
	for i := 0; i < len(order); i++ {
		m[order[i]] = index
		index++
	}
	b := true

	sort.Slice(words, func(i, j int) bool {
		li := len(words[i])
		lj := len(words[j])

		min := li
		if li > lj {
			min = lj
		}

		for k := 0; k < min; k++ {
			if m[words[i][k]] > m[words[j][k]] {
				return false
			} else if m[words[i][k]] < m[words[j][k]] {
				b = false
				return true
			} else {
				continue
			}
		}

		b = li >= lj
		return li < lj
	})
	return b
}

func isAlienSorted(words []string, order string) bool {
	m := make(map[uint8]int)
	for i := 0; i < len(order); i++ {
		m[order[i]] = i
	}

	b := true
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			li := len(words[i])
			lj := len(words[j])
			min := li
			if li > lj {
				min = lj
			}
			for k := 0; k < min; k++ {
				if words[i][k] != words[j][k] {
					b = m[words[i][k]] < m[words[j][k]]
					break
				}
				if k == min-1 {
					b = li <= lj
				}
			}
			if !b {
				return b
			}
		}
	}
	return true
}

//注意数组的尾巴部分和头部 还要再对比一次
func findMinDifference(timePoints []string) int {
	slice := make([]int, len(timePoints))
	for i, point := range timePoints {
		split := strings.Split(point, ":")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		slice[i] = a*60 + b
	}
	sort.Ints(slice)

	min := math.MaxInt64
	left := 0
	for left < len(slice) {
		right := left + 1
		if right < len(slice) {
			temp := slice[right] - slice[left]
			temp1 := 1440 - (slice[right] - slice[left])
			res := 0
			if temp < temp1 {
				res = temp
			} else {
				res = temp1
			}
			if res < min {
				min = res
			}
		}
		left++
	}

	temp := slice[len(slice)-1] - slice[0]
	temp1 := 1440 - (slice[len(slice)-1] - slice[0])
	res := 0
	if temp < temp1 {
		res = temp
	} else {
		res = temp1
	}
	if res < min {
		min = res
	}

	return min

}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			atoi, _ := strconv.Atoi(tokens[i])
			stack = append(stack, atoi)
		} else {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			res := 0
			if tokens[i] == "+" {
				res = a + b
			} else if tokens[i] == "-" {
				res = a - b
			} else if tokens[i] == "*" {
				res = a * b
			} else {
				res = a / b
			}
			stack = append(stack, res)
		}
	}
	return stack[0]
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	var travel func(x int)
	travel = func(x int) {
		//如果是栈为空 或者 栈中为负 什么都可以添加
		if len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, x)
		} else {
			//如果栈中是正 正的可以随便添加 但是负的就需要比较
			if x > 0 {
				stack = append(stack, x)
			} else {
				//如果 栈中正 遍历到的是负的
				if stack[len(stack)-1]+x > 0 {
					return
				} else if stack[len(stack)-1]+x == 0 {
					stack = stack[:len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
					travel(x)
				}
			}
		}
	}
	for i := 0; i < len(asteroids); i++ {
		travel(asteroids[i])
	}
	return stack

}

type MovingAverage struct {
	preSum []int
	size   int
	sum    int
}

///** Initialize your data structure here. */
//func Constructor(size int) MovingAverage {
//	return MovingAverage{
//		preSum: make([]int, size),
//		size:   size,
//		sum:    0,
//	}
//}

func (this *MovingAverage) Next(val int) float64 {
	//如果没有超过容量 就加到数组里 然后更新一下sum的值 减小一下size尺寸
	if this.size > 0 {
		this.preSum[len(this.preSum)-this.size] = val
		this.size--
		this.sum += val
		return float64(this.sum) / float64(len(this.preSum)-this.size)
	} else {
		//如果超过容量了 就更新一下数组 以及sum的值
		this.sum -= this.preSum[0]
		this.sum += val
		this.preSum = append(this.preSum[1:], val)
		return float64(this.sum) / float64(len(this.preSum))
	}

}

type RecentCounter struct {
	//记录最远一次pin的时间
	FurtherestPingTimeIndex int
	//记录总的pin的时间数组
	totalPintTime []int
}

//
//func Constructor() RecentCounter {
//	return RecentCounter{
//		FurtherestPingTimeIndex: 0,
//		totalPintTime:           []int{},
//	}
//}

func (this *RecentCounter) Ping(t int) int {
	//更新总的pin的时间数组
	this.totalPintTime = append(this.totalPintTime, t)
	//从最远pin的时间开始尝试
	left := this.FurtherestPingTimeIndex
	for {
		if this.totalPintTime[left] < t-3000 {
			//如果最远的不在范围内 就尝试下一个
			left++
		} else {
			break
		}
	}
	count := len(this.totalPintTime) - left
	this.FurtherestPingTimeIndex = left
	return count
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	CBTree *TreeNode
	//队列的头用来存储待插入的树节点
	queue []*TreeNode
}

//func Constructor(root *TreeNode) CBTInserter {
//	queue := make([]*TreeNode, 0)
//	queue = append(queue, root)
//
//	如果左右节点都有 则加入队列 把当前节点去掉
//for queue[0].Left != nil && queue[0].Right != nil {
//	queue = append(queue, queue[0].Left)
//	queue = append(queue, queue[0].Right)
//	queue = queue[1:]
//}
//
//return CBTInserter{
//	CBTree: root,
//	queue:  queue,
//}
//}

func (this *CBTInserter) Insert(v int) int {
	newNode := &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
	res := 0
	if this.queue[0].Left == nil {
		this.queue[0].Left = newNode
		res = this.queue[0].Val
	} else if this.queue[0].Right == nil {
		this.queue[0].Right = newNode
		res = this.queue[0].Val
	}

	//插入完再判断一下队列头是不是已经完全了 如果完全就把左右子节点加入 把当前节点去掉
	if this.queue[0].Left != nil && this.queue[0].Right != nil {
		this.queue = append(this.queue, this.queue[0].Left)
		this.queue = append(this.queue, this.queue[0].Right)
		this.queue = this.queue[1:]
	}
	return res
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.CBTree
}

func largestValues(root *TreeNode) []int {
	//特殊情况
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	//提前计算好当前队列长度
	l := len(queue)
	for l != 0 {
		//对长度中的进行遍历 添加到队列中
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

		}
		//但只对l长度中的进行排序 l代表某层的总的
		sort.Slice(queue[:l], func(i, j int) bool {
			return queue[i].Val > queue[j].Val
		})
		//把排序得到的最大值拿出来 添加进结果集
		res = append(res, queue[0].Val)
		//同时从队列中 去掉l层的值
		queue = queue[l:]
		//重新初始化l+1层的 队列长度
		l = len(queue)
	}

	return res
}

func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := 0
	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		//最左边就是第一个
		res = queue[0].Val
		//去掉l层的所有
		queue = queue[l:]
	}
	return res
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		//添加最尾巴的节点就是l层的最右侧的节点了
		res = append(res, queue[l-1].Val)
		//从队列中剔除掉上一层的所有节点
		queue = queue[l:]
	}
	return res
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)

	//如果节点的值是0 并且左右树都是空的 那就返回nil
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func sumNumbers(root *TreeNode) int {
	sum := 0

	var travel func(root *TreeNode, preSum int)
	travel = func(root *TreeNode, preSum int) {
		//递归的结束条件是 左右节点都为空 也就是叶子节点 然后把传过来的值*10 加上本身的值就构成了和
		if root.Left == nil && root.Right == nil {
			sum += preSum*10 + root.Val
			return
		}
		if root.Left != nil && root.Right != nil {
			travel(root.Left, preSum*10+root.Val)
			travel(root.Right, preSum*10+root.Val)
		} else if root.Left != nil {
			travel(root.Left, preSum*10+root.Val)
		} else if root.Right != nil {
			travel(root.Right, preSum*10+root.Val)
		}
		return
	}

	travel(root, 0)
	return sum
}

//前缀和思想
func pathSum1(root *TreeNode, targetSum int) int {
	res := 0
	preSum := make([]int, 0)
	preSum = append(preSum, 0)

	var travel func(root *TreeNode, tempSum int)
	travel = func(root *TreeNode, tempSum int) {
		if root == nil {
			return
		}
		tempSum += root.Val
		for i := 0; i < len(preSum); i++ {
			if targetSum == tempSum-preSum[i] {
				res++
			}
		}

		preSum = append(preSum, tempSum)
		travel(root.Left, tempSum)
		travel(root.Right, tempSum)
		preSum = preSum[:len(preSum)-1]
	}

	travel(root, 0)

	return res
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	m := make(map[int]int)
	m[0] = 1

	var travel func(root *TreeNode, tempSum int)
	travel = func(root *TreeNode, tempSum int) {
		if root == nil {
			return
		}
		tempSum += root.Val
		res += m[tempSum-targetSum]
		m[tempSum]++
		travel(root.Left, tempSum)
		travel(root.Right, tempSum)
		m[tempSum]--
	}

	travel(root, 0)

	return res
}

//递归 使用公共变量来记录待操作的节点
func increasingBST(root *TreeNode) *TreeNode {
	//初始化上一个待被添加节点
	lastNode := &TreeNode{}
	//res是lastNode的右子树
	res := lastNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}

		//中序遍历
		travel(root.Left)
		//把root左子树置为空
		root.Left = nil
		//把上一个待被添加的节点的右子树赋上节点
		lastNode.Right = root
		//更新下一个待被添加的节点
		lastNode = root
		travel(root.Right)
	}

	travel(root)

	return res.Right
}

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)

	var travel func(root *TreeNode) bool
	travel = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		if m[k-root.Val] {
			return true
		}
		m[root.Val] = true
		return travel(root.Left) || travel(root.Right)
	}

	return travel(root)
}

func minCostClimbingStairs(cost []int) int {
	m := make(map[int]int)

	m[len(cost)-1] = cost[len(cost)-1]
	m[len(cost)-2] = cost[len(cost)-2]
	for i := len(cost) - 3; i >= 0; i-- {
		if m[i+1] < m[i+2] {
			m[i] = m[i+1] + cost[i]
		} else {
			m[i] = m[i+2] + cost[i]
		}
	}

	if m[0] < m[1] {
		return m[0]
	}
	return m[1]
}

//中序遍历  中序遍历代码在中间位置 如果有时候分不清 记得移动移动代码书写的位置试一试
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//此标识记录是否找到目标节点 找到就开始干活了
	b := false
	var res *TreeNode
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		//中序遍历
		travel(root.Left)
		//如果找到了就返回
		if res != nil {
			return
		}
		//res为空并且标识为真的时候代表这时候的节点就是后继节点
		if b && res == nil {
			res = root
			return
		}

		//标志位更新
		if root == p {
			b = true
		}
		travel(root.Right)
	}
	travel(root)
	return res
}

func convertBST1(root *TreeNode) *TreeNode {
	preSum := make([]int, 0)
	preSum = append(preSum, 0)
	nodeSlice := make([]*TreeNode, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		nodeSlice = append(nodeSlice, root)
		preSum = append(preSum, root.Val+preSum[len(preSum)-1])
		travel(root.Right)
	}
	travel(root)

	for i := 0; i < len(nodeSlice); i++ {
		nodeSlice[i].Val = preSum[len(nodeSlice)] - preSum[i]
	}

	return root
}

func convertBST(root *TreeNode) *TreeNode {
	appendSum := 0
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Right)
		appendSum += root.Val
		root.Val = appendSum
		travel(root.Left)
	}
	travel(root)

	return root
}

type BSTIterator struct {
	root         *TreeNode
	lastIterator []*TreeNode
}

//func Constructor(root *TreeNode) BSTIterator {
//	var nodes []*TreeNode
//	var travel func(root *TreeNode)
//	travel = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//
//		travel(root.Left)
//		nodes = append(nodes, root)
//		travel(root.Right)
//	}
//	travel(root)
//	return BSTIterator{
//		root:         root,
//		lastIterator: nodes,
//	}
//}

func (this *BSTIterator) Next() int {
	res := this.lastIterator[0].Val
	this.lastIterator = this.lastIterator[1:]
	return res
}

func (this *BSTIterator) HasNext() bool {
	if len(this.lastIterator) > 0 {
		return true
	}
	return false
}

func containsNearbyAlmostDuplicate1(nums []int, k int, t int) bool {
	if k == 0 || len(nums) < 2 {
		return false
	}
	IsMatched := false
	var abs func(a int) int
	abs = func(a int) int {
		return int(math.Abs(float64(a)))
	}
	left := 0
	for left < len(nums)-1 {
		right := left + 1
		for right-left <= k && right < len(nums) {
			if abs(nums[right]-nums[left]) <= t {
				IsMatched = true
				break
			}
			right++
		}
		if IsMatched {
			return IsMatched
		}
		left++
	}
	return false
}

type myHeap []int

//比大小
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

//交换
func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

//取长度
func (h *myHeap) Len() int {
	return len(*h)
}

//弹出
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

//压入
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

type node struct {
	ch       [2]*node
	priority int
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val}
	}
	d := o.cmp(val)
	o.ch[d] = t._put(o.ch[d], val)
	if o.ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

func (t *treap) _delete(o *node, val int) *node {
	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._delete(o.ch[d], val)
		return o
	}
	if o.ch[1] == nil {
		return o.ch[0]
	}
	if o.ch[0] == nil {
		return o.ch[1]
	}
	d := 0
	if o.ch[0].priority > o.ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.ch[d] = t._delete(o.ch[d], val)
	return o
}

func (t *treap) delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

//自定义有序集合 人都麻了
func containsNearbyAlmostDuplicate3(nums []int, k, t int) bool {
	set := &treap{}
	for i, v := range nums {
		if lb := set.lowerBound(v - t); lb != nil && lb.val <= v+t {
			return true
		}
		set.put(v)
		if i >= k {
			set.delete(nums[i-k])
		}
	}
	return false
}

func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

//桶排序 人都麻了
func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	if k == 0 || len(nums) < 2 {
		return false
	}
	window := make([]int, 0)
	var abs func(a int) int
	abs = func(a int) int {
		return int(math.Abs(float64(a)))
	}

	left := 0
	for left < len(nums)-1 {
		right := left + 1
		for right-left <= k && right < len(nums) {
			window = append(window, nums[right])
			right++
		}
		sort.Ints(window)
		search := sort.Search(len(window)-1, func(i int) bool {
			if window[i] >= nums[left]-t {
				return true
			}
			return false
		})
		if abs(window[search]-nums[left]) <= t {
			return true
		}
		window = []int{}
		left++
	}
	return false
}

type MyCalendar struct {
	dateList [][2]int
}

//func Constructor() MyCalendar {
//	return MyCalendar{dateList: [][2]int{}}
//}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.dateList); i++ {
		if (start >= this.dateList[i][0] && start < this.dateList[i][1]) || (end > this.dateList[i][0] && end <= this.dateList[i][1]) || (start <= this.dateList[i][0] && end >= this.dateList[i][1]) {
			return false
		}
	}
	this.dateList = append(this.dateList, [2]int{start, end})
	return true
}

type myheap []int

func (m *myheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *myheap) Pop() (v interface{}) {
	*m, v = (*m)[:m.Len()-1], (*m)[m.Len()-1]
	return
}

func (m *myheap) Len() int {
	return len(*m)
}

func (m *myheap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap) Less(i, j int) bool {
	return (*m)[i] > (*m)[j]
}

type myNode struct {
	v    int
	next *myNode
	pre  *myNode
}

//type KthLargest struct {
//	k int
//	sort.IntSlice
//}
//
//func Constructor(k int, nums []int) KthLargest {
//	s := sort.IntSlice(nums)
//
//	return KthLargest{
//		k:        k,
//		IntSlice: s,
//	}
//}
//
//func (kl *KthLargest) Push(v interface{}) {
//	kl.IntSlice = append(kl.IntSlice, v.(int))
//}
//
//func (kl *KthLargest) Pop() interface{} {
//	a := kl.IntSlice
//	v := a[len(a)-1]
//	kl.IntSlice = a[:len(a)-1]
//	return v
//}
//
//func (this *KthLargest) Add(val int) int {
//	this.Push(val)
//	if len(this.IntSlice) >= this.k {
//		this.Pop()
//	}
//	return this.IntSlice[0]
//}

// KthLargest 使用到了堆 但是使用接口的时候需要这么用 heap.Push() 而不是 结构体.Push()
type KthLargest struct {
	sort.IntSlice
	k int
}

// Constructor 这题给整蒙了鸭
//func Constructor(k int, nums []int) KthLargest {
//	kl := KthLargest{k: k}
//	for _, val := range nums {
//		kl.Add(val)
//	}
//	return kl
//}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

type mheap struct {
	m [][2]int
}

func (m *mheap) Push(x interface{}) {
	(*m).m = append((*m).m, x.([2]int))
}

func (m *mheap) Pop() (v interface{}) {
	(*m).m, v = (*m).m[:m.Len()-1], (*m).m[m.Len()-1]
	return
}

func (m *mheap) Len() int {
	return len((*m).m)
}

func (m *mheap) Swap(i, j int) {
	(*m).m[i], (*m).m[j] = (*m).m[j], (*m).m[i]
}

func (m *mheap) Less(i, j int) bool {
	return (*m).m[i][1] < (*m).m[j][1]
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int, k)
	m := new(mheap)
	ma := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		ma[nums[i]]++
	}

	for key, value := range ma {
		heap.Push(m, [2]int{key, value})
		if len(m.m) > k {
			heap.Pop(m)
		}
	}

	for i := 0; i < k; i++ {
		res[i] = heap.Pop(m).([2]int)[0]
	}
	return res
}

//看不懂 艹
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

// Trie 字典树 想法学一下 实际上就是treenode 只不过有26个分支
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{
//		children: [26]*Trie{},
//		isEnd:    false,
//	}
//}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, v := range word {
		if cur.children[v-'a'] == nil {
			cur.children[v-'a'] = &Trie{
				children: [26]*Trie{},
				isEnd:    false,
			}
			cur = cur.children[v-'a']
		} else {
			cur = cur.children[v-'a']
		}
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, v := range word {
		if cur.children[v-'a'] == nil {
			return false
		} else {
			cur = cur.children[v-'a']
		}
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, v := range prefix {
		if cur.children[v-'a'] == nil {
			return false
		} else {
			cur = cur.children[v-'a']
		}
	}
	return true
}

func replaceWords(dictionary []string, sentence string) string {
	trie := &Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
	for i := 0; i < len(dictionary); i++ {
		cur := trie
		for _, v := range dictionary[i] {
			if cur.children[v-'a'] == nil {
				cur.children[v-'a'] = &Trie{
					children: [26]*Trie{},
					isEnd:    false,
				}
			}
			cur = cur.children[v-'a']
		}
		cur.isEnd = true
	}

	temp := make([]string, 0)
	split := strings.Split(sentence, " ")
	for _, s := range split {
		cur := trie
		for i, v := range s {
			if cur.children[v-'a'] == nil {
				temp = append(temp, s)
				break
			} else {
				cur = cur.children[v-'a']
				if cur.isEnd {
					temp = append(temp, s[:i+1])
					break
				} else if i == len(s)-1 {
					temp = append(temp, s)
					break
				}
			}
		}
	}

	res := strings.Join(temp, " ")
	return res
}

//type MagicDictionary1 struct {
//	m map[string]bool
//}
//
///** Initialize your data structure here. */
//func Constructor1() MagicDictionary {
//	return MagicDictionary{m: map[string]bool{}}
//}
//
//func (this *MagicDictionary) BuildDict1(dictionary []string) {
//	for i := 0; i < len(dictionary); i++ {
//		bytes := []byte(dictionary[i])
//		for j := 0; j < len(bytes); j++ {
//			old := bytes[j]
//			for k := 0; k < 26; k++ {
//				if byte(k+'a') == old {
//					continue
//				}
//				bytes[j] = byte('a' + k)
//				this.m[string(bytes)] = true
//			}
//			bytes[j] = old
//		}
//	}
//}
//
//func (this *MagicDictionary) Search1(searchWord string) bool {
//	return this.m[searchWord]
//}

//type MagicDictionary struct {
//	m map[int]map[string]bool
//}
//
///** Initialize your data structure here. */
//func Constructor() MagicDictionary {
//	return MagicDictionary{map[int]map[string]bool{}}
//}
//
//func (this *MagicDictionary) BuildDict(dictionary []string) {
//	for _, word := range dictionary {
//		mString := this.m[len(word)]
//		if mString != nil {
//			mString[word] = true
//		} else {
//			this.m[len(word)] = map[string]bool{word: true}
//		}
//	}
//}
//
//func (this *MagicDictionary) Search(searchWord string) bool {
//	if v, ok := this.m[len(searchWord)]; ok {
//		b := false
//		for i := range searchWord {
//			b = isExist(searchWord, i, v)
//			if b {
//				return true
//			}
//		}
//	}
//	return false
//}
//
//func isExist(word string, i int, m map[string]bool) bool {
//	old := word[i]
//	bytes := []byte(word)
//	for j := 0; j < 26; j++ {
//		if byte('a'+j) == old {
//			continue
//		}
//		bytes[i] = byte('a' + j)
//		if m[string(bytes)] {
//			return true
//		}
//	}
//	return false
//}

type MagicDictionary struct {
	mOld  map[string]bool
	mPair map[string]bool
}

/** Initialize your data structure here. */
//func Constructor() MagicDictionary {
//	return MagicDictionary{
//		mOld:  map[string]bool{},
//		mPair: map[string]bool{},
//	}
//}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.mOld[word] = true
		for i := range word {
			temp := []byte(word)
			temp1 := append(temp[:i], temp[i+1:]...)
			this.mPair[string(temp1)] = true
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if _, ok := this.mOld[searchWord]; ok {
		return false
	} else {
		temp := []byte(searchWord)
		res := false
		for i := range searchWord {
			temp1 := append(temp[:i], temp[i+1:]...)
			res = this.mPair[string(temp1)]
			if res {
				return true
			}
		}
	}
	return false
}

//思路有点糟糕蛤  使用后缀字典树存储 并使用map存储当前的单词  最后遍历map中的单词得到结果
func minimumLengthEncoding1(words []string) int {
	t := &Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
	m := make(map[string]bool)
	for _, word := range words {
		cur := t
		b := false
		for i := len(word) - 1; i >= 0; i-- {
			if cur.children[word[i]-'a'] == nil {
				delete(m, word[i+1:])
				b = true
				cur.children[word[i]-'a'] = &Trie{
					children: [26]*Trie{},
					isEnd:    false,
				}
			}
			cur.isEnd = false
			cur = cur.children[word[i]-'a']
		}
		cur.isEnd = true
		if b {
			m[word] = true
		}
	}

	count := 0
	for key := range m {
		count += len(key) + 1
	}
	return count
}

func minimumLengthEncoding2(words []string) int {
	t := &Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
	for _, word := range words {
		cur := t
		for i := len(word) - 1; i >= 0; i-- {
			if cur.children[word[i]-'a'] == nil {
				cur.children[word[i]-'a'] = &Trie{
					children: [26]*Trie{},
					isEnd:    false,
				}
			}
			cur = cur.children[word[i]-'a']
		}
		cur.isEnd = true
	}

	var travel func(trie *Trie, depth int) int
	travel = func(trie *Trie, depth int) int {
		if trie == nil {
			return 0
		}

		b := true
		for _, child := range trie.children {
			if child != nil {
				b = false
			}
		}
		if trie.isEnd && b {
			return depth + 1
		}

		count := depth
		for _, child := range trie.children {
			if child != nil {
				count += travel(child, 1)
			}
		}
		return count
	}
	count := 0
	count += travel(t, 0)
	return count
}

func minimumLengthEncoding(words []string) int {
	res := 0
	t := &Trie{
		children: [26]*Trie{},
		isEnd:    false,
	}
	for _, word := range words {
		cur := t
		for i := len(word) - 1; i >= 0; i-- {
			if cur.children[word[i]-'a'] == nil {
				cur.children[word[i]-'a'] = &Trie{
					children: [26]*Trie{},
					isEnd:    false,
				}
			}
			cur = cur.children[word[i]-'a']
		}
		cur.isEnd = true
	}

	var travel func(trie *Trie, depth int)
	travel = func(trie *Trie, depth int) {
		b := false

		for _, child := range trie.children {
			if child != nil {
				b = true
				travel(child, depth+1)
			}
		}
		if !b {
			res += depth + 1
		}
		return
	}
	travel(t, 0)
	return res
}

// MapSum1 m记录所有的前缀相关的map
//morigin 已经加进来过的单词 用于替换
type MapSum1 struct {
	m map[string]int

	mOrigin map[string]int
}

func Constructor1() MapSum1 {
	return MapSum1{m: map[string]int{}, mOrigin: map[string]int{}}
}

func (this *MapSum1) Insert1(key string, val int) {
	//如果这个单词曾被添加过 那就扣掉上次添加的值 加上这次最新的值
	if v, ok := this.mOrigin[key]; ok {
		for i := 0; i < len(key); i++ {
			this.m[key[:i+1]] -= v
			this.m[key[:i+1]] += val
			this.mOrigin[key] = val
		}
	} else {
		for i := 0; i < len(key); i++ {
			this.m[key[:i+1]] += val
			this.mOrigin[key] = val
		}
	}
}

func (this *MapSum1) Sum1(prefix string) int {
	return this.m[prefix]
}

// MapSum 使用字典树也行
type MapSum struct {
	trie    *Trie
	mOrigin map[string]int
}

//func Constructor() MapSum {
//	return MapSum{trie: &Trie{
//		children: [26]*Trie{},
//		isEnd:    false,
//	}, mOrigin: map[string]int{}}
//}

func (this *MapSum) Insert(key string, val int) {
	this.mOrigin[key] = val
	cur := this.trie
	for i := 0; i < len(key); i++ {
		if cur.children[key[i]-'a'] == nil {
			cur.children[key[i]-'a'] = &Trie{
				children: [26]*Trie{},
				isEnd:    false,
			}
		}
		cur = cur.children[key[i]-'a']
	}
	cur.isEnd = true
}

func (this *MapSum) Sum(prefix string) int {
	res := 0

	var travel1 func(prefix string, trie *Trie)
	travel1 = func(prefix string, trie *Trie) {
		if trie == nil {
			return
		}
		if trie.isEnd {
			res += this.mOrigin[prefix]
		}
		for key, child := range trie.children {
			travel1(prefix+fmt.Sprintf("%c", 'a'+key), child)
		}
	}

	cur := this.trie
	for i := 0; i < len(prefix); i++ {
		if cur.children[prefix[i]-'a'] == nil {
			return 0
		}
		cur = cur.children[prefix[i]-'a']
	}

	travel1(prefix, cur)

	return res
}

type Trie1 struct {
	left  *Trie1
	right *Trie1
}

func findMaximumXOR(nums []int) int {
	res := 0
	t := &Trie1{
		left:  nil,
		right: nil,
	}
	k := 30
	for i := 0; i < len(nums); i++ {
		cur := t
		for j := k; j >= 0; j-- {
			temp := nums[i] >> j & 1
			if temp == 0 {
				if cur.left == nil {
					cur.left = &Trie1{
						left:  nil,
						right: nil,
					}
				}
				cur = cur.left
			} else {
				if cur.right == nil {
					cur.right = &Trie1{
						left:  nil,
						right: nil,
					}
				}
				cur = cur.right
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		x := 0
		cur := t
		for j := k; j >= 0; j-- {
			temp := nums[i] >> j & 1
			if temp == 0 {
				if cur.right == nil {
					cur = cur.left
					x = 2 * x
				} else {
					cur = cur.right
					x = 2*x + 1
				}
			} else {
				if cur.left == nil {
					cur = cur.right
					x = 2 * x
				} else {
					cur = cur.left
					x = 2*x + 1
				}
			}
		}
		if res < x {
			res = x
		}
	}

	return res
}

//如何限定right 如果right==nums.len 那么 for中left<right 否则会超出 举个例子 跳出循环肯定是 left==right的时候 那么带个数字进去 也就是 3==3 就会超差 可见左闭右开
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid
		}
	}

	return left
}

//如何限定right 如果right==nums.len-1 那么 for中就需要left<=right 否则就漏掉东西了 跳出循环肯定是 left==right+1的时候 那么带个数字进去 也就是4<3 跳出循环刚刚好 可见左闭右闭
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}

	return left
}

//o(n)时间复杂度
func peakIndexInMountainArray1(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] >= arr[res] {
			res = i
		} else {
			return res
		}
	}
	return res
}

//o(log(n))时间复杂度
func peakIndexInMountainArray(arr []int) int {
	res := 0
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			left = mid + 1
		} else if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
			right = mid
		} else if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return mid
		}
	}
	return res
}

//利用map
func singleNonDuplicate1(nums []int) int {
	res := 0
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] == 2 {
			delete(m, num)
		}
	}
	for k := range m {
		res = k
	}
	return res
}

//下标 相等判断
func singleNonDuplicate2(nums []int) int {
	res := 0
	for i := 0; i < len(nums); {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			i += 2
		} else {
			res = nums[i]
			break
		}
	}
	return res
}

//位运算居然比相等还要快
func singleNonDuplicate3(nums []int) int {
	res := 0
	for i := 0; i < len(nums); {
		if i+1 < len(nums) && nums[i]^nums[i+1] == 0 {
			i += 2
		} else {
			res = nums[i]
			break
		}
	}
	return res
}

//位运算  如果两个数相等 那么异或^的结果就是0  任何数与0异或就是这个数本身 x^0 = x
func singleNonDuplicate4(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

//二分还不太会
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

type Solution struct {
	preSum []int
}

func Constructor(w []int) Solution {
	preSum := make([]int, len(w)+1)
	preSum[0] = 0
	for i := 0; i < len(w); i++ {
		preSum[i+1] = preSum[i] + w[i]
	}
	return Solution{preSum: preSum}
}

func (this *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(this.preSum[len(this.preSum)-1]) + 1
	return sort.SearchInts(this.preSum[1:], i)
}

//不太会啊 数组有序问题 最好先排个序 思路可能会开阔一点
func merge1(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		//如果这个左端点 比 前一个的端点小 说明有重合  有重合还得判断这个的右端点是否会比前一个的右端点大 大的话才会覆盖 小的话顶多算子集
		if intervals[i][0] <= res[len(res)-1][1] {
			if res[len(res)-1][1] < intervals[i][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

//哈希表 暴力解法
func relativeSortArray1(arr1 []int, arr2 []int) []int {
	//这样省内存 不存在底层数组移动问题
	res := make([]int, 0, len(arr1))
	//记录arr1中个数和值
	m := make(map[int]int)
	for _, v := range arr1 {
		m[v]++
	}

	for _, v := range arr2 {
		for i := 0; i < m[v]; i++ {
			res = append(res, v)
		}
		//遍历完删除
		delete(m, v)
	}

	//这样省内存 不存在底层数组移动问题
	temp := make([]int, 0, cap(res)-len(res))
	for k, v := range m {
		for i := 0; i < v; i++ {
			temp = append(temp, k)
		}
	}
	//排序
	sort.Ints(temp)
	res = append(res, temp...)
	return res
}

//使用指针 来遍历
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	left := 0
	for _, v := range arr2 {
		for i := left; i < len(arr1); i++ {
			if arr1[i] == v {
				if left == i {
					left++
					continue
				}
				arr1[i], arr1[left] = arr1[left], arr1[i]
				left++
			}
		}
	}

	//这里很致命 既然从arr1中取部分数组出来比较 那就得带上取出来的index
	sort.Slice(arr1[left:], func(i, j int) bool {
		return arr1[i+left] < arr1[j+left]
	})

	return arr1
}

//自定义排序 怎么比较??? 使用map映射的值来比较大小 进行交换
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	for i, v := range arr2 {
		m[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		a, ok := m[arr1[i]]
		b, ok1 := m[arr1[j]]
		if ok && ok1 {
			return a < b
		} else if ok {
			return true
		} else if !ok && !ok1 {
			return arr1[i] < arr1[j]
		}
		return false
	})

	return arr1
}

//怎么可以使用api呢
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

//冒泡排序 两两交换 直到最后  每次固定最后一个元素  每次固定最大值
func findKthLargest2(nums []int, k int) int {
	for count := 0; count < len(nums); count++ {
		for i := 0; i < len(nums)-1-count; i++ {
			if nums[i] > nums[i+1] {
				nums[i+1], nums[i] = nums[i], nums[i+1]
			}
		}
	}
	return nums[len(nums)-k]
}

//选择排序  选出小的 每次固定第一个元素  每次固定最小值
func findKthLargest3(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		//m 用来记录最小的索引位置
		m := i
		for j := i + 1; j < len(nums); j++ {
			if nums[m] > nums[j] {
				m = j
			}
			//if nums[i] > nums[j] {
			//不像冒泡排序  避免频繁交换
			//nums[i], nums[j] = nums[j], nums[i]
			//}
		}
		nums[i], nums[m] = nums[m], nums[i]
	}
	return nums[len(nums)-k]
}

//插入排序  速度一下子就上来了
func findKthLargest4(nums []int, k int) int {
	res := make([]int, 0, len(nums))
	res = append(res, nums[0])
	for i := 1; i < len(nums); i++ {
		//本来呢 插入排序是遍历的形式来选择插入的位置 但是由于已经排好序了 所以可以用二分法定位
		left, right := 0, len(res)
		pos := 0
		for left < right {
			mid := left + (right-left)/2
			if res[mid] == nums[i] {
				right = mid
				left = mid + 1
			} else if res[mid] < nums[i] {
				left = mid + 1
			} else if res[mid] > nums[i] {
				right = mid
			}
		}

		//pos 定义了待插入的位置
		pos = left

		//res的长度不够 要想将长度+1
		res = append(res, 0)
		//然后挪一个位置  挪位置通过copy函数 copy函数将挪两个切片里较短的切片个数量
		copy(res[pos+1:], res[pos:])
		//赋值
		res[pos] = nums[i]
	}
	return res[len(res)-k]
}

//归并排序  经典递归分治  速度一下子就上来了
func findKthLargest5(nums []int, k int) int {
	//合并
	var mergeTwofunc func(left, mid, right int) []int
	mergeTwofunc = func(left, mid, right int) []int {
		res := make([]int, 0, right-left+1)
		l, r := left, mid
		for l < mid || r <= right {
			if l == mid {
				res = append(res, nums[r])
				r++
			} else if r == right+1 {
				res = append(res, nums[l])
				l++
			} else if nums[l] < nums[r] {
				res = append(res, nums[l])
				l++
			} else {
				res = append(res, nums[r])
				r++
			}
		}
		return res
	}

	//拆开 + 合并  合并前各自集合有序
	var mergeSort func(left, right int)
	mergeSort = func(left, right int) {
		if left == right {
			return
		}
		mid := left + (right-left)/2
		mergeSort(left, mid)
		mergeSort(mid+1, right)
		temp := mergeTwofunc(left, mid+1, right)
		copy(nums[left:right+1], temp)
	}

	left, right := 0, len(nums)-1
	mergeSort(left, right)
	return nums[len(nums)-k]
}

//部分排序 冒泡排序  利用冒泡排序 排了几次就确定了几个最大值 我只要第k大的值 所以只需要排k次就好了
func findKthLargest6(nums []int, k int) int {
	for i := 0; i < k; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums[len(nums)-k]
}

//部分排序 选择排序 利用选择排序 这回将最大值弄到前面 置换了k次 就找到了第k大的值了鸭
func findKthLargest(nums []int, k int) int {
	for i := 0; i < k; i++ {
		temp := i
		for j := i; j < len(nums); j++ {
			if nums[j] > nums[temp] {
				temp = j
			}
		}
		nums[i], nums[temp] = nums[temp], nums[i]
	}
	return nums[k-1]
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([][]byte, 0, numRows)
	for i := 0; i < numRows; i++ {
		res = append(res, []byte{})
	}

	b := false

	j := 0
	for i := 0; i < len(s); i++ {
		if j == numRows {
			j -= 2
			b = true
		} else if j == -1 {
			j += 2
			b = false
		}

		if !b {
			res[j] = append(res[j], s[i])
			j++
		} else {
			res[j] = append(res[j], s[i])
			j--
		}
	}

	ans := []byte(s)
	index := 0
	for _, re := range res {
		index += copy(ans[index:], re)
	}
	return string(ans)
}

//时间复杂度太高了呜呜呜
func maxArea1(height []int) int {
	res := 0

	var min func(a, b int) int
	min = func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			product := (j - i) * min(height[i], height[j])
			if product > res {
				res = product
			}
		}
	}

	return res
}

//接雨水问题 类似的 要使用双指针
func maxArea(height []int) int {
	res := 0

	var minIndex func(a, b int) int
	minIndex = func(a, b int) int {
		if height[a] > height[b] {
			return b
		}
		return a
	}

	left, right := 0, len(height)-1
	for left < right {
		lowerIndex := minIndex(left, right)
		product := height[lowerIndex] * (right - left)
		if product > res {
			res = product
		}
		if lowerIndex == right {
			right--
		} else {
			left++
		}
	}
	return res
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] >= 2 {
			return true
		}
	}

	return false
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 10)
	queue := make([]*TreeNode, 0, 10)
	if root != nil {
		queue = append(queue, root)
	}
	length := len(queue)
	level := 0

	for length != 0 {
		res = append(res, []int{})
		for i := 0; i < length; i++ {
			treeNode := queue[0]
			res[level] = append(res[level], treeNode.Val)
			if treeNode.Left != nil {
				queue = append(queue, treeNode.Left)
			}
			if treeNode.Right != nil {
				queue = append(queue, treeNode.Right)
			}
			queue = queue[1:]
		}
		length = len(queue)
		level++
	}
	return res
}

//左闭右开区间的形式
func search1(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

//左闭右闭的形式
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

//简单使用map来做 如果不使用额外的数据结构更好
func isUnique1(astr string) bool {
	m := make(map[int32]bool)
	for _, s := range astr {
		if m[s] {
			return false
		}
		m[s] = true
	}
	return true
}

//位运算 定义一个26个0的数字 位集 作为集合 s-'a'的距离作为移动的长度 给某位置1 利用和该位的与运算 来判定该位之前是不是1 进而判断该位是否重复
func isUnique(astr string) bool {
	temp := int32(0)
	for _, s := range astr {
		//计算好要移动多少步
		move := s - 'a'
		//如果temp中move位的值是1的话 与运算的结果就是1 那就说明重复了
		if temp>>move&1 == 1 {
			return false
		}
		//如果不是1 那就赋1
		temp |= 1 << move
	}
	return true
}

//两数之和 对待数组问题 看看能否先排个序  然后双指针  看看能否进行优化跳过某些节点
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	sort.Ints(nums)
	res := make([]int, 0)
	left, right := 0, len(nums)-1
	for left < right {
		temp := nums[left]
		temp1 := nums[right]
		if nums[left]+nums[right] == target {
			res = append(res, left)
			res = append(res, right)
			//跳过哪些已经经历过的节点  防止重复
			for left < right && nums[left] == temp {
				left++
			}
			//跳过哪些已经经历过的节点  防止重复
			for right > left && nums[right] == temp1 {
				right--
			}
		} else if nums[left]+nums[right] > target {
			temp := nums[right]
			right--
			//如果不行的话 也进行跳过  防止重复判断
			for right >= 0 && nums[right] == temp {
				right--
			}
		} else if nums[left]+nums[right] < target {
			temp := nums[left]
			left++
			//如果不行的话 也进行跳过  防止重复判断
			for left < len(nums) && nums[left] == temp {
				left++
			}
		}
	}
	return res
}

//三数之和就是两数之和的递归版本  其实质上就是  定一个数  然后另外两个求两数之和
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	res := make([][]int, 0, 10)
	//排序
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; {
		//优化 从i+1开始 之前的都已经遍历过了
		left, right := i+1, len(nums)-1
		for left < right {
			temp := nums[left]
			temp1 := nums[right]
			//这不就是两数之和吗
			if temp+temp1 == -nums[i] {
				res = append(res, []int{nums[i], temp, temp1})
				//相同的跳过 防止重复
				for left < right && temp == nums[left] {
					left++
				}
				//相同的跳过 防止重复
				for right > left && temp1 == nums[right] {
					right--
				}
			} else if temp+temp1 > -nums[i] {
				//优化 相同的跳过 防止重复判断
				for left < right && temp1 == nums[right] {
					right--
				}
			} else if temp+temp1 < -nums[i] {
				//优化 相同的跳过 防止重复判断
				for right > left && temp == nums[left] {
					left++
				}
			}
		}
		tempi := nums[i]
		for i < len(nums)-2 && nums[i] == tempi {
			i++
		}
	}

	return res
}

//属实有点不会了 一次遍历
func maxProfit1(prices []int) int {
	res := 0
	min := math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			temp := prices[i] - min
			if temp > res {
				res = temp
			}
		}
	}
	return res
}

//单调栈 正序入栈 递增递减问题考虑单调栈或者单调队列
func maxProfit2(prices []int) int {
	res := 0
	stack := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		if len(stack) == 0 {
			//空栈就入栈
			stack = append(stack, prices[i])
			continue
		}

		//递增的就入栈 然后更新答案
		if stack[len(stack)-1] < prices[i] {
			stack = append(stack, prices[i])
			temp := prices[i] - stack[0]
			if temp > res {
				res = temp
			}
			continue
			//如果是递减的 去掉栈顶 直到恢复递增
		} else if stack[len(stack)-1] > prices[i] {
			for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, prices[i])
		}
	}
	return res
}

//单调栈 逆序入栈
func maxProfit(prices []int) int {
	res := 0
	stack := make([]int, 0)
	for i := len(prices) - 1; i >= 0; i-- {
		//空栈就入栈
		if len(stack) == 0 {
			stack = append(stack, prices[i])
			continue
		}

		//递减的就入栈 然后更新答案
		if stack[len(stack)-1] > prices[i] {
			temp := stack[0] - prices[i]
			if temp > res {
				res = temp
			}
			stack = append(stack, prices[i])
			continue
			//如果是递增的 去掉栈顶 直到恢复递减
		} else if stack[len(stack)-1] < prices[i] {
			for len(stack) > 0 && stack[len(stack)-1] < prices[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, prices[i])
		}
	}
	return res
}

func threeSumClosest(nums []int, target int) int {
	var abs func(x int, y int) int
	abs = func(x int, y int) int {
		if x > y {
			return x - y
		}
		return y - x
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if sum < target {
				temp := abs(sum, target)
				if abs(res, target) > temp {
					res = sum
				}

				tempLeft := nums[left]
				for left < right && tempLeft == nums[left] {
					left++
				}
			} else if sum > target {
				temp := abs(sum, target)
				if abs(res, target) > temp {
					res = sum
				}
				tempRight := nums[right]
				for right > left && tempRight == nums[right] {
					right--
				}
			}
		}
	}
	return res
}

func sortedSquares1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

func sortedSquares2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) < math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	return nums
}

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums), len(nums))
	i := len(res) - 1

	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
		i--
	}
	res[i] = nums[left] * nums[left]
	return res
}

//数组直接double 然后使用指针
func rotate1(nums []int, k int) {
	p := &nums
	l := len(nums)
	newNums := make([]int, 2*l, 2*l)
	copy(newNums, nums)
	copy(newNums[l:], nums)
	k %= l
	copy(*p, newNums[len(newNums)-k-l:len(newNums)-k])
}

//额外创建大小为k的数组 节省内存
func rotate2(nums []int, k int) {
	l := len(nums)
	k %= l
	ints := make([]int, k, k)
	copy(ints, nums[len(nums)-k:])
	copy(nums[k:], nums[:len(nums)-k])
	copy(nums[:k], ints)
	return
}

//连续三次翻转 即可得到答案
func rotate(nums []int, k int) {
	var swap func(nums []int)
	swap = func(nums []int) {
		for i := 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
	}
	k %= len(nums)
	swap(nums[:len(nums)-k])
	swap(nums[len(nums)-k:])
	swap(nums)
	return
}

//不太会 前缀和思想 主要还是先预处理一下 不要直接把数组累加起来啊啊啊啊  为了避免双重for循环 最好的方法是使用map存起来 备忘录思想怎么忘记了卧槽
func numberOfSubarrays1(nums []int, k int) int {
	res := 0

	//记录奇数为i的个数有几个
	m := make(map[int]int)

	preSum := make([]int, len(nums)+1, len(nums)+1)
	preSum[0] = 0
	m[0]++

	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			preSum[i+1] = preSum[i]
		} else {
			preSum[i+1] = preSum[i] + 1
		}
		m[preSum[i+1]]++
	}

	for i := 1; i < len(preSum); i++ {
		if preSum[i] < k {
			continue
		}
		res += m[preSum[i]-k]
	}

	return res
}

//属实不会 数学方法太强了 我擦
//思想 偶数不要的话 就直接去奇数的数组 但是偶数的个数会影响答案 所以奇数数组里存的是下标和
//奇数数组中两两的差就是中间偶数个数的数量 就可以和答案挂钩
func numberOfSubarrays(nums []int, k int) int {
	res := 0

	odd := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 0 {
			continue
		} else {
			odd = append(odd, i)
		}
	}

	for i := 0; i < len(odd); i++ {
		if i+1 < k {
			continue
		} else if i+1-k == 0 {
			if i == len(odd)-1 {
				res += (odd[i+1-k] + 1) * (len(nums) - odd[i])
			} else if i < len(odd)-1 {
				res += (odd[i+1-k] + 1) * (odd[i+1] - odd[i])
			}
		} else if i+1-k > 0 {
			if i == len(odd)-1 {
				res += (odd[i+1-k] - odd[i+1-k-1]) * (len(nums) - odd[i])
			} else if i < len(odd)-1 {
				res += (odd[i+1-k] - odd[i+1-k-1]) * (odd[i+1] - odd[i])
			}
		}
	}

	return res
}

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, []int{intervals[0][0], intervals[0][1]})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			if intervals[i][1] >= res[len(res)-1][1] {
				res[len(res)-1][1] = intervals[i][1]
			}
		} else {
			res = append(res, []int{intervals[i][0], intervals[i][1]})
		}
	}

	return res
}

//直接用copy修改底层数组
func moveZeroes1(nums []int) {
	j := 0
	for i := 0; i < len(nums)-j; {
		if nums[i] == 0 {
			copy(nums[i:], nums[i+1:len(nums)-j])
			nums[len(nums)-1-j] = 0
			i = 0
			j++
			continue
		}
		i++
	}
	return
}

func moveZeroes2(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, 1
	for right < len(nums) {
		if nums[left] != 0 && nums[right] != 0 {
			left++
			right++
		} else if nums[right] == 0 && nums[left] == 0 {
			right++
		} else if nums[right] == 0 {
			left++
			right++
		} else if nums[left] == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
		}
	}
}

//尽可能减少交换次数
func moveZeroes3(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, 1

	//同时限制左边和右边的指针 有时候只用一个for就够了
	for right < len(nums) {
		if nums[left] != 0 {
			left++
			right++
			continue
		}

		if nums[right] == 0 {
			right++
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right++
	}
}

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

//双指针  o(n)
func twoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else if numbers[left]+numbers[right] < target {
			left++
		}
	}

	return []int{}
}

//记得用二分法啊啊啊啊啊 属实不会 这里的二分法 不能直接使用 得这么想 固定某个数a 然后寻找target-a 是否存在数组中 o(nlog(n))
func twoSum3(numbers []int, target int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		temp := target - numbers[i]
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			if numbers[mid] == temp {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > temp {
				right = mid - 1
			} else if numbers[mid] < temp {
				left = mid + 1
			}
		}
		if left < 0 || left == len(numbers)-1 || numbers[left] != temp {
			continue
		}
	}

	return []int{}
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return
}

func reverseWords1(s string) string {
	res := make([]byte, 0, len(s))
	split := strings.Split(s, " ")
	for _, words := range split {
		for i := len(words) - 1; i >= 0; i-- {
			res = append(res, words[i])
		}
		res = append(res, ' ')
	}
	return string(res[:len(res)-1])
}

func reverseWords(s string) string {
	res := []byte(s)
	last := 0
	for i := 0; i < len(res); i++ {
		if res[i] == ' ' {
			for j := 0; j < (i-last)/2; j++ {
				res[last+j], res[i-1-j] = res[i-1-j], res[last+j]
			}
			last = i + 1
		}

		if i == len(res)-1 {
			for j := 0; j < (i-last)/2; j++ {
				res[last+j], res[i-j] = res[i-j], res[last+j]
			}
		}
	}
	return string(res)
}

//冒泡排序  超时
func sortArray1(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		//适当剪枝 len()-i到最后的部分都是有序区，避免再排
		for j := i + 1; j < len(nums)-i; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

//选择排序  超时
func sortArray2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		temp := i
		for j := i + 1; j < len(nums); j++ {
			if nums[temp] > nums[j] {
				temp = j
			}
		}
		nums[i], nums[temp] = nums[temp], nums[i]
	}
	return nums
}

//归并排序 没超时  先拆 拆到最后只剩一个了在排 排起来之后在两两排 合并
func sortArray3(nums []int) []int {
	var mergeNums func(leftNums []int, rightNums []int) []int
	mergeNums = func(leftNums []int, rightNums []int) []int {
		res := make([]int, 0, len(leftNums)+len(rightNums))
		l, r := 0, 0
		for l < len(leftNums) || r < len(rightNums) {
			if l < len(leftNums) && r < len(rightNums) {
				if leftNums[l] < rightNums[r] {
					res = append(res, leftNums[l])
					l++
				} else {
					res = append(res, rightNums[r])
					r++
				}
			} else if l < len(leftNums) {
				res = append(res, leftNums[l])
				l++
			} else if r < len(rightNums) {
				res = append(res, rightNums[r])
				r++
			}
		}
		return res
	}

	var split func([]int) []int
	split = func(nums []int) []int {
		if len(nums) == 1 {
			return nums
		}
		i, j := 0, len(nums)-1
		mid := i + (j-i)/2 + 1
		leftNums := split(nums[:mid])
		rightNums := split(nums[mid:])
		return mergeNums(leftNums, rightNums)
	}
	return split(nums)
}

//快速排序  谢了两小时你个废物草泥马的  思路是:随机选择出一个数 放到头部 然后使用左右指针把把数组分成两坨 一坨比这个数大 一坨比这个数小 然后再递归的调用这个函数 去解决两个数组各自的排序问题
func sortArray4(nums []int) []int {

	var quickSort func(nums []int)
	quickSort = func(nums []int) {
		if len(nums) == 0 || len(nums) == 1 {
			return
		}

		left, right := 0, len(nums)-1
		intn := rand.Intn(len(nums))
		nums[0], nums[intn] = nums[intn], nums[0]
		pivot := nums[0]
		//b站视频里直接取下标0的作为pivot 而随机的更好 防止出现复杂度不稳定的情况 如何修正呢
		//直接把这个随机的数和下标为0的换一下位置不就好了 傻逼艹
		for left < right {
			for right > left && nums[right] >= pivot {
				right--
			}

			if left < right {
				nums[left] = nums[right]
			}

			for left < right && nums[left] <= pivot {
				left++
			}

			if left < right {
				nums[right] = nums[left]
			}
			if left >= right {
				nums[left] = pivot
			}
		}

		//这里很关键 这里使用left 来区分 而不是intn 什么几把玩意儿
		quickSort(nums[:left])
		quickSort(nums[left+1:])

	}

	quickSort(nums)
	return nums
}

type myheap1 []int

func (m *myheap1) Len() int {
	return len(*m)
}

func (m *myheap1) Less(i, j int) bool {
	return (*m)[i] < (*m)[j]
}

func (m *myheap1) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myheap1) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *myheap1) Pop() interface{} {
	i := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return i
}

//堆排序 重要!!!
//1.使用堆heap的api的时候要实现相关方法
//2.使用push pop 方法的时候要带上heap 否则就只会执行自己定义的方案
//3.使用pop的时候 会先把堆顶元素和堆尾元素交换 然后把堆尾元素弹出来
//4.所以实现pop方法的时候 由于要把堆尾元素弹出来,因此就是把数组的最后一个元素弹出
//5.heap中的建堆规则是使用上浮操作 实际上下沉操作也是可以的  因此push的时候要在尾部push 然后判断是否需要上浮
//6.由于建堆是完全二叉树 完全二叉树的性质是 父节点i*2 =左子节点 i*2+1=右子节点 而子节点j/2 =父节点
//7.因此利用up上浮操作 或者down下沉操作 就可以完成子父节点的swap
func sortArray(nums []int) []int {
	//类型转换好像会使用相同底层数组 或者可能是因为该类型是引用类型
	m2 := new(myheap1)
	for i := 0; i < len(nums); i++ {
		//push的时候就会排序 但是排序只会找出堆顶的最大或者最小值
		heap.Push(m2, nums[i])
	}

	for i := 0; i < len(nums); i++ {
		//init的时候会再次排序
		//heap.Init(m2)

		//pop的时候会将 堆顶元素和尾巴交换 继而把尾巴元素弹出 pop的时候就会排序
		p := heap.Pop(m2).(int)
		nums[i] = p
	}

	return nums
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//直接用双指针 管他呢
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	l := new(ListNode)
	l.Next = head
	pre, slow, fast := new(ListNode), l, l
	for n > 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	return l.Next
}

//凡是涉及到倒序的问题 都要考虑栈这种先进后出的特殊数据结构
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s := make([]*ListNode, 0)
	cur := head
	for cur != nil {
		s = append(s, cur)
		cur = cur.Next
	}

	for n > 1 {
		s = s[:len(s)-1]
		n--
	}

	if len(s)-2 >= 0 {
		s[len(s)-2].Next = s[len(s)-1].Next
	} else {
		return s[len(s)-1].Next
	}
	return head
}

//超时 复杂度 m*n
func findRadius1(houses []int, heaters []int) int {
	m := make(map[int]int)
	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a-b <= 0 {
			return b - a
		}
		return a - b
	}

	for _, heater := range heaters {
		for _, house := range houses {
			if v, ok := m[house]; !ok {
				m[house] = abs(house, heater)
			} else if abs(house, heater) < v {
				m[house] = abs(house, heater)
			}
		}
	}

	res := 0
	for _, v := range m {
		if v > res {
			res = v
		}
	}
	return res
}

//还是超时了 中心拓展的思路都有了还是超时了
func findRadius2(houses []int, heaters []int) int {
	m := make(map[int]bool)

	sort.Ints(houses)
	sort.Ints(heaters)
	for _, heater := range heaters {
		m[heater] = true
	}

	res := 0
	for _, house := range houses {
		i := 0
		for !m[house+i] && !m[house-i] {
			i++
		}

		if i > res {
			res = i
		}
	}

	return res
}

//1.集合没事儿就排排序 复杂度最低只有n 复杂一点的只有n*log(n) 不排白不排
//2.题目难的直接用api  直接降低时间复杂度 否则搜索写不好o(n) 二分直接折半log(n)
func findRadius(houses []int, heaters []int) int {
	var abs func(a, b int) int
	abs = func(a, b int) int {
		if a-b <= 0 {
			return b - a
		}
		return a - b
	}
	sort.Ints(houses)
	sort.Ints(heaters)

	res := 0
	for _, house := range houses {
		//二分查找 如果找到相等的元素 则返回元素的下标 如果找不到相等的元素 会返回应该要被插入的下标位置
		ints := sort.SearchInts(heaters, house)
		temp := 0
		if ints == len(heaters) {
			//因此 如果返回的下标 超出了数组的索引 就说明在最后一位  那么只能与数组的最后一位进行比较 没有谁比它更接近了
			temp = abs(heaters[ints-1], house)
		} else if ints == 0 {
			//因此 如果返回的下标是0 说明应该被插入到0的位置  那么只能与数组的0位置进行比较 没有谁比它更接近了
			temp = abs(heaters[ints], house)
		} else {
			//因此 如果返回的下标不具有特殊性 说明应该被插入到 x 的位置上 那么有可能当前这个值 比原来数组里x索引 更近 例如 {1,3,6} 比较的值是5
			//也有可能当前这个值 比原来数组里 x-1索引 更近  例如 {1,3,6} 比较的值是4
			temp = abs(heaters[ints], house)
			temp1 := abs(heaters[ints-1], house)
			//两个结果进行比较 取较小的那一个
			if temp1 < temp {
				temp = temp1
			}
		}

		if temp > res {
			res = temp
		}
	}

	return res
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	res := 0
	m := make(map[uint8]int)
	left, right := 0, 1
	m[s[left]]++
	for right < len(s) && left <= right {
		for right < len(s) && m[s[right]] == 0 {
			m[s[right]]++
			right++
		}

		temp := right - left
		if res < temp {
			res = temp
		}

		if right == len(s) {
			return res
		}

		//优化 如果存在相连着的字符 那就直接一路跳到不连为止
		for m[s[right]] != 0 {
			m[s[left]]--
			left++
		}
	}

	return res
}

//万不得已下 别用深度比较 反射的包都太拉垮了 反射性能很差
func checkInclusion1(s1 string, s2 string) bool {
	m1 := make(map[int32]int)
	for _, v := range s1 {
		m1[v]++
	}

	l := len(s1)
	for i := 0; i < len(s2)-l+1; i++ {
		m2 := make(map[int32]int)
		for _, v := range s2[i : i+l] {
			m2[v]++
		}
		equal := reflect.DeepEqual(m1, m2)
		if equal {
			return true
		}
	}
	return false
}

//我说我这是滑动窗口我都不敢信
func checkInclusion2(s1 string, s2 string) bool {
	ints := [26]int{}
	for _, v := range s1 {
		ints[v-'a']++
	}

	l := len(s1)
	for i := 0; i < len(s2)-l+1; i++ {
		temp := ints
		for _, v := range s2[i : i+l] {
			if temp[v-'a'] == 0 {
				break
			} else {
				temp[v-'a']--
			}
		}
		if temp == [26]int{} {
			return true
		}
	}
	return false
}

//这才是 正统的滑动窗口好吗
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	ints := [26]int{}
	ints2 := [26]int{}
	for i, v := range s1 {
		ints[v-'a']++
		ints2[s2[i]-'a']++
	}
	if ints == ints2 {
		return true
	}

	n := len(s1)
	j := n

	//维护滑动窗口
	for j < len(s2) {
		ints2[s2[j-n]-'a']--
		ints2[s2[j]-'a']++
		j++
		if ints == ints2 {
			return true
		}
	}

	return false
}

//对角线的特征是 行索引和列索引和相等 可以利用和是奇数和偶数来判断 正向遍历还是反向遍历
func findDiagonalOrder(mat [][]int) []int {
	res := make([]int, 0, 0)
	if len(mat) == 1 {
		return mat[0]
	}
	if len(mat[0]) == 1 {
		for _, ints := range mat {
			res = append(res, ints[0])
		}
		return res
	}

	sum := len(mat) + len(mat[0])
	for i := 0; i <= sum; i++ {
		for j := 0; j <= i; j++ {
			if i%2 == 0 {
				if i-j >= 0 && i-j < len(mat) && j >= 0 && j < len(mat[0]) {
					res = append(res, mat[i-j][j])
				}
			} else {
				if j >= 0 && j < len(mat) && i-j >= 0 && i-j < len(mat[0]) {
					res = append(res, mat[j][i-j])
				}
			}
		}
	}
	return res
}

//横向比较
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	res := []byte(strs[0])
	for _, str := range strs[1:] {
		if len(res) == 0 {
			return ""
		}

		for i := 0; i < len(res); i++ {
			if res[i] != str[i] {
				res = res[:i]
				break
			}
		}
	}

	return string(res)
}

//纵向比较
func longestCommonPrefix(strs []string) string {
	res := []byte(strs[0])
	for i := range res {
		for _, str := range strs[1:] {
			if i == len(str) {
				return string(res[:i])
			}
			if str[i] != res[i] {
				res = res[:i]
				return string(res)
			}
		}
	}

	return string(res)
}
