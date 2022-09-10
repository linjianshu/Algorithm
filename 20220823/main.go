package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ints := [20]int{0, 50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 550, 600, 650, 700, 750, 800, 850, 900, 950}
	mm := map[int]string{
		0:   "redis_01",
		50:  "redis_02",
		100: "redis_03",
		150: "redis_04",
		200: "redis_05",
		250: "redis_06",
		300: "redis_07",
		350: "redis_08",
		400: "redis_09",
		450: "redis_10",
		500: "redis_11",
		550: "redis_12",
		600: "redis_13",
		650: "redis_14",
		700: "redis_15",
		750: "redis_16",
		800: "redis_17",
		850: "redis_18",
		900: "redis_19",
		950: "redis_20",
	}
	m_redis := map[string]bool{
		"redis_01": true,
		"redis_02": true,
		"redis_03": true,
		"redis_04": true,
		"redis_05": true,
		"redis_06": true,
		"redis_07": true,
		"redis_08": true,
		"redis_09": true,
		"redis_10": true,
		"redis_11": true,
		"redis_12": true,
		"redis_13": true,
		"redis_14": true,
		"redis_15": true,
		"redis_16": true,
		"redis_17": true,
		"redis_18": true,
		"redis_19": true,
		"redis_20": true,
	}
	add_ints := [20]int{25, 75, 125, 175, 225, 275, 325, 375, 425, 475, 525, 575, 625, 675, 725, 775, 825, 875, 925, 975}
	mm_add := map[int]string{
		25:  "add_redis_01",
		75:  "add_redis_02",
		125: "add_redis_04",
		175: "add_redis_05",
		225: "add_redis_06",
		275: "add_redis_07",
		325: "add_redis_08",
		375: "add_redis_09",
		425: "add_redis_10",
		475: "add_redis_11",
		525: "add_redis_12",
		575: "add_redis_13",
		625: "add_redis_14",
		675: "add_redis_15",
		725: "add_redis_16",
		775: "add_redis_17",
		825: "add_redis_18",
		875: "add_redis_19",
		925: "add_redis_20",
	}
	m_addredis := map[string]bool{
		"add_redis_01": false,
		"add_redis_02": false,
		"add_redis_03": false,
		"add_redis_04": false,
		"add_redis_05": false,
		"add_redis_06": false,
		"add_redis_07": false,
		"add_redis_08": false,
		"add_redis_09": false,
		"add_redis_10": false,
		"add_redis_11": false,
		"add_redis_12": false,
		"add_redis_13": false,
		"add_redis_14": false,
		"add_redis_15": false,
		"add_redis_16": false,
		"add_redis_17": false,
		"add_redis_18": false,
		"add_redis_19": false,
		"add_redis_20": false,
	}

	s := ""
	fmt.Scanf("%s\n", &s)
	split := strings.Split(s, ",")
	for i := 0; i < len(split); i++ {
		cmd := split[i][:1]
		content := split[i][2:]
		switch cmd {
		case "1":
			temp := content[len(content)-2:]
			atoi, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
			if m_redis[content] {
				fmt.Println(ints[atoi-1])
			} else {
				atoi = ints[atoi-1]
				for !m_addredis[mm_add[atoi+25]] && !m_redis[mm[atoi]] {
					atoi += 50
				}
				if m_addredis[mm_add[atoi+25]] {
					fmt.Println(atoi + 25)
				} else {
					//fmt.Println(ints[atoi])
					fmt.Println(atoi)
				}
			}
		case "2":
			sum := 0
			for j := 0; j < len(content); j++ {
				sum += int(content[j])
			}
			temp := sum % 999
			for k := temp; k < 1000; k++ {
				if m_addredis[mm_add[k]] {
					fmt.Println(k)
					break
				} else if m_redis[mm[k]] {
					fmt.Println(k)
					break
				}
			}
		case "3":
			i2 := strings.Split(content, ";")
			temptoken := i2[len(i2)-1]
			i3 := strings.Split(i2[0], ",")
			for q := 0; q < len(i3); q++ {
				s2 := i3[q]
				m_redis[s2] = false
				m_addredis[s2] = false
			}

			sum := 0
			for j := 0; j < len(temptoken); j++ {
				sum += int(temptoken[j])
			}
			temp := sum % 999
			for k := temp; k < 1000; k++ {
				if m_addredis[mm_add[k]] {
					fmt.Println(k)
					break
				} else if m_redis[mm[k]] {
					fmt.Println(k)
					break
				}
			}

		case "4":
			temp := content[len(content)-2:]
			atoi, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(add_ints[atoi-1])
		case "5":
			i2 := strings.Split(content, ";")
			temptoken := i2[len(i2)-1]
			i3 := strings.Split(i2[0], ",")
			for q := 0; q < len(i3); q++ {
				s2 := i3[q]
				m_redis[s2] = true
				m_addredis[s2] = true
			}

			sum := 0
			for j := 0; j < len(temptoken); j++ {
				sum += int(temptoken[j])
			}
			temp := sum % 999
			for k := temp; k < 1000; k++ {
				if m_addredis[mm_add[k]] {
					fmt.Println(k)
					break
				} else if m_redis[mm[k]] {
					fmt.Println(k)
					break
				}
			}
		}

	}

	//3:redis_10;huaweiuser1000000

	//fmt.Println(ints)
	//fmt.Println(add_ints)

	//fmt.Println(1)
	//num := 0
	//fmt.Scanf("%d\n", &num)
	//res := false
	//for i := 2; i <= 16; i++ {
	//	formatInt := strconv.FormatInt(int64(num), i)
	//	//fmt.Println(formatInt)
	//	b := false
	//	for j := 0; j < len(formatInt)/2; j++ {
	//		if formatInt[j] != formatInt[len(formatInt)-1-j] {
	//			b = true
	//		}
	//	}
	//	if b {
	//		continue
	//	} else {
	//		res = true
	//		fmt.Println(i)
	//	}
	//}
	//
	//if !res {
	//	fmt.Println(-1)
	//}
	//fmt.Println(res)

	////fmt.Println("Tom" > "Tomy")
	////fmt.Println("a" > "b")
	//
	//s := ""
	//_, err := fmt.Scanf("%s\n", &s)
	//if err != nil {
	//	fmt.Println("error.0001")
	//	return
	//}
	//split := strings.Split(s, ",")
	//m := make(map[string]int)
	//for i := 0; i < len(split); i++ {
	//	if split[i][0] >= 'A' && split[i][0] <= 'Z' {
	//		tempa := strings.ToLower(split[i])
	//		if tempa[1:] == split[i][1:] {
	//			m[split[i]]++
	//		} else {
	//			fmt.Println("error.0001")
	//			return
	//		}
	//	} else {
	//		fmt.Println("error.0001")
	//		return
	//	}
	//}
	//
	//max := 0
	//res := make([]string, 0, 0)
	//for key, value := range m {
	//	if value > max {
	//		res = []string{}
	//		res = append(res, key)
	//		max = value
	//	} else if value == max {
	//		res = append(res, key)
	//	}
	//}
	//
	//if len(res) == 1 {
	//	fmt.Println(res[0])
	//	return
	//}
	//
	//var travel func(i, j string, l, r int) bool
	//travel = func(i, j string, l, r int) bool {
	//	if l > len(i)-1 && r <= len(j)-1 {
	//		return true
	//	}
	//	if r > len(j)-1 && l <= len(i)-1 {
	//		return false
	//	}
	//
	//	if i[l] < j[r] {
	//		return true
	//	} else if i[l] > j[r] {
	//		return false
	//	} else {
	//		l++
	//		r++
	//		return travel(i, j, l, r)
	//	}
	//}
	//
	//sort.Slice(res, func(i, j int) bool {
	//	if strings.Contains(res[i], res[j]) {
	//		return false
	//	} else if strings.Contains(res[j], res[i]) {
	//		return true
	//	} else {
	//		l := 0
	//		r := 0
	//		return travel(res[i], res[j], l, r)
	//		//if res[i][l] < res[j][r] {
	//		//	return true
	//		//} else if res[i][l] == res[j][r] {
	//		//	l++
	//		//	r++
	//		//} else {
	//		//
	//		//}
	//	}
	//	return false
	//})
	//
	//fmt.Println(res[0])
}
