package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//obj := Constructor(2)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//param_1 := obj.Get(1)
	//fmt.Println(param_1)
	//obj.Put(3, 3)
	//get := obj.Get(2)
	//fmt.Println(get)
	//obj.Put(4, 4)
	//i := obj.Get(1)
	//fmt.Println(i)
	//i2 := obj.Get(3)
	//fmt.Println(i2)
	//i3 := obj.Get(4)
	//fmt.Println(i3)

	//obj := Constructor(2)
	//obj.Put(1, 0)
	//obj.Put(2, 2)
	//param_1 := obj.Get(1)
	//fmt.Println(param_1)
	//obj.Put(3, 3)
	//get := obj.Get(2)
	//fmt.Println(get)
	//obj.Put(4, 4)
	//i := obj.Get(1)
	//fmt.Println(i)
	//i2 := obj.Get(3)
	//fmt.Println(i2)
	//i3 := obj.Get(4)
	//fmt.Println(i3)

	//root := &TreeNode{
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

	//root := &TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val: 2,
	//	},
	//}

	//root := &TreeNode{
	//	Val:   4,
	//	Left:  nil,
	//	Right: nil,
	//}
	//ser := Constructor()
	//deser := Constructor()
	//data := ser.serialize(root)
	//ans := deser.deserialize(data)
	//fmt.Println(ans)

	//wuwuwu := &TreeNode{
	//	Val:   6,
	//	Left:  nil,
	//	Right: nil,
	//}
	//left := &TreeNode{
	//	Val:  5,
	//	Left: wuwuwu,
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   7,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
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
	//root := &TreeNode{
	//	Val:   3,
	//	Left:  left,
	//	Right: right,
	//}
	//ancestor := lowestCommonAncestor(root, left, wuwuwu)
	//fmt.Println(ancestor)
	//
	//a()
	//
	//element := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	//fmt.Println(element)

	//temperatures := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	//temperatures := dailyTemperatures([]int{30, 40, 50, 60})
	//temperatures := dailyTemperatures([]int{30, 60, 90})
	//fmt.Println(temperatures)

	//elements := nextGreaterElements([]int{1, 2, 1})
	//elements := nextGreaterElements([]int{1, 2, 3, 4, 3})
	//elements := nextGreaterElements([]int{5, 4, 3, 2, 1})
	//fmt.Println(elements)

	//ints := intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}})
	//ints := intersection([][]int{{1, 2, 3}, {4, 5, 6}})
	//fmt.Println(ints)

	//points := countLatticePoints([][]int{{2, 2, 1}})
	//points := countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}})
	//fmt.Println(points)

	//rectangles := countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}})
	//rectangles := countRectangles([][]int{{4, 7}, {4, 9}, {8, 5}, {6, 2}, {6, 4}}, [][]int{{4, 2}, {5, 6}})
	//rectangles := countRectangles([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}})
	//rectangles := countRectangles([][]int{{4773, 11}, {7353, 65}, {9905, 38}, {4016, 60}, {1569, 8}, {2814, 29}, {5479, 50}, {280, 11}, {2514, 31}, {4387, 79}, {2286, 32}, {9345, 61}, {9075, 63}, {2310, 43}, {2498, 4}, {2571, 59}, {6387, 46}, {9265, 13}, {7643, 20}, {7922, 78}, {2399, 21}, {9303, 83}, {114, 31}, {5903, 34}, {1780, 73}, {4615, 71}, {7343, 44}, {8134, 77}, {7010, 20}, {3291, 77}, {2273, 80}, {4332, 11}, {8345, 13}, {7634, 66}, {3945, 72}, {2887, 30}, {3025, 28}, {8686, 75}, {513, 32}, {2303, 5}, {1256, 1}, {230, 38}, {3544, 83}, {1320, 41}, {8790, 27}, {5119, 21}, {6470, 29}, {9730, 11}, {6071, 64}, {1599, 92}, {3845, 2}, {5651, 49}, {5182, 8}, {9925, 41}, {7761, 75}, {8597, 80}, {530, 94}, {2620, 86}, {5049, 28}, {2585, 9}, {3940, 40}, {6626, 78}, {5393, 85}, {9234, 83}, {3088, 3}, {9300, 31}, {330, 72}, {9906, 14}, {9491, 61}, {4750, 86}, {9024, 8}, {3277, 18}, {2753, 91}, {8564, 50}, {4501, 73}, {2221, 99}, {3154, 34}, {3946, 5}, {5497, 87}, {2666, 77}, {2101, 89}, {656, 67}, {7039, 30}, {8536, 23}, {2167, 23}, {3982, 7}, {4510, 43}, {9836, 63}, {734, 99}, {684, 12}, {93, 24}, {1895, 83}, {7175, 66}, {5533, 66}, {5133, 20}, {2738, 48}, {43, 72}, {1906, 40}, {9609, 98}, {7975, 35}, {2304, 89}, {1920, 98}, {1277, 31}, {4120, 98}, {1163, 23}, {5926, 88}, {163, 54}, {5804, 28}, {2349, 66}, {1083, 73}, {4589, 29}}, [][]int{{854, 53}, {237, 98}, {4604, 59}, {200, 38}, {5591, 10}, {443, 95}, {393, 69}, {7558, 8}, {9186, 17}, {7554, 67}, {8140, 14}, {6150, 59}, {9027, 45}, {3494, 34}, {1494, 65}, {5454, 60}, {5924, 19}, {9849, 5}, {2682, 34}, {9563, 61}, {4109, 13}, {4658, 49}, {271, 46}, {1008, 98}, {2951, 20}, {500, 93}, {1234, 4}, {501, 2}, {9749, 7}, {7545, 52}, {2020, 39}, {4685, 15}, {2239, 63}, {257, 27}, {4270, 88}, {9763, 60}, {8013, 95}, {2784, 68}, {3482, 40}, {8319, 63}, {6273, 54}, {3626, 100}, {7685, 42}, {3635, 9}, {3527, 89}, {2192, 64}, {7012, 90}, {4050, 83}, {7959, 13}, {9059, 74}, {3141, 73}, {5958, 75}, {9602, 62}, {3297, 58}, {254, 27}, {8777, 49}, {7289, 72}, {5417, 33}, {4796, 49}, {2409, 82}, {4444, 38}, {4645, 81}, {748, 96}, {8122, 92}, {2052, 64}, {6882, 52}, {5229, 14}, {9394, 42}, {6119, 100}, {3275, 30}, {8361, 80}, {3525, 59}, {2042, 2}, {4921, 43}, {2542, 60}, {6431, 51}, {3354, 21}, {3974, 67}, {60, 8}, {399, 82}, {943, 52}, {7273, 48}, {183, 51}, {4993, 69}, {884, 15}, {7488, 23}, {1443, 26}, {6543, 2}, {993, 51}, {9188, 5}, {2108, 32}, {1435, 74}, {8172, 49}, {3863, 42}, {216, 81}, {8714, 95}, {8022, 66}, {2102, 78}, {8521, 73}, {1667, 26}, {5894, 74}, {3340, 86}, {3051, 46}, {700, 47}, {5150, 68}, {6941, 33}, {8696, 64}, {1678, 76}, {7489, 36}, {1342, 91}, {8456, 22}, {2433, 47}, {8416, 5}, {9218, 81}, {4892, 16}, {6733, 12}, {1793, 89}, {6404, 85}, {3976, 16}, {3183, 12}, {9444, 91}, {5232, 57}, {3743, 80}, {7528, 34}, {5153, 24}, {2404, 44}, {8085, 80}, {9371, 65}, {8423, 76}, {6971, 75}, {5934, 9}, {6820, 6}, {9339, 35}, {4737, 83}, {6943, 82}, {4569, 84}, {7750, 2}, {6457, 75}, {7635, 58}, {9120, 39}, {6190, 29}, {4574, 84}, {2346, 52}, {5673, 79}, {4706, 73}, {2027, 43}, {4509, 6}, {7951, 29}, {3389, 14}, {7619, 38}, {6339, 89}, {419, 57}, {1564, 93}, {1439, 56}, {1833, 67}, {6692, 74}, {293, 3}, {8718, 6}, {2481, 75}, {2872, 35}, {7620, 75}, {7280, 13}, {6131, 96}, {9523, 23}, {1515, 27}, {3338, 47}, {6933, 2}, {3294, 50}, {89, 7}, {3746, 86}, {5030, 75}, {6953, 41}, {7991, 57}, {8061, 20}, {4239, 37}, {3237, 58}, {5773, 4}, {7318, 69}, {1691, 39}, {5334, 72}, {855, 58}, {1532, 81}, {746, 84}, {3643, 4}, {2370, 40}, {4914, 80}, {4187, 9}, {5899, 93}, {4600, 43}, {512, 57}, {267, 76}, {1925, 93}, {3231, 76}, {631, 70}, {9001, 94}, {4598, 86}, {2607, 44}, {2225, 29}, {5613, 26}, {7402, 33}, {6899, 74}, {9210, 65}, {3512, 98}, {7940, 83}, {7627, 81}, {8329, 100}, {252, 39}, {71, 55}, {2005, 53}, {3824, 21}, {4312, 15}, {2136, 36}, {3699, 98}, {604, 17}, {5138, 54}, {8435, 42}, {8845, 65}, {3343, 38}, {6203, 16}, {4607, 50}, {6317, 45}, {4999, 2}, {1988, 16}, {3514, 80}, {7596, 72}, {3486, 31}, {8139, 48}, {5752, 8}, {5063, 100}, {4776, 51}, {1641, 24}, {1738, 45}, {5476, 85}, {6735, 2}, {2188, 55}, {5195, 76}, {7779, 55}, {5444, 84}, {1184, 34}, {7873, 70}, {1026, 27}, {5009, 38}, {8610, 59}, {9465, 79}, {885, 34}, {3900, 4}, {2011, 63}, {7772, 60}, {9389, 78}, {9725, 8}, {7578, 75}, {5163, 74}, {526, 18}, {6305, 50}, {3079, 91}, {6490, 98}, {9077, 35}, {2350, 35}, {1003, 55}, {297, 38}, {6902, 61}, {4319, 17}, {795, 58}, {6875, 88}, {2999, 62}, {3704, 87}, {8253, 87}, {5729, 42}, {6980, 51}, {7041, 60}, {3660, 67}, {13, 50}, {1375, 10}, {2227, 10}, {5006, 24}, {1819, 40}, {9800, 89}, {5607, 19}, {244, 11}, {1798, 20}, {3647, 34}, {4414, 67}, {1145, 26}, {1654, 83}, {5359, 21}, {6414, 67}, {9100, 13}, {4264, 93}, {6928, 14}, {5273, 99}, {3599, 39}, {5690, 77}, {2487, 62}, {9656, 90}, {6486, 65}, {8259, 51}, {7722, 62}, {6290, 49}, {3324, 12}, {5275, 46}, {2546, 35}, {4647, 14}, {4252, 22}, {9983, 91}, {9513, 15}, {9830, 34}, {9723, 36}, {9724, 29}, {9466, 72}, {4164, 72}, {4175, 86}, {5834, 12}, {7131, 23}, {5945, 82}, {6610, 74}, {8514, 62}, {2472, 77}, {2857, 87}, {572, 1}, {1195, 35}, {9578, 35}, {416, 32}, {7520, 80}, {4694, 12}, {6538, 27}, {9255, 43}, {3084, 56}, {5385, 16}, {9633, 30}, {8856, 34}, {7831, 42}, {485, 23}, {1923, 75}, {4697, 40}, {1251, 20}, {5756, 51}, {6696, 73}, {1915, 8}, {2318, 14}, {9711, 78}, {3500, 21}, {9061, 25}, {8575, 17}, {4807, 96}, {6121, 9}, {8816, 31}, {7268, 87}, {9467, 7}, {351, 86}, {2278, 100}, {4838, 79}, {9013, 70}, {1071, 34}, {8713, 84}, {6777, 80}, {7819, 46}, {5817, 45}, {2593, 59}, {848, 17}, {2904, 19}, {7198, 42}, {2231, 58}, {9568, 64}, {1874, 97}, {2591, 23}, {4282, 69}, {134, 58}, {1692, 74}, {9143, 37}, {437, 9}, {8181, 93}, {4828, 35}, {5447, 14}, {3177, 6}, {7907, 45}, {1266, 44}, {6527, 54}, {8123, 46}, {9157, 13}, {9664, 57}, {9727, 85}, {9726, 92}, {1431, 62}, {1392, 8}, {28, 23}, {2073, 35}, {1987, 53}, {9779, 61}, {556, 95}, {5802, 62}, {7510, 23}, {1116, 98}, {6999, 12}, {2053, 49}, {4806, 67}, {5197, 66}, {9105, 79}, {3305, 34}, {6020, 90}, {2443, 1}, {3609, 23}, {962, 69}, {2553, 3}, {2476, 54}, {8072, 51}, {5840, 94}, {3672, 39}, {6564, 42}, {3453, 85}, {7175, 52}, {7429, 89}, {781, 14}, {4445, 53}, {6537, 51}, {2782, 48}, {9181, 83}, {2541, 93}, {1899, 69}, {6059, 33}, {2632, 61}, {6566, 53}, {3537, 98}, {9307, 58}, {4306, 50}, {2573, 8}, {7762, 85}, {5412, 24}, {112, 56}, {2417, 23}, {179, 58}, {8960, 45}, {5100, 17}, {1023, 24}, {8347, 27}, {3099, 42}, {6076, 15}, {5786, 85}, {9683, 50}, {6609, 35}, {542, 98}, {2969, 4}, {9298, 17}, {7862, 19}, {2669, 5}, {2208, 43}, {2512, 34}, {7444, 90}, {4395, 7}, {5956, 10}, {9994, 72}, {8961, 56}, {6120, 22}, {6637, 68}, {8580, 42}, {3977, 12}, {2063, 99}, {8332, 98}, {4602, 33}, {882, 82}, {4061, 73}, {7474, 62}, {5605, 92}, {5964, 23}, {3360, 54}, {6671, 95}, {3170, 79}, {8246, 59}, {6, 28}, {892, 4}, {6650, 47}, {421, 32}, {4617, 90}, {2297, 5}, {7718, 76}, {3188, 48}, {5059, 97}, {3326, 80}, {4611, 39}, {5481, 25}, {7684, 53}, {8535, 59}, {6681, 47}, {4819, 12}, {698, 75}, {3934, 49}, {844, 54}, {3250, 55}, {1611, 89}, {8054, 58}, {45, 16}, {8390, 84}, {6727, 76}, {7869, 19}, {5643, 98}, {1932, 67}, {5656, 50}, {1529, 66}, {434, 80}, {5816, 37}, {1004, 66}, {849, 68}, {9517, 22}, {7753, 70}, {2545, 27}, {596, 84}, {9915, 66}, {5402, 13}, {8094, 81}, {7526, 93}, {4456, 3}, {2471, 17}, {7231, 18}, {8784, 24}, {1176, 7}, {1019, 5}, {7182, 30}, {6417, 76}, {2219, 36}, {371, 82}, {3358, 30}, {1419, 13}, {256, 15}, {5226, 66}, {3349, 21}, {3566, 83}, {4115, 84}, {3656, 100}, {9996, 52}, {7004, 31}, {4373, 14}, {8190, 41}, {1099, 3}, {55, 38}, {5266, 96}})
	//fmt.Println(rectangles)

	//i := gohome([]int{1, 2, 3, 4}, []int{2, 4, 3, 4})
	//fmt.Println(i)

	//window := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	//window := maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4)
	//window := maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
	//fmt.Println(window)

	palindrome := isPalindrome(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	})

	//palindrome := isPalindrome(&ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val:  2,
	//		Next: nil,
	//	},
	//})

	fmt.Println(palindrome)

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
	//node := reserve(l2)
	//
	//fmt.Println(node)
}

