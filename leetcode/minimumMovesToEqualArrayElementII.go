package leetcode

import (
	"math"
	"sort"
)

func MinMovesToEqualArrayElementsII(nums []int) int {
	return minMoves2(nums)
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	median := nums[int(len(nums)/2)]

	sum := 0
	for _, num := range nums {
		sum += int(math.Abs(float64(num - median)))
	}

	return sum
}