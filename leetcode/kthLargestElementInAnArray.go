package leetcode

import (
	"sort"
)

// FindKthLargest function find kth largest element in an array
func FindKthLargest(nums []int, k int) int {
	return findKthLargest(nums, k)
}

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}