type Node struct {
	pre  *Node
	next *Node
	key  int
	val  int
}

func (n *Node) RemoveSelf() {
	n.pre.next = n.next
	n.next.pre = n.pre
	n.next = nil
	n.pre = nil
}

type DoubleLink struct {
	head *Node
	tail *Node
	size int
}

func (d *DoubleLink) RemoveFirst() int {
	if d.head.next == d.tail {
		return -1
	}
	temp := d.head.next.key
	d.head.next.RemoveSelf()

	d.size--
	return temp
}

func (d *DoubleLink) AddLast(newNode *Node) {
	newNode.pre = d.tail.pre
	d.tail.pre.next = newNode
	newNode.next = d.tail
	d.tail.pre = newNode

	d.size++
}

func (d *DoubleLink) RemoveNode(recentNode *Node) {
	recentNode.RemoveSelf()
	d.size--
}

func (d *DoubleLink) MakeRecently(recentNode *Node) {
	d.RemoveNode(recentNode)
	d.AddLast(recentNode)
}

// LRUCache 输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
type LRUCache struct {
	cache    map[int]*Node
	link     *DoubleLink
	capacity int
}

// Constructor 头尾都是虚指针 要在头部后面进行删除  要在尾部前面进行新增
//func Constructor(capacity int) LRUCache {
//
//	head := &Node{
//		pre:  nil,
//		next: nil,
//		key:  0,
//		val:  0,
//	}
//
//	tail := &Node{
//		pre:  nil,
//		next: nil,
//		key:  0,
//		val:  0,
//	}
//	head.next = tail
//	tail.pre = head
//
//	return LRUCache{
//		cache: make(map[int]*Node),
//		link: &DoubleLink{
//			head: head,
//			tail: tail,
//			size: 0,
//		},
//		capacity: capacity,
//	}
//}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.link.MakeRecently(v)
		return v.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.cache[key]; !ok {
		newNode := &Node{key: key, val: value}
		this.link.AddLast(newNode)
		this.cache[key] = newNode
		if this.link.size > this.capacity {
			old := this.link.RemoveFirst()
			delete(this.cache, old)
		}
	} else {
		v.val = value
		this.link.MakeRecently(v)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q == nil {
		return false
	} else if p == nil && q != nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//二叉搜索树的性质 中序遍历就可以从小到大输出了诶
func isValidBST1(root *TreeNode) bool {
	m := make([]int, 0, 0)
	var travel func(root *TreeNode)
	travel = func(root *TreeNode) {
		if root == nil {
			return
		}
		travel(root.Left)
		m = append(m, root.Val)
		travel(root.Right)
	}

	travel(root)

	for i := 0; i < len(m)-1; i++ {
		if m[i] >= m[i+1] {
			return false
		}
	}

	return true
}

//二叉搜索树 尝试带入参数 进而方便判断
func isValidBST(root *TreeNode) bool {
	var travel func(root *TreeNode, max int, min int) bool
	travel = func(root *TreeNode, max int, min int) bool {
		if root == nil {
			return true
		} else if root.Val >= max || root.Val <= min {
			return false
		}

		return travel(root.Left, root.Val, min) && travel(root.Right, max, root.Val)
	}

	return travel(root, math.MaxInt, math.MinInt)
}

//如果是二叉搜索树的话就有更好的解决方案 结合性质 左边小右边大 类似于二分法
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := searchBST(root.Left, val)
	right := searchBST(root.Right, val)
	if left == nil && right == nil {
		return nil
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

//删除 得先找到 然后找大于它的最小 然后交换 然后删掉最小的 删掉的时候 重复调用构造好的函数
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	var getMin func(root *TreeNode) *TreeNode
	getMin = func(root *TreeNode) *TreeNode {
		if root.Left == nil {
			return root
		} else {
			return getMin(root.Left)
		}
	}

	var travel func(root *TreeNode, key int) *TreeNode
	travel = func(root *TreeNode, key int) *TreeNode {
		if root == nil {
			return nil
		} else if root.Val == key {
			//如果左右都为空 直接赋值nil
			if root.Left == nil && root.Right == nil {
				return nil
			} else if root.Left == nil {
				//如果一边为空 就赋另一边
				return root.Right
			} else if root.Right == nil {
				//如果一边为空 就赋另一边
				return root.Left
			} else {
				//如果两边非空 就找大于它的最小值
				min := getMin(root.Right)
				//交换
				root.Val, min.Val = min.Val, root.Val
				//这里很重要!!!!!   利用现有api 删掉大于它的最小值的那个节点
				root.Right = travel(root.Right, min.Val)
			}
		} else if root.Val < key {
			root.Right = travel(root.Right, key)
		} else if root.Val > key {
			root.Left = travel(root.Left, key)
		}
		return root
	}

	return travel(root, key)
}

//完全二叉树的性质 一定有一边是满二叉树 满二叉树可以直接用2的(h-1)次方来
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l, r := root.Left, root.Right
	lheight, rheight := 1, 1
	for l != nil {
		l = l.Left
		lheight++
	}
	for r != nil {
		r = r.Right
		rheight++
	}

	if lheight == rheight {
		return int(math.Pow(2, float64(lheight)) - 1)
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

//前序遍历
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	s := make([]string, 0, 0)
//
//	var travel func(root *TreeNode)
//	travel = func(root *TreeNode) {
//		if root == nil {
//			s = append(s, "#")
//			return
//		}
//		s = append(s, fmt.Sprintf("%v", root.Val))
//		travel(root.Left)
//		travel(root.Right)
//	}
//	travel(root)
//	return strings.Join(s, ",")
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//
//	i := 0
//
//	s := make([]string, 0, 0)
//	s = append(s, strings.Split(data, ",")...)
//
//	var travel func() *TreeNode
//	travel = func() *TreeNode {
//
//		if len(s) == 0 {
//			return nil
//		}
//
//		if s[i] == "#" {
//			i++
//			return nil
//		}
//
//		u, _ := strconv.Atoi(s[i])
//		i++
//
//		newnode := &TreeNode{Val: u}
//		newnode.Left = travel()
//		newnode.Right = travel()
//		return newnode
//	}
//
//	return travel()
//}

// Codec 后序遍历
//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (this *Codec) serialize(root *TreeNode) string {
//	res := make([]string, 0, 0)
//	var travel func(root *TreeNode)
//	travel = func(root *TreeNode) {
//		if root == nil {
//			res = append(res, "#")
//			return
//		}
//
//		travel(root.Left)
//		travel(root.Right)
//		res = append(res, fmt.Sprintf("%v", root.Val))
//	}
//	travel(root)
//	return strings.Join(res, ",")
//}
//
//// Deserializes your encoded data to tree.
//func (this *Codec) deserialize(data string) *TreeNode {
//	split := strings.Split(data, ",")
//	var travel func() *TreeNode
//	travel = func() *TreeNode {
//		if len(split) == 0 {
//			return nil
//		}
//		if split[len(split)-1] == "#" {
//			split = split[:len(split)-1]
//			return nil
//		}
//		newNode := &TreeNode{}
//		atoi, _ := strconv.Atoi(split[len(split)-1])
//		split = split[:len(split)-1]
//		newNode.Val = atoi
//		newNode.Right = travel()
//		newNode.Left = travel()
//		return newNode
//	}
//
//	return travel()
//}

// Codec 层序遍历
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	q := make([]*TreeNode, 0, 0)
	res := make([]string, 0, 0)
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			if node == nil {
				res = append(res, "#")
				q = q[1:]
				continue
			}
			res = append(res, fmt.Sprintf("%v", node.Val))
			q = q[1:]
			q = append(q, node.Left, node.Right)
		}
	}

	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	split := strings.Split(data, ",")
	if split[0] == "#" {
		return nil
	}
	atoi, _ := strconv.Atoi(split[0])
	split = split[1:]
	root := &TreeNode{Val: atoi}

	q := make([]*TreeNode, 0, 0)
	q = append(q, root)
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			left := split[0]
			right := split[1]
			split = split[2:]
			if left == "#" {
				node.Left = nil
			} else {
				lnum, _ := strconv.Atoi(left)
				leftNode := &TreeNode{Val: lnum}
				node.Left = leftNode
				q = append(q, leftNode)
			}
			if right == "#" {
				node.Right = nil
			} else {
				rnum, _ := strconv.Atoi(right)
				rightNode := &TreeNode{Val: rnum}
				node.Right = rightNode
				q = append(q, rightNode)
			}
		}
	}
	return root
}

