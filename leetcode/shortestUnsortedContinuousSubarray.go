package leetcode

import "sort"

func FindUnsortedSubarray(nums []int) int {
	return findUnsortedSubarray(nums)
}

func findUnsortedSubarray(nums []int) int {
	back := make([]int, len(nums))
	copy(back, nums)
	sort.Ints(nums)

	start, end := len(nums), 0

	for i := 0 ; i < len(nums) ; i++ {
		if back[i] != nums[i] {
			start = i
			break
		}
	}

	if start == len(nums) {
		return 0
	}

	for i := len(nums)-1 ; i >= 0 ; i-- {
		if back[i] != nums[i] {
			end = i
			break
		}
	}

	return end - start + 1
}