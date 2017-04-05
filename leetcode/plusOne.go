package leetcode

import "fmt"

func PlusOne() {
	fmt.Println(plusOne([]int{1,9}))
	fmt.Println(plusOne([]int{2,0}))
	fmt.Println(plusOne([]int{9,9}))
	fmt.Println(plusOne([]int{1,2}))
}
func plusOne(digits []int) interface{} {
	carry := 0

	for i := len(digits) - 1 ; i >= 0 ; i-- {
		plus := 0

		if i == len(digits) - 1 {
			plus = digits[i] + carry + 1
		} else {
			plus = digits[i] + carry
		}


		if plus >= 10 {
			digits[i] = plus - 10
			carry = 1
		} else {
			digits[i] = plus
			carry = 0
		}
	}

	if carry == 1 {
		return append([]int{carry}, digits...)
	}
	return digits
}