//讲道理有点low 找到他们的公共祖先 然后逐个判断
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	parentP := make([]*TreeNode, 0, 0)
	parentQ := make([]*TreeNode, 0, 0)
	var travel func(root *TreeNode, target *TreeNode, path *[]*TreeNode) bool
	travel = func(root *TreeNode, target *TreeNode, path *[]*TreeNode) bool {
		if root == nil {
			*path = append(*path, nil)
			return false
		}

		if root == target {
			*path = append(*path, target)
			return true
		}

		*path = append(*path, root)
		if travel(root.Left, target, path) {
			return true
		}
		*path = (*path)[:len(*path)-1]

		if travel(root.Right, target, path) {
			return true
		}
		*path = (*path)[:len(*path)-1]
		return false
	}
	travel(root, p, &parentP)
	travel(root, q, &parentQ)
	m := make(map[*TreeNode]bool, 0)
	for i := 0; i < len(parentQ); i++ {
		m[parentQ[i]] = true
	}

	for i := len(parentP) - 1; i >= 0; i-- {
		if m[parentP[i]] {
			return parentP[i]
		}
	}
	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if root == p {
		return root
	} else if root == q {
		return root
	}
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
}

//单调栈  没想出来 还是要再学一学鸭
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				m[i] = j
				break
			}
		}
	}
	res := make([]int, len(nums2), len(nums2))
	stack := make([]int, 0, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	ans := make([]int, len(nums1), len(nums1))
	for i := 0; i < len(nums1); i++ {
		ans[i] = res[m[i]]
	}
	return ans

}

