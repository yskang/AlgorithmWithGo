package leetcode

import (
	"math"
)

// Search is a solution of the problem "33. Search in Rotated Sorted Array" in leetcode
func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if target == nums[0] {
		return 0
	}
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	start := 0
	end := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		check := int((end-start)/2 + start)
		numCheck := nums[check]

		if target > nums[0] {
			if numCheck < nums[0] {
				numCheck = math.MaxInt64
			}
		} else {
			if numCheck > nums[len(nums)-1] {
				numCheck = math.MinInt64
			}
		}

		if numCheck == target {
			return check
		} else if numCheck < target {
			start = check
		} else {
			end = check
		}
	}

	return -1
}
