package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {

	//indices := maxScoreIndices([]int{0, 0, 0})
	//fmt.Println(indices)

	//fmt.Println(len("pmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmm"))

	//hash := subStrHash("xqgfatvtlwnnkxipmipcpqwbxihxblaplpfckvxtihonijhtezdnkjmmk", 22, 51, 41, 9)
	//hash := subStrHash("leetcode", 7, 20, 3, 0)
	//hash := subStrHash("fbxzaad", 31, 100, 3, 32)
	//fmt.Println(hash)

	//sum := minimumSum(2932)
	//fmt.Println(sum)

	//array := pivotArray([]int{18, -13, -19, -11, 5, -17, 0, -18, -12, -6, -20, -8, -20, -4, 8}, -20)
	//fmt.Println(array)

	//time := minCostSetTime(0, 9, 18, 460)
	//fmt.Println(time)

	//steps := minSteps("cotxazilut", "nahrrmcchxwrieqqdwdpneitkxgnt")
	//fmt.Println(steps)

	//time := minimumTime([]int{1, 2, 3}, 6)
	//fmt.Println(time)

	time := minimumFinishTime([][]int{{45, 23}, {404, 3}, {735, 24}, {856, 3}, {249, 11}, {326, 7}, {589, 8}, {91, 17}, {126, 24}, {713, 21}, {606, 13}, {585, 5}, {861, 30}, {604, 4}, {822, 27}, {231, 6}, {507, 6}, {265, 5}, {912, 12}, {878, 29}, {46, 6}, {421, 20}, {941, 26}, {151, 25}, {490, 4}, {315, 14}, {630, 16}, {292, 24}, {214, 2}, {432, 18}, {520, 21}, {88, 26}, {4, 21}, {337, 28}, {780, 9}, {220, 29}, {721, 13}, {927, 25}, {67, 25}, {835, 21}, {646, 19}, {973, 26}, {235, 12}, {427, 9}, {471, 21}, {267, 16}, {388, 8}, {788, 13}, {937, 27}, {810, 26}, {288, 5}, {966, 28}, {698, 15}, {343, 12}, {648, 20}, {238, 4}, {436, 15}, {588, 4}, {373, 15}, {100, 12}, {180, 19}, {904, 15}, {854, 21}, {107, 2}, {822, 12}, {485, 24}, {196, 10}, {978, 24}, {178, 10}, {642, 29}, {455, 16}, {490, 17}, {455, 25}, {77, 7}, {456, 15}, {102, 25}, {767, 19}, {169, 5}, {461, 11}, {385, 25}, {896, 5}, {185, 26}, {885, 14}, {948, 26}, {907, 6}, {877, 18}, {421, 24}, {783, 15}, {999, 20}, {756, 10}, {308, 12}, {34, 12}, {339, 17}, {613, 15}, {270, 10}, {681, 3}, {385, 15}, {123, 4}, {10, 20}, {799, 28}, {506, 23}, {265, 24}, {193, 4}, {638, 23}, {144, 29}, {874, 19}, {470, 14}, {195, 16}, {77, 23}, {573, 28}, {559, 12}, {146, 10}, {538, 10}, {705, 4}, {592, 26}, {258, 24}, {900, 25}, {836, 8}, {353, 5}, {197, 26}, {572, 21}, {347, 17}, {763, 21}, {67, 20}, {927, 6}, {135, 4}, {392, 30}, {131, 15}, {650, 23}, {100, 30}, {848, 9}, {858, 27}, {203, 15}, {249, 4}, {884, 15}, {465, 18}, {316, 30}, {730, 15}, {310, 19}, {823, 21}, {785, 21}, {15, 16}}, 772, 502)
	//time := minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4)
	fmt.Println(time)
}

func findFinalValue(nums []int, original int) int {
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	for m[original] {
		original = original * 2
	}
	return original
}

//前缀和
func maxScoreIndices(nums []int) []int {
	res := make([]int, 0)
	preSum := make([]int, len(nums)+1)
	preSum[0] = 0
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	n := len(nums)
	sum := 0
	for i := 0; i < len(preSum); i++ {
		left := i - preSum[i]
		right := preSum[n] - preSum[i]
		count := left + right
		if count > sum {
			sum = count
			res = []int{i}
		} else if count == sum {
			sum = count
			res = append(res, i)
		} else {
			continue
		}
	}
	return res
}

//做不出来啊
func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	var calculate func(char uint8, p int, k int) int
	calculate = func(char uint8, p int, k int) int {
		m := 1
		for i := 0; i < k; i++ {
			m = m * (p % modulo)
		}
		return ((int(char-'a'+1) % modulo) * (m % modulo)) % modulo
	}

	for i := 0; i < len(s)-k+1; i++ {
		sum := 0
		temp := 0
		for j := i; j < i+k; j++ {
			sum += calculate(s[j], power, temp)
			temp++
		}
		if (sum % modulo) == hashValue {
			return s[i : i+k]
		}
	}
	return ""
}

func minimumSum(num int) int {
	a := num / 1000
	b := (num - a*1000) / 100
	c := (num - a*1000 - b*100) / 10
	d := num - a*1000 - b*100 - c*10
	m := make([]int, 4)
	m[0] = a
	m[1] = b
	m[2] = c
	m[3] = d
	sort.Ints(m)
	return m[0]*10 + m[3] + m[1]*10 + m[2]
}