func a() {
	var travel func(a []int)
	travel = func(a []int) {
		fmt.Printf("%p\n", a)
		a = append(a, 1, 2, 3)
		fmt.Println(a)
		fmt.Printf("%p\n", a)
		fmt.Println(cap(a))
		fmt.Println(len(a))
	}
	ints := make([]int, 0, 3)
	fmt.Printf("%p\n", ints)
	travel(ints)
	fmt.Printf("%p\n", ints)
	fmt.Println(len(ints))
	fmt.Println(cap(ints))
	fmt.Println(ints)
}

//单调栈  通过控制传入参数来控制 可以再优化一下
func dailyTemperatures1(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))

	q := make([][2]int, 0, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(q) != 0 && temperatures[i] >= q[len(q)-1][0] {
			q = q[:len(q)-1]
		}
		if len(q) == 0 {
			res[i] = 0
		} else {
			res[i] = q[len(q)-1][1] - i
		}
		q = append(q, [2]int{temperatures[i], i})
	}

	return res
}

//单调栈  通过直接传入index 作为索引下标
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures), len(temperatures))

	q := make([]int, 0, 0)
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(q) != 0 && temperatures[i] >= temperatures[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		if len(q) == 0 {
			res[i] = 0
		} else {
			res[i] = q[len(q)-1] - i
		}
		q = append(q, i)
	}

	return res
}

