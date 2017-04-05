package leetcode

import "fmt"

func HammingDistance() {
	fmt.Println(hammingDistance(1, 15))
}

func hammingDistance(x int, y int) int {
	z := x ^ y
	result := 0

	for i := 0 ; i < 32 ; i++{
		result += z & 1
		z >>= 1
	}

	return result
}