func pivotArray1(nums []int, pivot int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != pivot {
			res = append(res, nums[i])
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i] < pivot && res[j] < pivot {
			return i < j
		} else if res[i] > pivot && res[j] > pivot {
			return i < j
		} else if res[i] < pivot && res[j] > pivot {
			return i > j
		} else {
			//nums[i] > pivot && nums[j] < pivot
			return j > i
		}
	})

	//begin := 0
	result := make([]int, 0)
	c := len(nums) - len(res)

	for i := 0; i < len(res); i++ {
		if res[i] < pivot {
			//begin = i
			result = append(result, res[i])
		} else {
			for c != 0 {
				result = append(result, pivot)
				c--
			}
			result = append(result, res[i])
		}
	}

	for c != 0 {
		result = append(result, pivot)
		c--
	}

	return result
}

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, 0)
	max := make([]int, 0)
	min := make([]int, 0)
	equal := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			min = append(min, nums[i])
		} else if nums[i] == pivot {
			equal = append(equal, nums[i])
		} else {
			max = append(max, nums[i])
		}
	}
	res = append(min, equal...)
	res = append(res, max...)
	return res
}

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	min := targetSeconds / 100
	max := targetSeconds / 60

	temp := make([][]int, 0, 0)
	minute := min
	for minute <= max {
		seconds := targetSeconds - minute*60
		if seconds < 100 {
			if minute >= 10 {
				temp = append(temp, []int{minute / 10, minute % 10, seconds / 10, seconds % 10})
			} else if minute == 0 && seconds >= 10 {
				temp = append(temp, []int{seconds / 10, seconds % 10})
			} else if minute == 0 && seconds < 10 {
				temp = append(temp, []int{seconds % 10})
			} else {
				temp = append(temp, []int{minute, seconds / 10, seconds % 10})
			}
		}
		minute++
	}

	res := math.MaxInt64
	for _, ints := range temp {
		count := 0
		b := false
		b1 := false
		countmove := 0

		for i := 0; i < len(ints); i++ {
			count++
			if ints[i] == startAt {
				b = true
			}
			if i-1 >= 0 && ints[i-1] == ints[i] {
				b1 = true
				countmove++
			}
		}
		total := 0
		for count != 0 {
			total += moveCost + pushCost
			count--
		}
		if b {
			total -= moveCost
		}
		if b1 {
			total -= moveCost * countmove
		}
		if res > total {
			res = total
		}
	}

	return res
}

func prefixCount(words []string, pref string) int {
	res := 0
	for _, word := range words {
		prefix := strings.HasPrefix(word, pref)
		if prefix {
			res++
		}
	}
	return res
}

func minSteps(s string, t string) int {
	res := 0
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		//if m[c] > 0 {
		m[c]--
		//} else if m[c] == 0 {
		//	m[c]++
		//}
	}

	for _, v := range m {
		if v == 0 {
			continue
		}
		res += int(math.Abs(float64(v)))
	}
	return res
}

func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)
	left, right := 1, (totalTrips/len(time)+1)*time[len(time)-1]
	for left <= right {
		mid := left + (right-left)/2
		temp := 0
		for _, v := range time {
			temp += mid / v
		}
		if temp < totalTrips {
			left = mid + 1
		} else if temp >= totalTrips {
			right = mid - 1
		} else if temp == totalTrips {
			return int64(mid)
		}
	}
	return int64(left)
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	sort.Slice(tires, func(i, j int) bool {
		return tires[i][0] < tires[j][0]
	})
	run := make([]int, numLaps, numLaps)

	for _, tire := range tires {
		first := 0
		last := 0
		for j := 0; j < numLaps; j++ {
			cost := tire[0] * int(math.Pow(float64(tire[1]), float64(j)))
			if j == 0 {
				first = cost
				last = cost
			} else if j >= 1 {
				if first+changeTime <= cost {
					break
				}
				cost += last
				last = cost
			}
			if run[j] == 0 || run[j] > cost {
				run[j] = cost
			}
		}
	}

	for i := len(run) - 1; i >= 0; i-- {
		if run[i] != 0 {
			break
		} else {
			run = run[:len(run)-1]
		}
	}

	res := make([]int, numLaps+1, numLaps+1)
	res[0] = 0
	for i := 0; i < len(run); i++ {
		res[i+1] = run[i]
	}

	for i := len(run) + 1; i <= numLaps; i++ {
		for j := 1; j <= len(run) && j <= i; j++ {
			a := res[i-j] + changeTime + run[j-1]
			temp := a
			if res[i] == 0 || temp < res[i] {
				res[i] = temp
			}
		}
	}
	return res[len(res)-1]
}

func minimumFinishTime1(tires [][]int, changeTime, numLaps int) int {
	minSec := [18]int{}
	for i := range minSec {
		minSec[i] = math.MaxInt32 / 2
	}
	for _, tire := range tires {
		f, r := tire[0], tire[1]
		for x, time, sum := 1, f, 0; time <= changeTime+f; x++ {
			sum += time
			minSec[x] = min(minSec[x], sum)
			time *= r
		}
	}

	f := make([]int, numLaps+1)
	f[0] = -changeTime
	for i := 1; i <= numLaps; i++ {
		f[i] = math.MaxInt32
		for j := 1; j <= 17 && j <= i; j++ {
			f[i] = min(f[i], f[i-j]+minSec[j])
		}
		f[i] += changeTime
	}
	return f[numLaps]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