func nextGreaterElements1(nums []int) []int {
	temp := make([]int, len(nums), 2*cap(nums))
	copy(temp, nums)
	temp = append(temp, nums...)
	res := make([]int, len(temp), len(temp))
	q := make([]int, 0, 0)

	for i := len(temp) - 1; i >= 0; i-- {
		for len(q) != 0 && temp[i] >= q[len(q)-1] {
			q = q[:len(q)-1]
		}

		if len(q) == 0 {
			res[i] = -1
		} else {
			res[i] = q[len(q)-1]
		}

		q = append(q, temp[i])
	}

	return res[:len(nums)]

}

func intersection(nums [][]int) []int {
	res := make([]int, 0, 0)
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		ints := nums[i]
		for j := 0; j < len(ints); j++ {
			m[ints[j]]++
		}
	}

	for key, v := range m {
		if v == len(nums) {
			res = append(res, key)
		}
	}
	sort.Ints(res)
	return res
}

func countLatticePoints(circles [][]int) int {
	visited := make(map[[2]int]bool)
	for _, circle := range circles {
		for i := 0; i < len(circle); i++ {
			points := [2]int{circle[0], circle[1]}
			radius := circle[2]
			for row := points[0] - radius; row <= points[0]+radius; row++ {
				for column := points[1] - radius; column <= points[1]+radius; column++ {
					if (row-points[0])*(row-points[0])+(column-points[1])*(column-points[1]) <= radius*radius {
						if !visited[[2]int{row, column}] {
							visited[[2]int{row, column}] = true
						}
					}
				}
			}
		}
	}
	return len(visited)
}

