package leetcode

import (
	"fmt"
	"math"
)

func ReverseInteger() {
	fmt.Println(reverse(-12345))
}

func reverse(x int) int {
	negative := false
	if x == 0  || x < math.MinInt32 || x > math.MaxInt32{
		return 0
	}else if x < 0 {
		x = -x
		negative = true
	}
	str := ""
	for x > 0 {
		num := '0' + x % 10
		x /= 10
		str += string(num)
	}

	result := 0
	l := len(str)
	for i, num := range(str) {
		result += int(num - '0') * int(math.Pow10(l-i-1))
	}

	if negative {
		return -1 * result
	}
	return result
}
