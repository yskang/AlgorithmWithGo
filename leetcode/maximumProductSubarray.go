package leetcode

import (
	"math"
)

func MaxProduct(nums []int) int {
	return maxProduct(nums)
}

func maxProduct(nums []int) int {
	maxValue := math.MinInt64
	dp := make([]int, len(nums))
	for start := 0; start < len(nums); start++ {
		for end := start; end < len(nums); end++ {
			if dp[end] != 0 {
				dp[end] = dp[end] / nums[start-1]
			} else {
				dp[end] = 1
				for i := start; i <= end; i++ {
					dp[end] *= nums[i]
				}
			}
			if dp[end] > maxValue {
				maxValue = dp[end]
			}
		}
	}
	return maxValue
}