func countRectangles1(rectangles [][]int, points [][]int) []int {
	res := make([]int, 0, 0)

	m := make(map[[2]int]int)
	for _, rectangle := range rectangles {
		for i := 0; i <= rectangle[0]; i++ {
			for j := 0; j <= rectangle[1]; j++ {
				m[[2]int{i, j}]++
			}
		}
	}

	for i := 0; i < len(points); i++ {

		res = append(res, m[[2]int{points[i][0], points[i][1]}])
	}

	return res
}

func countRectangles2(rectangles [][]int, points [][]int) []int {
	res := make([]int, 0, 0)
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] == rectangles[j][0] {
			return rectangles[i][1] <= rectangles[j][1]
		} else {
			return rectangles[i][0] < rectangles[j][0]
		}
	})

	for _, point := range points {
		countx := 0
		for i := len(rectangles) - 1; i >= 0; i-- {
			if point[0] <= rectangles[i][0] && point[1] <= rectangles[i][1] {
				countx++
			} else if point[0] > rectangles[i][0] && point[1] > rectangles[i][1] {
				break
			}
		}
		res = append(res, countx)
	}

	return res
}

func countRectangles(rectangles [][]int, points [][]int) []int {
	res := make([]int, 0, 0)
	sort.Slice(rectangles, func(i, j int) bool {
		if rectangles[i][0] == rectangles[j][0] {
			return rectangles[i][1] <= rectangles[j][1]
		} else {
			return rectangles[i][0] >= rectangles[j][0]
		}
	})
	//temp11:=make([]int, len(rectangles), len(rectangles))
	//temp11=rectangles[:][0]
	for _, point := range points {
		countx := 0
		ints := sort.SearchInts(rectangles[:][0], point[0])

		for i := 0; i <= ints; i++ {
			if point[0] <= rectangles[i][0] && point[1] <= rectangles[i][1] {
				countx++
			} else if point[0] > rectangles[i][0] && point[1] > rectangles[i][1] {
				break
			}
		}
		res = append(res, countx)
	}

	return res
}

