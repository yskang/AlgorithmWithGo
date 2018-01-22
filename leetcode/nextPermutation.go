package leetcode

import (
	"math"
	"sort"
)

// NextPermutation is a solution of the problem "31. Next Permutation" in leetcode
func NextPermutation(nums []int) []int {
	nextPermutation(nums)
	return nums
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	// find the position, nums[i-1] < nums[i]
	pos := 0
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			pos = i
			break
		}
		if i == 1 {
			reverseList(nums)
			return
		}
	}
	// find next number of i-1 th and sort lest numbers
	minNext := math.MaxInt64
	minIndex := 0
	for j := pos; j < len(nums); j++ {
		if nums[pos-1] < nums[j] && minNext > nums[j] {
			minNext = nums[j]
			minIndex = j
		}
	}
	nums[minIndex], nums[pos-1] = nums[pos-1], nums[minIndex]
	sort.Ints(nums[pos:])
}

func reverseList(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
