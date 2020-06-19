package leetcode

import "math"

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

//执行用时：0 ms
func Reverse(x int) int {
	newNumber := 0
	negative := false
	if x < 0 {
		x = int(math.Abs(float64(x)))
		negative = true
	}
	for {
		newNumber = newNumber*10 + x%10
		if x < 10 {
			break
		}
		x = x / 10
	}

	if negative {
		if -float64(newNumber) < -(math.Pow(2, 31)) {
			return 0
		}
		return -newNumber
	} else {
		if float64(newNumber) > (math.Pow(2, 31) - 1) {
			return 0
		}
		return newNumber
	}
}
