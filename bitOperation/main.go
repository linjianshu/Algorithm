package main

import (
	"fmt"
)

func main() {
	fmt.Println(clearLastOne(4))
	fmt.Println(isTwoMi(8))
	fmt.Println(countBit(7))
	fmt.Println(countBit1(7))
	fmt.Println(DifferentBit(10, 12))
	fmt.Println(OnceNum([]int{10, 12, 10, 12, 2, 1, 2}))
	fmt.Println(OnceTowNums([]int{1, 2, 2, 3, 4, 4, 5, 3}))
	fmt.Println(Swap(4, 5))
	fmt.Println(IsJi(5))
	fmt.Println(SwapHighLow(34520))
	fmt.Println(BitNiXu(8))

}

//消去x的最后一位的1  举个例子就好了 x & (x - 1)就是用来消除x二进制最后一位1的
func clearLastOne(x int) int {
	return x & (x - 1)
}

//检测整数n是否是2的幂次  2的幂次 x=> 1的个数只有一个 其他的都是0 也就是说 x-1的话本位是0其余的都是1 所以如果&运算的话应该==0
func isTwoMi(x int) bool {
	return x&(x-1) == 0
}

//计算在一个 32 位的整数的二进制表示中有多少个 1
func countBit(x int) int {
	count := 0
	for x != 0 {
		//代表消除掉1 消除掉几次 就代表有几个1
		x = x & (x - 1)
		count++
	}
	return count
}

//计算在一个 32 位的整数的二进制表示中有多少个 1
func countBit1(x int) int {
	count := 0
	for i := 0; i < 32; i++ {
		if x>>i&1 == 1 {
			count++
		}
	}
	return count
}

// DifferentBit 将整数A转换为B，需要改变多少个bit位
//先用异或^来取 不同位数的结果 然后再数结果中1的个数 就是变化的个数
func DifferentBit(x int, y int) int {
	res := 0
	z := x ^ y
	for z != 0 {
		z = z & (z - 1)
		res++
	}
	return res
}

// OnceNum 数组中，只有一个数出现一次，剩下都出现两次，找出出现一次的
//异或运算 就是找出两个数之间差异的地方的运算
//如果两数相等 那么异或运算就=0
//0与任何数字异或就是等于那个数字
func OnceNum(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res ^ nums[i]
	}
	return res
}

// OnceTowNums 数组中，只有两个数出现一次，剩下都出现两次，找出出现一次的
func OnceTowNums(nums []int) (x int, y int) {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	//res = x ^ y
	//其中肯定有为1的位 找到那一位
	pos := 0
	temp := res
	for temp>>pos&1 == 0 {
		pos++
	}
	//第pos位为1
	for i := 0; i < len(nums); i++ {
		if (nums[i] >> pos & 1) == 1 {
			//其中 x肯定在第pos位为1的数字中
			x ^= nums[i]
		}
	}

	//x求出来了 满足交换律
	y = res ^ x
	return
}

// Swap 交换两数 位运算满足 c = a^b  ; a = c^b ; b = c ^ a
func Swap(x int, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	return x, y
}

// IsJi 判断奇偶 拿最后一位来做与运算 判断最后一位是不是1就好了
func IsJi(x int) bool {
	return x&1 != 0
}

// SwapHighLow 位操作进行高低位交换  把|看作逻辑加法  把&看作逻辑乘法
func SwapHighLow(x uint16) uint16 {
	x = x<<8 | x>>8
	return x
}

//位操作进行二进制逆序
func BitNiXu(x uint8) uint8 {
	y := x
	y = y<<1 | y>>1
	y = y<<2 | y>>2
	y = y<<4 | y>>4
	return y
}