//使用%来表示循环数组
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums)*2, len(nums)*2)
	n := len(nums)
	q := make([]int, 0, 0)
	//假装数组翻倍了 实际上没有
	for i := len(res) - 1; i >= 0; i-- {
		for len(q) != 0 && q[len(q)-1] <= nums[i%n] {
			q = q[:len(q)-1]
		}

		if len(q) == 0 {
			res[i] = -1
		} else {
			res[i] = q[len(q)-1]
		}
		q = append(q, nums[i%n])
	}
	return res[:len(res)/2]
}

//func cronTimeGenerator(expression string, targetTime string) []string {
//	res := make([]string, 0, 0)
//	cron := cron.New()
//	cron.AddFunc(expression, func() {
//		res = append(res, targetTime)
//	})
//	cron.Start()
//	time.Sleep(time.Second * 10)
//	cron.Stop()
//	return res
//}

func gohome(a []int, b []int) int {
	res := 0

	var travel func([]int, int)
	travel = func(ints []int, index int) {
		if len(ints) == 0 {
			res++
		}

		for i := 0; i < len(ints); i++ {
			if ints[0] < a[index] {
				continue
			}
			ints[0], ints[i] = ints[i], ints[0]
			travel(ints[1:], index+1)
			ints[0], ints[i] = ints[i], ints[0]
		}
	}
	travel(b, 0)
	return res
}

type queue []int

func (q *queue) AddLast(n int) {
	for len(*q) != 0 && (*q)[len(*q)-1] < n {
		*q = (*q)[:len(*q)-1]
	}
	*q = append(*q, n)
}

func (q *queue) Max() int {
	return (*q)[0]
}

func (q *queue) Pop(n int) {
	if (*q)[0] == n {
		*q = (*q)[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	q := new(queue)
	res := make([]int, 0, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			(*q).AddLast(nums[i])
		} else {
			(*q).AddLast(nums[i])

			res = append(res, (*q).Max())
			(*q).Pop(nums[i-k+1])
		}
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//回文链表 时空复杂度O(n) 利用后序遍历  间接拿到尾巴的值  然后和left来做比较
func isPalindrome1(head *ListNode) bool {

	left := head
	var travel func(node *ListNode) bool
	travel = func(node *ListNode) bool {
		if node == nil {
			return true
		}

		b := travel(node.Next)
		b = b && node.Val == left.Val
		left = left.Next
		return b
	}

	return travel(head)
}

//空间复杂度为O（1）
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	begin := head

	end := reserve(slow)

	for end != nil {
		if begin.Val != end.Val {
			return false
		}
		begin = begin.Next
		end = end.Next
	}
	return true
}

//递归反转链表  返回原来的头部 也就是最后的尾部
func reserve1(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}
	temp := reserve(head.Next)
	temp.Next = head
	head.Next = nil
	return head
}

//迭代反转链表  返回原来的尾部 也就是最后的头部
func reserve(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
